package errverify

import (
	"github.com/alimtvnetwork/errorwrapper-v4"
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
