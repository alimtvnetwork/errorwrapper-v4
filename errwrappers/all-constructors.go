package errwrappers

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errnew"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func New(capacity int) *Collection {
	wrappers := make(
		[]*errorwrapper.Wrapper,
		0,
		capacity)

	return &Collection{
		items: wrappers,
	}
}

func NewCap1() *Collection {
	wrappers := make(
		[]*errorwrapper.Wrapper,
		0,
		constants.ArbitraryCapacity1)

	return &Collection{
		items: wrappers,
	}
}

func NewCap2() *Collection {
	wrappers := make(
		[]*errorwrapper.Wrapper,
		0,
		constants.ArbitraryCapacity2)

	return &Collection{
		items: wrappers,
	}
}

func NewCap3() *Collection {
	wrappers := make(
		[]*errorwrapper.Wrapper,
		0,
		constants.ArbitraryCapacity3)

	return &Collection{
		items: wrappers,
	}
}

func NewCap4() *Collection {
	wrappers := make(
		[]*errorwrapper.Wrapper,
		0,
		constants.ArbitraryCapacity4)

	return &Collection{
		items: wrappers,
	}
}

func NewUsingErrors(errs ...error) *Collection {
	if errs == nil {
		return Empty()
	}

	return NewUsingErrorsPtr(&errs)
}

func NewUsingErrorsPtr(errs *[]error) *Collection {
	if errs == nil {
		return Empty()
	}

	wrappers := make(
		[]*errorwrapper.Wrapper,
		0,
		len(*errs))

	collection := &Collection{
		items: wrappers,
	}

	collection.
		AddErrorsPtr(errs)

	return collection
}

// NewUsingErrorWrappers
//
//  Don't clone the items, just adds those
func NewUsingErrorWrappers(
	errWrappers ...*errorwrapper.Wrapper,
) *Collection {
	if errWrappers == nil {
		return Empty()
	}

	return NewUsingErrorWrappersPtr(
		false,
		errWrappers,
	)
}

func NewUsingErrorWrapperOrFunc(
	existingErrorWrapper *errorwrapper.Wrapper,
	executorFunc func() *errorwrapper.Wrapper,
) *Collection {
	empty := Empty()

	empty.AddExistingErrorWrapperOrFunc(
		existingErrorWrapper,
		executorFunc)

	return empty
}

// NewUsingErrorWrappersClone
//
//  Clone and add items to the collection
func NewUsingErrorWrappersClone(
	errWrappers []*errorwrapper.Wrapper,
) *Collection {
	return NewUsingErrorWrappersPtr(
		true,
		errWrappers)
}

func NewUsingErrorWrappersPtr(
	isMakeClone bool,
	errWrappers []*errorwrapper.Wrapper,
) *Collection {
	hasCloned := false
	if errWrappers == nil {
		hasCloned = true
		errWrappers = []*errorwrapper.Wrapper{}
	}

	length := len(errWrappers)

	if isMakeClone && !hasCloned {
		wrappers := make(
			[]*errorwrapper.Wrapper,
			0,
			length)

		for _, wrapper := range errWrappers {
			if wrapper == nil || wrapper.IsEmpty() {
				continue
			}

			wrappers = append(wrappers, wrapper)
		}

		return &Collection{
			items: wrappers,
		}
	}

	// isMakeClone || hasCloned || !isMakeClone
	return &Collection{
		items: errWrappers,
	}
}

func NewWithItem(
	capacity int,
	variation errtype.Variation,
) *Collection {
	collection := New(capacity)

	return collection.Add(variation)
}

func NewWithType(variation errtype.Variation) *Collection {
	collection := NewCap1()

	return collection.AddWrappers(errorwrapper.NewPtr(variation))
}

func NewWithTypeUsingStackSkip(
	stackSkipIndex int,
	variation errtype.Variation,
) *Collection {
	collection := NewCap1()

	return collection.AddWrappers(
		errorwrapper.NewTypeUsingStackSkip(
			stackSkipIndex+codestack.Skip1,
			variation))
}

func NewWithMessage(variation errtype.Variation, msg string) *Collection {
	collection := NewCap1()

	return collection.AddWrappers(
		errorwrapper.NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			variation,
			msg))
}

func NewWithMessageUsingStackSkip(
	stackSkipIndex int,
	variation errtype.Variation, msg string,
) *Collection {
	collection := NewCap1()

	return collection.AddWrappers(
		errorwrapper.NewMsgDisplayErrorNoReference(
			stackSkipIndex+defaultSkipInternal,
			variation, msg))
}

func NewWithError(
	capacity int,
	variation errtype.Variation, err error,
) *Collection {
	if err == nil {
		return nil
	}

	collection := New(capacity)

	return collection.AddTypeError(
		variation, err)
}

func NewWithErrorUsingStackSkip(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
) *Collection {
	if err == nil {
		return nil
	}

	collection := New(constants.One)

	return collection.AddWrapperPtr(
		errnew.Messages.SingleUsingStackSkip(
			stackSkipIndex,
			errType,
			err.Error()))
}

func NewWithOnlyError(err error) *Collection {
	collection := NewCap2()

	collection.AddError(err)

	return collection
}

func NewUsingCollections(
	errorCollectionOfCollections ...*Collection,
) *Collection {
	if errorCollectionOfCollections == nil {
		return NewCap2()
	}

	return NewUsingCollectionsPtr(&errorCollectionOfCollections)
}

func NewUsingCollectionsPtr(
	errorCollectionOfCollections *[]*Collection,
) *Collection {
	if errorCollectionOfCollections == nil {
		return NewCap2()
	}

	allLengths := 0

	for _, errCollection := range *errorCollectionOfCollections {
		allLengths += errCollection.Length()
	}

	collection := New(allLengths)

	return collection.AddCollectionsPtr(
		errorCollectionOfCollections)
}

func NewWithOnlyCapError(capacity int, err error) *Collection {
	collection := New(capacity)

	collection.AddError(err)

	return collection
}

func Empty() *Collection {
	return New(0)
}

func MutexEmpty() *MutexCollection {
	collection := New(0)

	return &MutexCollection{
		internalCollection: collection,
	}
}

func MutexNew(capacity int) *MutexCollection {
	collection := New(capacity)

	return &MutexCollection{
		internalCollection: collection,
	}
}
