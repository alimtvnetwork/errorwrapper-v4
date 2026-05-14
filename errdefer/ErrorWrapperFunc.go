package errdefer

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errfunc"
)

func ErrorWrapperFunc(
	existingErrorWrapper *errorwrapper.Wrapper,
	errWrapperFunc errfunc.WrapperFunc,
) *errorwrapper.Wrapper {
	errWrapper := errWrapperFunc()

	return mergeErrorWrapper(
		existingErrorWrapper,
		errWrapper)
}
