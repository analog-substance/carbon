packer {
  required_version = ">= 1.7.0"
  required_plugins {
    virtualbox = {
      version = "~> 1"
      source  = "github.com/hashicorp/virtualbox"
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
  http_directory   = "./http/24.04/"
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

build {
  sources = [
    "sources.virtualbox-iso.carbon-vm-ubuntu",
  ]

  provisioner "shell" {
    inline = [
      "/usr/bin/cloud-init status --wait",
    ]
  }
}
