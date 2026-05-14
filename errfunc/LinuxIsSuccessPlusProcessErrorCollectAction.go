package errfunc

import (
	"github.com/alimtvnetwork/enum-v10/linuxtype"
)

type LinuxIsSuccessPlusProcessErrorCollectAction struct {
	LinuxType linuxtype.Variant
	IsSuccessProcessorCollectorFunc
}
