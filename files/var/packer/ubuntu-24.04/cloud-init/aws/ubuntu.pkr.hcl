packer {
  required_version = ">= 1.7.0"
  required_plugins {
    amazon = {
      version = ">= 1.2.8"
      source  = "github.com/hashicorp/amazon"
    }
  }
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
  user_data_file              = "./http/24.04/user-data"

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
    "sources.amazon-ebs.carbon-vm-ubuntu",
  ]

  provisioner "shell" {
    inline = [
      "/usr/bin/cloud-init status --wait",
    ]
  }
}
