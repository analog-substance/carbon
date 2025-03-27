# Copied from Carbon {{ .Version }}

variable "machines" {
  type = list(object({
    name = string
    image = optional(string, "desktop")
    type = optional(string, "t2.medium")
    profile = optional(string, "")
    purpose = optional(string, "util")
    volume_size = optional(number, 80)
    provider = optional(string, "aws")
    region = optional(string, "")
  }))
  default = []
}

variable "project" {
  type = string
  default = "carbon"
}
