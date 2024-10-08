
variable "machines" {
  type = list(object({
    name    = string
    image     = optional(string, "desktop")
    type    = optional(string, "t2.medium")
    profile = optional(string, "")
    purpose = optional(string, "util")
    volume_size = optional(number, 80)
    provider = optional(string, "aws")
  }))
  default = []
}

variable "disk_source" {
  type = string
  default = "../../../images/qemu/carbon-ubuntu-vm-20241004162421/carbon-ubuntu-vm-20241004162421"
}