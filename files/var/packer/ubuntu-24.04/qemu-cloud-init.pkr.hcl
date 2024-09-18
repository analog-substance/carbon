source "qemu" "carbon-vm-ubuntu" {

  iso_url      = var.iso_url
  iso_checksum = var.iso_checksum
  ssh_username = var.ssh_username
  ssh_password = var.ssh_password

  shutdown_command = "echo '${var.ssh_password}' | sudo -S shutdown -P now"
  disk_size        = "50G"
#   efi_boot         = true
  cpus             = 4
  memory           = 8128
  format           = "qcow2"
  accelerator      = "kvm"
  http_directory   = "./cloud-init/autoinstall/"
  ssh_timeout      = "20m"
  vm_name          = "carbon-ubuntu-vm-${local.timestamp}"
  disk_interface   = "virtio"
  boot_wait        = "5s"
  boot_command = [var.boot_command]
  output_directory = "../../../outputs/qemu-carbon-ubuntu-vm-${local.timestamp}"

}

build {
  sources = ["source.qemu.carbon-vm-ubuntu"]

  provisioner "shell" {
    inline = [
      "/usr/bin/cloud-init status --wait",
    ]
  }
}

