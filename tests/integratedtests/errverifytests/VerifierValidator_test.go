package errverifytests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/corevalidator"
	"gitlab.com/evatix-go/errorwrapper/tests/testwrappers/errverifytestwrappers"
)

func Test_VerifierValidator(t *testing.T) {
	for caseIndex, testCase := range errverifytestwrappers.VerifyValidatorTestCases {
		params := &corevalidator.ValidatorParamsBase{
			CaseIndex:                         caseIndex, // fixed
			IsIgnoreCompareOnActualInputEmpty: false,
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
