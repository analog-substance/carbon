package models

import (
	"fmt"
	builder "github.com/NoF0rte/cmd-builder"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

type ImageBuild struct {
	buildPath    string
	providerType string
	provisioner  string
}

func NewImageBuild(buildPath, provider, provisioner string) *ImageBuild {
	return &ImageBuild{
		buildPath:    buildPath,
		providerType: provider,
		provisioner:  provisioner,
	}
}

func (b *ImageBuild) Name() string {
	return filepath.Base(filepath.Dir(b.buildPath))
}

func (b *ImageBuild) ProviderType() string {
	return b.providerType
}

func (b *ImageBuild) Provisioner() string {
	return b.provisioner
}

func (b *ImageBuild) Build() error {
	packerPath, err := exec.LookPath("packer")
	if err != nil {
		return err
	}

	args := []string{
		"packer",
		"build",
	}

	var packerConfig PackerConfig
	err = hclsimple.DecodeFile(b.buildPath, nil, &packerConfig)
	if err != nil {
		return err
	}
	args = append(args, "-only", fmt.Sprintf("%s.%s", packerConfig.Source.Type, packerConfig.Source.Name), filepath.Dir(b.buildPath))
	if //goland:noinspection GoBoolExpressions
	runtime.GOOS == "windows" {
		return builder.Cmd(args[0], args[1:]...).Interactive().Run()
	}
	return syscall.Exec(packerPath, args, os.Environ())
}

type SourceBlock struct {
	Type   string   `hcl:"type,label"`
	Name   string   `hcl:"name,label"`
	Config hcl.Body `hcl:",remain"`
}

type BuildBlock struct {
	Name        string   `hcl:"name,optional"`
	Description string   `hcl:"description,optional"`
	FromSources []string `hcl:"sources,optional"`
	Config      hcl.Body `hcl:",remain"`
}

type PackerConfig struct {
	Source SourceBlock `hcl:"source,block"`
	Build  BuildBlock  `hcl:"build,block"`
}

func GetImageBuildsForProvider(provider string) ([]types.ImageBuild, error) {
	ret := []types.ImageBuild{}
	imageBuilds, err := getImageBuilds()
	if err != nil {
		return ret, err
	}
	for _, imageBuild := range imageBuilds {
		if imageBuild.ProviderType() == provider {
			ret = append(ret, imageBuild)
		}
	}
	return ret, nil
}

func getImageBuilds() ([]types.ImageBuild, error) {
	ret := []types.ImageBuild{}
	packerDir := common.PackerDir()
	listing, err := os.ReadDir(packerDir)
	if err != nil {
		return ret, err
	}
	for _, dirEntry := range listing {
		if !dirEntry.IsDir() {
			continue
		}

		buildFiles, err := os.ReadDir(filepath.Join(packerDir, dirEntry.Name()))
		if err != nil {
			return ret, err
		}
		for _, buildFile := range buildFiles {
			provider := ""
			provisioner := ""
			if strings.HasSuffix(buildFile.Name(), "-cloud-init.pkr.hcl") {
				provider = strings.TrimSuffix(buildFile.Name(), "-cloud-init.pkr.hcl")
				provisioner = "cloud-init"
			} else if strings.HasSuffix(buildFile.Name(), "-ansible.pkr.hcl") {
				provider = strings.TrimSuffix(buildFile.Name(), "-ansible.pkr.hcl")
				provisioner = "ansible"
			} else {
				continue
			}
			ret = append(ret, NewImageBuild(
				filepath.Join(packerDir, dirEntry.Name(), buildFile.Name()),
				provider,
				provisioner))
		}
	}
	return ret, nil
}
