locals {
  carbon_infra = yamldecode(file("${path.module}/carbon-config.yaml"))
}

module "{{.Name}}" {
  source   = "../../terraform/modules/carbon/aws"
  machines = local.carbon_infra["machines"]
  project = "{{.Name}}"
}