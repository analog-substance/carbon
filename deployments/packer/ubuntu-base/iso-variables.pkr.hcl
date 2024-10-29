variable "iso_url" {
  type    = string
  description = "URL to download the ISO file"
  #   default     = "https://releases.ubuntu.com/noble/ubuntu-24.04.1-live-server-amd64.iso"
  default = "https://releases.ubuntu.com/noble/ubuntu-24.04-live-server-amd64.iso"
}

variable "iso_checksum" {
  type    = string
  description = "Checksum to validate the downloaded ISO file"
  #   default     = "sha256:e240e4b801f7bb68c20d1356b60968ad0c33a41d00d828e74ceb3364a0317be9"
  default = "sha256:8762f7e74e4d64d72fceb5f70682e6b069932deedb4949c6975d0f0fe0a91be3"
}

variable "ssh_username" {
  type        = string
  description = "Username to use when connecting via SSH"
  default     = "isotope"
}

variable "ssh_password" {
  type        = string
  description = "Password to use when connecting via SSH"
  default     = "carbon"
}

variable "boot_command" {
  type        = string
  description = "Keyboard sequence to execute to properly boot the image."
  default     = "e<wait><down><down><down><end> autoinstall 'ds=nocloud-net;s=http://{{ .HTTPIP }}:{{ .HTTPPort }}/'<F10>"
}
