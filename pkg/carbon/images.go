package carbon

import (
	"fmt"
	"github.com/analog-substance/carbon/deployments"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/types"
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

func (c *Carbon) GetImageBuildTemplates() []string {
	var templates []string
	dirList, err := deployments.Files.ReadDir(common.DefaultPackerDirName)
	if err != nil {
		log.Debug("error getting packer templates", "err", err)
	}
	for _, dir := range dirList {
		templates = append(templates, dir.Name())
	}
	return templates
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

func (c *Carbon) GetImageBuild(name, provider, provisioner string) (types.ImageBuild, error) {
	imageBuilds, err := c.GetImageBuilds()
	if err != nil {
		return nil, err
	}

	for _, imageBuild := range imageBuilds {
		if imageBuild.Name() == name && imageBuild.ProviderType() == provider && imageBuild.Provisioner() == provisioner {
			return imageBuild, nil
		}
	}
	return nil, fmt.Errorf("image build  not found")
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
