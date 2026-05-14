package errdefer

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

func ErrorWithMessages(
	existingErrorWrapper *errorwrapper.Wrapper, // could be nil
	errType errtype.Variation,
	err error,
	messages ...string,
) *errorwrapper.Wrapper {
	closingErrorWrapper := errnew.Messages.ErrorWithManyUsingStackSkip(
		codestack.Skip1,
		errType,
		err,
		messages...)

	return mergeErrorWrapper(
		existingErrorWrapper,
		closingErrorWrapper)
}
