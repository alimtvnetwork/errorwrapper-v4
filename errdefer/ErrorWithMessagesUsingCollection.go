package errdefer

import (
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

func ErrorWithMessagesUsingCollection(
	errorCollection *errwrappers.Collection,
	errType errtype.Variation,
	err error,
	messages ...string,
) (isSuccess bool) {
	errorCollection.AddErrorWithMessages(
		errType,
		err,
		messages...)

	return err == nil
}
