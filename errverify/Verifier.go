package errverify

import (
	"strings"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/corevalidator"
	"github.com/alimtvnetwork/core-v9/enums/stringcompareas"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

type Verifier struct {
	Header, FunctionName string
	ExpectingMessage     string // It has no use in CollectionVerifier
	corevalidator.ValidatorCoreCondition
	VerifyAs                 stringcompareas.Variant
	IsCompareEmpty           bool // true means, seek for empty only
	IsVerifyErrorMessageOnly bool // don't verify full message
	IsPrintError             bool
}

func (it *Verifier) MethodName() string {
	return it.VerifyAs.Name()
}

func (it *Verifier) IsEmptyExpectingMessage() bool {
	return it == nil || it.ExpectingMessage == ""
}

func (it *Verifier) HasExpectingMessage() bool {
	return it != nil && it.ExpectingMessage != ""
}

func (it Verifier) NewTextValidator(
	searchTerm string,
) *corevalidator.TextValidator {
	return &corevalidator.TextValidator{
		Search:                 searchTerm,
		SearchAs:               it.VerifyAs,
		ValidatorCoreCondition: it.ValidatorCoreCondition,
	}
}

func (it Verifier) NewSliceValidator(
	inputLines, comparingLines []string,
) *corevalidator.SliceValidator {
	return &corevalidator.SliceValidator{
		ValidatorCoreCondition: it.ValidatorCoreCondition,
		ActualLines:            inputLines,
		ExpectedLines:          comparingLines,
		CompareAs:              it.VerifyAs,
	}
}

func (it Verifier) ValidateErrUsingTextValidator(
	isCompareIncludingReference bool,
	params *corevalidator.ValidatorParamsBase,
	actualErrorWrapper *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	if actualErrorWrapper.IsEmpty() &&
		it.IsEmptyExpectingMessage() {
		return nil
	}

	validator := it.NewTextValidator(
		it.ExpectingMessage)

	actual := actualErrorWrapper.FullOrErrorMessage(
		it.IsVerifyErrorMessageOnly,
		isCompareIncludingReference)

	err := validator.AllVerifyError(
		params,
		actual)

	if err == nil {
		return nil
	}

	return errnew.Type.Error(
		errtype.ValidationFailed,
		err)
}

func (it Verifier) ValidateErrUsingSliceValidator(
	params *corevalidator.ValidatorParamsBase,
	errorWrapper *errorwrapper.Wrapper,
	expectedLines []string,
) *errorwrapper.Wrapper {
	if errorWrapper.IsEmpty() && len(expectedLines) == 0 {
		return nil
	}

	var actualLines []string
	if it.IsVerifyErrorMessageOnly {
		actualLines = strings.Split(
			errorWrapper.ErrorString(),
			constants.NewLineUnix)
	} else {
		actualLines = errorWrapper.
			FullStringSplitByNewLine()
	}

	return it.ValidateActualLinesSliceValidator(
		params,
		actualLines,
		expectedLines)
}

func (it Verifier) ValidateErrCollectionUsingSliceValidator(
	isIncludeReferences bool,
	params *corevalidator.ValidatorParamsBase,
	errCollection *errwrappers.Collection,
	expectedLines []string,
) *errorwrapper.Wrapper {
	if errCollection.IsEmpty() && len(expectedLines) == 0 {
		return nil
	}

	actualLines := errCollection.LinesIf(isIncludeReferences)

	return it.ValidateActualLinesSliceValidator(
		params,
		actualLines,
		expectedLines)
}

func (it Verifier) ValidateActualLinesSliceValidator(
	params *corevalidator.ValidatorParamsBase,
	actualLines []string,
	expectedLines []string,
) *errorwrapper.Wrapper {
	if len(actualLines) == 0 && len(expectedLines) == 0 {
		return nil
	}

	validator := it.NewSliceValidator(
		actualLines,
		expectedLines)

	err := validator.AllVerifyError(
		params)

	if err == nil {
		return nil
	}

	return errnew.Type.Error(
		errtype.ValidationFailed,
		err)
}

func (it *Verifier) IsMatch(
	params *VerifyParams,
) bool {
	return IsMatch(
		params,
		it.IsVerifyErrorMessageOnly,
		it)
}
