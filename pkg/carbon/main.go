package carbon

import (
	"bytes"
	"embed"
	"fmt"
	builder "github.com/NoF0rte/cmd-builder"
	"github.com/analog-substance/carbon/deployments"
	"github.com/analog-substance/carbon/pkg/providers/aws"
	"github.com/analog-substance/carbon/pkg/providers/multipass"
	"github.com/analog-substance/carbon/pkg/providers/types"
	"github.com/analog-substance/carbon/pkg/providers/virtualbox"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"syscall"
)

type Options struct {
	Providers    []string
	Platforms    []string
	Environments []string
}

type Carbon struct {
	options      Options
	providers    []types.Provider
	platforms    []types.Platform
	environments []types.Environment
	machines     []types.VM
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

func (c *Carbon) Platforms() []types.Platform {
	if len(c.platforms) == 0 {
		c.platforms = []types.Platform{}
		for _, provider := range c.Providers() {
			c.platforms = append(c.platforms, provider.Platforms(c.options.Platforms...)...)
		}
	}

	return c.platforms
}

func (c *Carbon) GetVMs() []types.VM {
	if len(c.machines) == 0 {
		c.machines = []types.VM{}
		for _, platform := range c.Platforms() {
			for _, env := range platform.Environments(c.options.Environments...) {
				c.machines = append(c.machines, env.VMs()...)
			}
		}

	}
	return c.machines
}

func (c *Carbon) FindVMByID(id string) types.VM {
	for _, vm := range c.GetVMs() {
		if vm.ID() == id {
			return vm
		}
	}
	return nil
}

func (c *Carbon) FindVMByName(name string) types.VM {
	for _, vm := range c.GetVMs() {
		lowerName := strings.ToLower(vm.Name())
		name = strings.ToLower(name)
		if strings.Contains(lowerName, name) {
			return vm
		}
	}
	return nil
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
	autoinstall := false
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
		autoinstall = true
	}

	// determine user-data type (autoinstall or native cloud init)
	if autoinstall {
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

func (c *Carbon) BuildImage(name string) error {
	packerPath, err := exec.LookPath("packer")
	if err != nil {
		return err
	}

	args := []string{
		"packer",
		"build",
		path.Join(PackerDir, name),
	}

	//args = append(args, additionalArgs...)
	if //goland:noinspection GoBoolExpressions
	runtime.GOOS == "windows" {
		return builder.Cmd(args[0], args[1:]...).Interactive().Run()
	}
	return syscall.Exec(packerPath, args, os.Environ())
}

func (c *Carbon) GetImageBuilds() ([]string, error) {
	ret := []string{}
	listing, err := os.ReadDir(PackerDir)
	if err != nil {
		return ret, err
	}
	for _, file := range listing {
		ret = append(ret, file.Name())
	}
	return ret, nil
}

var availableProviders []types.Provider

func AvailableProviders() []types.Provider {
	if len(availableProviders) == 0 {
		allProviders := []types.Provider{
			aws.New(),
			//libvirt.New(),
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

	return bytes.ContainsAny(fileBytes, needle), nil
}

func copyFileFromEmbeddedFS(src, dest string, embeddedFS embed.FS) error {
	filebytes, err := embeddedFS.ReadFile(src)
	if err != nil {
		return err
	}
	err = os.WriteFile(dest, filebytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
