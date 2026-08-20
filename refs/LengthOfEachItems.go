package refs

import (
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
)

func LengthOfEachItems(manyCollections [][]ref.Value) int {
	if len(manyCollections) == 0 {
		return constants.Zero
	}

	length := constants.Zero

	for _, collection := range manyCollections {
		length += len(collection)
	}

	return length
}
