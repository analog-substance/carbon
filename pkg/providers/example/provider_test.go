package example

import (
	"github.com/analog-substance/carbon/pkg/types"
	"testing"
)

func TestProvider_IsAvailable(t *testing.T) {
	type fields struct {
		Provider       types.Provider
		profilesLoaded bool
		profiles       []types.Profile
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"Should always be enabled",
			fields{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Provider{
				Provider:       tt.fields.Provider,
				profilesLoaded: tt.fields.profilesLoaded,
				profiles:       tt.fields.profiles,
			}
			if got := p.IsAvailable(); got != tt.want {
				t.Errorf("IsAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}
