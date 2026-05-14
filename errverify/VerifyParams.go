package errverify

import (
	"gitlab.com/evatix-go/errorwrapper"
)

type VerifyParams struct {
	CaseIndex                 int
	FuncName                  string
	TestCaseName              string
	IsCompareWithoutReference bool
	ErrorWrapper              *errorwrapper.Wrapper
}

func (it VerifyParams) IsWithRef() bool {
	return !it.IsCompareWithoutReference
}
