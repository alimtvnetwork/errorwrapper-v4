package errjson

import (
	"errors"
	"io/ioutil"

	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

type newResultCreator struct{}

func (it *newResultCreator) Empty() *Result {
	return &Result{}
}

func (it *newResultCreator) Item(
	item *corejson.Result,
) *Result {
	return &Result{
		Result: item,
	}
}

func (it *newResultCreator) Error(
	errType errtype.Variation,
	err error,
) *Result {
	return &Result{
		ErrorWrapper: errnew.Type.ErrorUsingStackSkip(
			codestack.Skip1,
			errType,
			err),
	}
}

func (it *newResultCreator) Any(
	anyItem interface{},
) *Result {
	jsonResult := corejson.NewPtr(anyItem)

	if jsonResult.HasError() {
		return &Result{
			Result: jsonResult,
			ErrorWrapper: errnew.Type.ErrorUsingStackSkip(
				codestack.Skip1,
				errtype.JsonSyntaxIssue,
				jsonResult.MeaningfulError()),
		}
	}

	return &Result{
		Result: jsonResult,
	}
}

func (it *newResultCreator) Bytes(
	rawBytes []byte,
) *Result {
	if rawBytes == nil {
		jsonResult := corejson.Empty.ResultPtr()

		return &Result{
			Result: jsonResult,
		}
	}

	jsonResult := corejson.NewPtr(rawBytes)

	if jsonResult.IsEmptyError() {
		return &Result{
			Result: jsonResult,
		}
	}

	return &Result{
		Result:       jsonResult,
		ErrorWrapper: errnew.Error.Type(errtype.Marshalling, jsonResult.MeaningfulError()),
	}
}

func (it *newResultCreator) Marshal(anyObject interface{}) *Result {
	rs := corejson.NewPtr(anyObject)

	if rs.HasError() {
		return &Result{
			Result: rs,
			ErrorWrapper: errnew.Type.ErrorUsingStackSkip(
				codestack.Skip1,
				errtype.JsonSyntaxIssue,
				rs.MeaningfulError()),
		}
	}

	return &Result{
		Result:       rs,
		ErrorWrapper: nil,
	}
}

func (it *newResultCreator) BytesPtr(
	rawBytes *[]byte,
) *Result {
	if rawBytes == nil || len(*rawBytes) == 0 {
		jsonResult := corejson.Empty.ResultPtr()

		return &Result{
			Result: jsonResult,
		}
	}

	jsonResult := corejson.NewPtr(rawBytes)

	if jsonResult.IsEmptyError() {
		return &Result{
			Result: jsonResult,
		}
	}

	return &Result{
		Result:       jsonResult,
		ErrorWrapper: errnew.Error.Type(errtype.Marshalling, jsonResult.MeaningfulError()),
	}
}

func (it *newResultCreator) BytesWithError(
	rawBytes []byte,
	errWp *errorwrapper.Wrapper,
) *Result {
	if rawBytes == nil {
		jsonResult := corejson.
			Empty.
			ResultPtrWithErr(
				coredynamic.TypeName(Result{}),
				errWp.Error())

		return &Result{
			Result:       jsonResult,
			ErrorWrapper: errWp,
		}
	}

	jsonResult := corejson.NewPtr(rawBytes)

	return &Result{
		Result:       jsonResult,
		ErrorWrapper: errWp,
	}
}

func (it *newResultCreator) TypeName() string {
	return resultType
}

func (it *newResultCreator) File(
	filePath string,
) *Result {
	if filePath == "" {
		err := errors.New("filePath given as empty")

		return &Result{
			Result:       corejson.Empty.ResultPtrWithErr(it.TypeName(), err),
			ErrorWrapper: errnew.Type.Error(errtype.EmptyFilePath, err),
		}
	}

	allBytes, err := ioutil.ReadFile(filePath)
	typeName := "file:" + filePath

	if err == nil {
		return &Result{
			Result: corejson.NewResult.Ptr(allBytes, nil, typeName),
		}
	}

	// has error
	return &Result{
		Result:       corejson.NewResult.Ptr(allBytes, err, typeName),
		ErrorWrapper: errnew.Path.Error(errtype.FileRead, err, filePath),
	}
}

func (it *newResultCreator) FileLock(filePath string) *Result {
	writerLock.Lock()
	defer writerLock.Unlock()

	return it.File(filePath)
}

func (it *newResultCreator) ErrorWrapper(
	errorWrapper *errorwrapper.Wrapper,
) *Result {
	return &Result{
		ErrorWrapper: errorWrapper,
	}
}

func (it *newResultCreator) Create(
	result *corejson.Result,
	errorWrapper *errorwrapper.Wrapper,
) *Result {
	return &Result{
		Result:       result,
		ErrorWrapper: errorWrapper,
	}
}

func (it *newResultCreator) ValueOnly(
	result *corejson.Result,
) *Result {
	return &Result{
		Result: result,
	}
}

func (it *newResultCreator) Result(
	result *corejson.Result,
) *Result {
	return &Result{
		Result: result,
	}
}
