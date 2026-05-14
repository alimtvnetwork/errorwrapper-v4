package errverifytests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/alimtvnetwork/errorwrapper-v3/errverify"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	"github.com/alimtvnetwork/errorwrapper-v3/tests/testwrappers/errverifytestwrappers"
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
