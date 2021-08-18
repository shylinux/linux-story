package main

import (
	ice "shylinux.com/x/icebergs"
	_ "shylinux.com/x/icebergs/base"
	_ "shylinux.com/x/icebergs/core"
	_ "shylinux.com/x/icebergs/misc"

	_ "shylinux.com/x/linux-story/iso/centos"
	_ "shylinux.com/x/linux-story/iso/ubuntu"

	_ "shylinux.com/x/linux-story/src/gcc"
	_ "shylinux.com/x/linux-story/src/gdb"
	_ "shylinux.com/x/linux-story/src/glibc"

	_ "shylinux.com/x/linux-story/src/busybox"
	_ "shylinux.com/x/linux-story/src/kernel"
	_ "shylinux.com/x/linux-story/src/qemu"
)

func main() { print(ice.Run()) }
