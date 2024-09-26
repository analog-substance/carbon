package aws

import (
	"context"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
	"strings"
	"time"
)

type environment struct {
	name      string
	platform  *platform
	ec2Client *ec2.Client
	vpcId     string
}

func (e *environment) Name() string {
	return e.name
}

func (e *environment) Platform() types.Platform {
	return e.platform
}

func (e *environment) VMs() []types.VM {
	var vms []types.VM
	ec2Results, err := e.ec2Client.DescribeInstances(context.Background(), &ec2.DescribeInstancesInput{
		Filters: []ec2Types.Filter{
			ec2Types.Filter{
				Name:   aws.String("vpc-id"),
				Values: []string{e.vpcId},
			},
		},
	})
	if err != nil {
		log.Println("Error get VPCs", err)
		return vms
	}

	for _, reservations := range ec2Results.Reservations {
		for _, ec2Instance := range reservations.Instances {

			machine := awsInstanceToMachine(ec2Instance)
			machine.Env = e
			vms = append(vms, machine)
		}
	}
	return vms
}

func (e *environment) StartVM(id string) error {
	_, err := e.ec2Client.StartInstances(context.Background(), &ec2.StartInstancesInput{
		InstanceIds: []string{id},
	})
	return err
}

func (e *environment) StopVM(id string) error {
	hibernate := true
	_, err := e.ec2Client.StopInstances(context.Background(), &ec2.StopInstancesInput{
		InstanceIds: []string{id},
		Hibernate:   &hibernate,
	})
	return err
}

func (e *environment) RestartVM(id string) error {
	_, err := e.ec2Client.RebootInstances(context.Background(), &ec2.RebootInstancesInput{
		InstanceIds: []string{id},
	})
	return err
}

func awsInstanceToMachine(instance ec2Types.Instance) models.Machine {
	machine := models.Machine{
		InstanceID: *instance.InstanceId,
	}

	for _, tag := range instance.Tags {
		if *tag.Key == "Name" {
			machine.InstanceName = *tag.Value
			break
		}
	}

	if instance.PublicIpAddress != nil {
		machine.PublicIPAddresses = []string{*instance.PublicIpAddress}
	}

	if instance.PrivateIpAddress != nil {
		machine.PrivateIPAddresses = []string{*instance.PrivateIpAddress}
	}

	state := types.StateUnknown
	if instance.StateReason != nil {
		if strings.EqualFold(*instance.StateReason.Code, "Client.UserInitiatedHibernate") {
			state = types.StateSleeping
		}
	}

	if state == types.StateUnknown {
		state = stateFromEC2(instance.State.Name)
	}

	machine.CurrentState = state
	if instance.State.Name == ec2Types.InstanceStateNameRunning {
		machine.CurrentUpTime = time.Duration(time.Now().UnixNano() - instance.LaunchTime.UnixNano())
	}

	return machine
}

func stateFromEC2(state ec2Types.InstanceStateName) types.MachineState {
	if state == ec2Types.InstanceStateNameRunning {
		return types.StateRunning
	}
	if state == ec2Types.InstanceStateNameStopped {
		return types.StateStopped
	}
	if state == ec2Types.InstanceStateNameStopping {
		return types.StateStopping
	}
	if state == ec2Types.InstanceStateNameTerminated {
		return types.StateTerminating
	}

	log.Println("Unknown state for AWS VM:", state)
	return types.StateUnknown
}
