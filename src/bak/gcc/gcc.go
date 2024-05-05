package gcc

import (
	"shylinux.com/x/ice"
)

type compile struct {
	ice.Code
	source string `data:"http://mirrors.tencent.com/macports/distfiles/gcc48/gcc-4.8.4.tar.bz2"`
	list   string `name:"list path auto order build download" help:"编译器"`
}

func (s compile) Build(m *ice.Message, arg ...string) {
	s.Code.Build(m, "", "--enable-languages=c,c++", "--disable-multilib", "--disable-checking")
}
func (s compile) List(m *ice.Message, arg ...string) {
	s.Code.Source(m, "", arg...)
}

func init() { ice.CodeCtxCmd(compile{}) }
