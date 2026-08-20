package errverifytests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v4/errverify"
	"github.com/alimtvnetwork/errorwrapper-v4/tests/testwrappers/errverifytestwrappers"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_VerifierIsMatch(t *testing.T) {
	for caseIndex, testCase := range errverifytestwrappers.VerifyIsMatchTestCases {
		params := &errverify.VerifyParams{
			CaseIndex:    caseIndex,
			FuncName:     "Test_VerifierIsMatch",
			TestCaseName: "VerifyIsMatchTestCases",
			ErrorWrapper: testCase.ErrorWrapper,
		}

		isSuccess := testCase.
			Verifier.IsMatch(params)

		Convey(testCase.Verifier.Header, t, func() {
			So(isSuccess, ShouldBeTrue)
		})
	}
}
