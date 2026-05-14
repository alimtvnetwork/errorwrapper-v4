package errnew

import (
	"strings"

	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
	"github.com/alimtvnetwork/errorwrapper-v3/refs"
)

type newMergeToErrorWrapperCreator struct{}

func (it newMergeToErrorWrapperCreator) New(
	first,
	second *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if first.IsEmpty() && second.IsEmpty() {
		return nil
	}

	if first.IsEmpty() {
		return second
	}

	if second.IsEmpty() {
		return first
	}

	return first.
		ConcatNew().
		WrapperUsingStackSkip(
			defaultSkipInternal,
			second)
}

func (it newMergeToErrorWrapperCreator) UsingNewType(
	errType errtype.Variation,
	first,
	second *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if first.IsEmpty() && second.IsEmpty() {
		return nil
	}

	if first.IsEmpty() {
		return Type.Error(errType, second.CompiledError())
	}

	if second.IsEmpty() {
		return Type.Error(errType, first.CompiledError())
	}

	message := errorwrapper.MessagesJoined(
		first.FullString(),
		second.FullString())

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		message,
	)
}

func (it newMergeToErrorWrapperCreator) UsingStackSkip(
	stackSkipIndex int,
	first,
	second *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if first.IsEmpty() && second.IsEmpty() {
		return nil
	}

	if first.IsEmpty() {
		return second
	}

	if second.IsEmpty() {
		return first
	}

	return first.
		ConcatNew().
		WrapperUsingStackSkip(
			stackSkipIndex+defaultSkipInternal,
			second)
}

func (it newMergeToErrorWrapperCreator) Three(
	first,
	second,
	third *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if first.IsEmpty() && second.IsEmpty() && third.IsEmpty() {
		return nil
	}

	merged2 := it.New(first, second)

	return it.New(merged2, third)
}

func (it newMergeToErrorWrapperCreator) TwoWithRefs(
	variation errtype.Variation,
	first, second *errorwrapper.Wrapper,
	additionalReferences ...ref.Value,
) *errorwrapper.Wrapper {
	if first.IsEmpty() && second.IsEmpty() {
		return nil
	}

	if len(additionalReferences) == 0 {
		return it.UsingStackSkip(
			defaultSkipInternal,
			first,
			second)
	}

	return it.ManyAdditionalRefsUsingJoinerStackSkip(
		defaultSkipInternal,
		variation,
		errorwrapper.MessagesJoiner,
		refs.NewUsingValues(additionalReferences...),
		first,
		second)
}

func (it newMergeToErrorWrapperCreator) TwoWithRefsUsingStackSkip(
	stackSkipIndex int,
	variation errtype.Variation,
	first, second *errorwrapper.Wrapper,
	additionalReferences ...ref.Value,
) *errorwrapper.Wrapper {
	if first.IsEmpty() && second.IsEmpty() {
		return nil
	}

	if len(additionalReferences) == 0 {
		return it.UsingStackSkip(
			stackSkipIndex+defaultSkipInternal,
			first,
			second)
	}

	return it.ManyAdditionalRefsUsingJoinerStackSkip(
		stackSkipIndex+defaultSkipInternal,
		variation,
		errorwrapper.MessagesJoiner,
		refs.NewUsingValues(additionalReferences...),
		first,
		second)
}

// AddRefs
//
// Create new wrapper using the merged references
func (it newMergeToErrorWrapperCreator) AddRefs(
	currentWrapper *errorwrapper.Wrapper,
	additionalReferences ...ref.Value,
) *errorwrapper.Wrapper {
	if currentWrapper == nil || currentWrapper.IsEmpty() {
		return nil
	}

	references := currentWrapper.MergeNewReferences(
		additionalReferences...)

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		currentWrapper.Type(),
		currentWrapper.Error().Error(),
		references)
}

// AddRefsUsingStackSkip
//
// Create new wrapper using the merged references
func (it newMergeToErrorWrapperCreator) AddRefsUsingStackSkip(
	stackSkipIndex int,
	currentWrapper *errorwrapper.Wrapper,
	additionalReferences ...ref.Value,
) *errorwrapper.Wrapper {
	if currentWrapper == nil || currentWrapper.IsEmpty() {
		return nil
	}

	references := currentWrapper.MergeNewReferences(additionalReferences...)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		currentWrapper.Type(),
		currentWrapper.Error().Error(),
		references)
}

func (it newMergeToErrorWrapperCreator) Many(
	errType errtype.Variation,
	items ...*errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if len(items) == 0 {
		return nil
	}

	return it.ManyUsingJoinerStackSkip(
		defaultSkipInternal,
		errType,
		errorwrapper.MessagesJoiner,
		items...)
}

func (it newMergeToErrorWrapperCreator) ManyStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	items ...*errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if len(items) == 0 {
		return nil
	}

	return it.ManyUsingJoinerStackSkip(
		stackSkipIndex+defaultSkipInternal,
		errType,
		errorwrapper.MessagesJoiner,
		items...)
}

func (it newMergeToErrorWrapperCreator) ManyUsingJoinerStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	joiner string,
	items ...*errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if len(items) == 0 {
		return nil
	}

	compiledItems := make([]string, 0, len(items))
	references := refs.EmptyPtr()

	for _, item := range items {
		if item.IsEmpty() {
			continue
		}

		if item.HasCurrentError() {
			compiledItems = append(
				compiledItems,
				item.ErrorString())
		}

		references.AddCollection(
			item.References())
	}

	allMessageCompiled := strings.Join(compiledItems, joiner)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		errType,
		allMessageCompiled,
		references)
}

func (it newMergeToErrorWrapperCreator) ManyAdditionalRefsUsingJoinerStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	joiner string,
	additionalRefs *refs.Collection,
	items ...*errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if len(items) == 0 {
		return nil
	}

	compiledItems := make([]string, 0, len(items))
	references := refs.EmptyPtr()

	for _, item := range items {
		if item.IsEmpty() {
			continue
		}

		if item.HasCurrentError() {
			compiledItems = append(
				compiledItems,
				item.ErrorString())
		}

		references.AddCollection(
			item.References())
	}

	references.AddCollection(additionalRefs)

	allMessageCompiled := strings.Join(compiledItems, joiner)

	return errorwrapper.NewMsgDisplayError(
		stackSkipIndex+defaultSkipInternal,
		errType,
		allMessageCompiled,
		references)
}
