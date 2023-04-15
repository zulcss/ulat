#!/bin/bash
#

TARGET=$1
pushd $TARGET

mkdir -p var/rootdirs
dirs="home opt srv mnt media"
for d in $dirs; do \
    if [ -d $d ]l then
        echo "Moving $d to var/rootdirs"
        mv $d var/rootdirs
        echo "Symlinking $d to /$d"
        ln -s var/rootdirs/$d $d
    fi
done

mv usr/local var/rootdirs/usrlocal
mv root var/rootdirs/roothome
mv tmp sysroot/tmp

ln -sr var/rootdirs/usrlocal usr/local
ln -s var/rootdirs/roothome root
ln -s sysroot/tmp tmp

if [ -d var/lib/dpkg ]; then
    mv var/lib/dpkg usr/share/dpkg/db
    ln -sr usr/share/dpkg/db var/lib/dpkg
fi

if [ -d var/lib/apt ]; then
    mv var/lib/apt usr/share/apt/db
    ln -sr usr/share/apt/db var/lib/apt
fi

# Creating boot directories required for "ostree admin deploy"
mkdir -p boot/loader.0
mkdir -p boot/loader.1
ln -sf boot/loader.0 boot/loader

if [ -d /boot/efi ]; then
    cp -a boot/efi usr/lib/ostree-boot
fi
