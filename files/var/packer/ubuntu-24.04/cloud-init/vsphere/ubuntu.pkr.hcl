packer {
  required_version = ">= 1.7.0"
  required_plugins {
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

locals {
  timestamp = formatdate("YYYYMMDDhhmmss", timestamp())
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

  http_directory = "./http/24.04/"
  boot_command   = [var.boot_command]
  boot_wait      = "5s"

  vm_name = "carbon-ubuntu-vm-${local.timestamp}"

}


build {
  sources = [
    "sources.vsphere-iso.carbon-vm-ubuntu",
  ]

  provisioner "shell" {
    inline = [
      "/usr/bin/cloud-init status --wait",
    ]
  }
}
