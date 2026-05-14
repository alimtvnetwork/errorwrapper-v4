package refs

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/ref"
)

func LengthOfEachItemsPtr(manyCollections *[]*[]*ref.Value) int {
	length := constants.Zero

	for _, collection := range *manyCollections {
		if collection == nil {
			continue
		}

		length += len(*collection)
	}

	return length
}
