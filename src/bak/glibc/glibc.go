package glibc

import (
	"shylinux.com/x/ice"
)

type project struct {
	ice.Code
	source string `data:"https://mirrors.tencent.com/tinycorelinux/5.x/x86/tcz/src/glibc/glibc-2.17.tar.xz"`
	list   string `name:"list path auto order build download" help:"标准库"`
}

func (s project) Build(m *ice.Message, arg ...string) {
	s.Code.Build(m, "")
}
func (s project) List(m *ice.Message, arg ...string) {
	s.Code.Source(m, "", arg...)
}

func init() { ice.CodeCtxCmd(project{}) }
