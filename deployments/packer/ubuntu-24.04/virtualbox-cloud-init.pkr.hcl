
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
  http_directory = "${path.root}/cloud-init/autoinstall/"
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
