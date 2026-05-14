package errfunc

import (
	"gitlab.com/evatix-go/enum/linuxtype"
)

type LinuxErrorCollectorAction struct {
	LinuxType linuxtype.Variant
	CollectorFunc
}
