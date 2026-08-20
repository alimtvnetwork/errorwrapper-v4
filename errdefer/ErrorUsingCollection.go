package errdefer

import (
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
)

func ErrorUsingCollection(
	errorCollection *errwrappers.Collection,
	errType errtype.Variation,
	err error,
) {
	errorCollection.AddTypeError(
		errType,
		err)
}
