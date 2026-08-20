package errdefer

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
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
