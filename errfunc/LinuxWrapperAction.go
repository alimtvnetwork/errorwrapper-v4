package errfunc

import (
	"gitlab.com/evatix-go/enum/linuxtype"
)

type LinuxWrapperAction struct {
	LinuxType linuxtype.Variant
	WrapperFunc
}
