package refs

import (
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
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
