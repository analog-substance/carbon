package cmd

import (
	"reflect"
	"testing"
)

func TestBuildSSMArgs(t *testing.T) {
	tests := []struct {
		name       string
		instanceID string
		awsProfile string
		region     string
		want       []string
	}{
		{
			name:       "instance ID only",
			instanceID: "i-1234567890abcdef0",
			want:       []string{"aws", "ssm", "start-session", "--target", "i-1234567890abcdef0"},
		},
		{
			name:       "with aws profile",
			instanceID: "i-1234567890abcdef0",
			awsProfile: "staging",
			want:       []string{"aws", "ssm", "start-session", "--target", "i-1234567890abcdef0", "--profile", "staging"},
		},
		{
			name:       "with region",
			instanceID: "i-1234567890abcdef0",
			region:     "us-east-1",
			want:       []string{"aws", "ssm", "start-session", "--target", "i-1234567890abcdef0", "--region", "us-east-1"},
		},
		{
			name:       "with profile and region",
			instanceID: "i-1234567890abcdef0",
			awsProfile: "prod",
			region:     "eu-west-1",
			want:       []string{"aws", "ssm", "start-session", "--target", "i-1234567890abcdef0", "--profile", "prod", "--region", "eu-west-1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildSSMArgs(tt.instanceID, tt.awsProfile, tt.region)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildSSMArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
