package errorwrapper

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core-v9/corevalidator"
	"github.com/alimtvnetwork/core-v9/errcore"
	"github.com/alimtvnetwork/core-v9/iserror"
	"github.com/alimtvnetwork/errorwrapper-v3/errconsts"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
	"github.com/alimtvnetwork/errorwrapper-v3/refs"
)

type Wrapper struct {
	isDisplayableError           bool
	errorType                    errtype.Variation
	wrapperCompiledWithoutTraces corestr.SimpleStringOnce
	wholeCompiled                corestr.SimpleStringOnce
	typeName                     corestr.SimpleStringOnce
	typeCodeNameString           corestr.SimpleStringOnce
	stackTracesCompiled          corestr.SimpleStringOnce
	stackTraces                  codestack.TraceCollection
	references                   *refs.Collection
	currentError                 error
}

func (it *Wrapper) MustBeSafe() {
	it.HandleError()
}

func (it *Wrapper) StackTracesJsonResult() *corejson.Result {
	if it.IsEmpty() {
		return nil
	}

	return it.stackTraces.StackTracesJsonResult()
}

func (it *Wrapper) StackTraces() string {
	if it.IsEmpty() {
		return ""
	}

	return it.stackTraces.StackTraces()
}

func (it *Wrapper) NewStackTraces(stackSkip int) string {
	return codestack.
		NewStacksDefaultCount(stackSkip + defaultSkipInternal).
		CodeStacksString()
}

func (it *Wrapper) NewDefaultStackTraces() string {
	return codestack.
		NewStacksDefaultCount(defaultSkipInternal).
		CodeStacksString()
}

func (it *Wrapper) NewStackTracesJsonResult(stackSkip int) *corejson.Result {
	return codestack.
		NewStacksDefaultCount(defaultSkipInternal + stackSkip).
		JsonPtr()
}

func (it *Wrapper) NewDefaultStackTracesJsonResult() *corejson.Result {
	return codestack.
		NewStacksDefaultCount(defaultSkipInternal).
		JsonPtr()
}

func (it *Wrapper) IsBasicErrEqual(
	another errcoreinf.BasicErrWrapper,
) bool {
	if it.IsEmpty() && another == nil || another.IsEmpty() {
		return true
	}

	anotherCasted := CastBasicErrWrapperToWrapper(another)

	return it.IsEquals(anotherCasted)
}

func (it *Wrapper) MergeNewErrInf(right errcoreinf.BaseErrorOrCollectionWrapper) errcoreinf.BaseErrorOrCollectionWrapper {
	if right == nil || right.IsEmpty() {
		return it
	}

	return it.ConcatNew().ErrorInterface(right)
}

func (it *Wrapper) MergeNewMessage(newMessage string) errcoreinf.BaseErrorOrCollectionWrapper {
	return it.ConcatNew().Message(newMessage)
}

func (it *Wrapper) CloneInterface() errcoreinf.BasicErrWrapper {
	return it.ClonePtr()
}

func (it *Wrapper) GetAsBasicWrapperUsingTyper(
	errorTyper errcoreinf.BasicErrorTyper,
) errcoreinf.BasicErrWrapper {
	if it.IsEmpty() {
		return nil
	}

	return NewMsgDisplayError(
		defaultSkipInternal,
		errtype.NewUsingTyper(errorTyper),
		it.ErrorString(),
		it.references)
}

func (it *Wrapper) GetAsBasicWrapper() errcoreinf.BasicErrWrapper {
	return it
}

func (it *Wrapper) ReferencesList() []errcoreinf.Referencer {
	if it.IsReferencesEmpty() {
		return nil
	}

	return it.references.ReferencesList()
}

// ReflectSetToErrWrap
//
//  Reusing ReflectSetTo
func (it *Wrapper) ReflectSetToErrWrap(toPtr interface{}) *Wrapper {
	err := it.ReflectSetTo(toPtr)

	if err != nil {
		return NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			errtype.CastingFailed,
			"ReflectSet to cast failed!")
	}

	return nil
}

// IsCollect
//
//  ConcatNew() recommend to use instead.
//
// Warning :
//  mutates current error, recommended NOT use,
//  it is for the commonalities between error Wrapper and collection
func (it *Wrapper) IsCollect(
	another errcoreinf.BaseErrorOrCollectionWrapper,
) bool {
	if another == nil || another.IsEmpty() {
		return false
	}

	concatNew := it.ConcatNew()
	concatNewWrapper := concatNew.Wrapper

	asWrapper, isSuccess := another.(*Wrapper)
	if isSuccess {
		concatNewWrapper(asWrapper)

		return true
	}

	toWrapper := NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		it.errorTypeOrGeneric(),
		another.FullStringWithTraces(),
	)

	finalErrWp := concatNewWrapper(toWrapper)

	if finalErrWp != nil {
		*it = *finalErrWp
	}

	return true
}

// IsCollectedAny
//
//  ConcatNew() recommend to use instead.
//
// Warning :
//  mutates current error, recommended NOT use,
//  it is for the commonalities between error Wrapper and collection
func (it *Wrapper) IsCollectedAny(
	anotherItems ...errcoreinf.BaseErrorOrCollectionWrapper,
) bool {
	if len(anotherItems) == 0 {
		return false
	}

	toWrapper := NewUsingManyErrorInterfacesStackSkip(
		defaultSkipInternal,
		it.errorTypeOrGeneric(),
		anotherItems...)

	if toWrapper.IsEmpty() {
		return false
	}

	concatNew := it.ConcatNew()
	concatNewWrapper := concatNew.Wrapper
	finalErrWp := concatNewWrapper(toWrapper)

	if finalErrWp != nil {
		*it = *finalErrWp
	}

	return true
}

// IsCollectOn
//
//  ConcatNew() recommend to use instead.
//
// Warning :
//  mutates current error, recommended NOT use,
//  it is for the commonalities between error Wrapper and collection
func (it *Wrapper) IsCollectOn(
	isCollect bool,
	another errcoreinf.BaseErrorOrCollectionWrapper,
) bool {
	if !isCollect || another == nil || another.IsEmpty() {
		return false
	}

	return it.IsCollect(another)
}

func (it *Wrapper) IsEmptyAll(
	anotherItems ...errcoreinf.BaseErrorOrCollectionWrapper,
) bool {
	if len(anotherItems) == 0 {
		return true
	}

	return !it.IsCollectedAny(anotherItems...)
}

func (it *Wrapper) IsDefined() bool {
	return it.HasAnyError()
}

func (it *Wrapper) HasAnyIssues() bool {
	return it.HasAnyError()
}

func (it *Wrapper) CompiledJsonErrorWithStackTraces() error {
	if it == nil || it.IsEmpty() {
		return nil
	}

	jsonString := it.CompiledJsonStringWithStackTraces()

	return errors.New(jsonString)
}

func (it *Wrapper) CompiledJsonStringWithStackTraces() (jsonString string) {
	if it == nil || it.IsEmpty() {
		return ""
	}

	jsonResult := it.Json()

	if jsonResult.HasError() {
		finalJson := NewError(codestack.Skip1, jsonResult.MeaningfulError()).
			ConcatNew().
			Wrapper(it).
			Json()

		finalJson.MustBeSafe()

		return finalJson.SafeString()
	}

	newFinalJson := jsonResult.JsonString()
	go jsonResult.Dispose()

	return newFinalJson
}

func (it *Wrapper) MustBeEmptyError() {
	it.HandleError()
}

func (it *Wrapper) IsCollectionType() bool {
	return false
}

// ReflectSetTo
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
func (it *Wrapper) ReflectSetTo(toPtr interface{}) error {
	return coredynamic.ReflectSetFromTo(it, toPtr)
}

func (it Wrapper) ReferencesCollection() errcoreinf.ReferenceCollectionDefiner {
	return it.References()
}

func (it *Wrapper) IsNoError() bool {
	return it == nil || it.errorType == errtype.NoError
}

func (it *Wrapper) MustBeEmpty() {
	it.HandleError()
}

func (it *Wrapper) IsInvalid() bool {
	return it.HasError()
}

func (it *Wrapper) HasAnyError() bool {
	return it.HasError()
}

// Compile
//
//  Refers to the FullString
func (it *Wrapper) Compile() string {
	return it.FullString()
}

// CompileString
//
//  Refers to the FullString
func (it *Wrapper) CompileString() string {
	return it.FullString()
}

func (it *Wrapper) FullStringWithTracesIf(
	isStackTraces bool,
) string {
	if it.IsEmpty() {
		return ""
	}

	if isStackTraces {
		return it.FullStringWithTraces()
	}

	return it.FullString()
}

func (it *Wrapper) ReferencesCompiledString() string {
	if it.IsEmpty() || it.IsReferencesEmpty() {
		return ""
	}

	return it.References().String()
}

func (it *Wrapper) CompiledStackTracesString() string {
	if it.IsEmpty() {
		return ""
	}

	return it.StackTraceString()
}

func (it *Wrapper) SerializeWithoutTraces() ([]byte, error) {
	if it.IsEmpty() {
		return nil, nil
	}

	cloned := it.cloneWithoutStacktrace()

	return cloned.JsonPtr().Raw()
}

// Serialize
//
//  Returns json with stack-traces
func (it *Wrapper) Serialize() ([]byte, error) {
	if it.IsEmpty() {
		return nil, nil
	}

	return it.JsonPtr().Raw()
}

func (it *Wrapper) SerializeMust() []byte {
	if it.IsEmpty() {
		return nil
	}

	rawJsonBytes, err := it.JsonPtr().Raw()
	errcore.MustBeEmpty(err)

	return rawJsonBytes
}

func (it *Wrapper) LogIf(isLog bool) {
	if isLog {
		it.Log()
	}
}

func (it *Wrapper) HasReferences() bool {
	return it != nil && it.references.HasAnyItem()
}

func (it *Wrapper) TypeNameCodeMessage() string {
	if it.IsEmpty() {
		return ""
	}

	return it.errorType.TypeNameCodeMessage()
}

func (it *Wrapper) RawErrorTypeValue() uint16 {
	if it.IsEmpty() {
		return errtype.NoError.ValueUInt16()
	}

	return it.errorType.RawValue()
}

func (it *Wrapper) CodeTypeName() string {
	return it.TypeName()
}

func (it *Wrapper) ErrorTypeAsBasicErrorTyper() errcoreinf.BasicErrorTyper {
	return &it.errorType
}

func (it *Wrapper) StackTraceString() string {
	if it.stackTracesCompiled.IsInitialized() {
		return it.stackTracesCompiled.String()
	}

	if it.stackTraces.IsEmpty() {
		return it.
			stackTracesCompiled.
			GetPlusSetEmptyOnUninitialized()
	}

	toString := it.stackTraces.CodeStacksString()
	it.stackTracesCompiled.
		GetPlusSetOnUninitialized(toString)

	return it.stackTracesCompiled.String()
}

func (it *Wrapper) References() *refs.Collection {
	if it == nil {
		return nil
	}

	return it.references
}

func (it *Wrapper) CloneReferences() *refs.Collection {
	if it == nil || it.IsEmpty() {
		return nil
	}

	if it.HasReferences() {
		return it.references.ClonePtr()
	}

	return nil
}

func (it *Wrapper) MergeNewReferences(additionalReferences ...ref.Value) *refs.Collection {
	if it == nil || it.references.IsEmpty() {
		return refs.NewUsingValues(
			additionalReferences...)
	}

	return refs.NewExistingCollectionPlusAddition(
		it.references,
		additionalReferences...)
}

// HasError
//
// Returns true if not empty. Invert of IsEmpty()
//
// Conditions (true):
//  - if Wrapper is NOT nil, Or,
//  - if Wrapper is NOT StaticEmptyPtr, Or,
//  - if Wrapper .errorType is NOT IsNoError(), Or,
//  - if Wrapper .currentError is nil and Wrapper .references.IsEmpty()
func (it *Wrapper) HasError() bool {
	return !it.IsEmpty()
}

// HasCurrentError
//
// Refers to Wrapper error embedded or not.
//
// Best to use HasError
func (it *Wrapper) HasCurrentError() bool {
	return it != nil && it.currentError != nil
}

func (it *Wrapper) Type() errtype.Variation {
	if it == nil {
		return errtype.NoError
	}

	return it.errorType
}

// errorTypeOrGeneric
//
//  return errtype.Generic if empty or noError type,
//  or else returns the exact type
func (it *Wrapper) errorTypeOrGeneric() errtype.Variation {
	if it == nil || it.errorType == errtype.NoError {
		return errtype.Generic
	}

	return it.errorType
}

func (it *Wrapper) GetTypeVariantStruct() errtype.VariantStructure {
	if it == nil {
		return errtype.NoError.VariantStructure()
	}

	return it.errorType.VariantStructure()
}

// TypeString
//
//  Returns whole type string, should be refactored to whole-type string name
//
//  Format :
//   - errconsts.VariantStructStringFormat
//   - "%s (Code - %d) : %s" : "Name (Code - ValueInt) : Message from type string"
//   - Exact Example for errtype.Generic : "Generic (Code - 1) : Generic error"
func (it *Wrapper) TypeString() string {
	if it == nil {
		return constants.EmptyString
	}

	return it.errorType.VariantStructure().String()
}

// TypeNameCode
//
//  Returns error type Code number value with name,
//  on empty returns empty string.
//
// Format :
//  - "(#%d - %s)"
//  - "(#1 - Generic)"
func (it *Wrapper) TypeNameCode() string {
	if it == nil {
		return constants.EmptyString
	}

	return it.errorType.CodeWithTypeName()
}

// TypeName
//
//  Returns error type name,
//  on empty returns empty string.
//
// Example :
//  - For errtype.NoError : ""
//  - For errtype.Generic : "Generic"
func (it *Wrapper) TypeName() string {
	if it == nil || it.errorType == errtype.NoError {
		return constants.EmptyString
	}

	return it.errorType.Name()
}

// TypeNameWithCustomMessage
//
// 	errconsts.ErrorCodeHyphenTypeNameWithLineFormat = "(#%d - %s) %s"
func (it *Wrapper) TypeNameWithCustomMessage(customMessage string) string {
	if it == nil {
		return constants.EmptyString
	}

	return it.errorType.CodeTypeNameWithCustomMessage(customMessage)
}

// TypeCodeNameString
//
// Format : errconsts.ErrorCodeWithTypeNameFormat
//  "(Code - #%d) : %s"
func (it *Wrapper) TypeCodeNameString() string {
	if it == nil {
		return constants.EmptyString
	}

	if it.typeCodeNameString.IsInitialized() {
		return it.typeCodeNameString.String()
	}

	typeWithCode := it.
		errorType.
		ExplicitCodeValueName()

	return it.
		typeCodeNameString.
		GetPlusSetOnUninitialized(typeWithCode)
}

func (it *Wrapper) IsTypeOf(errType errtype.Variation) bool {
	return it.errorType == errType
}

func (it *Wrapper) IsEmptyError() bool {
	return it == nil || it.IsEmpty() || it.currentError == nil
}

func (it *Wrapper) Error() error {
	if it.IsEmpty() {
		return nil
	}

	return it.currentError
}

// ErrorString if empty error then returns ""
func (it *Wrapper) ErrorString() string {
	if it.IsEmpty() || it.currentError == nil {
		return constants.EmptyString
	}

	return it.currentError.Error()
}

func (it *Wrapper) IsReferencesEmpty() bool {
	return it == nil ||
		it.references == nil ||
		it.references.IsEmpty()
}

func (it *Wrapper) IsNull() bool {
	return it == nil
}

func (it *Wrapper) IsAnyNull() bool {
	return it == nil || it == StaticEmptyPtr || it.errorType.IsNoError()
}

// IsEmpty
//
// Refers to no error for print or doesn't treat this as error.
//
// Conditions (true):
//  - if Wrapper nil, Or,
//  - if Wrapper is StaticEmptyPtr, Or,
//  - if Wrapper .errorType is IsNoError(), Or,
//  - if Wrapper .currentError NOT nil and Wrapper .references.IsEmpty()
func (it *Wrapper) IsEmpty() bool {
	if it == nil || it == StaticEmptyPtr {
		return true
	}

	return it.errorType.IsNoError() ||
		it.currentError == nil &&
			it.references.IsEmpty()
}

func (it *Wrapper) IsSuccess() bool {
	return it.IsEmpty()
}

func (it *Wrapper) IsValid() bool {
	return it.IsEmpty()
}

func (it *Wrapper) IsFailed() bool {
	return it.HasError()
}

func (it *Wrapper) IsNotEquals(right *Wrapper) bool {
	return !it.IsEquals(right)
}

func (it *Wrapper) IsEquals(right *Wrapper) bool {
	if it == nil && right == nil {
		return true
	}

	if it == nil || right == nil {
		return false
	}

	if it == right {
		return true
	}

	if it.isDisplayableError != right.isDisplayableError {
		return false
	}

	if it.errorType != right.errorType {
		return false
	}

	if it.IsEmpty() && right.IsEmpty() {
		return true
	}

	// both are not nil confirmed, so if any nil returns false.
	if right.IsEmpty() || it.IsEmpty() {
		return false
	}

	if iserror.NotEqual(it.currentError, right.currentError) {
		return false
	}

	return it.
		references.
		IsEqual(right.references)
}

func (it *Wrapper) IsErrorEquals(err error) bool {
	if err == nil && it.IsEmpty() {
		return true
	}

	if err == nil || it.IsEmpty() {
		return false
	}

	if it.Error() == err {
		return true
	}

	if it.ErrorString() == err.Error() {
		return true
	}

	return false
}

func (it *Wrapper) IsErrorMessageEqual(msg string) bool {
	return it.IsErrorMessage(msg, true)
}

// IsErrorMessage
//
// If error IsEmpty then returns false regardless
func (it *Wrapper) IsErrorMessage(msg string, isCaseSensitive bool) bool {
	if it.IsEmpty() {
		return false
	}

	errMsg := it.ErrorString()

	if errMsg == msg {
		// same string or empty or also for case sensitivity
		return true
	}

	if !isCaseSensitive {
		return strings.EqualFold(errMsg, msg)
	}

	return false
}

// IsErrorMessageContains If error IsEmpty then returns false regardless
func (it *Wrapper) IsErrorMessageContains(
	msg string,
	isCaseSensitive bool,
) bool {
	if it.IsEmpty() && msg != "" {
		return false
	}

	errMsg := it.ErrorString()

	if errMsg == msg {
		// same string or empty or also for case sensitivity
		return true
	}

	if !isCaseSensitive {
		lowerErrorMsg := strings.ToLower(errMsg)
		lowerMsg := strings.ToLower(msg)

		return strings.Index(lowerErrorMsg, lowerMsg) > -1
	}

	return strings.Index(errMsg, msg) > -1
}

// HandleError
//
// Only call panic if Wrapper has currentError
func (it *Wrapper) HandleError() {
	if it.IsEmpty() {
		return
	}

	panic(it.FullStringWithTraces())
}

// HandleErrorWithMsg Only call panic if has currentError
func (it *Wrapper) HandleErrorWithMsg(newMessage string) {
	if it == nil || it.IsEmpty() {
		return
	}

	message := fmt.Sprintf(
		errconsts.DoubleStringTogetherFormat,
		newMessage,
		it.FullStringWithTraces())

	panic(message)
}

func (it *Wrapper) HandleErrorWithRefs(
	newMessage string,
	refVar,
	refVal interface{},
) {
	if it == nil || it.IsEmpty() {
		return
	}

	message := fmt.Sprintf(
		errconsts.DoubleStringTogetherWithReferenceFormat,
		newMessage,
		it.FullStringWithTraces(),
		refVar,
		refVal,
		refVal)

	panic(message)
}

func (it *Wrapper) Value() error {
	if it == nil {
		return nil
	}

	return it.currentError
}

func (it *Wrapper) FullStringSplitByNewLine() []string {
	if it == nil {
		return []string{}
	}

	return strings.Split(
		it.FullString(),
		constants.NewLineUnix)
}

func (it *Wrapper) FullString() string {
	if it == nil {
		return constants.EmptyString
	}

	if it.wrapperCompiledWithoutTraces.IsInitialized() {
		return it.wrapperCompiledWithoutTraces.String()
	}

	if !it.isDisplayableError {
		// only prints error string
		return it.ErrorString()
	}

	finalResult := it.
		wrapperCompiledWithoutTraces.
		GetPlusSetOnUninitializedFunc(
			it.finalMessageWithCategory)

	it.wrapperCompiledWithoutTraces = it.
		wrapperCompiledWithoutTraces

	return finalResult
}

func (it *Wrapper) FullStringWithoutReferences() string {
	if it == nil {
		return constants.EmptyString
	}

	if !it.isDisplayableError {
		// only prints error string
		return it.ErrorString()
	}

	return it.finalMessageWithoutRefPlusCategory(constants.EmptyString)
}

func (it *Wrapper) FullStringWithTraces() string {
	if it == nil {
		return constants.EmptyString
	}

	if it.wholeCompiled.IsInitialized() {
		return it.wholeCompiled.String()
	}

	if !it.isDisplayableError {
		// only prints error string
		return it.ErrorString()
	}

	fullString := it.FullString()

	if it.stackTraces.HasAnyItem() {
		toString := fullString +
			constants.NewLineUnix +
			it.StackTraceString()

		return it.
			wholeCompiled.
			GetPlusSetOnUninitialized(
				toString)
	}

	return it.
		wholeCompiled.
		GetPlusSetOnUninitialized(fullString)
}

func (it *Wrapper) StackTracesLimit(limit int) *codestack.TraceCollection {
	if it == nil {
		return codestack.EmptyTraceCollection()
	}

	if limit <= constants.TakeAllMinusOne {
		return &it.stackTraces
	}

	return it.stackTraces.SafeLimitCollection(limit)
}

func (it *Wrapper) FullStringWithLimitTraces(limit int) string {
	if it == nil {
		return constants.EmptyString
	}

	stacks := it.StackTracesLimit(limit)
	fullString := it.FullString()

	if stacks.IsEmpty() {
		return fullString
	}

	toString := fullString +
		constants.DefaultLine +
		stacks.CodeStacksString()

	return toString
}

func (it *Wrapper) getRefsCollectionCompiled() string {
	if it == nil || it.references == nil {
		return constants.EmptyString
	}

	return it.
		references.
		String()
}

func (it *Wrapper) finalMessageWithCategory() string {
	refsCompiled := it.getRefsCollectionCompiled()

	return it.finalMessageWithoutRefPlusCategory(refsCompiled)
}

func (it *Wrapper) finalMessageWithoutRefPlusCategory(referencesCompiled string) string {
	if it == nil {
		return ""
	}

	errorVariantStructure := it.GetTypeVariantStruct()
	errString := it.ErrorString()
	finalString := errorVariantStructure.
		CombineRefs(
			errString,
			referencesCompiled)

	return finalString
}

func (it *Wrapper) RawErrorTypeName() string {
	if it == nil {
		return constants.EmptyString
	}

	if it.typeName.IsInitialized() {
		return it.typeName.Value()
	}

	msg := fmt.Sprintf(
		constants.SprintTypeFormat,
		it.currentError)

	return it.
		typeName.
		GetPlusSetOnUninitialized(msg)
}

func (it Wrapper) String() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.FullString()
}

func (it Wrapper) StringIf(isWithRef bool) string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	if isWithRef {
		return it.FullString()
	}

	return it.FullStringWithoutReferences()
}

func (it *Wrapper) CompiledError() error {
	if it.IsEmpty() {
		return nil
	}

	return errors.New(it.FullString())
}

func (it *Wrapper) CompiledErrorWithStackTraces() error {
	if it == nil || it.IsEmpty() {
		return nil
	}

	return errors.New(
		it.FullStringWithTraces())
}

func (it *Wrapper) FullOrErrorMessage(
	isErrorMessage,
	isWithRef bool,
) string {
	if isErrorMessage {
		return it.ErrorString()
	}

	return it.StringIf(isWithRef)
}

func (it *Wrapper) Log() {
	if it.IsEmpty() {
		return
	}

	fmt.Println(it.FullString())
}

func (it *Wrapper) LogWithTraces() {
	if it.IsEmpty() {
		return
	}

	fmt.Println(it.FullStringWithTraces())
}

func (it *Wrapper) LogFatal() {
	if it.IsEmpty() {
		return
	}

	log.Fatalln(it.FullString())
}

func (it *Wrapper) LogFatalWithTraces() {
	if it.IsEmpty() {
		return
	}

	log.Fatalln(it.FullStringWithTraces())
}

// ConcatNew
//
//  It is safe to use for nil.
func (it *Wrapper) ConcatNew() ConcatNew {
	return ConcatNew{
		errWp: it,
	}
}

func (it *Wrapper) Dispose() {
	if it == nil {
		return
	}

	it.stackTraces.Dispose()
	it.references.Dispose()
	it.references = nil
	it.wrapperCompiledWithoutTraces.Dispose()
	it.wholeCompiled.Dispose()
	it.typeName.Dispose()
	it.typeCodeNameString.Dispose()
	it.stackTracesCompiled.Dispose()
	it.currentError = nil
}

func (it Wrapper) JsonModel() WrapperDataModel {
	return NewDataModel(&it)
}

func (it Wrapper) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it Wrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *Wrapper) UnmarshalJSON(data []byte) error {
	var dataModel WrapperDataModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		transpileModelToWrapper(
			&dataModel,
			it)
	}

	return err
}

func (it Wrapper) JsonResultWithoutTraces() *corejson.Result {
	return corejson.NewPtr(it.cloneWithoutStacktrace())
}

func (it Wrapper) Json() corejson.Result {
	return corejson.New(it)
}

func (it Wrapper) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *Wrapper) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Wrapper, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return it, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
//goland:noinspection GoLinterLocal
func (it *Wrapper) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Wrapper {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *Wrapper) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *Wrapper) ValidationErrUsingTextValidator(
	validator *corevalidator.TextValidator,
	params *corevalidator.ValidatorParamsBase,
) *Wrapper {
	err := validator.VerifyDetailError(
		params,
		it.FullString())

	if err != nil {
		return NewUsingError(
			defaultSkipInternal,
			errtype.ValidationMismatch,
			err)
	}

	return StaticEmptyPtr
}

func (it *Wrapper) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Wrapper) AsErrorWrapper() *Wrapper {
	return it
}

func (it Wrapper) NonPtr() Wrapper {
	return it
}

func (it Wrapper) Ptr() *Wrapper {
	return &it
}

func (it *Wrapper) ClonePtr() *Wrapper {
	if it == nil {
		return nil
	}

	err := it.currentError
	refCollection := it.CloneReferences()

	// those string nil fields will get generated in runtime.
	return &Wrapper{
		isDisplayableError: it.isDisplayableError,
		errorType:          it.errorType,
		references:         refCollection,
		stackTraces:        it.stackTraces.Clone(),
		currentError:       err,
	}
}

func (it Wrapper) Clone() Wrapper {
	if it.IsEmpty() {
		return Empty()
	}

	errString := errcore.ToString(it.currentError)
	refCollection := it.CloneReferences()

	// those string nil fields will get generated in runtime.
	return Wrapper{
		isDisplayableError: it.isDisplayableError,
		errorType:          it.errorType,
		references:         refCollection,
		stackTraces:        it.stackTraces.Clone(),
		currentError:       errors.New(errString),
	}
}

func (it Wrapper) cloneWithoutStacktrace() Wrapper {
	if it.IsEmpty() {
		return Empty()
	}

	errString := errcore.ToString(it.currentError)
	refCollection := it.CloneReferences()

	// those string nil fields will get generated in runtime.
	return Wrapper{
		isDisplayableError: it.isDisplayableError,
		errorType:          it.errorType,
		references:         refCollection,
		currentError:       errors.New(errString),
	}
}

func (it Wrapper) CloneNewStackSkipPtr(stackSkip int) *Wrapper {
	if it.IsEmpty() {
		return nil
	}

	errString := errcore.ToString(it.currentError)
	refCollection := it.CloneReferences()

	// those string nil fields will get generated in runtime.
	return &Wrapper{
		isDisplayableError: it.isDisplayableError,
		errorType:          it.errorType,
		references:         refCollection,
		stackTraces:        codestack.NewStacksDefaultCount(stackSkip + defaultSkipInternal),
		currentError:       errors.New(errString),
	}
}

func (it Wrapper) AsErrWrapperContractsBinder() ErrWrapperContractsBinder {
	return &it
}
