package errdefer

import (
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
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
