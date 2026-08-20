package errverifytests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v4/errverify"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
	"github.com/alimtvnetwork/errorwrapper-v4/tests/testwrappers/errverifytestwrappers"
	. "github.com/smartystreets/goconvey/convey"
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
