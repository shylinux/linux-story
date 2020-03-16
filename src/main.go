package main

import (
	"github.com/shylinux/icebergs"
	_ "github.com/shylinux/icebergs/base"
	_ "github.com/shylinux/icebergs/core"
	_ "github.com/shylinux/icebergs/misc"

	_ "github.com/shylinux/linux-story/cli/file"
	_ "github.com/shylinux/linux-story/cli/help"
	_ "github.com/shylinux/linux-story/cli/make"
	_ "github.com/shylinux/linux-story/cli/text"
	_ "github.com/shylinux/linux-story/cli/user"
)

func main() {
	println(ice.Run())
}
