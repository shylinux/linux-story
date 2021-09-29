package glibc

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/linux-story/src/gdb"
)

type Source struct {
	gdb.Source
	source string `data:"http://mirrors.tencent.com/tinycorelinux/5.x/x86/tcz/src/glibc/glibc-2.17.tar.xz"`
	list   string `name:"list path auto build download" help:"标准库"`
}

func init() { ice.CodeCmd(Source{}) }
