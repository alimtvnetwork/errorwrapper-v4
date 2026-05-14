package errdefer

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/errorwrapper"
)

func mergeErrorWrapper(
	existingErrorWrapper *errorwrapper.Wrapper,
	closingErrorWrapper *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if closingErrorWrapper == nil || closingErrorWrapper.IsEmpty() {
		return existingErrorWrapper
	}

	hasExistingError := existingErrorWrapper != nil &&
		existingErrorWrapper.HasError() &&
		!existingErrorWrapper.IsEquals(closingErrorWrapper)

	if hasExistingError {
		//goland:noinspection ALL
		return existingErrorWrapper.
			ConcatNew().
			WrapperUsingStackSkip(
				codestack.Skip1,
				closingErrorWrapper)
	}

	return closingErrorWrapper
}
