package linuxservicecmdtests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/errorwrapper/errverify"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/errorwrapper/tests/testwrappers/linuxservicecmdtestwrappers"
)

func Test_ServicesInstructionApply_ErrorValidation(t *testing.T) {
	coretests.SkipOnWindows(t)

	for caseIndex, testCase := range linuxservicecmdtestwrappers.ServicesErrorValidationTestCases {
		// Arrange
		errCollection := errwrappers.NewCap4()
		errVerifyCollection := errverify.CollectionVerifier{
			Verifier: errverify.Verifier{
				Header:       testCase.Header,
				FunctionName: "Test_ServicesInstructionApply",
				IsPrintError: true,
			},
			ExpectationLines: corestr.New.SimpleSlice.Direct(false, testCase.ErrorValidation),
			ErrorLength:      constants.InvalidIndex,
		}

		// Act
		testCase.Apply(errCollection)
		params := &errverify.VerifyCollectionParams{
			CaseIndex:       caseIndex,
			FuncName:        "Test_ServicesInstructionApply",
			TestCaseName:    "ServicesTestCases",
			ErrorCollection: errCollection,
		}

		// Assert
		Convey(testCase.Header, t, func() {
			isSuccess := errVerifyCollection.
				IsMatch(params)

			So(isSuccess, ShouldBeTrue)
		})
	}
}
