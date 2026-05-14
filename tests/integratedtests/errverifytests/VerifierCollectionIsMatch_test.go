package errverifytests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/errorwrapper/errverify"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/errorwrapper/tests/testwrappers/errverifytestwrappers"
)

func Test_VerifierCollectionIsMatch(t *testing.T) {
	for caseIndex, testCase := range errverifytestwrappers.VerifyCollectionIsMatchTestCases {
		params := &errverify.VerifyCollectionParams{
			CaseIndex:    caseIndex,
			FuncName:     testCase.Verifier.FunctionName,
			TestCaseName: "VerifyCollectionIsMatchTestCases",
			ErrorCollection: errwrappers.NewUsingErrorWrappersPtr(
				false,
				testCase.InputErrorCollections,
			),
		}

		isSuccess := testCase.
			Verifier.IsMatch(params)

		Convey(testCase.Verifier.Header, t, func() {
			So(isSuccess, ShouldBeTrue)
		})
	}
}
