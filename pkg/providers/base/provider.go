package base

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/analog-substance/carbon/deployments"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/spf13/viper"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

const CloudInitDir = "cloud-init"
const PackerFileSuffixCloudInit = "-cloud-init.pkr.hcl"
const PackerFileSuffixAnsible = "-ansible.pkr.hcl"
const PackerFileSuffixVariables = "-variables.pkr.hcl"
const PackerFilePrivateVarsExample = "private.auto.pkrvars.hcl.example"
const PackerFileIsoVars = "iso-variables.pkr.hcl"
const PackerFileLocalVars = "local-variables.pkr.hcl"
const PackerFilePacker = "packer.pkr.hcl"
const ISOVarUsage = "var.iso_url"

const providerName = "Base"

type Provider struct {
	name string
	//profiles []string
	config common.ProviderConfig
}

func New() types.Provider {
	return &Provider{
		name: providerName,
	}
}

func NewWithName(name string) types.Provider {
	return &Provider{
		name: name,
	}
}

func (p *Provider) IsAvailable() bool {
	return false
}

func (p *Provider) Name() string {
	return p.name
}

func (p *Provider) Type() string {
	return strings.ToLower(p.Name())
}

func (p *Provider) Profiles() []types.Profile {
	return []types.Profile{
		&Profile{
			profileName: fmt.Sprintf("%s Profile", p.name),
			provider:    p,
		},
	}
}

func (p *Provider) SetConfig(config common.ProviderConfig) {
	p.config = config
}

func (p *Provider) GetConfig() common.ProviderConfig {
	return p.config
}

func (p *Provider) NewImageBuild(name, tplDir string) (types.ImageBuild, error) {
	autoInstall := false
	cloudInitDir := ""
	userDataFile := ""

	// mkdir for new image build
	bootstrappedDir := filepath.Join(common.PackerDir(), name)
	err := os.MkdirAll(bootstrappedDir, 0755)
	if err != nil {
		log.Debug("failed to create new packer build dir", "dir", bootstrappedDir, "err", err)
		return nil, err
	}
	embeddedFS := deployments.Files
	tplPackerDir := filepath.Join(common.DefaultPackerDirName, tplDir)

	// copy packer file
	packerFilename := fmt.Sprintf("%s%s", p.Type(), PackerFileSuffixCloudInit)
	tplPackerFile := filepath.Join(tplPackerDir, packerFilename)
	bootstrappedPackerFile := filepath.Join(bootstrappedDir, packerFilename)
	err = copyFileFromEmbeddedFS(tplPackerFile, bootstrappedPackerFile, embeddedFS)
	if err != nil {
		log.Debug("failed to copy file from embedded fs", "tplPackerFile", tplPackerFile, "err", err)
		return nil, err
	}

	// don't care if it fails. file may not exist
	// copy packer vars
	packerVarsFilename := fmt.Sprintf("%s%s", p.Type(), PackerFileSuffixVariables)
	tplPackerVarsFile := filepath.Join(tplPackerDir, packerVarsFilename)
	bootstrappedPackerVarsFile := filepath.Join(bootstrappedDir, packerVarsFilename)
	_ = copyFileFromEmbeddedFS(tplPackerVarsFile, bootstrappedPackerVarsFile, embeddedFS)

	// copy local vars
	tplLocalVarsFile := filepath.Join(tplPackerDir, PackerFileLocalVars)
	bootstrappedLocalVarsFile := filepath.Join(bootstrappedDir, PackerFileLocalVars)
	err = copyFileFromEmbeddedFS(tplLocalVarsFile, bootstrappedLocalVarsFile, embeddedFS)
	if err != nil {
		return nil, err
	}

	// copy private vars example
	tplPrivateVarsExampleFile := filepath.Join(tplPackerDir, PackerFilePrivateVarsExample)
	bootstrappedPrivateVarsExampleFile := filepath.Join(bootstrappedDir, PackerFilePrivateVarsExample)
	err = copyFileFromEmbeddedFS(tplPrivateVarsExampleFile, bootstrappedPrivateVarsExampleFile, embeddedFS)
	if err != nil {
		return nil, err
	}

	// copy private vars example
	tplPackerFilePacker := filepath.Join(tplPackerDir, PackerFilePacker)
	bootstrappedPackerFilePacker := filepath.Join(bootstrappedDir, PackerFilePacker)
	err = copyFileFromEmbeddedFS(tplPackerFilePacker, bootstrappedPackerFilePacker, embeddedFS)
	if err != nil {
		return nil, err
	}

	// check for iso_vars in the packer file. so we can copy the variables over
	containsVars, err := fileContainsString(tplPackerFile, ISOVarUsage, embeddedFS)
	if err != nil {
		return nil, err
	}
	if containsVars {
		err = copyFileFromEmbeddedFS(filepath.Join(tplPackerDir, PackerFileIsoVars), filepath.Join(bootstrappedDir, PackerFileIsoVars), embeddedFS)
		if err != nil {
			return nil, err
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
		return nil, err
	}

	err = os.WriteFile(filepath.Join(cloudInitDir, "meta-data"), []byte{}, 0644)
	if err != nil {
		return nil, err
	}

	err = copyFileFromEmbeddedFS(userDataFile, filepath.Join(cloudInitDir, "user-data"), embeddedFS)
	if err != nil {
		return nil, err
	}

	return models.NewImageBuild(bootstrappedDir, p.Type(), "cloud-init"), nil
}

func (p *Provider) NewProject(name string, force bool) (types.Project, error) {
	baseName := filepath.Base(name)
	projectDir := filepath.Join(viper.GetString(common.ViperTerraformProjectDir), baseName)
	log.Debug("new project", "dir", projectDir, "service", p.Type())

	_, err := os.Stat(projectDir)
	if err == nil && !force {
		return nil, fmt.Errorf("project %s already exists", name)
	}

	err = os.MkdirAll(projectDir, 0755)
	if err != nil {
		log.Debug("failed to create project dir", "dir", projectDir, "service", p.Type(), "err", err)
		return nil, err
	}

	embeddedDir := filepath.Join(common.DefaultProjectsDirName, p.Type())
	dirListing, err := deployments.Files.ReadDir(embeddedDir)
	if err != nil {
		log.Debug("failed to read embedded project dir", "dir", common.DefaultProjectsDirName, "service", p.Type(), "err", err)
		return nil, err
	}
	for _, d := range dirListing {
		if !d.IsDir() {
			err = copyTemplateFromEmbeddedFS(
				filepath.Join(embeddedDir, d.Name()),
				filepath.Join(projectDir, d.Name()),
				deployments.Files,
				ImageBuildDate{name},
			)
			if err != nil {
				log.Debug("failed to copy embedded project file", "file", d.Name(), "service", p.Type(), "err", err)
				return nil, err
			}

		}
	}

	return models.NewProject(projectDir), nil
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

func copyTemplateFromEmbeddedFS(src, dest string, embeddedFS embed.FS, data any) error {
	fileBytes, err := embeddedFS.ReadFile(src)
	if err != nil {
		return err
	}

	tmpl, err := template.New(src).Parse(string(fileBytes))
	if err != nil {
		return err
	}

	var templateBytes bytes.Buffer
	err = tmpl.Execute(&templateBytes, data)
	if err != nil {
		return err
	}

	err = os.WriteFile(dest, templateBytes.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

type ImageBuildDate struct {
	Name string
}
