package gdb

import (
	"shylinux.com/x/ice"
)

type runtime struct {
	ice.Code
	source string `data:"http://mirrors.tencent.com/macports/distfiles/avr-gdb/gdb-7.6.1.tar.bz2"`
	list   string `name:"list path auto order build download" help:"调试器"`
}

func (s runtime) Build(m *ice.Message, arg ...string) {
	s.Code.Build(m, "")
}
func (s runtime) List(m *ice.Message, arg ...string) {
	s.Code.Source(m, "", arg...)
}

func init() { ice.CodeCtxCmd(runtime{}) }
