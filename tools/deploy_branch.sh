#!/bin/bash

sysroot=/os
os_name=master
branch=starlingx/master

rm -rf /os ; mkdir -p /os
ostree admin init-fs $sysroot
ostree admin os-init --sysroot=$sysroot $os_name
ostree --repo=$sysroot/ostree/repo pull-local /artifacts/config/starlingx/master/repo $branch
ostree admin deploy --sysroot=$sysroot --os=$os_name $branch
