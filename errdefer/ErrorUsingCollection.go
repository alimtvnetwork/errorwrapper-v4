package errdefer

import (
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
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
