package carbon

import "github.com/analog-substance/carbon/pkg/types"

func (c *Carbon) Profiles() []types.Profile {
	if len(c.profiles) == 0 {
		c.profiles = []types.Profile{}
		for _, provider := range c.Providers() {
			c.profiles = append(c.profiles, provider.Profiles()...)
		}
	}

	return c.profiles
}
