package gdb

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/linux-story/src/gcc"
)

type Source struct {
	gcc.Source
	source string `data:"http://mirrors.tencent.com/macports/distfiles/avr-gdb/gdb-7.6.1.tar.bz2"`
	list   string `name:"list path auto build download" help:"调试器"`
}

func (s Source) Build(m *ice.Message, arg ...string) {
	s.Code.Build(m)
}

func init() { ice.CodeCmd(Source{}) }
