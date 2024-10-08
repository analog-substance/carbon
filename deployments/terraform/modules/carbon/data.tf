
# data "aws_ami" "carbon_desktop" {
#   owners      = ["self"]
#   most_recent = true
#
#   filter {
#     name   = "virtualization-type"
#     values = ["hvm"]
#   }
#
#   filter {
#     name   = "root-device-type"
#     values = ["ebs"]
#   }
#
#   filter {
#     name   = "name"
#     values = ["carbon-ubuntu-desktop-*"]
#   }
# }
#
# locals {
#   amis = {
#     desktop  = data.aws_ami.carbon_desktop.id
#   }
# }


# variable qemu_images {
#   for_each = { for image_path in fileset(path.module, "../../../images/qemu/*/*") : image_path => image_path }
#
#   # other configuration using each.value
# }
