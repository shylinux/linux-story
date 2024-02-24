package main

import (
	"shylinux.com/x/ice"
	_ "shylinux.com/x/icebergs/base/aaa/portal"
	_ "shylinux.com/x/icebergs/core/chat/oauth"
	_ "shylinux.com/x/icebergs/misc/java"
	_ "shylinux.com/x/icebergs/misc/node"
	_ "shylinux.com/x/icebergs/misc/wx"

	_ "shylinux.com/x/linux-story/src/busybox"
	_ "shylinux.com/x/linux-story/src/gcc"
	_ "shylinux.com/x/linux-story/src/gdb"
	_ "shylinux.com/x/linux-story/src/glibc"
	_ "shylinux.com/x/linux-story/src/kernel"
	_ "shylinux.com/x/linux-story/src/qemu"
)

func main() { print(ice.Run()) }
