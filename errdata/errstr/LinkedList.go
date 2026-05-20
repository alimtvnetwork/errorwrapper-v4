package errstr

import (
	"github.com/alimtvnetwork/core-v9/coredata/corestr"

	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type LinkedList struct {
	*corestr.LinkedList
	ErrorWrapper *errorwrapper.Wrapper
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		LinkedList:   corestr.Empty.LinkedList(),
		ErrorWrapper: nil,
	}
}

func NewLinkedListUsingItemsError(
	errVariation errtype.Variation,
	err error,
	items []string,
) *LinkedList {
	errWrapper := errnew.Type.Error(errVariation, err)

	return &LinkedList{
		LinkedList:   corestr.New.LinkedList.Strings(items),
		ErrorWrapper: errWrapper,
	}
}

// NewLinkedListUsingItemsErrorWrapper wrapper nil will point to empty error wrapper
func NewLinkedListUsingItemsErrorWrapper(
	items *[]string,
	errorWrapper *errorwrapper.Wrapper,
) *LinkedList {
	var deref []string
	if items != nil {
		deref = *items
	}
	return &LinkedList{
		LinkedList:   corestr.New.LinkedList.Strings(deref),
		ErrorWrapper: errorWrapper,
	}
}

// EmptyLinkedListUsingError wrapper nil will point to empty error wrapper
func EmptyLinkedListUsingError(
	wrapper *errorwrapper.Wrapper,
) *LinkedList {
	return &LinkedList{
		LinkedList:   corestr.Empty.LinkedList(),
		ErrorWrapper: wrapper,
	}
}

func EmptyLinkedList() *LinkedList {
	return &LinkedList{
		LinkedList: corestr.Empty.LinkedList(),
	}
}

func NewLinkedListUsingPtrItemsError(
	ptrItems []*string,
	err error,
	errVariation errtype.Variation,
) *LinkedList {
	errWrapper := errnew.Type.Error(errVariation, err)

	return &LinkedList{
		LinkedList:   corestr.New.LinkedList.PointerStringsPtr(&ptrItems),
		ErrorWrapper: errWrapper,
	}
}
