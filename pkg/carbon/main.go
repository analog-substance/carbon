package carbon

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/analog-substance/carbon/deployments"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/providers/aws"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/providers/multipass"
	"github.com/analog-substance/carbon/pkg/providers/qemu"
	"github.com/analog-substance/carbon/pkg/providers/virtualbox"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type Options struct {
	Providers    []string
	Profiles     []string
	Environments []string
}

type Carbon struct {
	config       common.CarbonConfig
	providers    []types.Provider
	profiles     []types.Profile
	environments []types.Environment
	machines     []types.VM
	imageBuilds  []types.ImageBuild
	images       []types.Image
	projects     []types.Project
}

var log *slog.Logger

func init() {
	log = common.WithGroup("carbon")
}

func New(config common.CarbonConfig) *Carbon {
	carbon := &Carbon{config: config, providers: []types.Provider{}, profiles: []types.Profile{}, environments: []types.Environment{}}

	for _, provider := range AvailableProviders() {
		providerConfig, ok := config.Providers[provider.Type()]
		if !ok {
			providerConfig = common.ProviderConfig{
				Enabled: true,
			}
		}

		log.Debug("adding provider", "provider", provider.Type(), "config_exists", ok, "enabled", providerConfig.Enabled)
		if providerConfig.Enabled {
			// no config, or explicitly enabled
			carbon.providers = append(carbon.providers, provider)
			provider.SetConfig(providerConfig)
		}
	}

	return carbon
}

func (c *Carbon) Providers() []types.Provider {
	return c.providers
}

func (c *Carbon) Profiles() []types.Profile {
	if len(c.profiles) == 0 {
		c.profiles = []types.Profile{}
		for _, provider := range c.Providers() {
			c.profiles = append(c.profiles, provider.Profiles()...)
		}
	}

	return c.profiles
}

func (c *Carbon) GetVMs() []types.VM {
	if len(c.machines) == 0 {
		c.machines = []types.VM{}
		for _, profile := range c.Profiles() {
			for _, env := range profile.Environments() {
				c.machines = append(c.machines, env.VMs()...)
			}
		}

	}
	return c.machines
}

func (c *Carbon) FindVMByID(id string) []types.VM {
	for _, vm := range c.GetVMs() {
		if vm.ID() == id {
			return []types.VM{vm}
		}
	}
	return []types.VM{}
}

func (c *Carbon) FindVMByName(name string) []types.VM {

	vms := []types.VM{}

	for _, vm := range c.GetVMs() {
		lowerName := strings.ToLower(vm.Name())
		name = strings.ToLower(name)
		if strings.Contains(lowerName, name) {
			vms = append(vms, vm)
		}
	}
	return vms
}

func (c *Carbon) VMsFromHosts(hostnames []string) []types.VM {

	simpleProvider := base.New()
	profile := simpleProvider.Profiles()
	envs := profile[0].Environments()

	vms := []types.VM{}
	for _, hostname := range hostnames {
		vms = append(vms, &models.Machine{
			InstanceName:       hostname,
			CurrentState:       types.StateUnknown,
			InstanceID:         hostname,
			Env:                envs[0],
			PublicIPAddresses:  []string{hostname},
			PrivateIPAddresses: []string{hostname},
		})
	}
	return vms
}

const CloudInitDir = "cloud-init"
const PackerDir = "deployments/packer"
const PackerFileSuffixCloudInit = "-cloud-init.pkr.hcl"
const PackerFileSuffixAnsible = "-ansible.pkr.hcl"
const PackerFileSuffixVariables = "-variables.pkr.hcl"
const PackerFilePrivateVarsExample = "private.auto.pkrvars.hcl.example"
const PackerFileIsoVars = "iso-variables.pkr.hcl"
const PackerFileLocalVars = "local-variables.pkr.hcl"
const PackerFilePacker = "packer.pkr.hcl"
const ISOVarUsage = "var.iso_url"

func (c *Carbon) CreateImageBuild(name, tplDir, service string) error {
	autoInstall := false
	cloudInitDir := ""
	userDataFile := ""

	// mkdir for new image build
	bootstrappedDir := filepath.Join(PackerDir, name)
	err := os.MkdirAll(bootstrappedDir, 0755)
	if err != nil {
		log.Debug("failed to create new packer build dir", "dir", bootstrappedDir, "err", err)
		return err
	}
	embeddedFS := deployments.Files
	tplPackerDir := filepath.Join("packer", tplDir)

	// copy packer file
	packerFilename := fmt.Sprintf("%s%s", service, PackerFileSuffixCloudInit)
	tplPackerFile := filepath.Join(tplPackerDir, packerFilename)
	bootstrappedPackerFile := filepath.Join(bootstrappedDir, packerFilename)
	err = copyFileFromEmbeddedFS(tplPackerFile, bootstrappedPackerFile, embeddedFS)
	if err != nil {
		log.Debug("failed to copy file from embedded fs", "tplPackerFile", tplPackerFile, "err", err)
		return err
	}

	// don't care if it fails. file may not exist
	// copy packer vars
	packerVarsFilename := fmt.Sprintf("%s%s", service, PackerFileSuffixVariables)
	tplPackerVarsFile := filepath.Join(tplPackerDir, packerVarsFilename)
	bootstrappedPackerVarsFile := filepath.Join(bootstrappedDir, packerVarsFilename)
	_ = copyFileFromEmbeddedFS(tplPackerVarsFile, bootstrappedPackerVarsFile, embeddedFS)

	// copy local vars
	tplLocalVarsFile := filepath.Join(tplPackerDir, PackerFileLocalVars)
	bootstrappedLocalVarsFile := filepath.Join(bootstrappedDir, PackerFileLocalVars)
	err = copyFileFromEmbeddedFS(tplLocalVarsFile, bootstrappedLocalVarsFile, embeddedFS)
	if err != nil {
		return err
	}

	// copy private vars example
	tplPrivateVarsExampleFile := filepath.Join(tplPackerDir, PackerFilePrivateVarsExample)
	bootstrappedPrivateVarsExampleFile := filepath.Join(bootstrappedDir, PackerFilePrivateVarsExample)
	err = copyFileFromEmbeddedFS(tplPrivateVarsExampleFile, bootstrappedPrivateVarsExampleFile, embeddedFS)
	if err != nil {
		return err
	}

	// copy private vars example
	tplPackerFilePacker := filepath.Join(tplPackerDir, PackerFilePacker)
	bootstrappedPackerFilePacker := filepath.Join(bootstrappedDir, PackerFilePacker)
	err = copyFileFromEmbeddedFS(tplPackerFilePacker, bootstrappedPackerFilePacker, embeddedFS)
	if err != nil {
		return err
	}

	// check for iso_vars in the packer file. so we can copy the variables over
	containsVars, err := fileContainsString(tplPackerFile, ISOVarUsage, embeddedFS)
	if err != nil {
		return err
	}
	if containsVars {
		err = copyFileFromEmbeddedFS(filepath.Join(tplPackerDir, PackerFileIsoVars), filepath.Join(bootstrappedDir, PackerFileIsoVars), embeddedFS)
		if err != nil {
			return err
		}

		// if we need iso vars, we also need autoinstall
		autoInstall = true
	}

	// determine user-data type (autoinstall or native cloud init)
	if autoInstall {
		cloudInitDir = filepath.Join(bootstrappedDir, CloudInitDir, "autoinstall")
		userDataFile = filepath.Join(tplPackerDir, CloudInitDir, "autoinstall", "user-data")
	} else {
		cloudInitDir = filepath.Join(bootstrappedDir, CloudInitDir, "default")
		userDataFile = filepath.Join(tplPackerDir, CloudInitDir, "default", "user-data")
	}

	err = os.MkdirAll(cloudInitDir, 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(cloudInitDir, "meta-data"), []byte{}, 0644)
	if err != nil {
		return err
	}

	err = copyFileFromEmbeddedFS(userDataFile, filepath.Join(cloudInitDir, "user-data"), embeddedFS)
	if err != nil {
		return err
	}

	return nil
}

func (c *Carbon) BuildImage(name, provider, provisioner string) error {
	imageBuilds, err := c.GetImageBuilds()
	if err != nil {
		return err
	}

	for _, imageBuild := range imageBuilds {
		if imageBuild.Name() == name && imageBuild.ProviderType() == provider && imageBuild.Provisioner() == provisioner {
			return imageBuild.Build()
		}
	}
	return fmt.Errorf("image build not found")
}

func (c *Carbon) GetImage(imageID string) (types.Image, error) {
	images, err := c.GetImages()
	if err != nil {
		return nil, err
	}

	for _, image := range images {
		if image.ID() == imageID {
			return image, nil
		}
	}
	return nil, fmt.Errorf("image not found")
}

func (c *Carbon) LaunchImage(name, imageID string) error {
	images, err := c.GetImages()
	if err != nil {
		return err
	}

	for _, image := range images {
		if image.ID() == imageID {
			return image.Launch(types.ImageLaunchOptions{Name: name})
		}
	}
	return fmt.Errorf("image not found")
}

func (c *Carbon) GetImageBuilds() ([]types.ImageBuild, error) {
	if len(c.imageBuilds) == 0 {
		c.imageBuilds = []types.ImageBuild{}
		for _, profile := range c.Profiles() {
			for _, env := range profile.Environments() {
				imageBuilds, err := env.ImageBuilds()
				if err != nil {
					return nil, err
				}
				c.imageBuilds = append(c.imageBuilds, imageBuilds...)
			}
		}
	}

	return c.imageBuilds, nil
}

func (c *Carbon) GetImages() ([]types.Image, error) {
	if len(c.images) == 0 {
		c.images = []types.Image{}
		for _, profiles := range c.Profiles() {
			for _, env := range profiles.Environments() {
				images, err := env.Images()
				if err != nil {
					return nil, err
				}
				c.images = append(c.images, images...)
			}
		}
	}

	return c.images, nil
}

func (c *Carbon) GetProjects() ([]types.Project, error) {
	if len(c.projects) == 0 {
		projectsBaseDir := viper.GetString(common.ViperTerraformProjectDir)
		dirListing, err := os.ReadDir(projectsBaseDir)
		if err != nil {
			return nil, err
		}
		c.projects = []types.Project{}
		for _, projectDir := range dirListing {
			if projectDir.IsDir() {
				c.projects = append(c.projects, models.NewProject(filepath.Join(projectsBaseDir, projectDir.Name())))
			}
		}
	}

	return c.projects, nil
}

func (c *Carbon) GetProject(name string) (types.Project, error) {
	projects, err := c.GetProjects()
	if err != nil {
		return nil, err
	}
	for _, project := range projects {
		if project.Name() == name {
			return project, nil
		}
	}
	return nil, fmt.Errorf("project not found")
}

/*
func (c *Carbon) UpdateImageBuildCloudInit(imagebuildDir string) error {

		cloudInitDir := filepath.Join(imagebuildDir, "cloud-init")
		autoInstallFile := filepath.Join(cloudInitDir, "autoinstall", "user-data")
		//defaultCloudInit := filepath.Join(cloudInitDir, "default", "user-data")
		carbonInitFileBytes, err := deployments.Files.ReadFile("carbon/carbon-init")
		if err != nil {
			return err
		}

		autoInstallFileBytes, err := os.ReadFile(autoInstallFile)
		if err != nil {
			return err
		}

		var autoInstallCloudInit cloud_init.CloudConfig
		err = yaml.Unmarshal(autoInstallFileBytes, autoInstallCloudInit)
		if err != nil {
			return err
		}

		for i, runCmd := range autoInstallCloudInit.Runcmd {
			for ii, f := range runCmd {
				if strings.Contains(f, "CARBON_INIT_B64=") {
					autoInstallCloudInit.Runcmd[i][ii] = fmt.Sprintf("CARBON_INIT_B64=%s", base64.StdEncoding.EncodeToString(compress..carbonInitFileBytes))
					break
				}
			}
		}

		d, err := yaml.Marshal(&autoInstallCloudInit)
		if err != nil {
			return err
		}

		err = os.WriteFile(autoInstallFile, d, 0644)
		if err != nil {
			return err
		}

		//cloudInitListing, err := embeddedFS.ReadDir(baseCloudInitDir)
		//if err != nil {
		//	return err
		//}
		//packerListing, err := embeddedFS.ReadDir(tplPackerDir)
		//if err != nil {
		//	return err
		//}
		//
		//tpls := map[string]*cloud_init.CloudConfig{}
		//endResult := &cloud_init.CloudConfig{}
		//for _, d := range cloudInitListing {
		//	if strings.HasSuffix(d.Name(), ".yaml") {
		//		filebytes, err := embeddedFS.ReadFile(filepath.Join(baseCloudInitDir, d.Name()))
		//		if err != nil {
		//			log.Fatal(err)
		//		}
		//
		//		tpls[d.Name()] = &cloud_init.CloudConfig{}
		//
		//		err = yaml.Unmarshal(filebytes, tpls[d.Name()])
		//		if err != nil {
		//			return err
		//		}
		//
		//		endResult.MergeWith(tpls[d.Name()])
		//	}
		//}
		//d, err := yaml.Marshal(&endResult)
		//if err != nil {
		//	return err
		//}
		//
		//err = os.WriteFile(path.Join(cloudInitDir, "user-data"), d, 0644)
		//if err != nil {
		//	return err
		//}
		return nil
	}
*/
var availableProviders []types.Provider

func AvailableProviders() []types.Provider {
	if len(availableProviders) == 0 {
		allProviders := []types.Provider{
			aws.New(),
			qemu.New(),
			virtualbox.New(),
			multipass.New(),
		}

		for _, provider := range allProviders {
			isAvailable := provider.IsAvailable()
			log.Debug("assessing provider availability", "provider", provider.Type(), "isAvailable", isAvailable)
			if isAvailable {
				availableProviders = append(availableProviders, provider)
			}
		}
	}

	return availableProviders
}

func fileContainsString(path string, needle string, embeddedFS embed.FS) (bool, error) {
	fileBytes, err := embeddedFS.ReadFile(path)
	if err != nil {
		return false, err
	}

	return bytes.Contains(fileBytes, []byte(needle)), nil
}

func copyFileFromEmbeddedFS(src, dest string, embeddedFS embed.FS) error {
	fileBytes, err := embeddedFS.ReadFile(src)
	if err != nil {
		return err
	}
	err = os.WriteFile(dest, fileBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
