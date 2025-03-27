package aws

import (
	"context"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"strings"
	"time"
)

type Environment struct {
	name         string
	profile      *Profile
	ec2Client    *ec2.Client
	vpcId        string
	awsAccountId string
}

func (e *Environment) Name() string {
	return e.name
}

func (e *Environment) Profile() types.Profile {
	return e.profile
}

func (e *Environment) VMs() []types.VM {
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
		log.Debug("Error get instances", "profile", e.Name(), "err", err)
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

func (e *Environment) StartVM(id string) error {
	_, err := e.ec2Client.StartInstances(context.Background(), &ec2.StartInstancesInput{
		InstanceIds: []string{id},
	})
	return err
}

func (e *Environment) StopVM(id string) error {
	hibernate := false
	_, err := e.ec2Client.StopInstances(context.Background(), &ec2.StopInstancesInput{
		InstanceIds: []string{id},
		Hibernate:   &hibernate,
	})
	return err
}

func (e *Environment) SuspendVM(id string) error {
	hibernate := true
	_, err := e.ec2Client.StopInstances(context.Background(), &ec2.StopInstancesInput{
		InstanceIds: []string{id},
		Hibernate:   &hibernate,
	})
	return err
}

func (e *Environment) RestartVM(id string) error {
	_, err := e.ec2Client.RebootInstances(context.Background(), &ec2.RebootInstancesInput{
		InstanceIds: []string{id},
	})
	return err
}

func (e *Environment) DestroyVM(id string) error {
	_, err := e.ec2Client.TerminateInstances(context.Background(), &ec2.TerminateInstancesInput{
		InstanceIds: []string{id},
	})
	return err
}

func (e *Environment) CreateVM(options types.MachineLaunchOptions) error {

	//_, err := e.ec2Client.RunInstances(context.Background(), &ec2.RunInstancesInput{
	//
	//})
	//return err
	return nil
}

func (e *Environment) ImageBuilds() ([]types.ImageBuild, error) {
	return models.GetImageBuildsForProvider(e.profile.Provider().Type())
}
func (e *Environment) Images() ([]types.Image, error) {
	log.Debug("getting images", "env", e.Name(), "profile", e.profile.Name(), "provider", e.profile.Provider().Name())
	amis, err := e.ec2Client.DescribeImages(context.Background(), &ec2.DescribeImagesInput{
		Filters: []ec2Types.Filter{
			ec2Types.Filter{
				Name: aws.String("owner-id"),
				Values: []string{
					e.awsAccountId,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	images := []types.Image{}
	for _, image := range amis.Images {
		createdAt, err := time.Parse(time.RFC3339Nano, *image.CreationDate)
		if err != nil {
			log.Debug("failed to parse created time, setting it to now", "creationDate", *image.CreationDate)
			createdAt = time.Now()
		}
		images = append(images, models.NewImage(*image.ImageId, *image.Name, createdAt, e))
	}
	return images, nil
}

func awsInstanceToMachine(instance ec2Types.Instance) *models.Machine {
	machine := &models.Machine{
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

	machine.InstanceType = string(instance.InstanceType)
	return machine
}

func stateFromEC2(state ec2Types.InstanceStateName) types.MachineState {

	if state == ec2Types.InstanceStateNameRunning {
		return types.StateRunning
	}
	if state == ec2Types.InstanceStateNameStopped {
		return types.StateStopped
	}
	if state == ec2Types.InstanceStateNameStopping || state == ec2Types.InstanceStateNameShuttingDown {
		return types.StateStopping
	}
	if state == ec2Types.InstanceStateNameTerminated {
		return types.StateTerminating
	}
	if state == ec2Types.InstanceStateNamePending {
		return types.StateStarting
	}

	log.Debug("unknown state for vm", "state", state)
	return types.StateUnknown
}

func (e *Environment) DestroyImage(imageID string) error {
	return nil
}
