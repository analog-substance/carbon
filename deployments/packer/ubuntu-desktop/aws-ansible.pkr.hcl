
source "amazon-ebs" "carbon-ubuntu-desktop-ansible" {
  profile = var.aws_profile
  region  = var.aws_build_region

  ami_name                    = "carbon-ubuntu-desktop-ansible-${local.timestamp}"
  instance_type               = "t3.medium"
  ssh_username                = "ubuntu"
  ssh_interface               = "public_ip"
  ssh_timeout                 = "10m"
  encrypt_boot                = true
  associate_public_ip_address = true
  user_data_file              = "${path.root}/cloud-init/default-ansible/user-data"

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
    "sources.amazon-ebs.carbon-ubuntu-desktop-ansible",
  ]

  provisioner "ansible" {
    playbook_file = "../../ansible/ubuntu-desktop.yaml"
  }

  provisioner "shell" {
    inline = [
      "find /home/ -maxdepth 2 -type d -name '~*' -exec rm -rf {} \\;",
    ]
  }
}
