package errstr

import (
	"errors"

	"github.com/alimtvnetwork/core-v9/coredata/corestr"

	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
)

func linkedCollectionsErrorWrapper(errVariation errtype.Variation, err error) *errorwrapper.Wrapper {
	if err == nil {
		err = errors.New(errVariation.Name())
	}

	return errnew.Type.Error(errVariation, err)
}

type LinkedCollections struct {
	*corestr.LinkedCollections
	ErrorWrapper *errorwrapper.Wrapper
}

func NewLinkedCollections() *LinkedCollections {
	return &LinkedCollections{
		LinkedCollections: corestr.Empty.LinkedCollections(),
		ErrorWrapper:      nil,
	}
}

func NewLinkedCollectionsUsingItemsError(
	errVariation errtype.Variation,
	err error,
	items []string,
) *LinkedCollections {
	errWrapper := linkedCollectionsErrorWrapper(errVariation, err)

	return &LinkedCollections{
		LinkedCollections: corestr.New.LinkedCollection.Strings(items...),
		ErrorWrapper:      errWrapper,
	}
}

// NewLinkedCollectionsUsingItemsErrorWrapper wrapper nil will point to empty error wrapper
func NewLinkedCollectionsUsingItemsErrorWrapper(
	items *[]string,
	errorWrapper *errorwrapper.Wrapper,
) *LinkedCollections {
	var deref []string
	if items != nil {
		deref = *items
	}
	return &LinkedCollections{
		LinkedCollections: corestr.New.LinkedCollection.Strings(deref...),
		ErrorWrapper:      errorWrapper,
	}
}

// EmptyLinkedCollectionsUsingError wrapper nil will point to empty error wrapper
func EmptyLinkedCollectionsUsingError(
	wrapper *errorwrapper.Wrapper,
) *LinkedCollections {
	return &LinkedCollections{
		LinkedCollections: corestr.Empty.LinkedCollections(),
		ErrorWrapper:      wrapper,
	}
}

func EmptyLinkedCollections() *LinkedCollections {
	return &LinkedCollections{
		LinkedCollections: corestr.Empty.LinkedCollections(),
	}
}

func NewLinkedCollectionsUsingPtrItemsError(
	ptrItems []*string,
	err error,
	errVariation errtype.Variation,
) *LinkedCollections {
	errWrapper := linkedCollectionsErrorWrapper(errVariation, err)

	return &LinkedCollections{
		LinkedCollections: corestr.New.LinkedCollection.PointerStringsPtr(&ptrItems),
		ErrorWrapper:      errWrapper,
	}
}
