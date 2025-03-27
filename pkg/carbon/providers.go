package carbon

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/analog-substance/carbon/deployments"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/providers/aws"
	"github.com/analog-substance/carbon/pkg/providers/digitalocean"
	"github.com/analog-substance/carbon/pkg/providers/gcloud"
	"github.com/analog-substance/carbon/pkg/providers/multipass"
	"github.com/analog-substance/carbon/pkg/providers/qemu"
	"github.com/analog-substance/carbon/pkg/providers/virtualbox"
	"github.com/analog-substance/carbon/pkg/providers/vsphere"
	"github.com/analog-substance/carbon/pkg/types"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

var availableProviders []types.Provider

func AvailableProviders() []types.Provider {
	defer (common.Time("available providers"))()

	if len(availableProviders) == 0 {

		type providerAvailability struct {
			provider  types.Provider
			available bool
		}
		c := make(chan providerAvailability)
		for _, provider := range AllProviders {
			go func() {
				//defer (common.Time(fmt.Sprintf("(%s)provider.isAvailable", provider.Name())))()
				c <- providerAvailability{
					provider:  provider,
					available: provider.IsAvailable(),
				}
			}()
		}

		result := make([]providerAvailability, len(AllProviders))
		for i, _ := range result {
			result[i] = <-c
			log.Debug("assessing provider availability", "provider", result[i].provider.Type(), "isAvailable", result[i].available)
			if result[i].available {
				availableProviders = append(availableProviders, result[i].provider)
			}
		}
	}

	return availableProviders
}

func (c *Carbon) Providers() []types.Provider {
	return c.providers
}

func (c *Carbon) GetProvider(providerType string) (types.Provider, error) {
	for _, provider := range c.Providers() {
		if provider.Type() == providerType {
			return provider, nil
		}
	}

	return nil, fmt.Errorf("provider '%s' not found", providerType)
}

func (c *Carbon) NewProject(name string, providerType string, force bool) (types.Project, error) {

	if err := c.CopyTFModule(); err != nil {
		return nil, err
	}

	baseName := filepath.Base(name)
	projectDir := filepath.Join(common.GetConfig().Carbon.Dir[common.TerraformProjectConfigKey], baseName)
	log.Debug("new project", "dir", projectDir)

	_, err := os.Stat(projectDir)
	if err == nil && !force {
		return nil, fmt.Errorf("project %s already exists", name)
	}

	err = os.MkdirAll(projectDir, 0755)
	if err != nil {
		log.Debug("failed to create project dir", "dir", projectDir, "err", err)
		return nil, err
	}

	embeddedDir := path.Join(common.DefaultProjectsDirName, "example")
	err = copyTemplateDeploymentsFS(embeddedDir, projectDir, ImageBuildDate{name})
	if err != nil {
		log.Debug("failed to copy embedded terraform dir", "dir", embeddedDir, "err", err)
		return nil, err
	}

	return models.NewProject(projectDir), nil
}

func (c *Carbon) CopyTFModule() error {

	terraformDir := common.GetConfig().Carbon.Dir[common.DefaultTerraformDirName]
	log.Debug("copy terraform", "dir", terraformDir)

	err := os.MkdirAll(filepath.Join(terraformDir, "modules", "carbon"), 0755)
	if err != nil {
		log.Debug("failed to create terraformDir dir", "dir", terraformDir, "err", err)
		return err
	}

	embeddedDir := path.Join(common.DefaultTerraformDirName, "modules", "carbon")
	err = copyTemplateDeploymentsFS(embeddedDir, filepath.Join(terraformDir, "modules", "carbon"), struct {
		Version string
	}{
		Version: c.version,
	})
	if err != nil {
		log.Debug("failed to copy embedded terraform dir", "dir", embeddedDir, "err", err)
		return err
	}

	return nil
}

var AllProviders = []types.Provider{
	aws.New(),
	digitalocean.New(),
	gcloud.New(),
	multipass.New(),
	qemu.New(),
	virtualbox.New(),
	vsphere.New(),
}

func copyTemplateDeploymentsFS(embeddedDir string, dest string, templateData any) error {
	dirListing, err := deployments.Files.ReadDir(embeddedDir)
	if err != nil {
		log.Debug("failed to read embedded project dir", "dir", common.DefaultProjectsDirName, "err", err)
		return err
	}
	for _, d := range dirListing {
		if !d.IsDir() {
			err = copyTemplateFromEmbeddedFS(
				path.Join(embeddedDir, d.Name()),
				filepath.Join(dest, d.Name()),
				deployments.Files,
				templateData,
			)
			if err != nil {
				log.Debug("failed to copy embedded project file", "file", d.Name(), "err", err)
				return err
			}

		}
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
