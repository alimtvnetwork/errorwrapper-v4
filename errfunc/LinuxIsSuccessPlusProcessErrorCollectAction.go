package errfunc

import (
	"gitlab.com/evatix-go/enum/linuxtype"
)

type LinuxIsSuccessPlusProcessErrorCollectAction struct {
	LinuxType linuxtype.Variant
	IsSuccessProcessorCollectorFunc
}
