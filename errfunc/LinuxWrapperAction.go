package errfunc

import (
	"github.com/alimtvnetwork/enum-v10/linuxtype"
)

type LinuxWrapperAction struct {
	LinuxType linuxtype.Variant
	WrapperFunc
}
