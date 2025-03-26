
# Configure the Libvirt provider
provider "libvirt" {
  uri = "qemu:///system"
}


locals {
  qemu_images = fileset(path.module, "../../../../images/qemu/*/*")
}

# Create a new domain
resource "libvirt_network" "carbon_net" {
  # the name used by libvirt
  name = "${var.project}-net"

  # mode can be: "nat" (default), "none", "route", "open", "bridge"
  mode = "nat"

  autostart = true

  #  the domain used by the DNS server in this network
  domain = "${var.project}.carbon.local"

  #  list of subnets the addresses allowed for domains connected
  # also derived to define the host addresses
  # also derived to define the addresses served by the DHCP server
  addresses = ["10.17.3.0/24"]

  # (optional) the bridge device defines the name of a bridge device
  # which will be used to construct the virtual network.
  # (only necessary in "bridge" mode)
  # bridge = "br7"

  # (optional) the MTU for the network. If not supplied, the underlying device's
  # default is used (usually 1500)
  # mtu = 9000

  # (Optional) DNS configuration
  dns {
    # (Optional, default false)
    # Set to true, if no other option is specified and you still want to
    # enable dns.
    enabled    = true
    # (Optional, default false)
    # true: DNS requests under this domain will only be resolved by the
    # virtual network's cd own DNS server
    # false: Unresolved requests will be forwarded to the host's
    # upstream DNS server if the virtual network's DNS server does not
    # have an answer.
    local_only = false

    # (Optional) one or more DNS forwarder entries.  One or both of
    # "address" and "domain" must be specified.  The format is:
    # forwarders {
    #     address = "my address"
    #     domain = "my domain"
    #  }
    #

    # (Optional) one or more DNS host entries.  Both of
    # "ip" and "hostname" must be specified.  The format is:
    # hosts  {
    #     hostname = "my_hostname"
    #     ip = "my.ip.address.1"
    #   }
    # hosts {
    #     hostname = "my_hostname"
    #     ip = "my.ip.address.2"
    #   }
    #
  }

  # (Optional) one or more static routes.
  # "cidr" and "gateway" must be specified. The format is:
  # routes {
  #     cidr = "10.17.0.0/16"
  #     gateway = "10.18.0.2"
  #   }

  # (Optional) Dnsmasq options configuration
  dnsmasq_options {
    # (Optional) one or more option entries.
    # "option_name" muast be specified while "option_value" is
    # optional to also support value-less options.  The format is:
    # options  {
    #     option_name = "server"
    #     option_value = "/base.domain/my.ip.address.1"
    #   }
    # options  {
    #     option_name = "no-hosts"
    #   }
    # options {
    #     option_name = "address"
    #     ip = "/.api.base.domain/my.ip.address.2"
    #   }
    #
  }

}

# Base OS image to use to create a cluster of different
# nodes
resource "libvirt_volume" "os_image" {
  for_each = {for qimg in local.qemu_images : basename(qimg) => qimg}
  name     = each.key
  source   = "${path.module}/${each.value}"
  pool     = libvirt_pool.carbon.name
  format   = "qcow2"
}

resource libvirt_pool carbon {
  name = var.project
  type = "dir"
  path = pathexpand("~/.carbon/${var.project}-pool")
}

# volume to attach to the "master" domain as main disk
resource "libvirt_volume" "vm_vols" {
  for_each       = {for machine in var.machines : machine.name => machine if machine.provider == "qemu"}
  name           = "${each.value.name}.qcow2"
  base_volume_id = libvirt_volume.os_image[each.value.image].id
}

resource libvirt_domain vms {
  for_each = {for machine in var.machines : machine.name => machine if machine.provider == "qemu"}
  name     = each.value.name

  memory = 4096 #var.memory_mb
  vcpu = 2 #var.cpu

  #   cpu = {
  #     mode = "host-passthrough"
  #   }

  disk { volume_id = libvirt_volume.vm_vols[each.value.name].id }
  boot_device { dev = ["cdrom", "hd", "network"] }

  # filesystem {
  #   source   = pathexpand("~/")
  #   target   = "mnt"
  #   readonly = false
  # }

  # uses static  IP
  network_interface {
    network_name   = libvirt_network.carbon_net.name
    hostname       = each.value.name
    # addresses      = ["192.168.122.241"]
    # mac            = "AA:BB:CC:11:24:23"
    wait_for_lease = true
  }

  #  network_interface {
  #    bridge  = var.macvtap_iface
  #    hostname = var.hostname
  #  }


  # IMPORTANT
  # Ubuntu can hang is a isa-serial is not present at boot time.
  # If you find your CPU 100% and never is available this is why
  console {
    type        = "pty"
    target_port = "0"
    target_type = "serial"
  }
  # IMPORTANT
  # it will show no console otherwise
  video {
    type = "qxl"
  }

  graphics {
    type        = "spice"
    listen_type = "address"
    autoport    = "true"
  }

  #   connection {
  #     type             = "ssh"
  #     host             = self.network_interface[0].addresses[0]
  #     user             = "root"
  #     password         = "root"
  #     bastion_host     = "10.90.20.21"
  #     bastion_port     = "22"
  #     bastion_user     = "root"
  #     bastion_password = "password"
  #   }

  # Hostdev passtrough
  # provisioner "local-exec" {
  #   command = "virsh --connect ${var.provider_uri} attach-device ${var.hostname} --file passtrought-host.xml --live --persistent"
  # }

  # provisioner "remote-exec" {
  #   inline = ["echo first", "echo first second"]
  # }

  # provisioner "remote-exec" {
  #   inline = ["nmcli con mod \"$(nmcli -t -f NAME c s  | grep -v eth0)\" ipv4.addresses 10.0.0.60/24 ipv4.dns 10.0.0.1"]
  # }
}
