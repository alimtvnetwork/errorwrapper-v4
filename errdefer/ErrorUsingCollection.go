package errdefer

import (
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
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
