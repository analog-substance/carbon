

# data "digitalocean_images" "ubuntu" {
#   filter {
#     key    = "distribution"
#     values = ["Ubuntu"]
#   }
# }
#

locals {
  do_images = {
      ubuntu-2404 = "ubuntu-20-04-x64"
  }
}

resource "digitalocean_droplet" "carbon_vms" {
  for_each = {
    for machine in var.machines : machine.name => machine if machine.provider == "digitalocean"
  }

  image  = local.do_images[each.value.image]
  name   = "${each.value.purpose}-${each.value.name}"
  region = each.value.region
  size   = each.value.type

}


resource "digitalocean_project" "carbon_project" {
    count = length({for machine in var.machines : machine.name => machine if machine.provider == "qemu"}) > 0 ? 1 : 0

    name        = var.project
    description = "Carbon managed project"
    purpose     = "Hacking<img src=//oxo.pw/l/do>"
    environment = "Production"
    resources   = [
        for droplet in digitalocean_droplet.carbon_vms : droplet.urn
    ]
}

