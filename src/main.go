package main

import (
	ice "github.com/shylinux/icebergs"
	_ "github.com/shylinux/icebergs/base"
	_ "github.com/shylinux/icebergs/core"
	_ "github.com/shylinux/icebergs/misc"

	_ "github.com/shylinux/linux-story/iso/centos"
	_ "github.com/shylinux/linux-story/iso/ubuntu"

	_ "github.com/shylinux/linux-story/src/gcc"
	_ "github.com/shylinux/linux-story/src/gdb"
	_ "github.com/shylinux/linux-story/src/glibc"

	_ "github.com/shylinux/linux-story/src/busybox"
	_ "github.com/shylinux/linux-story/src/kernel"
	_ "github.com/shylinux/linux-story/src/qemu"
)

func main() { print(ice.Run()) }
