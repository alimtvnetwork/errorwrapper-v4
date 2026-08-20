package errverifytestwrappers

import (
	"github.com/alimtvnetwork/core-v9/enums/stringcompareas"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errverify"
)

var VerifyValidatorTestCases = []VerifyTestWrapper{
	{
		Verifier: errverify.Verifier{
			Header:       "Path location error wrapper validation",
			FunctionName: VerifierIsMatchFunc,
			ExpectingMessage: "[Error (PathMismatch - #298): " +
				"Path mismatch error, expectation didn't meet! " +
				"Additional : my demo message. Ref(s) " +
				"{[Path (string): \"location 1\"]}]",
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
