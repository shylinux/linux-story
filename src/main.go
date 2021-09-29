package main

import (
	"shylinux.com/x/ice"

	_ "shylinux.com/x/linux-story/src/gcc"
	_ "shylinux.com/x/linux-story/src/gdb"
	_ "shylinux.com/x/linux-story/src/glibc"

	_ "shylinux.com/x/linux-story/src/busybox"
	_ "shylinux.com/x/linux-story/src/kernel"
	_ "shylinux.com/x/linux-story/src/qemu"

	_ "shylinux.com/x/linux-story/iso/centos"
	_ "shylinux.com/x/linux-story/iso/ubuntu"
)

func main() { print(ice.Run()) }
