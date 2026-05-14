package errcmd

import (
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/stringslice"
)

func ProcessArgsJoinAppend(process string, args ...string) string {
	if len(args) == 0 {
		return process
	}

	slice := stringslice.AppendLineNew(args, process)

	return strings.Join(slice, constants.Space)
}
