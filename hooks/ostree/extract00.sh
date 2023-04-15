#!/bin/bash
#

TARGET=$1
pushd $TARGET
mkdir -p sysroot/ostree

# Move /etc to /usr/etc.
if [ -d $TARGET/usr/etc ]; then
	echo "ERROR: Non-empty /usr/etc found!" >&2
	ls -lR $TARGET/usr/etc
	exit 1
fi

# Need to create /usr since its not created in
# this stage.
echo "Moving /etc to usr/etc"
mkdir -p usr
mv etc usr
ln -s usr/etc etc
