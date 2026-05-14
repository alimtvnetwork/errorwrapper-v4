package errverify

import "github.com/alimtvnetwork/errorwrapper-v3/errwrappers"

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
