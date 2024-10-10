package cloud_init

type AptSource struct {
	Source string `yaml:"source"`
	Keyid  string `yaml:"keyid"`
}

type WriteFile struct {
	Path        string `yaml:"path"`
	Content     string `yaml:"content"`
	Owner       string `yaml:"owner"`
	Permissions string `yaml:"permissions"`
	Encoding    string `yaml:"encoding,omitempty"`
}

type CloudConfig struct {
	Timezone          string   `yaml:"timezone"`
	SSHDeletekeys     bool     `yaml:"ssh_deletekeys"`
	SSHAuthorizedKeys []string `yaml:"ssh_authorized_keys"`
	Apt               struct {
		Sources map[string]AptSource `yaml:"sources"`
	} `yaml:"apt"`
	WriteFiles     []WriteFile `yaml:"write_files"`
	PackageUpgrade bool        `yaml:"package_upgrade"`
	Packages       []string    `yaml:"packages"`
	Runcmd         [][]string  `yaml:"runcmd"`
}

func (c *CloudConfig) MergeWith(otherConfig *CloudConfig) {
	if otherConfig.Timezone != "" {
		c.Timezone = otherConfig.Timezone
	}

	if otherConfig.SSHDeletekeys {
		c.SSHDeletekeys = otherConfig.SSHDeletekeys
	}

	if otherConfig.SSHAuthorizedKeys != nil {
		c.SSHAuthorizedKeys = uniqStringSlice(append(c.SSHAuthorizedKeys, otherConfig.SSHAuthorizedKeys...)...)
	}

	if otherConfig.SSHAuthorizedKeys != nil {
		c.SSHAuthorizedKeys = uniqStringSlice(append(c.SSHAuthorizedKeys, otherConfig.SSHAuthorizedKeys...)...)
	}

	if otherConfig.Apt.Sources != nil {
		if c.Apt.Sources == nil {
			c.Apt.Sources = otherConfig.Apt.Sources
		} else {

			for sourceName, source := range otherConfig.Apt.Sources {
				c.Apt.Sources[sourceName] = source
			}
		}
	}

	if otherConfig.WriteFiles != nil {
		c.WriteFiles = uniqWriteFiles(c.WriteFiles, otherConfig.WriteFiles)
	}

	if otherConfig.Packages != nil {
		c.Packages = uniqStringSlice(append(c.Packages, otherConfig.Packages...)...)
	}
	if otherConfig.Runcmd != nil {
		c.Runcmd = append(c.Runcmd, otherConfig.Runcmd...)
	}

}

func uniqStringSlice(s ...string) []string {
	m := map[string]bool{}
	for _, v := range s {
		m[v] = true
	}
	r := []string{}
	for s := range m {
		r = append(r, s)
	}
	return r
}

func uniqWriteFiles(writeFile1, writeFile2 []WriteFile) []WriteFile {
	fileToWrite := map[string]WriteFile{}
	for _, wf1 := range writeFile1 {
		fileToWrite[wf1.Path] = wf1
	}
	for _, wf2 := range writeFile2 {
		fileToWrite[wf2.Path] = wf2
	}
	returnWriteFiles := []WriteFile{}
	for _, wf1 := range fileToWrite {
		returnWriteFiles = append(returnWriteFiles, wf1)
	}
	return returnWriteFiles
}
