package main

import (
	"shylinux.com/x/ice"
	_ "shylinux.com/x/ice/devops"

	_ "shylinux.com/x/linux-story/src/busybox"
	_ "shylinux.com/x/linux-story/src/gcc"
	_ "shylinux.com/x/linux-story/src/gdb"
	_ "shylinux.com/x/linux-story/src/glibc"
	_ "shylinux.com/x/linux-story/src/kernel"
	_ "shylinux.com/x/linux-story/src/qemu"
)

func main() { print(ice.Run()) }

func init() { ice.Info.NodeIcon = "src/kernel/linux.png" }
