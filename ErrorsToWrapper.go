package errorwrapper

import (
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/internal/consts"
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
