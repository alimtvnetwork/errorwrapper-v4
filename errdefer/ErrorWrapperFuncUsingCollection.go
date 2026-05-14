package errdefer

import (
	"gitlab.com/evatix-go/errorwrapper/errfunc"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

func ErrorWrapperFuncUsingCollection(
	errorCollection *errwrappers.Collection,
	errWrapperFunc errfunc.WrapperFunc,
) (isSuccess bool) {
	errWrapper := errWrapperFunc()
	errorCollection.AddWrapperPtr(errWrapper)

	return errWrapper.IsEmpty()
}
