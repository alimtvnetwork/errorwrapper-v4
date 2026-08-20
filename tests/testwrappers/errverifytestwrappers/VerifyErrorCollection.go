package errverifytestwrappers

import (
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/core-v9/enums/stringcompareas"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errverify"
)

var VerifyCollectionIsMatchTestCases = []VerifyErrorCollectionTestWrapper{
	{
		InputErrorCollections: []*errorwrapper.Wrapper{
			errnew.Type.Default(errtype.InvalidOption),
			errnew.Type.Default(errtype.NotSupportInWindows),
			errnew.Type.Default(errtype.NotSupportedOption),
		},
		Verifier: errverify.CollectionVerifier{
			Verifier: errverify.Verifier{
				Header:       "Collection has 3 elements with 3 wrappers, verify length and it's content",
				FunctionName: VerifyCollectionIsMatch,
				VerifyAs:     stringcompareas.Equal,
				IsPrintError: true,
			},
			ExpectationLines: &corestr.SimpleSlice{
				"[Error (InvalidOption - #469): Selected option is invalid!]",
				"[Error (NotSupportInWindows - #93): Current request is not supported in Windows Operating system.]",
				"[Error (NotSupportedOption - #107): None of the option is supported.]",
			},
			ErrorLength: 3,
		},
	},
	{
		InputErrorCollections: []*errorwrapper.Wrapper{
			errnew.Type.Default(errtype.InvalidOption),
			errnew.Type.Default(errtype.NotSupportInWindows),
			errnew.Type.Default(errtype.NotSupportedOption),
		},
		Verifier: errverify.CollectionVerifier{
			Verifier: errverify.Verifier{
				Header:       "Collection has 3 elements with 3 wrappers, verify it's content",
				FunctionName: VerifyCollectionIsMatch,
				VerifyAs:     stringcompareas.Equal,
				IsPrintError: true,
			},
			ExpectationLines: &corestr.SimpleSlice{
				"[Error (InvalidOption - #469): Selected option is invalid!]",
				"[Error (NotSupportInWindows - #93): Current request is not supported in Windows Operating system.]",
				"[Error (NotSupportedOption - #107): None of the option is supported.]",
			},
			ErrorLength: constants.InvalidValue,
		},
	},
}
