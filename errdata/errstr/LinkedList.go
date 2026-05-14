package errstr

import (
	"gitlab.com/evatix-go/core/coredata/corestr"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
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
	return &LinkedList{
		LinkedList:   corestr.New.LinkedList.StringsPtr(items),
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
