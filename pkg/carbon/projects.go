package carbon

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/types"
	"os"
	"path/filepath"
)

func (c *Carbon) GetProjects() ([]types.Project, error) {
	if len(c.projects) == 0 {
		projectsBaseDir := common.ProjectsDir()
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
