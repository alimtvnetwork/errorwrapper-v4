package errverify

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/core/errcore"
)

func IsMatch(
	params *VerifyParams,
	isVerifyErrorMsg bool,
	verifier *Verifier,
) bool {
	errWp := params.ErrorWrapper
	isEmptyMsg := verifier.IsEmptyExpectingMessage()
	isErrEmpty := errWp.IsEmpty()

	if isEmptyMsg && isErrEmpty {
		return true
	}

	isPrint := verifier.IsPrintError
	isEmptyCompareFailed := verifier.IsCompareEmpty &&
		errWp.HasError()
	actualCompareValue := errWp.FullOrErrorMessage(isVerifyErrorMsg, params.IsWithRef())
	isSuccess := !isEmptyCompareFailed

	if isEmptyCompareFailed {
		def := &errcore.ExpectationMessageDef{
			CaseIndex:      params.CaseIndex,
			FuncName:       params.FuncName,
			TestCaseName:   params.TestCaseName,
			When:           verifier.Header + " - ErrorWrapper was expecting to be empty but there was error!",
			Expected:       "",
			IsNonWhiteSort: false,
		}

		def.PrintIf(
			isPrint,
			actualCompareValue)
	}

	isFailedEmptyChecking := isEmptyMsg ||
		isErrEmpty

	isSuccess = isSuccess && !isFailedEmptyChecking

	if isFailedEmptyChecking {
		def := &errcore.ExpectationMessageDef{
			CaseIndex:      params.CaseIndex,
			FuncName:       params.FuncName,
			TestCaseName:   params.TestCaseName,
			When:           verifier.Header + " - ErrorWrapper doesn't match in terms of given data!\n",
			Expected:       verifier.ExpectingMessage,
			IsNonWhiteSort: false,
		}

		def.PrintIf(
			isPrint,
			constants.EmptyString)
	}

	isFullMessageEqual := true

	if verifier.HasExpectingMessage() {
		expectationDef := &errcore.ExpectationMessageDef{
			CaseIndex:      params.CaseIndex,
			FuncName:       params.FuncName,
			TestCaseName:   params.TestCaseName,
			When:           verifier.Header,
			Expected:       verifier.ExpectingMessage,
			IsNonWhiteSort: true,
		}

		isFullMessageEqual = coretests.IsStrMsgNonWhiteSortedEqual(
			isPrint,
			actualCompareValue,
			expectationDef)
	}

	return isSuccess && isFullMessageEqual
}
