package errdefer

import (
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errfunc"
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
