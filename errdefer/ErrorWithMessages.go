package errdefer

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
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
