package errdefer

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ErrorMessagesUsingFunc(
	existingErrorWrapper *errorwrapper.Wrapper, // could be nil
	errType errtype.Variation,
	errorFunc func() error,
	messages ...string,
) *errorwrapper.Wrapper {
	if errorFunc == nil {
		return existingErrorWrapper
	}

	closerErr := errorFunc()
	closingErrorWrapper := errnew.Messages.ErrorWithManyUsingStackSkip(
		codestack.Skip1,
		errType,
		closerErr,
		messages...,
	)

	return mergeErrorWrapper(
		existingErrorWrapper,
		closingErrorWrapper)
}
