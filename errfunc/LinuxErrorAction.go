package errfunc

import (
	"gitlab.com/evatix-go/enum/linuxtype"
)

type LinuxErrorAction struct {
	LinuxType linuxtype.Variant
	SimpleErrorFunc
}
