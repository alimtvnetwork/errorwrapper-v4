package linuxservicecmdtests

import (
	"runtime"
	"testing"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/core-v9/coretests"
	"github.com/alimtvnetwork/errorwrapper-v4/errverify"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
	"github.com/alimtvnetwork/errorwrapper-v4/tests/testwrappers/linuxservicecmdtestwrappers"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ServicesInstructionApply_ErrorValidation(t *testing.T) {
	coretests.SkipOnWindows(t)
	if runtime.GOOS == "darwin" {
		t.Skip("requires linux systemctl/service")
	}

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
