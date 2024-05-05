package kernel

import (
	"shylinux.com/x/ice"
)

type runtime struct {
	ice.Code
	source string `data:"http://mirrors.tencent.com/osmc/osmc/download/buildroot/2014.05/linux-3.12.20.tar.xz"`
	list   string `name:"list path auto order build download" help:"源代码"`
}

func (s runtime) Build(m *ice.Message, arg ...string) {
	s.Code.Build(m, "", func(p string) { s.Code.System(m, p, "make", "defconfig") })
}
func (s runtime) List(m *ice.Message, arg ...string) {
	s.Code.Source(m, "", arg...)
}

func init() { ice.CodeCtxCmd(runtime{}) }
