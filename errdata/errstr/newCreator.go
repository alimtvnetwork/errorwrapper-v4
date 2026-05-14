package errstr

type newCreator struct {
	Hashset                    *newHashsetCreator
	Hashmap                    *newHashmapCreator
	HashsetsCollection         *newHashsetsCollectionCreator
	CharHashsetMap             *newCharHashsetMapCreator
	CharCollectionMap          *newCharCollectionMapCreator
	Collection                 *newCollectionCreator
	Result                     *newResultCreator
	Results                    *newResultsCreator
	ResultsWithErrorCollection *newResultsWithErrorCollectionCreator
	Result2                    *newResultTwoCreator
	ResultWithApplicable       *newResultApplicableCreator
	ResultWithApplicable2      *newResultApplicable2Creator
}
