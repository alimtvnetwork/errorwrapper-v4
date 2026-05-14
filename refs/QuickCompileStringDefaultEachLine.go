package refs

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func QuickCompileStringDefaultEachLine(
	quickReferences ...QuickReference,
) (line string) {
	if len(quickReferences) == 0 {
		return constants.EmptyString
	}

	slice := QuickCompileStrings(quickReferences...)

	return strings.Join(slice, constants.NewLineUnix)
}
