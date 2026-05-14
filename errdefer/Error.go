package errdefer

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func Error(
	existingErrorWrapper *errorwrapper.Wrapper, // could be nil
	errType errtype.Variation,
	err error,
) *errorwrapper.Wrapper {
	closingErrorWrapper := errnew.Type.Error(
		errType,
		err)

	return mergeErrorWrapper(
		existingErrorWrapper,
		closingErrorWrapper)
}
