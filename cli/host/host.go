package host

import "shylinux.com/x/ice"

type Host struct {
}

func init() {
	ice.Cmd("web.code.linux.cli.host", Host{})
}
