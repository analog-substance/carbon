
locals {
  carbon_infra = yamldecode(file("${path.module}/carbon.yaml"))
}

module "example-qemu-carbon" {
  source = "../../modules/carbon/"

  machines = local.carbon_infra["machines"]
}