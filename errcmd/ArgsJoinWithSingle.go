package errcmd

import (
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/stringslice"
)

func ArgsJoinWithSingle(arg1 string, args ...string) string {
	if len(args) == 0 {
		return constants.EmptyString
	}

	newSlice := stringslice.AppendLineNew(
		args,
		arg1)

	return strings.Join(newSlice, constants.Space)
}
