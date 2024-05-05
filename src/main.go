package main

import (
	"shylinux.com/x/ice"
	_ "shylinux.com/x/linux-story/src/system"
)

func main() { print(ice.Run()) }

func init() {
	ice.Info.NodeIcon = "src/system/linux.png"
	ice.Info.NodeMain = "web.code.system.studio"
}
