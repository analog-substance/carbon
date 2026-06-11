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

func TestBuildSSMPortForwardArgs(t *testing.T) {
	tests := []struct {
		name       string
		instanceID string
		awsProfile string
		region     string
		localPort  string
		remotePort string
		remoteHost string
		want       []string
	}{
		{
			name:       "instance port forward",
			instanceID: "i-1234567890abcdef0",
			localPort:  "8080",
			remotePort: "80",
			want: []string{
				"aws", "ssm", "start-session",
				"--target", "i-1234567890abcdef0",
				"--document-name", "AWS-StartPortForwardingSession",
				"--parameters", `{"portNumber":["80"],"localPortNumber":["8080"]}`,
			},
		},
		{
			name:       "remote host port forward",
			instanceID: "i-1234567890abcdef0",
			localPort:  "5432",
			remotePort: "5432",
			remoteHost: "db.internal",
			want: []string{
				"aws", "ssm", "start-session",
				"--target", "i-1234567890abcdef0",
				"--document-name", "AWS-StartPortForwardingSessionToRemoteHost",
				"--parameters", `{"host":["db.internal"],"portNumber":["5432"],"localPortNumber":["5432"]}`,
			},
		},
		{
			name:       "instance port forward with profile and region",
			instanceID: "i-1234567890abcdef0",
			awsProfile: "prod",
			region:     "us-east-1",
			localPort:  "8080",
			remotePort: "80",
			want: []string{
				"aws", "ssm", "start-session",
				"--target", "i-1234567890abcdef0",
				"--document-name", "AWS-StartPortForwardingSession",
				"--parameters", `{"portNumber":["80"],"localPortNumber":["8080"]}`,
				"--profile", "prod",
				"--region", "us-east-1",
			},
		},
		{
			name:       "remote host port forward with profile and region",
			instanceID: "i-1234567890abcdef0",
			awsProfile: "staging",
			region:     "eu-west-1",
			localPort:  "3306",
			remotePort: "3306",
			remoteHost: "mysql.internal",
			want: []string{
				"aws", "ssm", "start-session",
				"--target", "i-1234567890abcdef0",
				"--document-name", "AWS-StartPortForwardingSessionToRemoteHost",
				"--parameters", `{"host":["mysql.internal"],"portNumber":["3306"],"localPortNumber":["3306"]}`,
				"--profile", "staging",
				"--region", "eu-west-1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildSSMPortForwardArgs(tt.instanceID, tt.awsProfile, tt.region, tt.localPort, tt.remotePort, tt.remoteHost)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildSSMPortForwardArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
