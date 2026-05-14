package errfunc

import (
	"gitlab.com/evatix-go/enum/linuxtype"
)

type LinuxIsSuccessPlusErrorCollectAction struct {
	LinuxType linuxtype.Variant
	IsSuccessCollectorFunc
}
