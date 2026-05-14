package refs

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func QuickCompileString(
	joiner string,
	quickReferences ...QuickReference,
) (line string) {
	if len(quickReferences) == 0 {
		return constants.EmptyString
	}

	slice := QuickCompileStrings(quickReferences...)

	return strings.Join(slice, joiner)
}
