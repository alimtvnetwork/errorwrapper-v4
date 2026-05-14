package errcmd

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
)

func ArgsJoinPrepend(argPrepend string, args ...string) string {
	if len(args) == 0 {
		return argPrepend
	}

	slice := stringslice.PrependLineNew(argPrepend, args)

	return strings.Join(slice, constants.Space)
}
