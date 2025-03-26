package gcloud

import (
	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"context"
	"errors"
	"fmt"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/providers/base"
	"github.com/analog-substance/carbon/pkg/types"
	"google.golang.org/api/iterator"
	"strings"
)

type Environment struct {
	name                 string
	gProject             string
	zone                 string
	profile              types.Profile
	gCloudInstanceClient *compute.InstancesClient
	ctx                  context.Context
}

func (e *Environment) Name() string {
	return e.name
}

func (e *Environment) Profile() types.Profile {
	return e.profile
}

func (e *Environment) VMs() []types.VM {
	var vms []types.VM

	req := &computepb.ListInstancesRequest{
		Project: e.gProject,
		Zone:    e.zone,
	}

	it := e.gCloudInstanceClient.List(e.ctx, req)
	for {
		_, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Debug("Error listing instances", "error", err)
			break
		}
		instanceList := it.Response.(*computepb.InstanceList)
		for _, instance := range instanceList.Items {
			var publicIPs = []string{}
			var privateIPs = []string{}

			for _, iface := range instance.GetNetworkInterfaces() {
				privateIPs = append(privateIPs, iface.GetNetworkIP())
			}

			mts := strings.Split(*instance.MachineType, "/")
			mt := mts[len(mts)-1]

			log.Debug("last started", "timestamp", instance.LastStartTimestamp)
			vms = append(vms, &models.Machine{
				InstanceName: instance.GetName(),
				CurrentState: stateFromStatus(instance.GetStatus()),
				InstanceID:   fmt.Sprintf("%d", instance.Id),
				Env:          e,
				//CurrentUpTime:      d,
				InstanceType:       mt,
				PrivateIPAddresses: privateIPs,
				PublicIPAddresses:  publicIPs,
			})
		}
		if instanceList.NextPageToken == nil {
			break
		}
	}

	return vms
}

func (e *Environment) StartVM(id string) error {
	req := &computepb.StartInstanceRequest{
		Project:  e.gProject,
		Zone:     e.zone,
		Instance: id,
	}
	_, err := e.gCloudInstanceClient.Start(e.ctx, req)
	//if err != nil {
	return err
	//}
	//err = op.Wait(e.ctx)
	//return err
}

func (e *Environment) StopVM(id string) error {
	req := &computepb.StopInstanceRequest{
		Project:  e.gProject,
		Zone:     e.zone,
		Instance: id,
	}
	_, err := e.gCloudInstanceClient.Stop(e.ctx, req)
	//if err != nil {
	return err
	//}
	//err = op.Wait(e.ctx)
	//return err
}

func (e *Environment) SuspendVM(id string) error {
	req := &computepb.SuspendInstanceRequest{
		Project:  e.gProject,
		Zone:     e.zone,
		Instance: id,
	}
	_, err := e.gCloudInstanceClient.Suspend(e.ctx, req)
	//if err != nil {
	return err
	//}
	//err = op.Wait(e.ctx)
	//return err
}

func (e *Environment) RestartVM(id string) error {
	req := &computepb.ResetInstanceRequest{
		Project:  e.gProject,
		Zone:     e.zone,
		Instance: id,
	}
	_, err := e.gCloudInstanceClient.Reset(e.ctx, req)
	//if err != nil {
	return err
	//}
	//err = op.Wait(e.ctx)
	//return err
}

func (e *Environment) DestroyVM(id string) error {
	return errors.New("not yet implemented")
}

func (e *Environment) DestroyImage(imageID string) error {
	return errors.New("not yet implemented")
}

func (e *Environment) CreateVM(options types.MachineLaunchOptions) error {
	return errors.New("not yet implemented")
}

func (e *Environment) ImageBuilds() ([]types.ImageBuild, error) {
	return models.GetImageBuildsForProvider(e.profile.Provider().Type())
}

func (e *Environment) Images() ([]types.Image, error) {
	return base.GetImagesForFileBasedProvider(e.Profile().Provider().Type(), e)
}

func stateFromStatus(state string) types.MachineState {
	/*
			0:         "UNDEFINED_STATUS",
		428935662: "DEPROVISIONING",
		290896621: "PROVISIONING",
		413483285: "REPAIRING",
		121282975: "RUNNING",
		431072283: "STAGING",
		444276141: "STOPPED",
		350791796: "STOPPING",
		51223995:  "SUSPENDED",
		514206246: "SUSPENDING",
		250018339: "TERMINATED",

	*/

	if state == "STOPPED" {
		return types.StateStopped
	}
	if state == "STOPPING" || state == "SUSPENDING" {
		return types.StateStopping
	}
	if state == "PROVISIONING" {
		return types.StateStarting
	}
	if state == "SUSPENDED" {
		return types.StateSleeping
	}
	if state == "DEPROVISIONING" {
		return types.StateTerminating
	}
	if state == "TERMINATED" {
		return types.StateTerminated
	}
	if state == "RUNNING" {
		return types.StateRunning
	}

	log.Debug("unknown state for VM", "state", state)
	return types.StateUnknown
}
