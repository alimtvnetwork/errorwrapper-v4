package errverify

import "gitlab.com/evatix-go/errorwrapper/errwrappers"

type VerifyCollectionParams struct {
	CaseIndex                 int
	FuncName                  string
	TestCaseName              string
	IsCompareWithoutReference bool
	ErrorCollection           *errwrappers.Collection
}

func (it VerifyCollectionParams) IsWithRef() bool {
	return !it.IsCompareWithoutReference
}
