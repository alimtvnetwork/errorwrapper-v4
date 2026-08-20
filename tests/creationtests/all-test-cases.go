package creationtests

import (
	"errors"

	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core-v9/coretests"
	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errany"
	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errbool"
	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errbyte"
	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errfloat"
	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errfloat64"
	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errint"
	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errjson"
	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errstr"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
)

var testCases = []TestCaseWrapper{
	{
		BaseTestCase: coretests.BaseTestCase{
			Title: "ErrorWrapper creation tests",
			ArrangeInput: []errcoreinf.BaseErrorOrCollectionWrapper{
				errnew.Null.Error(errors.New("something is null")),
				errnew.Null.Simple(nil),
				errnew.Null.Simple(nilErr),
				errnew.Null.WithMessage("error is null", nilErr),
				errnew.Type.Default(errtype.AccountExpired),
				passwordCrudErr,
				errnew.Type.DirectRefs(errtype.IdMissing, "some ids missing", "id1", "id2"),
				errnew.Refs.TypeQuick(errtype.FailedProcess, "process 1", nilErr, 1, "id1", "id2"),
				errwrappers.
					Empty().
					AddWrapperPtr(
						errnew.NotFound.Invalid(
							"something is invalid", "s1", "s2")),
				errwrappers.
					Empty().
					NullSimple(nilErr),
				errany.New.Result.ErrorWrapper(passwordCrudErr).ErrorWrapper,
				errbool.New.Result.ErrorWrapper(passwordCrudErr).ErrorWrapper,
				errbyte.New.Result.ErrorWrapper(passwordCrudErr).ErrorWrapper,
				errfloat.New.Result.ErrorWrapper(passwordCrudErr).ErrorWrapper,
				errfloat64.New.Result.ErrorWrapper(passwordCrudErr).ErrorWrapper,
				errint.New.Result.ErrorWrapper(passwordCrudErr).ErrorWrapper,
				errjson.New.Result.ErrorWrapper(passwordCrudErr).ErrorWrapper,
				errstr.New.Result.ErrorWrapper(passwordCrudErr).ErrorWrapper,
			},
			ExpectedInput: []string{
				"Case Index : 0, Error Index: 0, Full string: [Error (Null - #3): Null reference. something is null.], Json : {\"IsDisplayableError\":true,\"CurrentError\":\"something is null\",\"ErrorType\":{\"Category\":\"Null\",\"Id\":3},\"StackTraces\":{},\"References\":null,\"HasError\":true}",
				"Case Index : 0, Error Index: 1, Full string: [Error (Null - #3): Null reference. Ref(s) {[Type (string): \"interface{}.(nil)\"]}], Json : {\"IsDisplayableError\":true,\"CurrentError\":\"\",\"ErrorType\":{\"Category\":\"Null\",\"Id\":3},\"StackTraces\":{},\"References\":[{\"VariableName\":\"Type\",\"ValueString\":\"interface{}.(nil)\"}],\"HasError\":true}",
				"Case Index : 0, Error Index: 2, Full string: [Error (Null - #3): Null reference. Ref(s) {[Type (string): \"interface{}.(nil)\"]}], Json : {\"IsDisplayableError\":true,\"CurrentError\":\"\",\"ErrorType\":{\"Category\":\"Null\",\"Id\":3},\"StackTraces\":{},\"References\":[{\"VariableName\":\"Type\",\"ValueString\":\"interface{}.(nil)\"}],\"HasError\":true}",
				"Case Index : 0, Error Index: 3, Full string: [Error (Null - #3): Null reference. Additional : error is null. Ref(s) {[Type (string): \"interface{}.(nil)\"]}], Json : {\"IsDisplayableError\":true,\"CurrentError\":\"error is null\",\"ErrorType\":{\"Category\":\"Null\",\"Id\":3},\"StackTraces\":{},\"References\":[{\"VariableName\":\"Type\",\"ValueString\":\"interface{}.(nil)\"}],\"HasError\":true}",
				"Case Index : 0, Error Index: 4, Full string: [Error (AccountExpired - #896): Account(s) expired!], Json : ",
				"Case Index : 0, Error Index: 5, Full string: [Error (PasswordCrud - #914): Password CRUD failed! some password.], Json : {\"IsDisplayableError\":true,\"CurrentError\":\"some password\",\"ErrorType\":{\"Category\":\"PasswordCrud\",\"Id\":914},\"StackTraces\":{},\"References\":null,\"HasError\":true}",
				"Case Index : 0, Error Index: 6, Full string: [Error (IdMissing - #569): Id not found or missing or undefined! Additional : some ids missing. Ref(s) {[References (string): \"\"id1\", \"id2\"\"]}], Json : {\"IsDisplayableError\":true,\"CurrentError\":\"some ids missing\",\"ErrorType\":{\"Category\":\"IdMissing\",\"Id\":569},\"StackTraces\":{},\"References\":[{\"VariableName\":\"References\",\"ValueString\":\"\\\"id1\\\", \\\"id2\\\"\"}],\"HasError\":true}",
				"Case Index : 0, Error Index: 7, Full string: [Error (FailedProcess - #119): Requested process or task is failed. Additional : \"process 1\", \"<nil>\", \"1\", \"id1\", \"id2\". Ref(s) {[: (string): \"\"process 1\", \"<nil>\", \"1\", \"id1\", \"id2\"\"]}], Json : {\"IsDisplayableError\":true,\"CurrentError\":\"\\\"process 1\\\", \\\"\\u003cnil\\u003e\\\", \\\"1\\\", \\\"id1\\\", \\\"id2\\\"\",\"ErrorType\":{\"Category\":\"FailedProcess\",\"Id\":119},\"StackTraces\":{},\"References\":[{\"VariableName\":\":\",\"ValueString\":\"\\\"process 1\\\", \\\"\\u003cnil\\u003e\\\", \\\"1\\\", \\\"id1\\\", \\\"id2\\\"\"}],\"HasError\":true}",
				"Case Index : 0, Error Index: 8, Full string: # Error Wrappers - Collection - Length[1]\n\n- [Error (Invalid - #271): Invalid! Additional : something is invalid. Ref(s) {[References ([]interface {}): \"[s1 s2]\"]}], Json : []",
				"Case Index : 0, Error Index: 9, Full string: # Error Wrappers - Collection - Length[1]\n\n- [Error (Null - #3): Null reference. Ref(s) {[Type (string): \"interface{}.(nil)\"]}], Json : []",
				"Case Index : 0, Error Index: 10, Full string: [Error (PasswordCrud - #914): Password CRUD failed! some password.], Json : {\"IsDisplayableError\":true,\"CurrentError\":\"some password\",\"ErrorType\":{\"Category\":\"PasswordCrud\",\"Id\":914},\"StackTraces\":{},\"References\":null,\"HasError\":true}",
				"Case Index : 0, Error Index: 11, Full string: [Error (PasswordCrud - #914): Password CRUD failed! some password.], Json : {\"IsDisplayableError\":true,\"CurrentError\":\"some password\",\"ErrorType\":{\"Category\":\"PasswordCrud\",\"Id\":914},\"StackTraces\":{},\"References\":null,\"HasError\":true}",
				"Case Index : 0, Error Index: 12, Full string: [Error (PasswordCrud - #914): Password CRUD failed! some password.], Json : {\"IsDisplayableError\":true,\"CurrentError\":\"some password\",\"ErrorType\":{\"Category\":\"PasswordCrud\",\"Id\":914},\"StackTraces\":{},\"References\":null,\"HasError\":true}",
				"Case Index : 0, Error Index: 13, Full string: [Error (PasswordCrud - #914): Password CRUD failed! some password.], Json : {\"IsDisplayableError\":true,\"CurrentError\":\"some password\",\"ErrorType\":{\"Category\":\"PasswordCrud\",\"Id\":914},\"StackTraces\":{},\"References\":null,\"HasError\":true}",
				"Case Index : 0, Error Index: 14, Full string: [Error (PasswordCrud - #914): Password CRUD failed! some password.], Json : {\"IsDisplayableError\":true,\"CurrentError\":\"some password\",\"ErrorType\":{\"Category\":\"PasswordCrud\",\"Id\":914},\"StackTraces\":{},\"References\":null,\"HasError\":true}",
				"Case Index : 0, Error Index: 15, Full string: [Error (PasswordCrud - #914): Password CRUD failed! some password.], Json : {\"IsDisplayableError\":true,\"CurrentError\":\"some password\",\"ErrorType\":{\"Category\":\"PasswordCrud\",\"Id\":914},\"StackTraces\":{},\"References\":null,\"HasError\":true}",
				"Case Index : 0, Error Index: 16, Full string: [Error (PasswordCrud - #914): Password CRUD failed! some password.], Json : {\"IsDisplayableError\":true,\"CurrentError\":\"some password\",\"ErrorType\":{\"Category\":\"PasswordCrud\",\"Id\":914},\"StackTraces\":{},\"References\":null,\"HasError\":true}",
				"Case Index : 0, Error Index: 17, Full string: [Error (PasswordCrud - #914): Password CRUD failed! some password.], Json : {\"IsDisplayableError\":true,\"CurrentError\":\"some password\",\"ErrorType\":{\"Category\":\"PasswordCrud\",\"Id\":914},\"StackTraces\":{},\"References\":null,\"HasError\":true}",
			},
		},
	},
}
