package qemu

import (
	"shylinux.com/x/ice"
)

type runtime struct {
	ice.Code
	source string `data:"http://mirrors.tencent.com/tinycorelinux/10.x/x86_64/tcz/src/qemu/qemu-3.1.0.tar.xz"`
	list   string `name:"list path auto order build download" help:"虚拟机"`
}

func (s runtime) Build(m *ice.Message, arg ...string) {
	s.Code.Build(m, "")
}
func (s runtime) List(m *ice.Message, arg ...string) {
	s.Code.Source(m, "", arg...)
}

func init() { ice.CodeCtxCmd(runtime{}) }
