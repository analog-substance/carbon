
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
  http_directory = "./cloud-init/autoinstall/"
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
