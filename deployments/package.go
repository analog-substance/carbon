package deployments

import "embed"

//go:embed ansible
//go:embed carbon
//go:embed cloud-init
//go:embed packer
//go:embed projects/aws
//go:embed projects/qemu
var Files embed.FS
