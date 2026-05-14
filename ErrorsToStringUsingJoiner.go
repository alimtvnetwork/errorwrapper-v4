package errorwrapper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// ErrorsToStringUsingJoiner nil items will be ignored.
func ErrorsToStringUsingJoiner(
	joiner string,
	errItems ...error,
) string {
	length := len(errItems)

	if length == 0 {
		return constants.EmptyString
	}

	list := make([]string, 0, length)

	for _, err := range errItems {
		if err == nil {
			continue
		}

		list = append(
			list,
			err.Error())
	}

	return strings.Join(list, joiner)
}
