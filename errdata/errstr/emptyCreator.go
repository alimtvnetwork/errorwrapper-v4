package errstr

import "gitlab.com/evatix-go/errorwrapper"

type emptyCreator struct{}

func (it *emptyCreator) Result() *Result {
	return &Result{}
}

func (it *emptyCreator) Results() *Results {
	return &Results{}
}

func (it *emptyCreator) Result2() *Result2 {
	return &Result2{}
}

func (it *emptyCreator) ResultsWithErrorCollection() *ResultsWithErrorCollection {
	return &ResultsWithErrorCollection{}
}

func (it *emptyCreator) ResultWithApplicable() *ResultWithApplicable {
	return &ResultWithApplicable{}
}

func (it *emptyCreator) ResultWithApplicable2() *ResultWithApplicable2 {
	return &ResultWithApplicable2{}
}

func (it *emptyCreator) Hashset() *Hashset {
	return &Hashset{}
}

func (it *emptyCreator) Hashmap() *Hashmap {
	return &Hashmap{}
}

func (it *emptyCreator) LinkedList() *LinkedList {
	return &LinkedList{}
}

func (it *emptyCreator) HashsetsCollection() *HashsetsCollection {
	return &HashsetsCollection{}
}

func (it *emptyCreator) CharCollectionMap() *CharCollectionMap {
	return &CharCollectionMap{}
}

func (it *emptyCreator) CharHashsetMap() *CharHashsetMap {
	return &CharHashsetMap{}
}

func (it *emptyCreator) SimpleStringOnce() *SimpleStringOnce {
	return &SimpleStringOnce{}
}

func (it *emptyCreator) ResultWithError(errorWrapper *errorwrapper.Wrapper) *Result {
	return &Result{
		ErrorWrapper: errorWrapper,
	}
}

func (it *emptyCreator) ResultsWithError(errorWrapper *errorwrapper.Wrapper) *Results {
	return &Results{
		ErrorWrapper: errorWrapper,
	}
}

func (it *emptyCreator) ResultWithValue(value string) *Result {
	return &Result{
		Value: value,
	}
}

func (it *emptyCreator) ResultsWithValue(values []string) *Results {
	return &Results{
		Values: values,
	}
}
