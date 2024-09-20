



source "qemu" "carbon-vm-ubuntu-ansible" {

  iso_url      = var.iso_url
  iso_checksum = var.iso_checksum
  ssh_username = var.ssh_username
  ssh_password = var.ssh_password

  shutdown_command = "echo '${var.ssh_password}' | sudo -S shutdown -P now"
  disk_size      = "50G"
  format         = "qcow2"
  accelerator    = "kvm"
  http_directory = "${path.root}/cloud-init/autoinstall-ansible/"
  ssh_timeout    = "20m"
  vm_name        = "carbon-ubuntu-vm-${local.timestamp}"
  disk_interface = "virtio"
  boot_wait      = "5s"
  boot_command = [var.boot_command]
  output_directory = "outputs/qemu-ansible-carbon-ubuntu-vm-${local.timestamp}"
}

build {
  sources = ["source.qemu.carbon-vm-ubuntu-ansible"]

  provisioner "ansible" {
    playbook_file = "../../ansible/ubuntu-desktop.yaml"
  }

  provisioner "shell" {
    inline = [
      "find /home/ -maxdepth 2 -type d -name '~*' -exec rm -rf {} \\;",
    ]
  }
}

