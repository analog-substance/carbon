#! /bin/bash

VM_DIR="carbon-vm-ubuntu-$(date +'%y%m%d_%H%M')"
mkdir "$VM_DIR"
VBoxManage createvm --name $VM_DIR --ostype "Ubuntu_64" --register --basefolder `pwd`/$VM_DIR/
mv output-carbon-vm-ubuntu/*-disk001.vmdk $VM_DIR/${VM_DIR}_disk01.vmdk
VBoxManage storagectl  $VM_DIR --name "SATA Controller" --add sata --controller IntelAhci
VBoxManage storageattach $VM_DIR --storagectl "SATA Controller" --port 0 --device 0 --type hdd --medium `pwd`/$VM_DIR/${VM_DIR}_disk01.vmdk
VBoxManage modifyvm $VM_DIR --ioapic on
VBoxManage modifyvm $VM_DIR --memory 8096 --vram 128
VBoxManage modifyvm $VM_DIR --firmware efi
VBoxManage modifyvm $VM_DIR --nic1 nat
VBoxManage modifyvm $VM_DIR --nat-pf1 "guestssh,tcp,,2222,,22"
VBoxManage startvm $VM_DIR