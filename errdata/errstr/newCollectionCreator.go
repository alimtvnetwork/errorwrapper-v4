package errstr

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

type newCollectionCreator struct{}

func (it *newCollectionCreator) Empty() *Collection {
	return &Collection{}
}

func (it *newCollectionCreator) Strings(
	items []string,
) *Collection {
	return &Collection{
		Collection: corestr.New.Collection.Strings(items),
	}
}

func (it *newCollectionCreator) Error(
	errType errtype.Variation,
	err error,
) *Collection {
	return &Collection{
		ErrorWrapper: errnew.Type.ErrorUsingStackSkip(
			codestack.Skip1,
			errType,
			err),
	}
}

func (it *newCollectionCreator) ErrorWrapper(
	errorWrapper *errorwrapper.Wrapper,
) *Collection {
	return &Collection{
		ErrorWrapper: errorWrapper,
	}
}

func (it *newCollectionCreator) Create(
	hashset *corestr.Collection,
	errorWrapper *errorwrapper.Wrapper,
) *Collection {
	return &Collection{
		Collection:   hashset,
		ErrorWrapper: errorWrapper,
	}
}

func (it *newCollectionCreator) StringsErrorWrapper(
	items []string,
	errorWrapper *errorwrapper.Wrapper,
) *Collection {
	return &Collection{
		Collection:   corestr.New.Collection.Strings(items),
		ErrorWrapper: errorWrapper,
	}
}

func (it *newCollectionCreator) ValueOnly(
	items []string,
) *Collection {
	return &Collection{
		Collection: corestr.New.Collection.Strings(items),
	}
}

func (it *newCollectionCreator) Cap(capacity int) *Collection {
	return &Collection{
		Collection: corestr.New.Collection.Cap(capacity),
	}
}

func (it *newCollectionCreator) ErrorWrapperWithSpreadStrings(
	wrapper *errorwrapper.Wrapper,
	items ...string,
) *Collection {
	return &Collection{
		Collection: corestr.New.Collection.Strings(
			items),
		ErrorWrapper: wrapper,
	}
}

func (it *newCollectionCreator) LenCap(len, capacity int) *Collection {
	return &Collection{
		Collection:   corestr.New.Collection.LenCap(len, capacity),
		ErrorWrapper: nil,
	}
}

func (it *newCollectionCreator) ItemsWithError(
	err error,
	errVariation errtype.Variation,
	items []string,
) *Collection {
	errWrapper := errnew.Type.Error(errVariation, err)

	return &Collection{
		Collection: corestr.New.Collection.Strings(
			items,
		),
		ErrorWrapper: errWrapper,
	}
}

// ItemsWithErrorWrapper
//
// wrapper nil will point to empty error wrapper
func (it *newCollectionCreator) ItemsWithErrorWrapper(
	isMakeClone bool,
	errorWrapper *errorwrapper.Wrapper,
	items *[]string,
) *Collection {
	return &Collection{
		Collection: corestr.New.Collection.StringsPtrOption(
			isMakeClone,
			items),
		ErrorWrapper: errorWrapper,
	}
}
