package models

import (
	"testing"
)

func TestBuildSSMProxyCommand(t *testing.T) {
	tests := []struct {
		name       string
		instanceID string
		awsProfile string
		awsRegion  string
		want       string
	}{
		{
			name:       "instance ID only",
			instanceID: "i-1234567890abcdef0",
			want:       "aws ssm start-session --target i-1234567890abcdef0 --document-name AWS-StartSSHSession --parameters portNumber=%p",
		},
		{
			name:       "with aws profile",
			instanceID: "i-1234567890abcdef0",
			awsProfile: "staging",
			want:       "aws ssm start-session --target i-1234567890abcdef0 --document-name AWS-StartSSHSession --parameters portNumber=%p --profile staging",
		},
		{
			name:       "with region",
			instanceID: "i-1234567890abcdef0",
			awsRegion:  "us-east-1",
			want:       "aws ssm start-session --target i-1234567890abcdef0 --document-name AWS-StartSSHSession --parameters portNumber=%p --region us-east-1",
		},
		{
			name:       "with profile and region",
			instanceID: "i-1234567890abcdef0",
			awsProfile: "prod",
			awsRegion:  "eu-west-1",
			want:       "aws ssm start-session --target i-1234567890abcdef0 --document-name AWS-StartSSHSession --parameters portNumber=%p --profile prod --region eu-west-1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildSSMProxyCommand(tt.instanceID, tt.awsProfile, tt.awsRegion)
			if got != tt.want {
				t.Errorf("buildSSMProxyCommand() = %q, want %q", got, tt.want)
			}
		})
	}
}
