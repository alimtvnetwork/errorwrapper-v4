package errstr

var (
	New = &newCreator{
		Hashset:                    &newHashsetCreator{},
		Hashmap:                    &newHashmapCreator{},
		HashsetsCollection:         &newHashsetsCollectionCreator{},
		CharHashsetMap:             &newCharHashsetMapCreator{},
		Collection:                 &newCollectionCreator{},
		Result:                     &newResultCreator{},
		Results:                    &newResultsCreator{},
		ResultsWithErrorCollection: &newResultsWithErrorCollectionCreator{},
		Result2:                    &newResultTwoCreator{},
		ResultWithApplicable:       &newResultApplicableCreator{},
		ResultWithApplicable2:      &newResultApplicable2Creator{},
	}
	Empty = &emptyCreator{}
)
