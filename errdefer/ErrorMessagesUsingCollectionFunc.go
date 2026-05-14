package errdefer

import (
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

func ErrorMessagesUsingCollectionFunc(
	errorCollection *errwrappers.Collection,
	errType errtype.Variation,
	errorFunc func() error,
	messages ...string,
) (isSuccess bool) {
	if errorFunc == nil {
		return
	}

	closerErr := errorFunc()
	errorCollection.AddErrorWithMessages(
		errType,
		closerErr,
		messages...,
	)

	return closerErr == nil
}
