package errdefer

import (
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

func ErrorUsingFunc(
	existingErrorWrapper *errorwrapper.Wrapper, // could be nil
	errType errtype.Variation,
	errorFunc func() error,
) *errorwrapper.Wrapper {
	if errorFunc == nil {
		return existingErrorWrapper
	}

	closerErr := errorFunc()
	closingErrorWrapper := errnew.Type.Error(
		errType,
		closerErr,
	)

	return mergeErrorWrapper(
		existingErrorWrapper,
		closingErrorWrapper)
}
