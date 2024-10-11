package models

import (
	"encoding/json"
	"fmt"
	builder "github.com/NoF0rte/cmd-builder"
	"github.com/analog-substance/carbon/pkg/types"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Project struct {
	buildPath           string
	config              *types.ProjectConfig
	lastKnownGoodConfig []byte
}

func NewProject(buildPath string) *Project {
	return &Project{
		buildPath: buildPath,
	}
}

func (d *Project) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name string `json:"name"`
	}{
		Name: d.Name(),
	})
}

func (d *Project) Name() string {
	return filepath.Base(d.buildPath)
}

func (d *Project) TerraformApply() error {
	terraformPath, err := exec.LookPath("terraform")
	if err != nil {
		return err
	}

	args := []string{
		"apply",
	}

	return builder.Cmd(terraformPath, args...).Dir(d.buildPath).Interactive().Run()
}

func (d *Project) AddMachine(machine *types.ProjectMachine, noApply bool) error {
	_, err := d.GetConfig()
	if err != nil {
		return err
	}

	for _, existingMachine := range d.config.Machines {
		if existingMachine.Name == machine.Name {
			return fmt.Errorf("machine '%s' already exists", machine.Name)
		}
	}

	d.config.Machines = append(d.config.Machines, machine)
	err = d.SaveConfig()
	if err != nil {
		return err
	}

	factory := builder.NewFactory(builder.CmdFactoryOptions{
		Stdout: os.Stdout,
		Dir:    d.buildPath,
	})

	tfPlanTempFile, err := os.CreateTemp("", "carbon-tf-apply-add-machine-")
	if err != nil {
		return fmt.Errorf("CreateTemp error: %v", err)
	}
	defer os.Remove(tfPlanTempFile.Name())
	tfPlanTempFile.Close()

	err = factory.Cmd("terraform", "plan", "-out", tfPlanTempFile.Name()).Run()
	if err != nil {
		return fmt.Errorf("error running terraform plan: %v", err)
	}

	terraformShowOutput, err := factory.Cmd("terraform", "show", tfPlanTempFile.Name()).Output()
	if err != nil {
		return fmt.Errorf("error running terraform show: %v", err)
	}

	if !strings.Contains(terraformShowOutput, " 0 to change, 0 to destroy") {
		_ = d.rollBackConfig()
		return fmt.Errorf("terraform plan had other changes")
	}

	err = factory.Cmd("terraform", "apply", tfPlanTempFile.Name()).Run()
	if err != nil {
		return fmt.Errorf("error running terraform apply: %v", err)
	}

	return nil
}

func (d *Project) GetConfig() (*types.ProjectConfig, error) {
	if d.config == nil {
		configFilePath := filepath.Join(d.buildPath, "carbon-config.yaml")
		fileBytes, err := os.ReadFile(configFilePath)
		if err != nil {
			return nil, err
		}

		d.config = &types.ProjectConfig{}
		if err := yaml.Unmarshal(fileBytes, d.config); err != nil {
			return nil, err
		}
		d.lastKnownGoodConfig = fileBytes
	}
	return d.config, nil
}

func (d *Project) SaveConfig() error {
	if d.config != nil {
		yamlConfig, err := yaml.Marshal(d.config)
		if err != nil {
			return err
		}
		return d.saveConfigBytes(yamlConfig)
	}
	return fmt.Errorf("no config loaded")
}

func (d *Project) rollBackConfig() error {
	if d.lastKnownGoodConfig != nil {
		return d.saveConfigBytes(d.lastKnownGoodConfig)
	}
	return fmt.Errorf("no config loaded")
}

func (d *Project) saveConfigBytes(filebytes []byte) error {
	configFilePath := filepath.Join(d.buildPath, "carbon-config.yaml")
	if err := os.WriteFile(configFilePath, filebytes, 0644); err != nil {
		return err
	}
	return nil
}
