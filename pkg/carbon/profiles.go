package carbon

import (
	"github.com/analog-substance/carbon/pkg/types"
	"sync"
)

func (c *Carbon) Profiles() []types.Profile {
	if len(c.profiles) == 0 {
		c.profiles = []types.Profile{}
		mu := sync.Mutex{}
		wait := sync.WaitGroup{}
		for _, provider := range c.Providers() {
			wait.Add(1)
			go func() {
				profiles := provider.Profiles()
				mu.Lock()
				c.profiles = append(c.profiles, profiles...)
				mu.Unlock()
				wait.Done()
			}()
		}
		wait.Wait()
	}
	return c.profiles
}
