# Copied from Carbon {{ .Version }}

data "aws_ami" "carbon_desktop" {
  owners = ["self"]
  most_recent = true

  filter {
    name = "virtualization-type"
    values = ["hvm"]
  }

  filter {
    name = "root-device-type"
    values = ["ebs"]
  }

  filter {
    name = "name"
    values = ["carbon-ubuntu-desktop-*"]
  }
}

locals {
  amis = {
    desktop = data.aws_ami.carbon_desktop.id
  }
}

variable "environment" {
  default = "carbon-default"
}

variable "vpc_cidr" {
  description = "CIDR range for whole VPC"
  default     = "10.42.0.0/16"
}

variable "availability_zone" {
  description = "availability zone for subnets"
  default     = "us-east-1a"
}

variable "bastion_subnet_cidr" {
  type = list(string)
  description = "Bastion subnet CIDR values"
  default = ["10.42.0.0/24"]
}

variable "operations_subnet_cidr" {
  type = list(string)
  description = "Operations subnet CIDR values"
  default = ["10.42.1.0/24"]
}

resource "aws_instance" "carbon_vm" {
  for_each                             = {
    for machine in var.machines :machine.name => machine if machine.provider == "aws"
  }
  ami                                  = local.amis[each.value.image]
  instance_type                        = each.value.type
  availability_zone                    = var.availability_zone
  subnet_id                            = var.operations_subnet_cidr[0]
  instance_initiated_shutdown_behavior = "stop"
  hibernation                          = true
  root_block_device {
    volume_size = each.value.volume_size
    encrypted   = true

    tags = {
      "user:environment" = var.environment
      "user:purpose"     = each.value.purpose
    }
  }

  #   vpc_security_group_ids = [
  #     aws_security_group.
  #   ]
  associate_public_ip_address = true

  lifecycle {
    ignore_changes = [ami, associate_public_ip_address, root_block_device]
  }

  tags = {
    Name               = "${each.value.purpose}-${each.value.name}"
    "user:environment" = var.environment
    "user:purpose"     = each.value.purpose
  }
}
