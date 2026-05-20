package errstr

import (
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/converters"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type newCharHashsetMapCreator struct{}

func (it *newCharHashsetMapCreator) Cap(
	capacity int,
	eachCollectionCapacity int,
) *CharHashsetMap {
	return &CharHashsetMap{
		CharHashsetMap: corestr.
			New.
			CharHashsetMap.
			Cap(
				capacity,
				eachCollectionCapacity),
	}
}

func (it *newCharHashsetMapCreator) NewCharHashsetMapUsingItemsError(
	err error,
	errVariation errtype.Variation,
	items []string,
) *CharHashsetMap {
	errWrapper := errnew.Type.Error(
		errVariation,
		err)
	length := 0

	if items != nil {
		length = len(items) / constants.ArbitraryCapacity3
	}

	return &CharHashsetMap{
		CharHashsetMap: corestr.
			New.
			CharHashsetMap.
			CapItems(length, length, items...),
		ErrorWrapper: errWrapper,
	}
}

// ErrorWrapperItemsPtr wrapper nil will point to empty error wrapper
func (it *newCharHashsetMapCreator) ErrorWrapperItemsPtr(
	eachHashsetCapacity int,
	errorWrapper *errorwrapper.Wrapper,
	items *[]string,
) *CharHashsetMap {
	var deref []string
	if items != nil {
		deref = *items
	}
	return &CharHashsetMap{
		CharHashsetMap: corestr.
			New.
			CharHashsetMap.
			CapItems(
				constants.ArbitraryCapacity30,
				eachHashsetCapacity,
				deref...,
			),
		ErrorWrapper: errorWrapper,
	}
}

// ErrorWrapperItems wrapper nil will point to empty error wrapper
func (it *newCharHashsetMapCreator) ErrorWrapperItems(
	eachHashsetCapacity int,
	errorWrapper *errorwrapper.Wrapper,
	items []string,
) *CharHashsetMap {
	return &CharHashsetMap{
		CharHashsetMap: corestr.
			New.
			CharHashsetMap.
			CapItems(
				constants.ArbitraryCapacity30,
				eachHashsetCapacity,
				items,
			),
		ErrorWrapper: errorWrapper,
	}
}

// EmptyCharHashsetMapUsingError wrapper nil will point to empty error wrapper
func (it *newCharHashsetMapCreator) EmptyCharHashsetMapUsingError(
	errorWrapper *errorwrapper.Wrapper,
) *CharHashsetMap {
	return &CharHashsetMap{
		CharHashsetMap: corestr.Empty.CharHashsetMap(),
		ErrorWrapper:   errorWrapper,
	}
}

func (it *newCharHashsetMapCreator) Empty() *CharHashsetMap {
	return &CharHashsetMap{
		CharHashsetMap: corestr.Empty.CharHashsetMap(),
	}
}

func (it *newCharHashsetMapCreator) ErrorWithTypePointerStringsPtr(
	errVariation errtype.Variation,
	err error,
	ptrItems *[]*string,
) *CharHashsetMap {
	errWrapper := errnew.Type.Error(errVariation, err)
	items := converters.StringsTo.PtrOfPtrToPtrStrings(ptrItems)
	length := len(*items)

	return &CharHashsetMap{
		CharHashsetMap: corestr.New.CharHashsetMap.CapItemsPtr(length, length/constants.ArbitraryCapacity3, items),
		ErrorWrapper:   errWrapper,
	}
}
