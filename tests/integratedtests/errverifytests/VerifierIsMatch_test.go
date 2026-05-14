package errverifytests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/errorwrapper/errverify"
	"gitlab.com/evatix-go/errorwrapper/tests/testwrappers/errverifytestwrappers"
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
