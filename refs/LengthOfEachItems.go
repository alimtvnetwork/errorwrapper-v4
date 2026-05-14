package refs

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/ref"
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
