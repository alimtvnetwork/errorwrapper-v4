package errdefer

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
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
