package errdefer

import (
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errfunc"
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
