package vsphere

import (
	"context"
	"fmt"
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/vmware/govmomi/session/cache"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"net/url"
)

type Profile struct {
	types.Profile
	apiClient *vim25.Client
}

func NewProfile(name string, providerInstance *Provider, config common.ProfileConfig) *Profile {

	vClient, err := newClient(config.URL, config.Username, config.GetPassword())
	if err != nil {
		log.Debug("Error creating vSphere client", "error", err)
	}

	return &Profile{
		Profile:   base.NewProfile(name, providerInstance, config),
		apiClient: vClient,
	}
}

func (p *Profile) Environments() []types.Environment {
	var environments []types.Environment

	// Create a view of HostSystem objects
	m := view.NewManager(p.apiClient)
	ctx := context.Background()
	v, err := m.CreateContainerView(ctx, p.apiClient.ServiceContent.RootFolder, []string{"HostSystem"}, true)
	if err != nil {
		panic(err)
	}

	defer v.Destroy(ctx)

	// Retrieve summary property for all hosts
	// Reference: http://pubs.vmware.com/vsphere-60/topic/com.vmware.wssdk.apiref.doc/vim.HostSystem.html
	var hss []mo.HostSystem
	err = v.Retrieve(ctx, []string{"HostSystem"}, []string{"summary"}, &hss)
	if err != nil {
		panic(err)

	}

	for _, hs := range hss {
		totalCPU := int64(hs.Summary.Hardware.CpuMhz) * int64(hs.Summary.Hardware.NumCpuCores)
		freeCPU := int64(totalCPU) - int64(hs.Summary.QuickStats.OverallCpuUsage)
		freeMemory := int64(hs.Summary.Hardware.MemorySize) - (int64(hs.Summary.QuickStats.OverallMemoryUsage) * 1024 * 1024)
		fmt.Printf("%s\t", hs.Summary.Config.Name)
		fmt.Printf("%d\t", hs.Summary.QuickStats.OverallCpuUsage)
		fmt.Printf("%d\t", totalCPU)
		fmt.Printf("%d\t", freeCPU)
		fmt.Printf("%s\t", (units.ByteSize(hs.Summary.QuickStats.OverallMemoryUsage))*1024*1024)
		fmt.Printf("%s\t", units.ByteSize(hs.Summary.Hardware.MemorySize))
		fmt.Printf("%d\t", freeMemory)
		fmt.Printf("\n")
	}

	environments = append(environments, &Environment{
		name:    "default",
		profile: p,
	})

	return environments
}

func newClient(host, user, pass string) (*vim25.Client, error) {
	endpoint, err := url.Parse(fmt.Sprintf("https://%s", host))
	endpoint.User = url.UserPassword(user, pass)

	s := &cache.Session{
		URL:      endpoint,
		Insecure: true,
	}

	c := new(vim25.Client)
	err = s.Login(context.Background(), c, nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}
