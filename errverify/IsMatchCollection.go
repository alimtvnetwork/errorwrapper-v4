package errverify

import (
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/corevalidator"
	"github.com/alimtvnetwork/core-v9/errcore"
)

func IsMatchCollection(
	params *VerifyCollectionParams,
	verifier *CollectionVerifier,
) bool {
	errCollection := params.ErrorCollection
	isErrEmpty := errCollection.IsEmpty()

	if verifier.ErrorLength == 0 &&
		isErrEmpty {
		return true
	}

	isPrint := verifier.IsPrintError
	isEmptyCompareFailed := verifier.IsCompareEmpty &&
		errCollection.HasError()
	isSuccess := !isEmptyCompareFailed

	if isEmptyCompareFailed {
		def := &errcore.ExpectationMessageDef{
			CaseIndex:      params.CaseIndex,
			FuncName:       params.FuncName,
			TestCaseName:   params.TestCaseName,
			When:           verifier.Header + " - ErrorCollection was expecting to be empty but there was error!",
			Expected:       "",
			IsNonWhiteSort: false,
		}

		def.PrintIf(
			isPrint,
			errCollection.String())
	}

	isLengthMismatch := verifier.ErrorLength > constants.InvalidValue &&
		verifier.ErrorLength != errCollection.Length()
	isSuccess = isSuccess && !isLengthMismatch

	if isLengthMismatch {
		def := &errcore.ExpectationMessageDef{
			CaseIndex:      params.CaseIndex,
			FuncName:       params.FuncName,
			TestCaseName:   params.TestCaseName,
			When:           verifier.Header + " - ErrorCollection Length Mismatch!",
			Expected:       verifier.ErrorLength,
			IsNonWhiteSort: false,
		}

		def.PrintIf(
			isPrint,
			errCollection.Length())
	}

	coreValidatorParams := corevalidator.Parameter{
		CaseIndex:                         params.CaseIndex,
		IsIgnoreCompareOnActualInputEmpty: false,
		IsAttachUserInputs:                true,
		IsCaseSensitive:                   true,
	}

	verifyError := verifier.ValidateErrCollectionUsingSliceValidator(
		params.IsWithRef(),
		&coreValidatorParams,
		errCollection,
		verifier.ExpectationLines.Items)

	verifyError.Log()

	return isSuccess && verifyError.IsSuccess()
}
