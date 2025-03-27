package deployments

import "embed"

//go:embed ansible
//go:embed carbon
//go:embed cloud-init
//go:embed packer
//go:embed projects/example
//go:embed terraform
var Files embed.FS
