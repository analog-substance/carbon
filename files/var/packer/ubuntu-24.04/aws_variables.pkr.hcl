
variable "aws_profile" {
  default     = "default"
  description = "The aws credentials profile to use."
  type        = string
}

variable "aws_build_region" {
  default     = "us-east-1"
  description = "The region in which to retrieve the base AMI from and build the new AMI."
  type        = string
}

variable "aws_tags" {
  default = {
    Architecture = "x86_64"
    OS_Version   = "Ubuntu Noble Numbat"
  }

  type = map(string)
}

variable "aws_run_tags" {
  default = {
    Environment = "testing"
    Type        = "builder"
  }

  type = map(string)
}

variable "aws_vpc_filters" {
  default = {
    "Name" : "AMI Builds"
  }

  type = map(string)
}

variable "aws_subnet_filters" {
  default = {
    "Name" : "AMI Builds"
  }

  type = map(string)
}
