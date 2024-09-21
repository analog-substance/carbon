package aws

import (
	"context"
	types3 "github.com/analog-substance/carbon/pkg/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	types2 "github.com/aws/aws-sdk-go-v2/service/ec2/types"
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

func (e *environment) Platform() types3.Platform {
	return e.platform
}

func (e *environment) VMs() []types3.VM {
	var vms []types3.VM
	ec2Results, err := e.ec2Client.DescribeInstances(context.Background(), &ec2.DescribeInstancesInput{
		Filters: []types2.Filter{
			types2.Filter{
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

func awsInstanceToMachine(instance types2.Instance) types3.Machine {
	machine := types3.Machine{
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

	state := types3.StateUnknown
	if instance.StateReason != nil {
		if strings.EqualFold(*instance.StateReason.Code, "Client.UserInitiatedHibernate") {
			state = types3.StateSleeping
		}
	}

	if state == types3.StateUnknown {
		state = stateFromEC2(instance.State.Name)
	}

	machine.CurrentState = state
	if instance.State.Name == types2.InstanceStateNameRunning {
		machine.CurrentUpTime = time.Duration(time.Now().UnixNano() - instance.LaunchTime.UnixNano())
	}

	return machine
}

func stateFromEC2(state types2.InstanceStateName) types3.MachineState {
	if state == types2.InstanceStateNameRunning {
		return types3.StateRunning
	}
	if state == types2.InstanceStateNameStopped {
		return types3.StateOff
	}
	if state == types2.InstanceStateNameStopping {
		return types3.StateStopping
	}
	if state == types2.InstanceStateNameTerminated {
		return types3.StateTerminating
	}
	return types3.StateUnknown
}
