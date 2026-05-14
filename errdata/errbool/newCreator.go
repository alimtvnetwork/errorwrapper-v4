package errbool

type newCreator struct {
	Result                     *newResultCreator
	Results                    *newResultsCreator
	ResultsWithErrorCollection *newResultsWithErrorCollectionCreator
	Result2                    *newResultTwoCreator
	ResultWithApplicable       *newResultApplicableCreator
	ResultWithApplicable2      *newResultApplicable2Creator
}
