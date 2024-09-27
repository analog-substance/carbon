package carbon

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/analog-substance/carbon/deployments"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/providers/aws"
	"github.com/analog-substance/carbon/pkg/providers/multipass"
	"github.com/analog-substance/carbon/pkg/providers/qemu"
	"github.com/analog-substance/carbon/pkg/providers/simple"
	"github.com/analog-substance/carbon/pkg/providers/virtualbox"
	"github.com/analog-substance/carbon/pkg/types"
	"os"
	"path"
	"strings"
)

type Options struct {
	Providers    []string
	Profiles     []string
	Environments []string
}

type Carbon struct {
	options      Options
	providers    []types.Provider
	profiles     []types.Profile
	environments []types.Environment
	machines     []types.VM
	imageBuilds  []types.ImageBuild
	images       []types.Image
}

func New(options Options) *Carbon {

	carbon := &Carbon{options: options}

	if options.Providers == nil || len(options.Providers) == 0 {
		carbon.providers = AvailableProviders()
	} else {
		provs := []types.Provider{}
		for _, provider := range AvailableProviders() {
			for _, providerStr := range options.Providers {
				if strings.ToLower(providerStr) == strings.ToLower(provider.Name()) {
					provs = append(provs, provider)
				}
			}
		}
		carbon.providers = provs
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
			c.profiles = append(c.profiles, provider.Profiles(c.options.Profiles...)...)
		}
	}

	return c.profiles
}

func (c *Carbon) GetVMs() []types.VM {
	if len(c.machines) == 0 {
		c.machines = []types.VM{}
		for _, profile := range c.Profiles() {
			for _, env := range profile.Environments(c.options.Environments...) {
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

	simpleProvider := simple.New()
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
	bootstrappedDir := path.Join(PackerDir, name)
	err := os.MkdirAll(bootstrappedDir, 0755)
	if err != nil {
		return err
	}
	embeddedFS := deployments.Files
	tplPackerDir := path.Join("packer", tplDir)

	// copy packer file
	packerFilename := fmt.Sprintf("%s%s", service, PackerFileSuffixCloudInit)
	tplPackerFile := path.Join(tplPackerDir, packerFilename)
	bootstrappedPackerFile := path.Join(bootstrappedDir, packerFilename)
	err = copyFileFromEmbeddedFS(tplPackerFile, bootstrappedPackerFile, embeddedFS)
	if err != nil {
		return err
	}

	// don't care if it fails. file may not exist
	// copy packer vars
	packerVarsFilename := fmt.Sprintf("%s%s", service, PackerFileSuffixVariables)
	tplPackerVarsFile := path.Join(tplPackerDir, packerVarsFilename)
	bootstrappedPackerVarsFile := path.Join(bootstrappedDir, packerVarsFilename)
	_ = copyFileFromEmbeddedFS(tplPackerVarsFile, bootstrappedPackerVarsFile, embeddedFS)

	// copy local vars
	tplLocalVarsFile := path.Join(tplPackerDir, PackerFileLocalVars)
	bootstrappedLocalVarsFile := path.Join(bootstrappedDir, PackerFileLocalVars)
	err = copyFileFromEmbeddedFS(tplLocalVarsFile, bootstrappedLocalVarsFile, embeddedFS)
	if err != nil {
		return err
	}

	// copy private vars example
	tplPrivateVarsExampleFile := path.Join(tplPackerDir, PackerFilePrivateVarsExample)
	bootstrappedPrivateVarsExampleFile := path.Join(bootstrappedDir, PackerFilePrivateVarsExample)
	err = copyFileFromEmbeddedFS(tplPrivateVarsExampleFile, bootstrappedPrivateVarsExampleFile, embeddedFS)
	if err != nil {
		return err
	}

	// copy private vars example
	tplPackerFilePacker := path.Join(tplPackerDir, PackerFilePacker)
	bootstrappedPackerFilePacker := path.Join(bootstrappedDir, PackerFilePacker)
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
		err = copyFileFromEmbeddedFS(path.Join(tplPackerDir, PackerFileIsoVars), path.Join(bootstrappedDir, PackerFileIsoVars), embeddedFS)
		if err != nil {
			return err
		}

		// if we need iso vars, we also need autoinstall
		autoInstall = true
	}

	// determine user-data type (autoinstall or native cloud init)
	if autoInstall {
		cloudInitDir = path.Join(bootstrappedDir, CloudInitDir, "autoinstall")
		userDataFile = path.Join(tplPackerDir, CloudInitDir, "autoinstall", "user-data")
	} else {
		cloudInitDir = path.Join(bootstrappedDir, CloudInitDir, "default")
		userDataFile = path.Join(tplPackerDir, CloudInitDir, "default", "user-data")
	}

	err = os.MkdirAll(cloudInitDir, 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(cloudInitDir, "meta-data"), []byte{}, 0644)
	if err != nil {
		return err
	}

	err = copyFileFromEmbeddedFS(userDataFile, path.Join(cloudInitDir, "user-data"), embeddedFS)
	if err != nil {
		return err
	}

	// Eventually we need want to allow customizing the user-data
	// this could be useful for the future
	//
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
	//		filebytes, err := embeddedFS.ReadFile(path.Join(baseCloudInitDir, d.Name()))
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

func (c *Carbon) GetImageBuilds() ([]types.ImageBuild, error) {
	if len(c.imageBuilds) == 0 {
		c.imageBuilds = []types.ImageBuild{}
		for _, profile := range c.Profiles() {
			for _, env := range profile.Environments(c.options.Environments...) {
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
			for _, env := range profiles.Environments(c.options.Environments...) {
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
			if provider.IsAvailable() {
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
