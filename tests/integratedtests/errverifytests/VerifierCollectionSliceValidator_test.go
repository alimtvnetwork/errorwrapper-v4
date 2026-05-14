package errverifytests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/alimtvnetwork/core-v9/corevalidator"
	"github.com/alimtvnetwork/errorwrapper-v3/errverify"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	"github.com/alimtvnetwork/errorwrapper-v3/tests/testwrappers/errverifytestwrappers"
)

func Test_VerifierCollectionSliceValidator(t *testing.T) {
	for caseIndex, testCase := range errverifytestwrappers.VerifyCollectionSliceValidatorTestCases {
		params := &errverify.VerifyCollectionParams{
			CaseIndex:    caseIndex,
			FuncName:     "Test_VerifierCollectionSliceValidator",
			TestCaseName: "VerifyCollectionSliceValidatorTestCases",
			ErrorCollection: errwrappers.NewUsingErrorWrappersPtr(
				false,
				testCase.InputErrorCollections,
			),
		}

		validatorParamsBase := &corevalidator.Parameter{
			CaseIndex:                         caseIndex,
			IsIgnoreCompareOnActualInputEmpty: false,
			IsAttachUserInputs:                true,
			IsCaseSensitive:                   true,
		}

		verifyErrWp := testCase.
			Verifier.ValidateErrCollectionUsingSliceValidator(
			params.IsWithRef(),
			validatorParamsBase,
			params.ErrorCollection,
			(*testCase.Verifier.ExpectationLines))

		sliceValidator := testCase.Verifier.NewSliceValidator(
			params.ErrorCollection.StringsWithoutHeader(),
			(*testCase.Verifier.ExpectationLines))

		// Redundant but verifies the same thing words in the collection
		verifySliceErr := params.ErrorCollection.ValidationErrUsingSliceValidator(
			sliceValidator,
			validatorParamsBase)

		convey.Convey(testCase.Verifier.Header, t, func() {
			verifyErrWp.Log()
			verifySliceErr.Log()

			convey.So(verifyErrWp.IsSuccess(), convey.ShouldBeTrue)
			convey.So(verifySliceErr.IsSuccess(), convey.ShouldBeTrue)
		})
	}
}
