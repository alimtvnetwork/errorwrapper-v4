package errcmd

import (
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
)

func ArgsJoinSlice(args []string) string {
	if len(args) == 0 {
		return constants.EmptyString
	}

	return strings.Join(args, constants.Space)
}
