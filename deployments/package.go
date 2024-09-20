package deployments

import "embed"

//go:embed ansible
//go:embed carbon
//go:embed cloud-init
//go:embed packer
var Files embed.FS
