package errorwrapper

import (
	"errors"
	"strings"

	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/corecsv"
	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
	"github.com/alimtvnetwork/errorwrapper-v4/refs"
)

func New(errType errtype.Variation) Wrapper {
	return *NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		"",
	)
}

func NewTypeUsingStackSkip(stackSkipIndex int, errType errtype.Variation) *Wrapper {
	return NewMsgDisplayErrorNoReference(
		stackSkipIndex+defaultSkipInternal,
		errType,
		"",
	)
}

func NewPtr(errType errtype.Variation) *Wrapper {
	return NewMsgDisplayError(
		defaultSkipInternal,
		errType,
		"",
		nil,
	)
}

func NewPtrUsingStackSkip(
	stackStartIndex int,
	errType errtype.Variation,
) *Wrapper {
	return NewMsgDisplayErrorNoReference(
		stackStartIndex+defaultSkipInternal,
		errType,
		constants.EmptyString,
	)
}

func NewUsingBasicErr(
	basicErrWrapper errcoreinf.BasicErrWrapper,
) *Wrapper {
	if basicErrWrapper == nil || basicErrWrapper.IsEmpty() {
		return nil
	}

	casted, isSuccess := basicErrWrapper.(*Wrapper)

	if isSuccess {
		return casted
	}

	errTypeVal := basicErrWrapper.
		ErrorTypeAsBasicErrorTyper().
		Value()

	return NewMsgDisplayError(
		defaultSkipInternal,
		errtype.Variation(errTypeVal),
		basicErrWrapper.ErrorString(),
		refs.NewUsingBasicErrWrap(basicErrWrapper))
}

func ErrorType(
	basicErrWrapper errcoreinf.BasicErrWrapper,
) errtype.Variation {
	if basicErrWrapper == nil || basicErrWrapper.IsEmpty() {
		return errtype.NoError
	}

	errTypeVal := basicErrWrapper.
		ErrorTypeAsBasicErrorTyper().
		Value()

	return errtype.Variation(errTypeVal)
}

func NewUsingBasicErrStackSkip(
	stackStartIndex int,
	basicErrWrapper errcoreinf.BasicErrWrapper,
) *Wrapper {
	if basicErrWrapper == nil || basicErrWrapper.IsEmpty() {
		return nil
	}

	casted, isSuccess := basicErrWrapper.(*Wrapper)

	if isSuccess {
		return casted.CloneNewStackSkipPtr(stackStartIndex)
	}

	errTypeVal := basicErrWrapper.
		ErrorTypeAsBasicErrorTyper().
		Value()

	return NewMsgDisplayError(
		stackStartIndex+defaultSkipInternal,
		errtype.Variation(errTypeVal),
		basicErrWrapper.ErrorString(),
		refs.NewUsingBasicErrWrap(basicErrWrapper))
}

func CastBasicErrWrapperToWrapper(
	basicErrWrapper errcoreinf.BasicErrWrapper,
) *Wrapper {
	if basicErrWrapper == nil || basicErrWrapper.IsEmpty() {
		return nil
	}

	casted, isSuccess := basicErrWrapper.(*Wrapper)

	if isSuccess {
		return casted
	}

	return nil
}

func NewUsingManyErrorInterfacesStackSkip(
	stackStartIndex int,
	errType errtype.Variation,
	errorInterfaces ...errcoreinf.BaseErrorOrCollectionWrapper,
) *Wrapper {
	if len(errorInterfaces) == 0 {
		return EmptyPtr()
	}

	errorWrappers := NewInterfacesToErrorWrappers(
		errorInterfaces...)

	if len(errorWrappers) == 0 {
		return nil
	}

	slice := make([]string, len(errorWrappers))
	for i, errWp := range errorWrappers {
		slice[i] = errWp.CompiledErrorWithStackTraces().Error()
	}

	return NewMessagesUsingJoiner(
		stackStartIndex+defaultSkipInternal,
		errType,
		constants.DefaultLine,
		slice...)
}

func NewInterfacesToErrorWrappers(
	errorInterfaces ...errcoreinf.BaseErrorOrCollectionWrapper,
) []*Wrapper {
	if len(errorInterfaces) == 0 {
		return nil
	}

	newSlice := make([]*Wrapper, 0, len(errorInterfaces))
	for _, errInf := range errorInterfaces {
		if errInf == nil || errInf.IsEmpty() {
			continue
		}

		casted := CastInterfaceToErrorWrapper(errInf)
		if casted != nil {
			newSlice = append(
				newSlice,
				casted)

			continue
		}

		newSlice = append(
			newSlice,
			NewMsgDisplayErrorNoReference(
				defaultSkipInternal,
				errtype.Generic,
				errInf.FullStringWithTraces(),
			))
	}

	return newSlice
}

// CastInterfaceToErrorWrapper
//
//   - if empty then returns nil or empty then
//   - cast to Wrapper on success returns
//   - On fail, usages BaseErrorOrCollectionWrapper.GetAsBasicWrapper() then cast
//   - cast to Wrapper on success return if failed
func CastInterfaceToErrorWrapper(
	errorInterface errcoreinf.BaseErrorOrCollectionWrapper,
) *Wrapper {
	if errorInterface == nil || errorInterface.IsEmpty() {
		return EmptyPtr()
	}

	asWrapper, isSuccess := errorInterface.(*Wrapper)

	if isSuccess {
		return asWrapper
	}

	return CastBasicErrWrapperToWrapper(
		errorInterface.GetAsBasicWrapper())
}

func CastInterfaceToErrorWrapperUsingStackSkip(
	stackStartIndex int,
	errorInterface errcoreinf.BaseErrorOrCollectionWrapper,
) *Wrapper {
	if errorInterface == nil || errorInterface.IsEmpty() {
		return EmptyPtr()
	}

	asWrapper, isSuccess := errorInterface.(*Wrapper)
	if isSuccess {
		return asWrapper.
			CloneNewStackSkipPtr(
				stackStartIndex + defaultSkipInternal)
	}

	return EmptyPtr()
}

func InterfaceToErrorWrapperUsingStackSkip(
	stackStartIndex int,
	errType errtype.Variation,
	errorInterface errcoreinf.BaseErrorOrCollectionWrapper,
) *Wrapper {
	if errorInterface == nil || errorInterface.IsEmpty() {
		return EmptyPtr()
	}

	asWrapper, isSuccess := errorInterface.(*Wrapper)

	if isSuccess {
		return asWrapper.
			CloneNewStackSkipPtr(
				stackStartIndex + defaultSkipInternal)
	}

	return NewMsgDisplayErrorNoReference(
		stackStartIndex+defaultSkipInternal,
		errType,
		errorInterface.FullString(),
	)
}

// InterfaceToErrorWrapper
//
//	first tries to cast to Wrapper if not successful then create new one.
func InterfaceToErrorWrapper(
	errType errtype.Variation,
	errorInterface errcoreinf.BaseErrorOrCollectionWrapper,
) *Wrapper {
	if errorInterface == nil || errorInterface.IsEmpty() {
		return EmptyPtr()
	}

	asWrapper, isSuccess := errorInterface.(*Wrapper)

	if isSuccess {
		return asWrapper
	}

	return NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		errorInterface.FullString(),
	)
}

func NewMsgUsingAllParams(
	stackSkipIndex int,
	errType errtype.Variation,
	isDisplayableError bool,
	message string,
	references *refs.Collection,
) *Wrapper {
	stacks := codestack.New.StackTrace.DefaultCount(
		stackSkipIndex + defaultSkipInternal)

	if message == "" {
		return &Wrapper{
			isDisplayableError: isDisplayableError,
			errorType:          errType,
			stackTraces:        stacks,
			references:         references,
			currentError:       nil,
		}
	}

	err := errors.New(message)

	return &Wrapper{
		isDisplayableError: isDisplayableError,
		errorType:          errType,
		references:         references,
		currentError:       err,
		stackTraces:        stacks,
	}
}

func NewMsgDisplayError(
	stackSkipIndex int,
	errType errtype.Variation,
	message string,
	references *refs.Collection,
) *Wrapper {
	stacks := codestack.New.StackTrace.DefaultCount(
		stackSkipIndex + defaultSkipInternal)

	if message == "" {
		return &Wrapper{
			isDisplayableError: true,
			errorType:          errType,
			currentError:       nil,
			references:         references,
			stackTraces:        stacks,
		}
	}

	err := errors.New(message)

	return &Wrapper{
		isDisplayableError: true,
		errorType:          errType,
		references:         references,
		currentError:       err,
		stackTraces:        stacks,
	}
}

func NewMsgDisplayErrorNoReference(
	stackSkipIndex int,
	errType errtype.Variation,
	message string,
) *Wrapper {
	stacks := codestack.New.StackTrace.DefaultCount(
		stackSkipIndex + defaultSkipInternal)

	if message == "" {
		return &Wrapper{
			isDisplayableError: true,
			errorType:          errType,
			stackTraces:        stacks,
		}
	}

	err := errors.New(message)

	return &Wrapper{
		isDisplayableError: true,
		errorType:          errType,
		currentError:       err,
		stackTraces:        stacks,
	}
}

func NewMsgDisplayErrorReferencesPtr(
	stackSkipIndex int,
	errType errtype.Variation,
	message string,
	references ...ref.Value,
) *Wrapper {
	stacks := codestack.New.StackTrace.DefaultCount(
		stackSkipIndex + defaultSkipInternal)

	refsCollection := refs.NewUsingRefsOrNil(
		references...)

	if message == "" {
		return &Wrapper{
			isDisplayableError: true,
			errorType:          errType,
			stackTraces:        stacks,
			references:         refsCollection,
		}
	}

	err := errors.New(message)

	return &Wrapper{
		isDisplayableError: true,
		errorType:          errType,
		references:         refsCollection,
		currentError:       err,
		stackTraces:        stacks,
	}
}

func NewMsgDisplayErrorUsingStackTracesPtr(
	errType errtype.Variation,
	message string,
	stackTraces codestack.TraceCollection,
	references ...ref.Value,
) *Wrapper {
	refsCollection := refs.NewUsingRefsOrNil(
		references...)

	if message == "" {
		return &Wrapper{
			isDisplayableError: true,
			errorType:          errType,
			stackTraces:        stackTraces,
			references:         refsCollection,
		}
	}

	err := errors.New(message)

	return &Wrapper{
		isDisplayableError: true,
		errorType:          errType,
		references:         refsCollection,
		currentError:       err,
		stackTraces:        stackTraces,
	}
}

func NewMsgUsingCodeStacksPtr(
	errType errtype.Variation,
	isDisplayableError bool,
	message string,
	codeStacks codestack.TraceCollection,
	references *refs.Collection,
) *Wrapper {
	if message == "" {
		return &Wrapper{
			isDisplayableError: isDisplayableError,
			errorType:          errType,
			stackTraces:        codeStacks,
			references:         references,
		}
	}

	err := errors.New(message)

	return &Wrapper{
		isDisplayableError: isDisplayableError,
		errorType:          errType,
		references:         references,
		currentError:       err,
		stackTraces:        codeStacks,
	}
}

func EmptyPtr() *Wrapper {
	return StaticEmptyPtr
}

func Empty() Wrapper {
	return Wrapper{
		isDisplayableError: false,
		errorType:          errtype.NoError,
	}
}

func EmptyPrint() Wrapper {
	return Wrapper{
		isDisplayableError: true,
		errorType:          errtype.NoError,
	}
}

func NewErrorPlusMsgUsingAllParamsPtr(
	stackSkipIndex int,
	errType errtype.Variation,
	isDisplayableError bool,
	err error,
	message string,
	references *refs.Collection,
) *Wrapper {
	if err == nil {
		return NewMsgUsingAllParams(
			stackSkipIndex+defaultSkipInternal,
			errType,
			isDisplayableError,
			message,
			references)
	}

	codeStacks := codestack.New.StackTrace.DefaultCount(
		stackSkipIndex + defaultSkipInternal)

	if message == "" {
		return &Wrapper{
			isDisplayableError: isDisplayableError,
			currentError:       err,
			references:         references,
			errorType:          errType,
			stackTraces:        codeStacks,
		}
	}

	err2 := errors.New(message + err.Error())

	return &Wrapper{
		isDisplayableError: isDisplayableError,
		currentError:       err2,
		references:         references,
		errorType:          errType,
		stackTraces:        codeStacks,
	}
}

func NewErrUsingAllParams(
	stackSkipIndex int,
	errType errtype.Variation,
	isDisplayableError bool,
	err error,
	references *refs.Collection,
) Wrapper {
	return *NewErrUsingAllParamsPtr(
		stackSkipIndex,
		errType,
		isDisplayableError,
		err,
		references,
	)
}

func NewErrUsingAllParamsPtr(
	stackSkipIndex int,
	errType errtype.Variation,
	isDisplayableError bool,
	err error,
	references *refs.Collection,
) *Wrapper {
	codeStacks := codestack.New.StackTrace.DefaultCount(
		stackSkipIndex + defaultSkipInternal)

	if err == nil {
		return &Wrapper{
			isDisplayableError: isDisplayableError,
			errorType:          errType,
			stackTraces:        codeStacks,
			references:         references,
		}
	}

	return &Wrapper{
		errorType:          errType,
		isDisplayableError: isDisplayableError,
		currentError:       err,
		stackTraces:        codeStacks,
		references:         references,
	}
}

func NewError(
	stackSkipIndex int,
	err error,
) *Wrapper {
	if err == nil {
		return nil
	}

	codeStacks := codestack.New.StackTrace.DefaultCount(
		stackSkipIndex + defaultSkipInternal)

	return &Wrapper{
		isDisplayableError: true,
		currentError:       err,
		errorType:          errtype.Generic,
		stackTraces:        codeStacks,
	}
}

func NewUsingError(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
) *Wrapper {
	codeStacks := codestack.New.StackTrace.DefaultCount(
		stackSkipIndex + defaultSkipInternal)

	if err == nil {
		return &Wrapper{
			isDisplayableError: true,
			errorType:          errType,
			stackTraces:        codeStacks,
		}
	}

	return &Wrapper{
		isDisplayableError: true,
		currentError:       err,
		errorType:          errType,
		stackTraces:        codeStacks,
	}
}

func NewUsingErrorWithoutTypeDisplay(
	errType errtype.Variation,
	err error,
) Wrapper {
	if err == nil {
		return Wrapper{
			isDisplayableError: false,
			errorType:          errType,
		}
	}

	return Wrapper{
		isDisplayableError: false,
		currentError:       err,
		errorType:          errType,
	}
}

func NewUsingErrorWithoutTypeDisplayPtr(
	errType errtype.Variation,
	err error,
) *Wrapper {
	if err == nil {
		return &Wrapper{
			errorType: errType,
		}
	}

	return &Wrapper{
		currentError: err,
		errorType:    errType,
	}
}

func NewUsingTypeErrorAndMessage(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
	msg string,
) *Wrapper {
	codeStacks := codestack.New.StackTrace.DefaultCount(
		stackSkipIndex + defaultSkipInternal)

	if err == nil {
		return &Wrapper{
			isDisplayableError: true,
			errorType:          errType,
			stackTraces:        codeStacks,
		}
	}

	newErr := ErrorMessageToError(
		err,
		msg)

	return &Wrapper{
		isDisplayableError: true,
		currentError:       newErr,
		errorType:          errType,
		stackTraces:        codeStacks,
	}
}

func NewUsingErrorAndMessage(
	stackSkipIndex int,
	err error,
	msg string,
) *Wrapper {
	return NewUsingTypeErrorAndMessage(
		stackSkipIndex+defaultSkipInternal,
		errtype.Generic,
		err,
		msg)
}

func NewMessagesUsingJoiner(
	stackSkipIndex int,
	errType errtype.Variation,
	joiner string,
	messages ...string,
) *Wrapper {
	codeStacks := codestack.New.StackTrace.DefaultCount(
		stackSkipIndex + defaultSkipInternal)

	if len(messages) == 0 {
		return &Wrapper{
			isDisplayableError: true,
			errorType:          errType,
			stackTraces:        codeStacks,
		}
	}

	toStr := strings.Join(messages, joiner)
	err := errors.New(toStr)

	return &Wrapper{
		isDisplayableError: true,
		currentError:       err,
		errorType:          errType,
		stackTraces:        codeStacks,
	}
}

func NewGeneric(
	stackSkipIndex int,
	error error,
) *Wrapper {
	return NewUsingError(
		stackSkipIndex+defaultSkipInternal,
		errtype.Generic,
		error)
}

func NewUnknownMessage(
	stackSkipIndex int,
	isDisplayableError bool,
	message string,
) *Wrapper {
	return NewMsgUsingAllParams(
		stackSkipIndex+defaultSkipInternal,
		errtype.Unknown,
		isDisplayableError,
		message,
		nil)
}

func NewUsingWrapper(
	stackSkipIndex int,
	currentWrapper *Wrapper,
	additionalReferences ...ref.Value,
) *Wrapper {
	if currentWrapper == nil || currentWrapper.IsEmpty() {
		return nil
	}

	codeStacks := codestack.New.StackTrace.Default(
		stackSkipIndex+defaultSkipInternal,
		codestack.DefaultStackCount*2)

	references := refs.
		NewExistingCollectionPlusAddition(
			currentWrapper.references,
			additionalReferences...)

	return NewMsgUsingCodeStacksPtr(
		currentWrapper.Type(),
		true,
		currentWrapper.Error().Error(),
		codeStacks,
		references)
}

func NewRef(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
	varName string,
	val interface{},
) *Wrapper {
	if err == nil {
		return nil
	}

	reference := refs.
		NewWithItem(constants.One, varName, val)

	return NewMsgUsingAllParams(
		stackSkipIndex+defaultSkipInternal,
		errType,
		true,
		err.Error(),
		reference)
}

func NewRefOne(
	stackSkipIndex int,
	errType errtype.Variation,
	var1 string,
	val1 interface{},
) *Wrapper {
	return NewOnlyRefs(
		stackSkipIndex+defaultSkipInternal,
		errType,
		ref.Value{
			Value:    val1,
			Variable: var1,
		},
	)
}

func NewRef1Msg(
	stackSkipIndex int,
	errType errtype.Variation,
	msg string,
	var1 string,
	val1 interface{},
) *Wrapper {
	if msg == constants.EmptyString {
		return nil
	}

	return NewRefWithMessage(
		stackSkipIndex+defaultSkipInternal,
		errType,
		msg,
		ref.Value{
			Value:    val1,
			Variable: var1,
		},
	)
}

func NewRef2Msg(
	stackSkipIndex int,
	errType errtype.Variation,
	msg string,
	var1 string,
	val1 interface{},
	var2 string,
	val2 interface{},
) *Wrapper {
	return NewRefWithMessage(
		stackSkipIndex+defaultSkipInternal,
		errType,
		msg,
		ref.Value{
			Value:    val1,
			Variable: var1,
		},
		ref.Value{
			Value:    val2,
			Variable: var2,
		},
	)
}

// TypeReferenceQuick - errorTypeName - (...., items)...
func TypeReferenceQuick(
	stackSkipIndex int,
	errType errtype.Variation,
	referencesValues ...interface{},
) *Wrapper {
	compiledMessage := corecsv.AnyItemsToStringDefault(
		referencesValues...)

	return NewMsgDisplayErrorReferencesPtr(
		stackSkipIndex+defaultSkipInternal,
		errType,
		compiledMessage,
		ref.Value{
			Variable: ":",
			Value:    compiledMessage,
		},
	)
}

// NewErrorRef1 alias for NewRef
func NewErrorRef1(
	stackSkipIndex int,
	errType errtype.Variation,
	err error,
	var1 string,
	val1 interface{},
) *Wrapper {
	if err == nil {
		return nil
	}

	reference := refs.
		NewWithItem(constants.Capacity1, var1, val1)

	return NewMsgUsingAllParams(
		stackSkipIndex+defaultSkipInternal,
		errType,
		true,
		err.Error(),
		reference)
}

func NewRefWithMessage(
	stackSkipIndex int,
	errType errtype.Variation,
	message string,
	refValues ...ref.Value,
) *Wrapper {
	reference := refs.
		New(len(refValues)).
		Adds(refValues...)

	return NewMsgUsingAllParams(
		stackSkipIndex+defaultSkipInternal,
		errType,
		true,
		message,
		reference)
}

func NewOnlyRefs(
	stackSkipIndex int,
	errType errtype.Variation,
	refValues ...ref.Value,
) *Wrapper {
	reference := refs.
		New(len(refValues)).
		Adds(refValues...)

	return NewMsgUsingAllParams(
		stackSkipIndex+defaultSkipInternal,
		errType,
		true,
		"",
		reference)
}

func NewRefs(
	stackSkipIndex int,
	errType errtype.Variation,
	error error,
	refValues ...ref.Value,
) *Wrapper {
	if error == nil {
		return nil
	}

	return NewRefWithMessage(
		stackSkipIndex+defaultSkipInternal,
		errType,
		error.Error(),
		refValues...)
}

func NewRef2(
	stackSkipIndex int,
	errType errtype.Variation,
	error error,
	varName string,
	val interface{},
	varName2 string,
	val2 interface{},
) *Wrapper {
	if error == nil {
		return nil
	}

	reference := refs.
		NewWithItem(constants.Capacity2, varName, val).
		Add(varName2, val2)

	return NewMsgUsingAllParams(
		stackSkipIndex+defaultSkipInternal,
		errType,
		true,
		error.Error(),
		reference)
}

func NewPath(
	stackSkipIndex int,
	errType errtype.Variation,
	error error,
	filePath string,
) *Wrapper {
	if error == nil {
		return nil
	}

	reference := refs.
		NewWithItem(
			defaultSkipInternal,
			"Path",
			filePath)

	return NewMsgUsingAllParams(
		stackSkipIndex+defaultSkipInternal,
		errType,
		true,
		error.Error(),
		reference)
}

func NewPathMsg(
	stackSkipIndex int,
	errType errtype.Variation,
	message,
	filePath string,
) *Wrapper {
	reference := refs.
		NewWithItem(
			defaultSkipInternal,
			"Path",
			filePath)

	return NewMsgUsingAllParams(
		stackSkipIndex+defaultSkipInternal,
		errType,
		true,
		message,
		reference)
}

func NewPathMessages(
	stackSkipIndex int,
	errType errtype.Variation,
	filePath string,
	messages ...string,
) *Wrapper {
	reference := refs.
		NewWithItem(
			constants.Capacity1,
			"Path",
			filePath)

	messagesCompiled := MessagesJoined(messages...)

	return NewMsgUsingAllParams(
		stackSkipIndex+defaultSkipInternal,
		errType,
		true,
		messagesCompiled,
		reference)
}
