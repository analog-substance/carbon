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
