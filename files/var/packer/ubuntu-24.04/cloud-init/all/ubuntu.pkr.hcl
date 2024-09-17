packer {
  required_version = ">= 1.7.0"
  required_plugins {
    amazon = {
      version = ">= 1.2.8"
      source  = "github.com/hashicorp/amazon"
    }
    ansible = {
      version = ">= 1.1.1"
      source  = "github.com/hashicorp/ansible"
    }
    virtualbox = {
      version = "~> 1"
      source  = "github.com/hashicorp/virtualbox"
    }
    vsphere = {
      version = "~> 1"
      source  = "github.com/hashicorp/vsphere"
    }
  }
}

variable "iso_url" {
  type        = string
  description = "URL to download the ISO file"
  default     = "https://releases.ubuntu.com/noble/ubuntu-24.04-live-server-amd64.iso"
}

variable "iso_checksum" {
  type        = string
  description = "Checksum to validate the downloaded ISO file"
  default     = "sha256:8762f7e74e4d64d72fceb5f70682e6b069932deedb4949c6975d0f0fe0a91be3"
}

variable "ssh_username" {
  type        = string
  description = "Username to use when connecting via SSH"
  default     = "isotope"

}

variable "ssh_password" {
  type        = string
  description = "Password to use when connecting via SSH"
  default     = "carbon"

}

variable "boot_command" {
  type        = string
  description = "Keyboard sequence to execute to properly boot the image."
  default     = "e<wait><down><down><down><end> autoinstall 'ds=nocloud-net;s=http://{{ .HTTPIP }}:{{ .HTTPPort }}/'<F10>"
}

variable "vsphere_endpoint" {
  type        = string
  description = "vSphere URL."
  default     = ""

}

variable "vsphere_username" {
  type        = string
  description = "Username to authenticate to vShpere."
  default     = ""

}

variable "vsphere_password" {
  type        = string
  description = "Password to authenticate with vSphere."
  default     = ""

}

variable "vsphere_host" {
  type        = string
  description = "Host in vShere to provision VM on."
  default     = ""

}

variable "vsphere_datastore" {
  type        = string
  description = "Datastore in vSphere to provision VM on."
  default     = ""

}

variable "aws_profile" {
  default     = "default"
  description = "The aws credentials profile to use."
  type        = string
}

variable "aws_build_region" {
  default     = "us-east-1"
  description = "The region in which to retrieve the base AMI from and build the new AMI."
  type        = string
}

variable "aws_tags" {
  default = {
    Architecture = "x86_64"
    OS_Version   = "Ubuntu Noble Numbat"
  }

  type = map(string)
}

variable "aws_run_tags" {
  default = {
    Environment = "testing"
    Type        = "builder"
  }

  type = map(string)
}

variable "aws_vpc_filters" {
  default = {
    "Name" : "AMI Builds"
  }

  type = map(string)
}

variable "aws_subnet_filters" {
  default = {
    "Name" : "AMI Builds"
  }

  type = map(string)
}

locals {
  timestamp = formatdate("YYYYMMDDhhmmss", timestamp())
}

source "virtualbox-iso" "carbon-vm-ubuntu" {
  guest_os_type    = "Ubuntu_64"
  vm_name          = "carbon-ubuntu-vm-${local.timestamp}"
  iso_url          = var.iso_url
  iso_checksum     = var.iso_checksum
  ssh_username     = var.ssh_username
  ssh_password     = var.ssh_password
  disk_size        = 80000
  ssh_timeout      = "15m"
  shutdown_command = "echo '${var.ssh_password}' | sudo -S shutdown -P now"
  headless         = false
  firmware         = "efi"
  http_directory   = "../cloud-init/"
  boot_command     = [var.boot_command]
  boot_wait        = "5s"
  vboxmanage = [
    [
      "modifyvm",
      "{{.Name}}",
      "--memory",
      "4096"
    ],
    [
      "modifyvm",
      "{{.Name}}",
      "--cpus",
      "4"
    ],
    [
      "modifyvm",
      "{{.Name}}",
      "--nat-localhostreachable1",
      "on"
    ]
  ]
}

source "vsphere-iso" "carbon-vm-ubuntu" {
  vcenter_server      = var.vsphere_endpoint
  username            = var.vsphere_username
  password            = var.vsphere_password
  host                = var.vsphere_host
  datastore           = var.vsphere_datastore
  insecure_connection = "true"
  iso_url             = var.iso_url
  iso_checksum        = var.iso_checksum
  ssh_username        = var.ssh_username
  ssh_password        = var.ssh_password
  communicator        = "ssh"
  ssh_timeout         = "30m"

  guest_os_type = "ubuntu64Guest"
  RAM           = 8128
  CPUs          = 4
  cpu_cores     = 2
  storage {
    disk_size             = 80000
    disk_thin_provisioned = true
  }
  network_adapters {
    network_card = "vmxnet3"
    network      = "VM Network"
  }

  shutdown_command = "echo '${var.ssh_password}' | sudo -S systemctl poweroff"

  http_directory   = "../cloud-init/"
  boot_command   = [var.boot_command]
  boot_wait      = "5s"

  vm_name = "carbon-ubuntu-vm-${local.timestamp}"

}

source "amazon-ebs" "carbon-vm-ubuntu" {
  profile = var.aws_profile
  region  = var.aws_build_region

  ami_name                    = "carbon-vm-ami-${local.timestamp}"
  instance_type               = "t3.medium"
  ssh_username                = "ubuntu"
  ssh_interface               = "public_ip"
  ssh_timeout                 = "10m"
  encrypt_boot                = true
  associate_public_ip_address = true
  user_data_file              = "../cloud-init/"

  temporary_key_pair_type = "ed25519"

  launch_block_device_mappings {
    delete_on_termination = true
    device_name           = "/dev/sda1"
    encrypted             = true
    volume_size           = 80
    volume_type           = "gp3"
  }

  tags     = var.aws_tags
  run_tags = var.aws_run_tags

  source_ami_filter {
    filters = {
      architecture        = "x86_64"
      name                = "ubuntu/images/hvm-ssd-gp3/ubuntu-noble-24.04-amd64-server*"
      root-device-type    = "ebs"
      virtualization-type = "hvm"
    }
    most_recent = true
    owners      = ["099720109477"]
  }

  subnet_filter {
    filters   = var.aws_subnet_filters
    most_free = true
    random    = false
  }

  vpc_filter {
    filters = var.aws_vpc_filters
  }
}

build {
  sources = [
    "sources.vsphere-iso.carbon-vm-ubuntu",
    "sources.virtualbox-iso.carbon-vm-ubuntu",
    "sources.amazon-ebs.carbon-vm-ubuntu",
  ]

  provisioner "shell" {
    inline = [
      "/usr/bin/cloud-init status --wait",
    ]
  }
}
