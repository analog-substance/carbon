package models

import (
	"encoding/json"
	"testing"

	"github.com/analog-substance/carbon/pkg/types"
)

// fakeProvider implements types.Provider. Only Name is used by the tests; the
// embedded interface satisfies the rest of the method set (and would panic if
// any unimplemented method were called).
type fakeProvider struct {
	types.Provider
	name string
}

func (p *fakeProvider) Name() string { return p.name }

// fakeProfile implements types.Profile.
type fakeProfile struct {
	types.Profile
	name     string
	provider types.Provider
}

func (p *fakeProfile) Name() string             { return p.name }
func (p *fakeProfile) Provider() types.Provider { return p.provider }

// fakeEnv implements types.Environment.
type fakeEnv struct {
	types.Environment
	name    string
	profile types.Profile
}

func (e *fakeEnv) Name() string           { return e.name }
func (e *fakeEnv) Profile() types.Profile { return e.profile }

func TestMachineMarshalJSON(t *testing.T) {
	env := &fakeEnv{
		name: "prod",
		profile: &fakeProfile{
			name:     "default",
			provider: &fakeProvider{name: "AWS"},
		},
	}

	m := &Machine{
		InstanceName:       "web-1",
		InstanceID:         "i-123",
		InstanceType:       "t3.small",
		PublicIPAddresses:  []string{"1.2.3.4"},
		PrivateIPAddresses: []string{"10.0.0.5"},
		CurrentState:       types.StateRunning,
		Env:                env,
	}

	out, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	var got map[string]interface{}
	if err := json.Unmarshal(out, &got); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	tests := map[string]string{
		"name":        "web-1",
		"id":          "i-123",
		"type":        "t3.small",
		"environment": "prod",
		"profile":     "default",
		"provider":    "AWS",
	}
	for key, want := range tests {
		if got[key] != want {
			t.Errorf("marshalled %q = %v, want %q", key, got[key], want)
		}
	}
}

func TestMachineMarshalJSONNilEnv(t *testing.T) {
	m := &Machine{InstanceName: "web-1"}

	out, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	var got map[string]interface{}
	if err := json.Unmarshal(out, &got); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	for _, key := range []string{"environment", "profile", "provider"} {
		if got[key] != "" {
			t.Errorf("marshalled %q = %v, want empty string", key, got[key])
		}
	}
}

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
