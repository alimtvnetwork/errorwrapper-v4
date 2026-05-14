package errdefer

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
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
