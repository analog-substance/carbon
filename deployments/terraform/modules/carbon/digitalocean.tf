terraform {
    required_providers {
        digitalocean = {
            source  = "digitalocean/digitalocean"
            version = "~> 2.0"
        }
    }
}

data "digitalocean_images" "ubuntu" {
    filter {
        key    = "distribution"
        values = ["Ubuntu"]
    }
}


locals {
    do_images = {
        desktop = data.digitalocean_images.ubuntu.id
    }
}

resource "digitalocean_droplet" "carbon_vm" {
    for_each                             = {
    for machine in var.machines :machine.name => machine if machine.provider == "digitalocean"
    }

    image   = local.do_images[each.value.image]
    name    = "${each.value.purpose}-${each.value.name}"
    region  = "nyc2"
    size                        = each.value.type
}
