packer {
  required_version = ">= 1.7.0"
  required_plugins {
    qemu = {
      version = "~> 1"
      source  = "github.com/hashicorp/qemu"
    }
  }
}

variable "iso_url" {
  type        = string
  description = "URL to download the ISO file"
  default     = "https://releases.ubuntu.com/noble/ubuntu-24.04.1-live-server-amd64.iso"
}

variable "iso_checksum" {
  type        = string
  description = "Checksum to validate the downloaded ISO file"
  default     = "sha256:e240e4b801f7bb68c20d1356b60968ad0c33a41d00d828e74ceb3364a0317be9"
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
  default     = "e<wait><down><down><down><end> autoinstall 'ds=nocloud-net;s=/cidata/'<F10>"
}


locals {
  timestamp = formatdate("YYYYMMDDhhmmss", timestamp())
}

source "qemu" "carbon-vm-ubuntu" {

  iso_url      = var.iso_url
  iso_checksum = var.iso_checksum
  ssh_username = var.ssh_username
  ssh_password = var.ssh_password

  shutdown_command = "echo '${var.ssh_password}' | sudo -S shutdown -P now"
  #   output_directory  = "output_centos_tdhtest"
  disk_size      = "50G"
  format         = "qcow2"
  accelerator    = "kvm"
  cd_files = ["../cloud-init/meta-data", "../cloud-init/user-data"]
  cd_label       = "CIDATA"
  http_directory = "../cloud-init/"
#   http_interface = "wlp2s0"
  ssh_timeout    = "20m"
  vm_name        = "carbon-ubuntu-vm-${local.timestamp}"
  #   net_device     = "virtio-net"
  #   network_bridge = "net_bridge"
  disk_interface = "virtio"
  boot_wait      = "5s"
  boot_command = [var.boot_command]
}

build {
  sources = ["source.qemu.carbon-vm-ubuntu"]

  provisioner "shell" {
    inline = [
      "/usr/bin/cloud-init status --wait",
    ]
  }
}

