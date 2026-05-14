package errcmd

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
)

func ProcessArgsJoinAppend(process string, args ...string) string {
	if len(args) == 0 {
		return process
	}

	slice := stringslice.AppendLineNew(args, process)

	return strings.Join(slice, constants.Space)
}
