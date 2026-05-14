package errcmd

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func ArgsJoinSlicePtr(args *[]string) string {
	if args == nil || len(*args) == 0 {
		return constants.EmptyString
	}

	return strings.Join(*args, constants.Space)
}
