package errdefer

import (
	"github.com/alimtvnetwork/errorwrapper-v3/errfunc"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

func ErrorWrapperFuncUsingCollection(
	errorCollection *errwrappers.Collection,
	errWrapperFunc errfunc.WrapperFunc,
) (isSuccess bool) {
	errWrapper := errWrapperFunc()
	errorCollection.AddWrapperPtr(errWrapper)

	return errWrapper.IsEmpty()
}
