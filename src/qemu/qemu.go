package qemu

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/linux-story/src/gdb"
)

type Source struct {
	gdb.Source
	source string `data:"http://mirrors.tencent.com/tinycorelinux/10.x/x86_64/tcz/src/qemu/qemu-3.1.0.tar.xz"`
	list   string `name:"list path auto build download" help:"虚拟机"`
}

func init() { ice.CodeCmd(Source{}) }
