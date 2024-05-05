package ctags

import "shylinux.com/x/ice"

type ctags struct {
	ice.Code
	source string `data:"http://prdownloads.sourceforge.net/ctags/ctags-5.8.tar.gz"`
	list   string `name:"list path auto order build download" help:"索引"`
}

func (s ctags) Build(m *ice.Message, arg ...string) {
	s.Code.Build(m, "")
}
func (s ctags) List(m *ice.Message, arg ...string) {
	s.Code.Source(m, "", arg...)
}

func init() { ice.CodeCtxCmd(ctags{}) }
