package errstr

import (
	"github.com/alimtvnetwork/core-v9/coredata/corestr"

	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

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
	errWrapper := errnew.Type.Error(errVariation, err)

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
	errWrapper := errnew.Type.Error(errVariation, err)

	return &LinkedCollections{
		LinkedCollections: corestr.New.LinkedCollection.PointerStringsPtr(&ptrItems),
		ErrorWrapper:      errWrapper,
	}
}
