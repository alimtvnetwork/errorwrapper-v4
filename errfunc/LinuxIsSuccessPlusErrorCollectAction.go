package errfunc

import (
	"github.com/alimtvnetwork/enum-v10/linuxtype"
)

type LinuxIsSuccessPlusErrorCollectAction struct {
	LinuxType linuxtype.Variant
	IsSuccessCollectorFunc
}
