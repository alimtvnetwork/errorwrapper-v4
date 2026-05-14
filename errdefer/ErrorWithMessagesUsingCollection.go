package errdefer

import (
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
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
