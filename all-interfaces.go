package errorwrapper

import (
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/coreinterface"
	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
	"github.com/alimtvnetwork/errorwrapper-v4/refs"
)

type ResultsContractsBinder interface {
	ValuesWithErrorWrapper
	AsResultsContractsBinder() ResultsContractsBinder
}

type ErrorTypeStringer interface {
	TypeCodeNameString() string
	TypeNameCode() string
	TypeString() string
}

type TypeOfErrorWrapperGetter interface {
	Type() errtype.Variation
}

type ValuesWithErrorWrapper interface {
	ValueWithErrorWrapper
	coreinterface.Clearer
	coreinterface.HasAnyItemChecker
	coreinterface.LengthGetter
	coreinterface.IsAnyNullChecker
}

type ValueWithErrorWrapper interface {
	coreinterface.HasErrorChecker
	coreinterface.IsEmptyChecker
	coreinterface.IsSuccessValidator
	coreinterface.IsInvalidChecker
	coreinterface.HasIssuesOrEmptyChecker
	coreinterface.Stringer
	corejson.Jsoner
	corejson.JsonParseSelfInjector
	corejson.JsonContractsBinder
	coreinterface.HasSafeItemsChecker
	InterfaceGetter
	coreinterface.SafeStringer
	coreinterface.Disposer
}

type ValuesWithErrorWrapperCollectionBinder interface {
	ValuesWithErrorWrapper
	AsValuesWithErrorWrapperCollectionBinder() ValuesWithErrorWrapperCollectionBinder
}

type InterfaceGetter interface {
	ErrorWrapperInf() ErrWrapper
}

type ValueWithErrorWrapperBinder interface {
	ValueWithErrorWrapper
	AsValueWithErrorWrapperBinder() ValueWithErrorWrapperBinder
}

type Getter interface {
	ErrorWrapper() *Wrapper
}

type Cloner interface {
	Clone() Wrapper
	ClonePtr() *Wrapper
}

type ErrorStringer interface {
	ErrorString() string
}

// ReflectSetToErrorWrapper
//
// # Set any object from to toPointer object
//
// Valid Inputs or Supported (https://t.ly/1Lpt):
//   - From, To: (null, null)                          -- do nothing
//   - From, To: (sameTypePointer, sameTypePointer)    -- try reflection
//   - From, To: (sameTypeNonPointer, sameTypePointer) -- try reflection
//   - From, To: ([]byte or *[]byte, otherType)        -- try unmarshal, reflect
//   - From, To: (otherType, *[]byte)                  -- try marshal, reflect
//
// Validations:
//   - Check null, if both null no error return quickly.
//   - NotSupported returns as error.
//   - NotSupported: (from, to) - (..., not pointer)
//   - NotSupported: (from, to) - (null, notNull)
//   - NotSupported: (from, to) - (notNull, null)
//   - NotSupported: (from, to) - not same type and not bytes on any
//   - `From` null or nil is not supported and will return error.
//
// Reference:
//   - Reflection String Set Example : https://go.dev/play/p/fySLYuOvoRK.go?download=true
//   - Method document screenshot    : https://prnt.sc/26dmf5g
type ReflectSetToErrorWrapper interface {
	// ReflectSetToErrWrap
	//
	// Set any object from to toPointer object
	//
	// Valid Inputs or Supported (https://t.ly/1Lpt):
	//  - From, To: (null, null)                          -- do nothing
	//  - From, To: (sameTypePointer, sameTypePointer)    -- try reflection
	//  - From, To: (sameTypeNonPointer, sameTypePointer) -- try reflection
	//  - From, To: ([]byte or *[]byte, otherType)        -- try unmarshal, reflect
	//  - From, To: (otherType, *[]byte)                  -- try marshal, reflect
	//
	// Validations:
	//  - Check null, if both null no error return quickly.
	//  - NotSupported returns as error.
	//      - NotSupported: (from, to) - (..., not pointer)
	//      - NotSupported: (from, to) - (null, notNull)
	//      - NotSupported: (from, to) - (notNull, null)
	//      - NotSupported: (from, to) - not same type and not bytes on any
	//  - `From` null or nil is not supported and will return error.
	//
	// Reference:
	//  - Reflection String Set Example : https://go.dev/play/p/fySLYuOvoRK.go?download=true
	//  - Method document screenshot    : https://prnt.sc/26dmf5g
	ReflectSetToErrWrap(toPtr interface{}) *Wrapper
}

type BasicErrWrapper interface {
	errcoreinf.BasicErrWrapper

	StackTraceString() string
	AsErrorWrapper() *Wrapper
	IsNoError() bool

	coreinterface.Stringer
}

type ConcatNewGetter interface {
	ConcatNew() ConcatNew
}

type MustBeEmptyError interface {
	MustBeEmpty()
}

type ErrWrapper interface {
	BasicErrWrapper

	TypeOfErrorWrapperGetter
	// MustBeEmptyError
	//
	//  Panics if error exist
	MustBeEmptyError

	References() *refs.Collection
	MergeNewReferences(additionalReferences ...ref.Value) *refs.Collection
	IsTypeOf(errType errtype.Variation) bool
	ReflectSetToErrorWrapper

	corejson.JsonContractsBinder
}

type ErrWrapperContractsBinder interface {
	ErrWrapper
	AsErrWrapperContractsBinder() ErrWrapperContractsBinder
}
