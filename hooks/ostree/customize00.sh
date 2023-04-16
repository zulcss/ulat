#!/bin/bash
#

tmpfiles_conf=usr/etc/tmpfiles.d/00ostree-tmpfiles.conf
TARGET=$1
cd $TARGET

mkdir -p var/rootdirs
touch $tmpfiles_conf

dirs="home opt srv mnt media"
for d in $dirs; do \
    if [ -d ${d} ] && [ ! -L ${d} ]; then
        mv $d var/rootdirs
        ln -s var/rootdirs/$d $d
        echo "d /var/rootdirs/${dir} 0755 root root -" >>${tmpfiles_conf}
    fi
done

mkdir -p usr/etc/tmpfiles.d
echo "d /var/rootdirs 0755 root root -" >>${tmpfiles_conf}
# disable the annoying logs on the console
echo "w /proc/sys/kernel/printk - - - - 3" >> ${tmpfiles_conf}
echo "d /var/home 0755 root root -" >>${tmpfiles_conf}

echo "d /var/rootdirs/opt 0755 root root -" >>${tmpfiles_conf}
if [ -d var/opt ]; then
    rm -rf var/opt
    mv opt var/rootdirs/opt
    ln -s var/rootdirs/opt opt
fi

if [ -d /var/lib/dpkg ]; then
    mv var/lib/dpkg usr/share/dpkg/db
    echo "L /var/lib/dpkg - - - - /usr/share/dpkg/db" >>${tmpfiles_conf}
fi

if [ -d root ] && [ ! L root ]; then
    echo "d /var/rootdirs/root 0755 root root -" >>${tmpfiles_conf}
    mv root var/rootdirs
    ln -sf var/rootdirs/root root
fi

mkdir -p boot/loader.0
mkdir -p boot/loader.1
ln -sf boot/loader.0 boot/loader
