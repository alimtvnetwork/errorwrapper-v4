package errcmd

import (
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/stringslice"
)

func ArgsJoinPrepend(argPrepend string, args ...string) string {
	if len(args) == 0 {
		return argPrepend
	}

	slice := stringslice.PrependLineNew(argPrepend, args)

	return strings.Join(slice, constants.Space)
}
