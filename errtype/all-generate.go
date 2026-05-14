package errtype

import (
	"github.com/alimtvnetwork/core-v9/corecsv"
	"github.com/alimtvnetwork/core-v9/coreimpl/enumimpl"
)

func generateAllNamesSlice() []string {
	length := len(errTypesMap)
	sliceNames := make([]string, length)

	index := 0
	for variation := range errTypesMap {
		sliceNames[index] = variation.Name()
		index++
	}

	return sliceNames
}

func generateRangesCsvString() string {
	sliceNames := generateAllNamesSlice()

	return corecsv.RangeNamesWithValuesIndexesCsvString(
		sliceNames...)
}

func generateAllErrorTypeMap() map[string]Variation {
	length := len(errTypesMap)
	newMap := make(map[string]Variation, length)

	for variation, detail := range errTypesMap {
		newMap[detail.Name] = variation
	}

	return newMap
}

func generateAllNameWithValues() []string {
	slice := make([]string, maxValue+1)

	for i := 0; i <= maxValue; i++ {
		slice[i] = enumimpl.NameWithValue(Variation(i))
	}

	return slice
}
