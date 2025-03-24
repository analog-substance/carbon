package aws

import (
	"github.com/analog-substance/carbon/pkg/types"
	"testing"
)

func TestProvider_IsAvailable(t *testing.T) {
	type fields struct {
		Provider        types.Provider
		awsProfileNames []string
		profiles        []types.Profile
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "should return false if providers are available",
			fields: fields{},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Provider{
				Provider:        tt.fields.Provider,
				awsProfileNames: tt.fields.awsProfileNames,
				profiles:        tt.fields.profiles,
			}
			if got := p.IsAvailable(); got != tt.want {
				t.Errorf("IsAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}
