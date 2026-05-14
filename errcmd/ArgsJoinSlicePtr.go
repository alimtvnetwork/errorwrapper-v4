package errcmd

import (
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
)

func ArgsJoinSlicePtr(args *[]string) string {
	if args == nil || len(*args) == 0 {
		return constants.EmptyString
	}

	return strings.Join(*args, constants.Space)
}
