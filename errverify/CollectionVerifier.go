package errverify

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/corevalidator"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

// CollectionVerifier helps to compare lines in error collection
type CollectionVerifier struct {
	Verifier
	ExpectationLines *corestr.SimpleSlice
	ErrorLength      int // -1 represents no checking or length
}

func (it *CollectionVerifier) IsMatch(
	params *VerifyCollectionParams,
) bool {
	return IsMatchCollection(
		params,
		it)
}

func (it *CollectionVerifier) IsMatchTestCase(
	caseIndex int,
	errCollection *errwrappers.Collection,
) bool {
	params := VerifyCollectionParams{
		CaseIndex:       caseIndex,
		ErrorCollection: errCollection,
	}

	return IsMatchCollection(
		&params,
		it)
}

func (it CollectionVerifier) ValidateErrUsingTextValidator(
	params *corevalidator.ValidatorParamsBase,
	isUseStrings bool,
	errorCollectionActual *errwrappers.Collection,
) *errorwrapper.Wrapper {
	return it.ExpectedValidateErrUsingTextValidator(
		params,
		isUseStrings,
		errorCollectionActual,
		it.ExpectingMessage)
}

func (it CollectionVerifier) ExpectedValidateErrUsingTextValidator(
	params *corevalidator.ValidatorParamsBase,
	isUseStrings bool,
	errorCollectionActual *errwrappers.Collection,
	expected string,
) *errorwrapper.Wrapper {
	if errorCollectionActual.IsEmpty() &&
		expected == constants.EmptyString {
		return nil
	}

	errCollectionLen := errorCollectionActual.Length()
	isLengthMismatch := it.ErrorLength > constants.InvalidValue &&
		it.ErrorLength != errCollectionLen

	if isLengthMismatch {
		def := &errcore.ExpectationMessageDef{
			CaseIndex:      params.CaseIndex,
			FuncName:       "",
			TestCaseName:   "",
			When:           it.Header + " - ErrorCollection Length Mismatch!",
			Expected:       it.ErrorLength,
			IsNonWhiteSort: false,
		}

		def.PrintIf(
			it.IsPrintError,
			errCollectionLen)

		return errnew.WasExpecting(
			errtype.LengthMismatch,
			"errorCollection length mismatch",
			it.ErrorLength,
			errCollectionLen)
	}

	validator := it.NewTextValidator(
		expected,
	)

	var validationErr error

	if isUseStrings {
		validationErr = validator.AllVerifyError(
			params,
			errorCollectionActual.StringsWithoutHeader()...)
	} else {
		validationErr = validator.AllVerifyError(
			params,
			errorCollectionActual.String())
	}

	if validationErr == nil {
		return nil
	}

	return errnew.Type.Error(
		errtype.ValidationFailed,
		validationErr)
}

func (it CollectionVerifier) ValidateErrUsingSliceValidator(
	params *corevalidator.ValidatorParamsBase,
	errorCollectionActual *errwrappers.Collection,
) *errorwrapper.Wrapper {
	return it.ExpectingLinesValidateErrUsingSliceValidator(
		params,
		errorCollectionActual,
		it.ExpectationLines.Items)
}

func (it CollectionVerifier) ExpectingLinesValidateErrUsingSliceValidator(
	params *corevalidator.ValidatorParamsBase,
	errorCollectionActual *errwrappers.Collection,
	expectedLines []string,
) *errorwrapper.Wrapper {
	if errorCollectionActual.IsEmpty() && len(expectedLines) == 0 {
		return nil
	}

	errCollectionLen := errorCollectionActual.Length()
	isLengthMismatch := it.ErrorLength > constants.InvalidValue &&
		it.ErrorLength != errCollectionLen

	var lengthErr *errorwrapper.Wrapper

	if isLengthMismatch {
		def := &errcore.ExpectationMessageDef{
			CaseIndex:      params.CaseIndex,
			FuncName:       "",
			TestCaseName:   "",
			When:           it.Header + " - ErrorCollection Length Mismatch!",
			Expected:       it.ErrorLength,
			IsNonWhiteSort: false,
		}

		def.PrintIf(
			it.IsPrintError,
			errCollectionLen)

		lengthErr = errnew.WasExpecting(
			errtype.LengthMismatch,
			"errorCollection length mismatch",
			it.ErrorLength,
			errCollectionLen)
	}

	actualLines := errorCollectionActual.StringsWithoutHeader()
	validator := it.NewSliceValidator(
		actualLines,
		expectedLines)

	err := validator.AllVerifyError(
		params)

	if err == nil {
		return lengthErr
	}

	if lengthErr.IsEmpty() {
		return errnew.Type.Error(
			errtype.ValidationFailed,
			err)
	}

	return errnew.Messages.Many(
		errtype.ValidationFailed,
		err.Error(),
		lengthErr.FullString())
}
