terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

variable "bucket_prefix" {
  default = "tf-state"
}

resource "aws_s3_bucket" "backend" {
  bucket_prefix = var.bucket_prefix
}

## cant use ACLs :(
## We may need to get that fixed for future fun things, but probably not required for this
# resource "aws_s3_bucket_acl" "backend" {
#   bucket = aws_s3_bucket.backend.id
#   acl    = "private"
# }

resource "aws_s3_bucket_versioning" "backend" {
  bucket = aws_s3_bucket.backend.id
  versioning_configuration {
    status = "Enabled"
  }
}

