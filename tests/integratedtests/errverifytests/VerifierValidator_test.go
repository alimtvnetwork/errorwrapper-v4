package errverifytests

import (
	"testing"

	"github.com/alimtvnetwork/core-v9/corevalidator"
	"github.com/alimtvnetwork/errorwrapper-v4/tests/testwrappers/errverifytestwrappers"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_VerifierValidator(t *testing.T) {
	for caseIndex, testCase := range errverifytestwrappers.VerifyValidatorTestCases {
		params := &corevalidator.Parameter{
			CaseIndex:                  caseIndex, // fixed
			IsSkipCompareOnActualEmpty: false,
			IsAttachUserInputs:         true,
			IsCaseSensitive:            true,
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
