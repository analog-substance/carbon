terraform {
  required_providers {
    libvirt = {
      source = "dmacvicar/libvirt"
      version = "0.7.6"
    }
  }
}
# Configure the Libvirt provider
provider "libvirt" {
  uri = "qemu:///system"
}

# Create a new domain
resource "libvirt_network" "kube_network" {
  # the name used by libvirt
  name = "k8snet"

  # mode can be: "nat" (default), "none", "route", "open", "bridge"
  mode = "nat"

  #  the domain used by the DNS server in this network
  domain = "k8s.local"

  #  list of subnets the addresses allowed for domains connected
  # also derived to define the host addresses
  # also derived to define the addresses served by the DHCP server
  addresses = ["10.17.3.0/24", "2001:db8:ca2:2::1/64"]

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
    enabled = true
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