variable "vsphere_endpoint" {
  type        = string
  description = "vSphere URL."
  default     = ""

}

variable "vsphere_username" {
  type        = string
  description = "Username to authenticate to vShpere."
  default     = ""

}

variable "vsphere_password" {
  type        = string
  description = "Password to authenticate with vSphere."
  default     = ""

}

variable "vsphere_host" {
  type        = string
  description = "Host in vShere to provision VM on."
  default     = ""

}

variable "vsphere_datastore" {
  type        = string
  description = "Datastore in vSphere to provision VM on."
  default     = ""

}
