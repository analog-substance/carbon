#
# resource "aws_instance" "engagement_vm" {
#   for_each                              = { for machine in var.machines:machine.name => instance if machine.provider == "aws" }
#   ami                                  = local.amis[each.value.image]
#   instance_type                        = each.value.type
#   availability_zone                    = element(flatten(data.aws_availability_zones.available.*.names), index(var.machines,each.value))
#   subnet_id                            = element(aws_subnet.cluster_subnet.*.id, index(var.machines,each.value))
#   instance_initiated_shutdown_behavior = "stop"
#   hibernation                          = each.value.size == "g4dn.xlarge" ? false : true
#   root_block_device {
#     volume_size = each.value.volume_size
#     encrypted   = true
#
#     tags = {
#       "user:engagement" = var.engagement_name
#       "user:environment" = var.environment
#       "user:purpose" = each.value.purpose
#     }
#   }
#
#   #iam_instance_profile         = aws_iam_instance_profile.redteam_profile.id
#   vpc_security_group_ids = [
#     aws_security_group.generic_acl.id,
#   ]
#   associate_public_ip_address = true
#   #user_data                    = data.template_cloudinit_config.readteam_desktop[count.index].rendered
#
#   lifecycle {
#     ignore_changes = [ami, associate_public_ip_address, root_block_device]
#   }
#
#   tags = {
#     Name              = "${var.engagement_name}:${each.value.purpose}:${each.value.name}"
#     "user:engagement" = var.engagement_name
#     "user:environment" = var.environment
#     "user:purpose" = each.value.purpose
#   }
# }
