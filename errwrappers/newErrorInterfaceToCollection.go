package errwrappers

import (
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core-v9/errcore"
	"github.com/alimtvnetwork/core-v9/isany"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
)

type newErrorInterfaceToCollection struct{}

// AnyType
//
//	tries may ways to get to the error wrapper.
//	on fail returns parsed error
//
// Steps:
//   - *errorwrapper.Wrapper
//   - errorwrapper.Wrapper
//   - errorwrapper.ErrWrapper
//   - *errcore.RawErrCollection
//   - errcore.RawErrCollection
//   - errcoreinf.BaseRawErrCollectionDefiner
//   - errcoreinf.BasicErrWrapper
//   - errcoreinf.BaseErrorWrapperCollectionDefiner
//   - errcoreinf.BaseErrorOrCollectionWrapper
//   - corejson.Result
//   - *corejson.Result
//   - []byte
//   - string
func (it newErrorInterfaceToCollection) AnyType(
	variation errtype.Variation,
	errInf interface{},
) (errCollection *Collection, parsedErrWp *errorwrapper.Wrapper) {
	empty := Empty()
	if isany.Null(errInf) {
		return empty, nil
	}

	switch casted := errInf.(type) {
	case *Collection:
		if casted == nil || casted.IsEmpty() {
			return empty, nil
		}

		// has existing items
		return casted, nil
	case Collection:
		return casted.ToPtr(), nil
	case *errorwrapper.Wrapper:
		return empty.AddWrapperPtr(casted), nil
	case errorwrapper.Wrapper:
		return empty.AddWrapperPtr(casted.Ptr()), nil
	case errorwrapper.ErrWrapper:
		return empty.AddWrapperPtr(casted.AsErrorWrapper()), nil

	case *errcore.RawErrCollection:
		return empty.AddWrapperPtr(
			errnew.ErrInterface.ActualRawErrCollection(
				variation,
				casted)), nil

	case errcoreinf.BaseRawErrCollectionDefiner:
		return empty.AddWrapperPtr(
			errnew.ErrInterface.RawErrCollection(
				variation,
				casted)), nil

	case errcoreinf.BasicErrWrapper:
		return empty.AddWrapperPtr(
			errnew.ErrInterface.BasicErr(casted)), nil
	case errcoreinf.BaseErrorWrapperCollectionDefiner:
		return empty.AddWrapperPtr(
			errnew.ErrInterface.BasicErr(
				casted.GetAsBasicWrapperUsingTyper(
					variation.AsBasicErrorTyper()))), nil
	case errcoreinf.BaseErrorOrCollectionWrapper:
		return empty.AddWrapperPtr(
			errnew.ErrInterface.Default(variation, casted)), nil
	case corejson.Result:
		return Deserialize.UsingJsonResult(&casted)
	case *corejson.Result:
		return Deserialize.UsingJsonResult(casted)
	case []byte:
		return Deserialize.UsingBytes(casted)
	case string:
		return Deserialize.UsingBytes([]byte(casted))
	}

	return empty, errnew.FromTo.Message(errtype.CastingFailed,
		"cannot cast or make to error collection",
		errInf, empty)
}

func (it newErrorInterfaceToCollection) Create(
	baseErrOrCollection errcoreinf.BaseErrorOrCollectionWrapper,
) *Collection {
	if baseErrOrCollection == nil || baseErrOrCollection.IsEmpty() {
		return Empty()
	}

	casted := it.Cast(baseErrOrCollection)

	if casted != nil {
		return casted
	}

	return Empty()
}

func (it newErrorInterfaceToCollection) UsingCollectionDefiner(
	collectionDefiner errcoreinf.BaseErrorWrapperCollectionDefiner,
) *Collection {
	if collectionDefiner == nil || collectionDefiner.IsEmpty() {
		return Empty()
	}

	casted := it.Cast(collectionDefiner)

	if casted != nil {
		return casted
	}

	return Empty()
}

// Cast
//
//	returns nil on casting failed or not supported.
func (it newErrorInterfaceToCollection) Cast(
	baseErrOrCollection errcoreinf.BaseErrorOrCollectionWrapper,
) *Collection {
	if baseErrOrCollection == nil || baseErrOrCollection.IsEmpty() {
		return nil
	}

	switch casted := baseErrOrCollection.(type) {
	case *Collection:
		return casted
	case *errorwrapper.Wrapper:
		return Empty().AddWrapperPtr(casted)
	case errorwrapper.ErrWrapper:
		return Empty().AddWrapperPtr(casted.AsErrorWrapper())
	case errcoreinf.BasicErrWrapper:
		collection := NewCap1()
		collection.AddBasicErrWrapper(casted)

		return collection
	}

	return nil
}
