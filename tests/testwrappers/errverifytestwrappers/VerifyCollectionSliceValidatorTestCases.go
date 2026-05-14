package errverifytestwrappers

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/enums/stringcompareas"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errverify"
)

var VerifyCollectionSliceValidatorTestCases = []VerifyErrorCollectionTestWrapper{
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
				Items: []string{
					"[Error (InvalidOption - #470): Selected option is invalid!]",
					"[Error (NotSupportInWindows - #93): Current request is not supported in Windows Operating system.]",
					"[Error (NotSupportedOption - #107): None of the option is supported.]",
				},
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
				Items: []string{
					"[Error (InvalidOption - #470): Selected option is invalid!]",
					"[Error (NotSupportInWindows - #93): Current request is not supported in Windows Operating system.]",
					"[Error (NotSupportedOption - #107): None of the option is supported.]",
				},
			},
			ErrorLength: constants.InvalidValue,
		},
	},
}
