package errfunc

import (
	"github.com/alimtvnetwork/enum-v10/linuxtype"
)

type LinuxErrorCollectorAction struct {
	LinuxType linuxtype.Variant
	CollectorFunc
}
