package ubuntu

import (
	"shylinux.com/x/ice"
)

type Source struct {
	ice.Code
	source string `data:"http://mirrors.tencent.com/macports/distfiles/gcc48/gcc-4.8.4.tar.bz2"`
	list   string `name:"list path auto build download" help:"操作系统"`
}

func init() { ice.CodeCmd(Source{}) }
