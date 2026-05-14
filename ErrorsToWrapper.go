package errorwrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/internal/consts"
)

// ErrorsToWrapper nil items will be ignored.
func ErrorsToWrapper(
	errType errtype.Variation,
	errItems ...error,
) *Wrapper {
	toString := ErrorsToStringUsingJoiner(
		consts.DefaultErrorLineSeparator,
		errItems...)

	return NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		toString)
}
