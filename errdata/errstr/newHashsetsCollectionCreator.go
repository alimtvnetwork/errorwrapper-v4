package errstr

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
)

type newHashsetsCollectionCreator struct{}

func (it *newHashsetsCollectionCreator) UsingHashsets(sets []corestr.Hashset) *HashsetsCollection {
	return &HashsetsCollection{
		HashsetsCollection: corestr.New.HashsetsCollection.UsingHashsets(sets...),
	}
}

func (it *newHashsetsCollectionCreator) UsingSpreadHashsets(
	items ...corestr.Hashset,
) *HashsetsCollection {
	return &HashsetsCollection{
		HashsetsCollection: corestr.New.HashsetsCollection.UsingHashsets(items...),
		ErrorWrapper:       nil,
	}
}

func (it *newHashsetsCollectionCreator) UsingSpreadHashsetsPtr(
	items ...*corestr.Hashset,
) *HashsetsCollection {
	return &HashsetsCollection{
		HashsetsCollection: corestr.New.HashsetsCollection.UsingHashsetsPointers(items...),
		ErrorWrapper:       nil,
	}
}

func (it *newHashsetsCollectionCreator) LenCap(len, capacity int) *HashsetsCollection {
	return &HashsetsCollection{
		HashsetsCollection: corestr.New.HashsetsCollection.LenCap(len, capacity),
	}
}

// ErrorWrapperHashsets wrapper nil will point to empty error wrapper
func (it *newHashsetsCollectionCreator) ErrorWrapperHashsets(
	errorWrapper *errorwrapper.Wrapper,
	items ...*corestr.Hashset,
) *HashsetsCollection {
	return &HashsetsCollection{
		HashsetsCollection: corestr.
			New.
			HashsetsCollection.
			UsingHashsetsPointers(items...),
		ErrorWrapper: errorWrapper,
	}
}

// ErrorWrapper wrapper nil will point to empty error wrapper
func (it *newHashsetsCollectionCreator) ErrorWrapper(
	errorWrapper *errorwrapper.Wrapper,
) *HashsetsCollection {
	return &HashsetsCollection{
		HashsetsCollection: corestr.New.HashsetsCollection.Empty(),
		ErrorWrapper:       errorWrapper,
	}
}

func (it *newHashsetsCollectionCreator) Empty() *HashsetsCollection {
	return &HashsetsCollection{
		HashsetsCollection: corestr.New.HashsetsCollection.Empty(),
		ErrorWrapper:       nil,
	}
}
