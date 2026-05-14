package errverifytests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/alimtvnetwork/core-v9/corevalidator"
	"github.com/alimtvnetwork/errorwrapper-v3/tests/testwrappers/errverifytestwrappers"
)

func Test_VerifierValidator(t *testing.T) {
	for caseIndex, testCase := range errverifytestwrappers.VerifyValidatorTestCases {
		params := &corevalidator.Parameter{
			CaseIndex:                         caseIndex, // fixed
			IsSkipCompareOnActualEmpty: false,
			IsAttachUserInputs:                true,
			IsCaseSensitive:                   true,
		}

		validationErr := testCase.
			Verifier.
			ValidateErrUsingTextValidator(
				true,
				params,
				testCase.ErrorWrapper)

		Convey(testCase.Verifier.Header, t, func() {
			validationErr.Log()
			So(validationErr.IsSuccess(), ShouldBeTrue)
		})
	}
}
