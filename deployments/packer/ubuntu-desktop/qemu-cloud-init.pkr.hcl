source "qemu" "carbon-ubuntu-desktop" {

  iso_url      = var.iso_url
  iso_checksum = var.iso_checksum
  ssh_username = var.ssh_username
  ssh_password = var.ssh_password

  shutdown_command = "echo '${var.ssh_password}' | sudo -S shutdown -P now"
  disk_size = "50G"
  #   efi_boot         = true
  cpus             = 4
  memory           = 8128
  format           = "qcow2"
  accelerator      = "kvm"
  http_directory   = "${path.root}/cloud-init/autoinstall/"
  ssh_timeout      = "20m"
  vm_name          = "carbon-ubuntu-desktop-${local.timestamp}"
  disk_interface   = "virtio"
  boot_wait        = "5s"
  boot_command = [var.boot_command]
  output_directory = "deployments/images/qemu/carbon-ubuntu-desktop-${local.timestamp}"

}

build {
  sources = ["source.qemu.carbon-ubuntu-desktop"]

  provisioner "shell" {
    inline = [
      # without sudo
      #util.py[WARNING]: REDACTED config part /etc/cloud/cloud.cfg.d/99-installer.cfg, insufficient permissions
      #util.py[WARNING]: REDACTED config part /etc/cloud/cloud.cfg.d/90-installer-network.cfg, insufficient permissions
      #"/usr/bin/cloud-init status --wait",
      "sudo /usr/bin/cloud-init status --wait",
    ]
  }
}

