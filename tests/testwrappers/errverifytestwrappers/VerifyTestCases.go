package errverifytestwrappers

import (
	"gitlab.com/evatix-go/core/enums/stringcompareas"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errverify"
)

var VerifyValidatorTestCases = []VerifyTestWrapper{
	{
		Verifier: errverify.Verifier{
			Header:       "Path location error wrapper validation",
			FunctionName: VerifierIsMatchFunc,
			ExpectingMessage: "[Error (PathMismatch - #299): " +
				"Path mismatch error, expectation didn't meet! " +
				"Additional : my demo message. Ref(s) " +
				"{[File Path (string): \"location 1\"]}]",
			VerifyAs:                 stringcompareas.Equal,
			IsCompareEmpty:           false,
			IsVerifyErrorMessageOnly: false,
			IsPrintError:             true,
		},
		ErrorWrapper: errnew.Path.Messages(
			errtype.PathMismatch,
			"location 1",
			"my demo message"),
	},
	{
		Verifier: errverify.Verifier{
			Header:                   "Path location error validation",
			FunctionName:             VerifierIsMatchFunc,
			ExpectingMessage:         "my demo message",
			VerifyAs:                 stringcompareas.Equal,
			IsCompareEmpty:           false,
			IsVerifyErrorMessageOnly: true,
			IsPrintError:             true,
		},
		ErrorWrapper: errnew.Path.Messages(
			errtype.PathMismatch,
			"location 1",
			"my demo message"),
	},
	{
		Verifier: errverify.Verifier{
			Header:                   "Path location error validation",
			FunctionName:             VerifierIsMatchFunc,
			ExpectingMessage:         "my demo message",
			VerifyAs:                 stringcompareas.Equal,
			IsCompareEmpty:           false,
			IsVerifyErrorMessageOnly: true,
			IsPrintError:             true,
		},
		ErrorWrapper: errnew.Path.Messages(
			errtype.PathMismatch,
			"location 1",
			"my demo message"),
	},
	{
		Verifier: errverify.Verifier{
			Header:                   "Path location error validation",
			FunctionName:             VerifierIsMatchFunc,
			ExpectingMessage:         "my demo message",
			VerifyAs:                 stringcompareas.Equal,
			IsCompareEmpty:           true, // it is not working which make sense because we are using validator in test
			IsVerifyErrorMessageOnly: true,
			IsPrintError:             true,
		},
		ErrorWrapper: errnew.Path.Messages(
			errtype.PathMismatch,
			"location 1",
			"my demo message"),
	},
}
