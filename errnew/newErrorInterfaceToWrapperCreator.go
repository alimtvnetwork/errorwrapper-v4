package errnew

import (
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core-v9/errcore"
	"github.com/alimtvnetwork/core-v9/isany"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type newErrorInterfaceToWrapperCreator struct{}

func (it newErrorInterfaceToWrapperCreator) Default(
	variation errtype.Variation,
	errInf errcoreinf.BaseErrorOrCollectionWrapper,
) *errorwrapper.Wrapper {
	if errInf == nil || errInf.IsEmpty() {
		return nil
	}

	return errorwrapper.InterfaceToErrorWrapper(
		variation,
		errInf)
}

// AnyType
//
//  tries may ways to get to the error wrapper.
//  on fail returns parsed error
//
// Steps:
//  - *errorwrapper.Wrapper
//  - errorwrapper.Wrapper
//  - errorwrapper.ErrWrapper
//  - *errcore.RawErrCollection
//  - errcore.RawErrCollection
//  - errcoreinf.BaseRawErrCollectionDefiner
//  - errcoreinf.BasicErrWrapper
//  - errcoreinf.BaseErrorWrapperCollectionDefiner
//  - errcoreinf.BaseErrorOrCollectionWrapper
//  - corejson.Result
//  - *corejson.Result
//  - []byte
//  - error
//  - string
func (it newErrorInterfaceToWrapperCreator) AnyType(
	variation errtype.Variation,
	errInf interface{},
) (convertedErrWrapper *errorwrapper.Wrapper, parsedErrWp *errorwrapper.Wrapper) {
	if isany.Null(errInf) {
		return nil, nil
	}

	switch casted := errInf.(type) {
	case *errorwrapper.Wrapper:
		return casted, nil
	case errorwrapper.Wrapper:
		return casted.Ptr(), nil
	case errorwrapper.ErrWrapper:
		return casted.AsErrorWrapper(), nil

	case *errcore.RawErrCollection:
		return it.ActualRawErrCollection(
			variation,
			casted), nil
	case errcore.RawErrCollection:
		return it.ActualRawErrCollection(
			variation,
			casted.ToRawErrCollection()), nil

	case errcoreinf.BaseRawErrCollectionDefiner:
		return it.RawErrCollection(
			variation,
			casted), nil

	case errcoreinf.BasicErrWrapper:
		return it.BasicErr(casted), nil
	case errcoreinf.BaseErrorWrapperCollectionDefiner:
		return it.BasicErr(
			casted.GetAsBasicWrapperUsingTyper(
				variation.AsBasicErrorTyper())), nil
	case errcoreinf.BaseErrorOrCollectionWrapper:
		return it.Default(variation, casted), nil
	case corejson.Result:
		return DeserializeTo.JsonResultToWrapperUsingStackSkip(
			defaultSkipInternal,
			&casted)
	case *corejson.Result:
		return DeserializeTo.JsonResultToWrapperUsingStackSkip(
			defaultSkipInternal,
			casted)
	case []byte:
		return DeserializeTo.BytesToWrapperUsingStackSkip(
			defaultSkipInternal,
			casted)
	case error:
		convErrWp, parsedFailed := DeserializeTo.JsonErrToWrapper(
			casted)

		if parsedFailed.HasAnyIssues() {
			// create general error
			return Type.Error(variation, casted), nil
		}

		return convErrWp, nil
	case string:
		convErrWp, parsedFailed := DeserializeTo.JsonStringToWrapper(
			false,
			casted)

		if parsedFailed.HasAnyIssues() {
			// create general error
			return Message.Default(variation, casted), nil
		}

		return convErrWp, nil
	}

	return nil, FromTo.Message(errtype.CastingFailed,
		"cannot cast or make to error collection",
		errInf, &errorwrapper.Wrapper{})
}

func (it newErrorInterfaceToWrapperCreator) ActualRawErrCollection(
	variation errtype.Variation,
	rawErrCollection *errcore.RawErrCollection,
) *errorwrapper.Wrapper {
	if rawErrCollection == nil || rawErrCollection.IsEmpty() {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		variation,
		rawErrCollection.CompiledError().Error())
}

// NoType
//
//  tries to cast to error wrapper first if not possible then creates new one.
//  using no type
func (it newErrorInterfaceToWrapperCreator) NoType(
	errInf errcoreinf.BaseErrorOrCollectionWrapper,
) *errorwrapper.Wrapper {
	if errInf == nil || errInf.IsEmpty() {
		return nil
	}

	return errorwrapper.InterfaceToErrorWrapper(
		errtype.Unknown,
		errInf)
}

func (it newErrorInterfaceToWrapperCreator) BasicErr(
	basicErr errcoreinf.BasicErrWrapper,
) *errorwrapper.Wrapper {
	if basicErr == nil || basicErr.IsEmpty() {
		return nil
	}

	return errorwrapper.NewUsingBasicErr(basicErr)
}

func (it newErrorInterfaceToWrapperCreator) RawErrCollection(
	variation errtype.Variation,
	rawErrCollection errcoreinf.BaseRawErrCollectionDefiner,
) *errorwrapper.Wrapper {
	if rawErrCollection == nil || rawErrCollection.IsEmpty() {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		variation,
		rawErrCollection.CompiledStackTracesString())
}

func (it newErrorInterfaceToWrapperCreator) ErrorWrapperCollectionDefiner(
	variation errtype.Variation,
	collection errcoreinf.BaseErrorWrapperCollectionDefiner,
) *errorwrapper.Wrapper {
	if collection == nil || collection.IsEmpty() {
		return nil
	}

	basicErr := collection.GetAsBasicWrapperUsingTyper(
		variation.AsBasicErrorTyper())

	return errorwrapper.NewUsingBasicErr(basicErr)
}

func (it newErrorInterfaceToWrapperCreator) ErrorWrapperCollectionsDefiner(
	variation errtype.Variation,
	collections ...errcoreinf.BaseErrorWrapperCollectionDefiner,
) *errorwrapper.Wrapper {
	if collections == nil || len(collections) == 0 {
		return nil
	}

	rawErrCollection := errcore.RawErrCollection{}

	for _, collection := range collections {
		basicErr := collection.GetAsBasicWrapper()

		if basicErr == nil || basicErr.IsEmpty() {
			continue
		}

		rawErrCollection.AddError(
			basicErr.CompiledErrorWithStackTraces())
	}

	if rawErrCollection.IsEmpty() {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		variation,
		rawErrCollection.CompiledStackTracesString())
}
