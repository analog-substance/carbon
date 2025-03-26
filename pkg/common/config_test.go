package common

import (
	"reflect"
	"slices"
	"testing"
)

func TestCarbonConfigFile_Keys(t *testing.T) {
	type fields struct {
		Carbon CarbonConfig
	}

	SetProvidersTypes([]string{"aws", "digitalocean", "multipass", "qemu", "virtualbox"})

	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "keys are correct",
			fields: fields{},
			want: []string{

				"carbon.dir.terraform",
				"carbon.dir.projects",
				"carbon.dir.instance",
				"carbon.dir.deployments",
				"carbon.dir.packer",
				"carbon.dir.images",
				"carbon.providers.aws.enabled",
				"carbon.providers.aws.auto_discover",
				"carbon.providers.aws.profiles.default.enabled",
				"carbon.providers.aws.profiles.default.username",
				"carbon.providers.aws.profiles.default.password",
				"carbon.providers.aws.profiles.default.password_command",
				"carbon.providers.aws.profiles.default.use_1pass_cli",
				"carbon.providers.aws.profiles.default.url",
				"carbon.providers.aws.profiles.default.environments",
				"carbon.providers.aws",
				"carbon.providers.digitalocean.enabled",
				"carbon.providers.digitalocean.auto_discover",
				"carbon.providers.digitalocean.profiles.default.enabled",
				"carbon.providers.digitalocean.profiles.default.username",
				"carbon.providers.digitalocean.profiles.default.password",
				"carbon.providers.digitalocean.profiles.default.password_command",
				"carbon.providers.digitalocean.profiles.default.use_1pass_cli",
				"carbon.providers.digitalocean.profiles.default.url",
				"carbon.providers.digitalocean.profiles.default.environments",
				"carbon.providers.digitalocean",
				"carbon.providers.multipass.enabled",
				"carbon.providers.multipass.auto_discover",
				"carbon.providers.multipass.profiles.default.enabled",
				"carbon.providers.multipass.profiles.default.username",
				"carbon.providers.multipass.profiles.default.password",
				"carbon.providers.multipass.profiles.default.password_command",
				"carbon.providers.multipass.profiles.default.use_1pass_cli",
				"carbon.providers.multipass.profiles.default.url",
				"carbon.providers.multipass.profiles.default.environments",
				"carbon.providers.multipass",
				"carbon.providers.qemu.enabled",
				"carbon.providers.qemu.auto_discover",
				"carbon.providers.qemu.profiles.default.enabled",
				"carbon.providers.qemu.profiles.default.username",
				"carbon.providers.qemu.profiles.default.password",
				"carbon.providers.qemu.profiles.default.password_command",
				"carbon.providers.qemu.profiles.default.use_1pass_cli",
				"carbon.providers.qemu.profiles.default.url",
				"carbon.providers.qemu.profiles.default.environments",
				"carbon.providers.qemu",
				"carbon.providers.virtualbox.enabled",
				"carbon.providers.virtualbox.auto_discover",
				"carbon.providers.virtualbox.profiles.default.enabled",
				"carbon.providers.virtualbox.profiles.default.username",
				"carbon.providers.virtualbox.profiles.default.password",
				"carbon.providers.virtualbox.profiles.default.password_command",
				"carbon.providers.virtualbox.profiles.default.use_1pass_cli",
				"carbon.providers.virtualbox.profiles.default.url",
				"carbon.providers.virtualbox.profiles.default.environments",
				"carbon.providers.virtualbox",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Keys()

			if len(got) != len(tt.want) {
				t.Errorf("GetConfig() len(got) = %v, len(want) %v", len(got), len(tt.want))
			}

			for _, v := range tt.want {
				if !slices.Contains(got, v) {
					t.Errorf("Keys() = did not contain %v", v)
				}
			}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "gets proper default enabled",
			args: args{
				"carbon.providers.qemu.enabled",
			},
			want: true,
		},
		{
			name: "gets proper default enabled",
			args: args{
				"carbon.providers.aws.auto_discover",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet(t *testing.T) {
	type args struct {
		s string
		v any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "disables provider",
			args: args{
				"carbon.providers.aws.enabled",
				"false",
			},
			want: false,
		},

		{
			name: "disables profile",
			args: args{
				"carbon.providers.aws.profiles.default.enabled",
				"false",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Set(tt.args.s, tt.args.v)
			if got := Get(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}
