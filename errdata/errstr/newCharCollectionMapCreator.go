package errstr

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

type newCharCollectionMapCreator struct{}

func (it *newCharCollectionMapCreator) Cap(capacity int, eachCollectionCapacity int) *CharCollectionMap {
	return &CharCollectionMap{
		CharCollectionMap: corestr.
			New.
			CharCollectionMap.
			CapSelfCap(capacity, eachCollectionCapacity),
	}
}

func (it *newCharCollectionMapCreator) ErrorWithItems(
	errVariation errtype.Variation,
	err error,
	items []string,
) *CharCollectionMap {
	errWrapper := errnew.Type.Error(errVariation, err)

	return &CharCollectionMap{
		CharCollectionMap: corestr.
			New.
			CharCollectionMap.
			Items(items),
		ErrorWrapper: errWrapper,
	}
}

// ItemsWithErrorWrapper wrapper nil will point to empty error wrapper
func (it *newCharCollectionMapCreator) ItemsWithErrorWrapper(
	items []string,
	errorWrapper *errorwrapper.Wrapper,
) *CharCollectionMap {
	length := len(items)

	return &CharCollectionMap{
		CharCollectionMap: corestr.New.CharCollectionMap.ItemsPtrWithCap(
			length,
			constants.ArbitraryCapacity30,
			&items,
		),
		ErrorWrapper: errorWrapper,
	}
}

// ErrorWrapper wrapper nil will point to empty error wrapper
func (it *newCharCollectionMapCreator) ErrorWrapper(
	errorWrapper *errorwrapper.Wrapper,
) *CharCollectionMap {
	return &CharCollectionMap{
		CharCollectionMap: corestr.Empty.CharCollectionMap(),
		ErrorWrapper:      errorWrapper,
	}
}

func (it *newCharCollectionMapCreator) Empty() *CharCollectionMap {
	return &CharCollectionMap{}
}

func (it *newCharCollectionMapCreator) ErrorWithPointerItemsPtr(
	errVariation errtype.Variation,
	err error,
	ptrItems *[]*string,
) *CharCollectionMap {
	errWrapper := errnew.Type.Error(errVariation, err)
	items := converters.
		PointerStringsToStrings(ptrItems)

	return &CharCollectionMap{
		CharCollectionMap: corestr.
			New.
			CharCollectionMap.
			ItemsPtr(items),
		ErrorWrapper: errWrapper,
	}
}
