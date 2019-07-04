Minimal Go container creator

run: go run containerize.go run /bin/sh

and you will have a minimal container to play around inside.
Only the most basic of restrictions are there, like:

chroot
namespace
process space

