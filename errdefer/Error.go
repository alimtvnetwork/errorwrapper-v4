package errdefer

import (
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
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
