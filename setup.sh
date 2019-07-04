mkdir -p root/bin
mkdir -p root/proc
mkdir -p root/tmp
wget https://busybox.net/downloads/binaries/1.31.0-i686-uclibc/busybox
mv busybox root/bin
chmod +x root/bin/busybox
pushd root/bin
for i in $(./busybox --list)
 do
   ln -s busybox ./$i
 done
popd

