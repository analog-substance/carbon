#! /bin/bash

VMDIR="carbon-vm-ubuntu-$(date +'%y%m%d_%H%M')"
mkdir "$VMDIR"
VBoxManage createvm --name $VMDIR --ostype "Ubuntu_64" --register --basefolder `pwd`/$VMDIR/
mv output-carbon-vm-ubuntu/*-disk001.vmdk $VMDIR/${VMDIR}_disk01.vmdk
VBoxManage storagectl  $VMDIR --name "SATA Controller" --add sata --controller IntelAhci
VBoxManage storageattach $VMDIR --storagectl "SATA Controller" --port 0 --device 0 --type hdd --medium `pwd`/$VMDIR/${VMDIR}_disk01.vmdk
VBoxManage modifyvm $VMDIR --ioapic on
VBoxManage modifyvm $VMDIR --memory 8096 --vram 128
VBoxManage modifyvm $VMDIR --firmware efi
VBoxManage modifyvm $VMDIR --nic1 nat
VBoxManage modifyvm $VMDIR --nat-pf1 "guestssh,tcp,,2222,,22"
VBoxManage startvm $VMDIR