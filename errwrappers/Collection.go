package errwrappers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coredata/corepayload"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/corevalidator"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/errorwrapper/internal/messages"
	"gitlab.com/evatix-go/errorwrapper/ref"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/refs"
)

type Collection struct {
	items []*errorwrapper.Wrapper
}

func (it *Collection) IsNull() bool {
	return it == nil
}

func (it *Collection) IsAnyNull() bool {
	return it == nil || it.items == nil
}

func (it *Collection) ErrorString() string {
	return it.String()
}

func (it *Collection) IsInvalid() bool {
	return it.HasError()
}

func (it *Collection) HasAnyError() bool {
	return it.HasError()
}

func (it *Collection) HasAnyIssues() bool {
	return it.HasError()
}

func (it *Collection) IsDefined() bool {
	return it.HasError()
}

func (it *Collection) Compile() string {
	return it.String()
}

func (it *Collection) HandleErrorWithRefs(
	newMessage string,
	refVar, refVal interface{},
) {
	if it.IsEmpty() {
		return
	}

	finalErrWp := errnew.Ref.Message(
		errtype.Generic,
		converters.AnyToValueString(refVar),
		refVal,
		newMessage)

	panic(finalErrWp.FullStringWithTraces() + it.FullStringWithTraces())
}

func (it *Collection) HandleErrorWithMsg(newMessage string) {
	if it.IsEmpty() {
		return
	}

	panic(newMessage + it.FullStringWithTraces())
}

func (it *Collection) FullString() string {
	if it.IsEmpty() {
		return ""
	}

	return it.String()
}

func (it *Collection) CompiledError() error {
	if it.IsEmpty() {
		return nil
	}

	return errors.New(it.String())
}

func (it *Collection) FullStringWithTraces() string {
	if it.IsEmpty() {
		return ""
	}

	slice := it.StringsIf(
		true)

	return strings.Join(
		slice,
		constants.DefaultLine)
}

func (it *Collection) FullStringWithTracesIf(
	isStackTraces bool,
) string {
	if it.IsEmpty() {
		return ""
	}

	slice := it.StringsIf(
		isStackTraces)

	return strings.Join(
		slice,
		constants.DefaultLine)
}

func (it *Collection) AllReferences() *refs.Collection {
	if it.IsEmpty() {
		return refs.EmptyPtr()
	}

	allRefs := refs.New(it.Length())

	for _, item := range it.items {
		allRefs.AddCollection(item.References())
	}

	return allRefs
}

func (it *Collection) ReferencesCompiledString() string {
	allRefs := it.AllReferences()

	return allRefs.String()
}

func (it *Collection) CompiledErrorWithStackTraces() error {
	if it.IsEmpty() {
		return nil
	}

	return errors.New(it.FullStringWithTraces())
}

func (it *Collection) CompiledStackTracesString() string {
	if it.IsEmpty() {
		return ""
	}

	traces := corestr.New.SimpleSlice.Cap(it.Length())

	for _, item := range it.items {
		traces.Add(item.CompiledStackTracesString())
	}

	return traces.JoinLine()
}

func (it *Collection) CompiledJsonErrorWithStackTraces() error {
	if it.IsEmpty() {
		return nil
	}

	jsonResult := it.Json()
	jsonString := jsonResult.JsonString()
	err := errors.New(jsonString)
	go jsonResult.Dispose()

	return err
}

func (it *Collection) CompiledJsonStringWithStackTraces() (jsonString string) {
	if it.IsEmpty() {
		return ""
	}

	jsonResult := it.Json()
	jsonString = jsonResult.JsonString()
	go jsonResult.Dispose()

	return jsonString
}

func (it *Collection) FullStringSplitByNewLine() []string {
	if it.IsEmpty() {
		return nil
	}

	return strings.Split(
		it.Compile(),
		constants.DefaultLine)
}

func (it *Collection) FullStringWithoutReferences() string {
	if it.IsEmpty() {
		return ""
	}

	slice := corestr.New.SimpleSlice.Cap(it.Length())

	for _, item := range it.items {
		slice.Add(item.FullStringWithoutReferences())
	}

	return slice.JoinLine()
}

func (it *Collection) MustBeEmptyError() {
	if it.IsEmpty() {
		return
	}

	panic(it.FullStringWithTraces())
}

func (it *Collection) MustBeSafe() {
	if it.IsEmpty() {
		return
	}

	panic(it.FullStringWithTraces())
}

func (it *Collection) SerializeWithoutTraces() ([]byte, error) {
	if it.IsEmpty() {
		return nil, nil
	}

	slice := corestr.New.SimpleSlice.Cap(it.Length())
	rawErrCollection := errcore.RawErrCollection{}

	for _, item := range it.items {
		allBytes, err := item.SerializeWithoutTraces()
		rawErrCollection.Add(err)

		if len(allBytes) > 0 && err != nil {
			slice.Add(string(allBytes))
		}
	}

	finalBytes, err := slice.MarshalJSON()

	if rawErrCollection.IsEmpty() {
		return finalBytes, err
	}

	rawErrCollection.Add(err)

	return finalBytes, rawErrCollection.CompiledError()
}

func (it *Collection) Serialize() ([]byte, error) {
	return it.JsonPtr().Raw()
}

func (it *Collection) SerializeMust() []byte {
	return it.JsonPtr().RawMust()
}

func (it *Collection) Value() error {
	return it.CompiledError()
}

func (it *Collection) LogFatalWithTraces() {
	if it.IsEmpty() {
		return
	}

	log.Fatalln(it.FullStringsWithTraces())
}

func (it *Collection) LogIf(isLog bool) {
	if !isLog || it.IsEmpty() {
		return
	}

	fmt.Println(it.FullString())
}

func (it *Collection) IsCollectionType() bool {
	return true
}

func (it *Collection) ReflectSetTo(
	toPtr interface{},
) error {
	return coredynamic.ReflectSetFromTo(it, toPtr)
}

// IsCollect
//
//  first tries to cast to errorwrapper.Wrapper
//  if not successful then creates new one
func (it *Collection) IsCollect(
	another errcoreinf.BaseErrorOrCollectionWrapper,
) bool {
	if another == nil || another.IsEmpty() {
		return false
	}

	if !another.IsCollectionType() {
		errWp := errnew.ErrInterface.NoType(another)
		it.AddWrapperPtr(errWp)

		return errWp.HasError()
	}

	// error collection type
	castedCollection, isSuccess := another.(*Collection)

	if isSuccess {
		it.AddCollections(castedCollection)

		return castedCollection.HasAnyItem()
	}

	panic("couldn't manage to collect error:" + another.Compile() + it.Compile())
}

func (it *Collection) IsCollectedAny(
	anotherItems ...errcoreinf.BaseErrorOrCollectionWrapper,
) bool {
	return !it.IsEmptyAll(anotherItems...)
}

func (it *Collection) IsPayloadWrapperCollected(
	payloadWrapper *corepayload.PayloadWrapper,
) (isCollected bool) {
	return it.IsCollected(
		errnew.Payload.Create(payloadWrapper))
}

// IsCollectOn
//
//  Only returns true if condition is true + error exist
func (it *Collection) IsCollectOn(
	isCollect bool,
	another errcoreinf.BaseErrorOrCollectionWrapper,
) bool {
	if !isCollect {
		return false
	}

	return it.IsCollect(another)
}

// IsEmptyAll
//
//  true represents no error occurred.
//
// To check error occurred check IsCollected or IsCollectedAny
func (it *Collection) IsEmptyAll(
	anotherItems ...errcoreinf.BaseErrorOrCollectionWrapper,
) (isAllEmpty bool) {
	if len(anotherItems) == 0 {
		return true
	}

	isAllEmpty = true

	for _, errInf := range anotherItems {
		if it.IsCollect(errInf) {
			isAllEmpty = false
		}
	}

	return isAllEmpty
}

func (it *Collection) GetAsBasicWrapper() errcoreinf.BasicErrWrapper {
	return it.GetAsBasicWrapper()
}

func (it *Collection) GetAsBasicWrapperUsingTyper(
	errorTyper errcoreinf.BasicErrorTyper,
) errcoreinf.BasicErrWrapper {
	errType := errtype.NewUsingTyper(errorTyper)

	return it.GetAsErrorWrapperUsingType(errType)
}

func (it *Collection) StackTracesJsonResult() *corejson.Result {
	if it.IsEmpty() {
		return nil
	}

	return it.AllStackTraces().JsonPtr()
}

// AllStackTraces
//
//  it will return all stack-traces string
func (it *Collection) AllStackTraces() *codestack.TraceCollection {
	if it.IsEmpty() {
		return nil
	}

	traces := codestack.NewTraceCollection(
		it.Length() * codestack.DefaultStackCount)

	for _, errorWrapper := range it.items {
		if errorWrapper.IsEmpty() {
			continue
		}

		currentTraces := errorWrapper.
			StackTracesLimit(constants.TakeAllMinusOne)

		traces.Adds(currentTraces.Items...)
	}

	return traces
}

// StackTraces
//
//  it will return all stack-traces string
func (it *Collection) StackTraces() string {
	if it.IsEmpty() {
		return ""
	}

	return it.AllStackTraces().CodeStacksString()
}

func (it *Collection) NewStackTraces(stackSkip int) string {
	return codestack.
		NewStacksDefaultCount(stackSkip + codestack.Skip1).
		CodeStacksString()
}

func (it *Collection) NewDefaultStackTraces() string {
	return codestack.
		NewStacksDefaultCountSkip1().
		CodeStacksString()
}

func (it *Collection) NewStackTracesJsonResult(
	stackSkip int,
) *corejson.Result {
	return codestack.
		NewStacksDefaultCount(stackSkip + codestack.Skip1).
		JsonPtr()
}

// NewDefaultStackTracesJsonResult
//
//  creates new stack-traces and returns as json
func (it *Collection) NewDefaultStackTracesJsonResult() *corejson.Result {
	return codestack.NewStacksDefaultCountSkip1().JsonPtr()
}

func (it *Collection) CompiledToGenericBasicErrWrapper() errcoreinf.BasicErrWrapper {
	return it.GetAsBasicWrapper()
}

func (it *Collection) CompiledToBasicErrWrapper(
	errType errcoreinf.BasicErrorTyper,
) errcoreinf.BasicErrWrapper {
	if it.IsEmpty() {
		return nil
	}

	convType := errtype.NewUsingTyper(errType)

	if convType.IsNoError() {
		return nil
	}

	return it.
		GetAsErrorWrapperUsingType(convType)
}

func (it *Collection) CompiledToErrorWithTraces(
	errType errcoreinf.BasicErrorTyper,
) error {
	if it.IsEmpty() {
		return nil
	}

	convType := errtype.NewUsingTyper(errType)

	if convType.IsNoError() {
		return nil
	}

	return it.
		GetAsErrorWrapperUsingType(convType).
		CompiledErrorWithStackTraces()
}

func (it *Collection) AddErrorUsingBasicType(
	errType errcoreinf.BasicErrorTyper,
	err error,
) errcoreinf.BaseErrorWrapperCollectionDefiner {
	if err == nil {
		return nil
	}

	errWp := errnew.Type.ErrorUsingStackSkip(
		codestack.Skip1,
		errtype.NewUsingTyper(errType),
		err)

	return it.AddWrapperPtr(errWp)
}

func (it *Collection) AddBasicErrWrapper(
	basicErrWrapper errcoreinf.BasicErrWrapper,
) errcoreinf.BaseErrorWrapperCollectionDefiner {
	if basicErrWrapper == nil || basicErrWrapper.IsEmpty() {
		return it
	}

	errWp := errnew.
		ErrInterface.
		BasicErr(basicErrWrapper)

	return it.AddWrapperPtr(errWp)
}

func (it *Collection) First() *errorwrapper.Wrapper {
	if it.IsEmpty() {
		panic(errtype.
			EmptyCollection.
			ErrorNoRefs("No error wrapper present to retrieve first."))
	}

	return it.items[constants.Zero]
}

func (it *Collection) FirstDynamic() interface{} {
	return it.First()
}

func (it *Collection) Last() *errorwrapper.Wrapper {
	if it.IsEmpty() {
		panic(errtype.
			EmptyCollection.
			ErrorNoRefs("No error wrapper present to retrieve last."))
	}

	return it.items[it.LastIndex()]
}

func (it *Collection) LastDynamic() interface{} {
	return it.Last()
}

func (it *Collection) FirstOrDefault() *errorwrapper.Wrapper {
	if it.IsEmpty() {
		return nil
	}

	return it.items[constants.Zero]
}

func (it *Collection) FirstOrDefaultCompiledError() error {
	if it.IsEmpty() {
		return nil
	}

	return it.First().CompiledError()
}

func (it *Collection) FirstOrDefaultError() error {
	if it.IsEmpty() {
		return nil
	}

	return it.First().Error()
}

func (it *Collection) FirstOrDefaultFullMessage() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.First().FullString()
}

func (it *Collection) LastOrDefaultCompiledError() error {
	if it.IsEmpty() {
		return nil
	}

	return it.Last().CompiledError()
}

func (it *Collection) LastOrDefaultError() error {
	if it.IsEmpty() {
		return nil
	}

	return it.Last().Error()
}

func (it *Collection) LastOrDefaultFullMessage() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return it.Last().FullString()
}

func (it *Collection) FirstOrDefaultDynamic() interface{} {
	return it.FirstOrDefault()
}

func (it *Collection) LastOrDefault() *errorwrapper.Wrapper {
	if it.IsEmpty() {
		return nil
	}

	return it.items[it.LastIndex()]
}

func (it *Collection) LastOrDefaultDynamic() interface{} {
	return it.LastOrDefault()
}

func (it *Collection) Skip(skippingItemsCount int) *Collection {
	items := it.items[skippingItemsCount:]

	return NewUsingErrorWrappers(items...)
}

func (it *Collection) SkipDynamic(skippingItemsCount int) interface{} {
	return it.Skip(skippingItemsCount)
}

func (it *Collection) TakeFromTo(fromIndex, toIndex int) *Collection {
	items := it.items[fromIndex:toIndex]

	return NewUsingErrorWrappers(items...)
}

func (it *Collection) Take(takeDynamicItems int) *Collection {
	items := it.items[:takeDynamicItems]

	return NewUsingErrorWrappers(items...)
}

func (it *Collection) TakeDynamic(takeDynamicItems int) interface{} {
	return it.Take(takeDynamicItems)
}

func (it *Collection) LimitDynamic(limit int) interface{} {
	return it.Take(limit)
}

func (it *Collection) Count() int {
	return it.Length()
}

func (it *Collection) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *Collection) LastIndex() int {
	return it.Length() - 1
}

func (it *Collection) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *Collection) StateTracker() StateCounter {
	return NewStateCount(it)
}

// AddError no error then skip adding
func (it *Collection) AddError(err error) {
	if err == nil {
		return
	}

	it.items = append(
		it.items,
		errorwrapper.NewUsingError(
			defaultSkipInternal,
			defaultErrorType,
			err),
	)
}

// AddErrorChain
//
//  no error then skip adding
func (it *Collection) AddErrorChain(err error) *Collection {
	if err == nil {
		return it
	}

	it.items = append(
		it.items,
		errorwrapper.NewUsingError(
			defaultSkipInternal,
			defaultErrorType,
			err),
	)

	return it
}

// AddErrorWithMessages no error then skip adding
func (it *Collection) AddErrorWithMessages(
	errType errtype.Variation,
	err error,
	messages ...string,
) *Collection {
	if err == nil {
		return it
	}

	it.items = append(
		it.items,
		errnew.Messages.ErrorWithManyUsingStackSkip(
			defaultSkipInternal,
			errType,
			err,
			messages...),
	)

	return it
}

// AddErrors no error then skip adding
func (it *Collection) AddErrors(errs ...error) {
	if errs == nil || len(errs) == 0 {
		return
	}

	it.AddErrorsPtr(&errs)
}

// AddErrorsPtr no error then skip adding
func (it *Collection) AddErrorsPtr(errs *[]error) {
	if errs == nil || len(*errs) == 0 {
		return
	}

	for _, err := range *errs {
		if err == nil {
			continue
		}

		it.items = append(
			it.items,
			errorwrapper.NewUsingError(
				defaultSkipInternal,
				defaultErrorType,
				err),
		)
	}
}

// ConditionalAddError adds error if isAdd and error not nil.
func (it *Collection) ConditionalAddError(
	isAdd bool,
	err error,
) {
	if !isAdd {
		return
	}

	it.AddError(err)
}

// ExecuteCollectFuncIf
// if true then executes the function and
// add to error collection if function returns error
func (it *Collection) ExecuteCollectFuncIf(
	isExecute bool,
	f func() *errorwrapper.Wrapper,
) (isAddedAny, isNothingAdded bool) {
	if !isExecute {
		return false, true
	}

	return it.AddAllFunctions(f)
}

func (it *Collection) AddFmt(
	variation errtype.Variation,
	format string,
	v ...interface{},
) *Collection {
	return it.AddWrapperPtr(errnew.Fmt.Fmt(
		variation,
		format,
		v...))
}

func (it *Collection) AddFmtErr(
	variation errtype.Variation,
	err error,
	format string,
	v ...interface{},
) *Collection {
	return it.AddWrapperPtr(
		errnew.Fmt.Error(
			variation,
			err,
			format,
			v...))
}

func (it *Collection) AddFmtMsg(
	variation errtype.Variation,
	msg string,
	format string,
	v ...interface{},
) *Collection {
	return it.AddWrapperPtr(
		errnew.Fmt.Message(
			variation,
			msg,
			format,
			v...))
}

func (it *Collection) AddFmtIf(
	isAdd bool,
	variation errtype.Variation,
	format string,
	v ...interface{},
) *Collection {
	if !isAdd {
		return it
	}

	return it.AddWrapperPtr(errnew.Fmt.Fmt(
		variation,
		format,
		v...))
}

func (it *Collection) AddRawErrCollection(
	variation errtype.Variation,
	message string,
	rawErrCollection errcoreinf.BaseRawErrCollectionDefiner,
) *Collection {
	if rawErrCollection == nil || rawErrCollection.IsEmpty() {
		return it
	}

	compiled := rawErrCollection.FullString()

	return it.AddWrapperPtr(
		errnew.Message.NewUsingStackSkip(
			codestack.Skip1,
			variation,
			message+" "+compiled))
}

func (it *Collection) AddSlice(
	variation errtype.Variation,
	message string,
	slice ...string,
) *Collection {
	if message == "" && len(slice) == 0 {
		return it
	}

	return it.AddWrapperPtr(
		errnew.Ref.MessageUsingStackSkip(
			codestack.Skip1,
			variation,
			"Slice",
			slice,
			message))
}

func (it *Collection) AddIf(
	isAdd bool,
	errorWrapper *errorwrapper.Wrapper,
) *Collection {
	if !isAdd {
		return it
	}

	return it.AddWrapperPtr(errorWrapper)
}

func (it *Collection) AddsIf(
	isAdd bool,
	errorWrappers ...*errorwrapper.Wrapper,
) *Collection {
	if !isAdd || len(errorWrappers) == 0 {
		return it
	}

	return it.AddWrappersPtr(&errorWrappers)
}

// ConditionalAddErrorWrapper adds error if isAdd and error not nil.
func (it *Collection) ConditionalAddErrorWrapper(
	isAdd bool,
	errorWrapper *errorwrapper.Wrapper,
) *Collection {
	if !isAdd {
		return it
	}

	return it.AddWrapperPtr(errorWrapper)
}

// AddAllErrorFunctions
//
//  Runs until all execution is done.
//  All errors wrappers will be collected.
//
//  Finally, returns status if any error is collected.
//
//  Once error is returned process halted and exits out the loop
//  Only the first occurred error will be collected.
//
//  Warning:
//   - Getting an error will NOT halt the process and continue running.
//
//  To exit out on first error see AddAnyErrorFunctions
func (it *Collection) AddAllErrorFunctions(
	errorType errtype.Variation,
	executeAllFunctions ...func() error,
) (isAnyAdded bool) {
	if len(executeAllFunctions) == 0 {
		return false
	}

	for _, executeFunc := range executeAllFunctions {
		if executeFunc == nil {
			continue
		}

		err := executeFunc()

		if err == nil {
			continue
		}

		isAnyAdded = true
		it.items = append(
			it.items,
			errnew.Type.Error(errorType, err))
	}

	return isAnyAdded
}

// AddAnyErrorFunctions
//
//  Runs until first function returns error.
//
//  Once error is returned process halted and exits out the loop
//  Only the first occurred error will be collected.
//
//  Warning:
//   - After first error, rest will not be executed.
//
//  To run all functions see AddAllErrorFunctions
func (it *Collection) AddAnyErrorFunctions(
	errorType errtype.Variation,
	executeFunctions ...func() error,
) (isAnyAdded bool) {
	if len(executeFunctions) == 0 {
		return false
	}

	for _, executeFunc := range executeFunctions {
		if executeFunc == nil {
			continue
		}

		err := executeFunc()

		if err != nil {
			it.items = append(
				it.items,
				errnew.Type.Error(errorType, err))

			return true
		}
	}

	return false
}

func (it *Collection) AddFunction(
	successExecute func() *errorwrapper.Wrapper,
) (isCollected bool) {
	if successExecute == nil {
		return false
	}

	errWrapper := successExecute()

	if errWrapper.IsEmpty() {
		return false
	}

	// has error
	it.items = append(
		it.items,
		errWrapper)

	return true
}

// AddAnyFunctions
//
//  Runs until first function returns error.
//
//  Once error is returned process halted and exits out the loop
//  Only the first occurred error will be collected.
//
//  Warning:
//   - After first error, rest will not be executed.
//
//  To run all functions see AddAllFunctions
func (it *Collection) AddAnyFunctions(
	successExecutes ...func() *errorwrapper.Wrapper,
) (isAnyAdded, isNothingAdded bool) {
	if len(successExecutes) == 0 {
		return false, true
	}

	for _, executeFunc := range successExecutes {
		if executeFunc == nil {
			continue
		}

		errWrapper := executeFunc()

		if errWrapper == nil {
			continue
		}

		if errWrapper.HasError() {
			it.items = append(
				it.items,
				errWrapper)

			return true, false
		}
	}

	return false, true
}

// AddAllFunctions
//
//  Runs until all execution is done.
//  All errors wrappers will be collected.
//
//  Finally, returns status if any error is collected.
//
//  Once error is returned process halted and exits out the loop
//  Only the first occurred error will be collected.
//
//  Warning:
//   - Getting an error will NOT halt the process and continue running.
//
//  To exit out on first error see AddAnyFunctions
func (it *Collection) AddAllFunctions(
	executeFunctions ...func() *errorwrapper.Wrapper,
) (isAddedAny, isNothingAdded bool) {
	if len(executeFunctions) == 0 {
		return false, true
	}

	for _, executeFunc := range executeFunctions {
		if executeFunc == nil {
			continue
		}

		errWrapper := executeFunc()

		if errWrapper == nil {
			continue
		}

		if errWrapper.HasError() {
			isAddedAny = true
			it.items = append(
				it.items,
				errWrapper)
		}
	}

	return isAddedAny, !isAddedAny
}

// IsSuccessFunc
//
//  returns true  if error is not returned  nor collected.
//  returns false if error is returned      and collected.
//
// Usages:
// - if IsSuccessFunc() { do this }
//
// See also IsFailedFunc
//  invert of IsFailedFunc
func (it *Collection) IsSuccessFunc(
	executeFunction func() *errorwrapper.Wrapper,
) (isSuccess bool) {
	if executeFunction == nil {
		return true
	}

	errWrapper := executeFunction()
	if errWrapper.HasError() {
		it.items = append(
			it.items,
			errWrapper)

		return false
	}

	return true
}

// IsFailedFunc
//
//  returns true  if error is returned     and collected.
//  returns false if error is not returned nor collected.
//
// Usages:
// - if IsFailedFunc() { return failed exit }
//
// See also IsSuccessFunc
//  invert of IsSuccessFunc
func (it *Collection) IsFailedFunc(
	executeFunction func() *errorwrapper.Wrapper,
) (isFailed bool) {
	return !it.IsSuccessFunc(executeFunction)
}

func (it *Collection) AddIsSuccessCollectorWithKey(
	key string,
	isSuccessCollector func(errorCollection *Collection) (isSuccess bool),
) *Collection {
	return it.AddIsSuccessCollectorWithReferences(
		isSuccessCollector,
		ref.Value{
			Variable: "Key",
			Value:    key,
		})
}

func (it *Collection) AddIsSuccessCollectorWithReferences(
	isSuccessCollector func(errorCollection *Collection) (isSuccess bool),
	references ...ref.Value,
) *Collection {
	if len(references) == 0 {
		return it.AddAllIsSuccessCollectorFunctions(isSuccessCollector)
	}

	if isSuccessCollector == nil {
		return it
	}

	emptyCollection := Empty()
	isSuccessCollector(emptyCollection)

	if emptyCollection.HasError() {
		it.AddWrapperWithAdditionalRefs(
			emptyCollection.First(),
			references...)
	}

	return it
}

func (it *Collection) AddAllIsSuccessCollectorFunctions(
	isSuccessCollectors ...func(errorCollection *Collection) (isSuccess bool),
) *Collection {
	if len(isSuccessCollectors) == 0 {
		return it
	}

	for _, isSuccessCollectorFunc := range isSuccessCollectors {
		if isSuccessCollectorFunc == nil {
			continue
		}

		isSuccessCollectorFunc(it)
	}

	return it
}

func (it *Collection) AddAllIsSuccessProcessCollectorFunc(
	dynamicInput coredynamic.Dynamic,
	isSuccessProcessCollector func(
	dynamicIn coredynamic.Dynamic,
	errorCollection *Collection,
) (isSuccess bool),
) *Collection {
	isSuccessProcessCollector(dynamicInput, it)

	return it
}

func (it *Collection) AddRef(
	errType errtype.Variation,
	err error,
	varName string,
	val interface{},
) *Collection {
	if err == nil {
		return it
	}

	errW := errorwrapper.
		NewErrorPlusMsgUsingAllParamsPtr(
			defaultSkipInternal,
			errType,
			true,
			err,
			constants.EmptyString,
			refs.NewDirectItem(varName, val))

	if errW.IsEmpty() {
		return it
	}

	it.items = append(
		it.items,
		errW,
	)

	return it
}

func (it *Collection) AddRefsCollection(
	errType errtype.Variation,
	refsCollection *refs.Collection,
	err error,
) *Collection {
	if err == nil {
		return it
	}

	errW := errorwrapper.
		NewErrorPlusMsgUsingAllParamsPtr(
			defaultSkipInternal,
			errType,
			true,
			err,
			constants.EmptyString,
			refsCollection)

	if errW.IsEmpty() {
		return it
	}

	it.items = append(
		it.items,
		errW,
	)

	return it
}

func (it *Collection) GetErrorWrapperWithoutHeader(
	errType errtype.Variation,
) *errorwrapper.Wrapper {
	if it == nil || it.IsEmpty() {
		return nil
	}

	toStrings := it.ToStrings(false, false)
	toString := strings.Join(toStrings, constants.NewLineUnix)

	return errnew.Type.Error(errType, errors.New(toString))
}

func (it *Collection) ConcatNewOptions(
	isCloneAsWellOnEmpty bool,
	errCollections ...*Collection,
) *Collection {
	if len(errCollections) == 0 {
		return NewUsingErrorWrappersPtr(
			isCloneAsWellOnEmpty,
			it.items,
		)
	}

	newCollection := NewUsingErrorWrappersPtr(
		isCloneAsWellOnEmpty,
		it.items,
	)

	return newCollection.AddCollectionsPtr(&errCollections)
}

func (it *Collection) ConcatNewClone(
	errCollections ...*Collection,
) *Collection {
	return it.ConcatNewOptions(
		true,
		errCollections...)
}

// ConcatNew
//
//  no clone copies to new list
func (it *Collection) ConcatNew(
	errCollections ...*Collection,
) *Collection {
	return it.ConcatNewOptions(
		false,
		errCollections...)
}

func (it *Collection) AddCollections(
	errCollections ...*Collection,
) *Collection {
	if errCollections == nil {
		return it
	}

	for _, errCollection := range errCollections {
		if errCollection == nil || errCollection.IsEmpty() {
			continue
		}

		it.AddWrappersPtr(&errCollection.items)
	}

	return it
}

func (it *Collection) AddCollectionsPtr(
	errCollections *[]*Collection,
) *Collection {
	if errCollections == nil {
		return it
	}

	for _, errCollection := range *errCollections {
		if errCollection == nil || errCollection.IsEmpty() {
			continue
		}

		it.AddWrappersPtr(&errCollection.items)
	}

	return it
}

// AddWrapper Skip on empty
func (it *Collection) AddWrapper(
	errorWrapper errorwrapper.Wrapper,
) *Collection {
	if errorWrapper.IsEmpty() {
		return it
	}

	it.items = append(
		it.items,
		&errorWrapper,
	)

	return it
}

func (it *Collection) IsCollected(
	errorWrapper *errorwrapper.Wrapper,
) (isCollected bool) {
	if errorWrapper.IsEmpty() {
		return false
	}

	it.items = append(
		it.items,
		errorWrapper,
	)

	return true
}

func (it *Collection) IsAnyWrappersCollected(
	errorWrappers ...*errorwrapper.Wrapper,
) (isCollected bool) {
	if len(errorWrappers) == 0 {
		return false
	}

	for _, errW := range errorWrappers {
		if errW.IsEmpty() {
			continue
		}

		it.items = append(
			it.items,
			errW,
		)

		isCollected = true
	}

	return isCollected
}

// IsNoneCollected
//
//  refers to IsSuccess, no errors found
//
//  Invert of IsAnyWrappersCollected
//
//  alias of IsEmptyAll, IsSuccessAll
//
// Usages:
//  - if IsNoneCollected() { do success }
func (it *Collection) IsNoneCollected(
	errorWrappers ...*errorwrapper.Wrapper,
) (isNoneCollected bool) {
	if len(errorWrappers) == 0 {
		return true
	}

	return !it.IsAnyWrappersCollected(
		errorWrappers...)
}

// IsSuccessAll
//
//  refers to IsSuccess, no errors found
//
//  Invert of IsAnyWrappersCollected
//
//  alias of IsEmptyAll, IsSuccessAll
//
// Usages:
//  - if IsSuccessAll() { do success }
func (it *Collection) IsSuccessAll(
	errorWrappers ...*errorwrapper.Wrapper,
) (isSuccessAll bool) {
	if len(errorWrappers) == 0 {
		return true
	}

	return !it.IsAnyWrappersCollected(
		errorWrappers...)
}

// IsFailedAny
//
//  refers to if any errors found and collected
//
//  alias of  IsAnyWrappersCollected
//  invert of IsEmptyAll, IsSuccessAll
//
// Usages:
//  - if IsFailedAny() { return exit }
func (it *Collection) IsFailedAny(
	errorWrappers ...*errorwrapper.Wrapper,
) (isSuccessAll bool) {
	if len(errorWrappers) == 0 {
		return true
	}

	return !it.IsAnyWrappersCollected(
		errorWrappers...)
}

func (it *Collection) CollectNull(
	anyItem interface{},
) *Collection {
	errWp := errnew.Null.UsingStackSkip(
		defaultSkipInternal,
		"",
		anyItem,
	)

	return it.AddWrapperPtr(errWp)
}

func (it *Collection) CollectNullMany(
	anyItems ...interface{},
) *Collection {
	errWp := errnew.Null.ManyByCheckingUsingStackSkip(
		defaultSkipInternal,
		anyItems...,
	)

	return it.AddWrapperPtr(errWp)
}

func (it *Collection) NullSimple(
	anyItem interface{},
) *Collection {
	errWp := errnew.Null.UsingStackSkip(
		defaultSkipInternal,
		"",
		anyItem,
	)

	return it.AddWrapperPtr(errWp)
}

func (it *Collection) Null(
	message string,
	anyItem interface{},
) *Collection {
	errWp := errnew.Null.UsingStackSkip(
		defaultSkipInternal,
		message,
		anyItem,
	)

	return it.AddWrapperPtr(errWp)
}

// IsAnyOfNull
//
//  returns true if any is null and also will be collected as error
//  Usages : errnew.Null .ManyByCheckingUsingStackSkip
func (it *Collection) IsAnyOfNull(
	anyItems ...interface{},
) (isNoneCollected bool) {
	if len(anyItems) == 0 {
		panic("must provide testing element!")
	}

	errWp := errnew.Null.ManyByCheckingUsingStackSkip(
		codestack.Skip1,
		anyItems...)

	return it.IsCollected(errWp)
}

// IsAnyOfNullMessage
//
//  returns true if any is null and also will be collected as error
//  Usages : errnew.Null .ManyWithMessage
func (it *Collection) IsAnyOfNullMessage(
	message string,
	anyItems ...interface{},
) (isNoneCollected bool) {
	if len(anyItems) == 0 {
		panic("must provide testing element!")
	}

	errWp := errnew.Null.ManyWithMessage(
		message,
		anyItems...)

	return it.IsCollected(errWp)
}

// IsNullItem
//
//  returns true if the item is null and
//  new error will be created and collected
//
//  Usages : errnew.Null .UsingStackSkip
func (it *Collection) IsNullItem(
	anyItem interface{},
) (isNoneCollected bool) {
	errWp := errnew.Null.UsingStackSkip(
		defaultSkipInternal,
		"",
		anyItem,
	)

	return it.IsCollected(errWp)
}

// IsNullItemMessage
//
//  returns true if the item is null and
//  new error will be created and collected
//
//  Usages : errnew.Null .UsingStackSkip
func (it *Collection) IsNullItemMessage(
	msg string,
	anyItem interface{},
) (isNoneCollected bool) {
	errWp := errnew.Null.UsingStackSkip(
		defaultSkipInternal,
		msg,
		anyItem,
	)

	return it.IsCollected(errWp)
}

// Append
//
//  Skip on empty or nil
//
//  alias for AddWrapperPtr
func (it *Collection) Append(
	errorWrapper *errorwrapper.Wrapper,
) *Collection {
	return it.AddWrapperPtr(errorWrapper)
}

// AddWrapperPtr Skip on empty or nil
func (it *Collection) AddWrapperPtr(
	errorWrapper *errorwrapper.Wrapper,
) *Collection {
	if errorWrapper.IsEmpty() {
		return it
	}

	it.items = append(
		it.items,
		errorWrapper,
	)

	return it
}

// IsCollectedErrorWrapperOrFunc
//
//  On existing error, errWrapperFunc will NOT be executed.
//
//  If existingErrWrapper doesn't contain error
//  then return status based on errWrapperFunc
//
//  alias for AddExistingErrorWrapperOrFunc
func (it *Collection) IsCollectedErrorWrapperOrFunc(
	existingErrWrapper *errorwrapper.Wrapper,
	errWrapperFunc func() *errorwrapper.Wrapper,
) (isAddedAny bool) {
	if existingErrWrapper.IsEmpty() && errWrapperFunc == nil {
		return false
	}

	if existingErrWrapper.HasError() {
		return it.IsCollected(existingErrWrapper)
	}

	return it.IsCollected(errWrapperFunc())
}

// AddExistingErrorWrapperOrFunc
//
//  On existing error, errWrapperFunc will NOT be executed.
//
//  If existingErrWrapper doesn't contain error
//  then return status based on errWrapperFunc
//
//  alias for IsCollectedErrorWrapperOrFunc
func (it *Collection) AddExistingErrorWrapperOrFunc(
	existingErrWrapper *errorwrapper.Wrapper,
	errWrapperFunc func() *errorwrapper.Wrapper,
) (isAddedAny bool) {
	if existingErrWrapper.IsEmpty() && errWrapperFunc == nil {
		return false
	}

	if existingErrWrapper.HasError() {
		return it.IsCollected(existingErrWrapper)
	}

	return it.IsCollected(errWrapperFunc())
}

func (it *Collection) AddWrapperPlusFuncPtr(
	errorWrapper *errorwrapper.Wrapper,
	errWrapperFunc func() *errorwrapper.Wrapper,
) (isAddedAny, isNothingAdded bool) {
	isAddedAny = false
	isNothingAdded = true

	if errorWrapper.HasError() {
		isAddedAny = true
		isNothingAdded = false
	}

	it.AddWrapperPtr(errorWrapper)
	nextWrapper := errWrapperFunc()

	if nextWrapper.HasError() {
		isAddedAny = true
		isNothingAdded = false

		it.AddWrapperPtr(nextWrapper)
	}

	return isAddedAny, isNothingAdded
}

// AddWrappers Skip on empty or nil
func (it *Collection) AddWrappers(
	errWrappers ...*errorwrapper.Wrapper,
) *Collection {
	for _, errW := range errWrappers {
		if errW == nil || errW.IsEmpty() {
			continue
		}

		it.items = append(
			it.items,
			errW,
		)
	}

	return it
}

func (it *Collection) AddTypeRefQuick(
	errType errtype.Variation,
	referencesValues ...interface{},
) *Collection {
	errWrapper := errorwrapper.TypeReferenceQuick(
		defaultSkipInternal,
		errType,
		referencesValues...,
	)

	return it.AddWrapperPtr(errWrapper)
}

// AddWrappersPtr Skip on empty or nil
func (it *Collection) AddWrappersPtr(
	errWrappers *[]*errorwrapper.Wrapper,
) *Collection {
	if errWrappers == nil {
		return it
	}

	for _, errW := range *errWrappers {
		if errW == nil || errW.IsEmpty() {
			continue
		}

		it.items = append(
			it.items,
			errW,
		)
	}

	return it
}

// AddTypeError Skip on empty or nil
func (it *Collection) AddTypeError(
	variation errtype.Variation,
	err error,
) *Collection {
	if err == nil {
		return it
	}

	it.items = append(
		it.items,
		errorwrapper.NewUsingError(
			defaultSkipInternal,
			variation,
			err),
	)

	return it
}

func (it *Collection) Add(variation errtype.Variation) *Collection {
	it.items = append(
		it.items,
		errorwrapper.NewPtr(variation),
	)

	return it
}

// AddUsingMsg if additionalMessage empty nothing added
func (it *Collection) AddUsingMsg(
	variation errtype.Variation,
	message string,
) *Collection {
	if message == "" {
		return it
	}

	it.items = append(
		it.items,
		errorwrapper.NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			variation,
			message),
	)

	return it
}

// AddPathIssueMessages if no error then doesn't add to the collection.
func (it *Collection) AddPathIssueMessages(
	variation errtype.Variation,
	pathLocation string,
	messages ...string,
) *Collection {
	errWrapper := errorwrapper.NewPathMessages(
		defaultSkipInternal,
		variation,
		pathLocation,
		messages...)

	it.items = append(
		it.items,
		errWrapper)

	return it
}

// AddPathIssue if no error then doesn't add to the collection.
func (it *Collection) AddPathIssue(
	variation errtype.Variation,
	err error,
	pathLocation string,
) *Collection {
	if err == nil {
		return it
	}

	errWrapper := errorwrapper.NewPath(
		defaultSkipInternal,
		variation,
		err,
		pathLocation)

	it.items = append(
		it.items,
		errWrapper)

	return it
}

func (it *Collection) AddOnlyRefs(
	variation errtype.Variation,
	refValues ...ref.Value,
) *Collection {
	if len(refValues) == 0 {
		return it
	}

	errWrapper := errorwrapper.NewOnlyRefs(
		defaultSkipInternal,
		variation,
		refValues...)

	it.items = append(
		it.items,
		errWrapper)

	return it
}

func (it *Collection) AddRef1Msg(
	variation errtype.Variation,
	msg string,
	var1 string,
	val1 interface{},
) *Collection {
	errWrapper := errorwrapper.NewRef1Msg(
		defaultSkipInternal,
		variation,
		msg,
		var1,
		val1,
	)

	it.items = append(
		it.items,
		errWrapper)

	return it
}

func (it *Collection) AddRef2Msg(
	variation errtype.Variation,
	msg string,
	var1 string,
	val1 interface{},
	var2 string,
	val2 interface{},
) *Collection {
	errWrapper := errorwrapper.NewRef2Msg(
		defaultSkipInternal,
		variation,
		msg,
		var1,
		val1,
		var2,
		val2,
	)

	it.items = append(
		it.items,
		errWrapper)

	return it
}

func (it *Collection) AddJsonResultAsDeserializedWrapper(
	jsonResult *corejson.Result,
) (isAdded bool) {
	if jsonResult == nil {
		return false
	}

	errWp, parsedErr := errnew.
		DeserializeTo.JsonResultToWrapper(
		jsonResult)

	it.AddWrappers(errWp, parsedErr)

	return errWp.HasError() || parsedErr.HasError()
}

func (it *Collection) AddExpectation(
	variation errtype.Variation,
	title string,
	expecting, actual interface{},
) *Collection {
	errWrapper := errnew.WasExpecting(
		variation,
		title,
		expecting,
		actual)

	it.items = append(
		it.items,
		errWrapper)

	return it
}

func (it *Collection) WasExpecting(
	variation errtype.Variation,
	title string,
	expecting, actual interface{},
) *Collection {
	errWrapper := errnew.WasExpecting(
		variation,
		title,
		expecting,
		actual)

	it.items = append(
		it.items,
		errWrapper)

	return it
}

func (it *Collection) AddRefs(
	variation errtype.Variation,
	err error,
	refValues ...ref.Value,
) *Collection {
	if len(refValues) == 0 {
		return it
	}

	errWrapper := errorwrapper.NewRefs(
		defaultSkipInternal,
		variation,
		err,
		refValues...)

	it.items = append(
		it.items,
		errWrapper)

	return it
}

func (it *Collection) AddWrapperWithAdditionalRefs(
	wrapper *errorwrapper.Wrapper,
	refValues ...ref.Value,
) *Collection {
	if len(refValues) == 0 {
		it.items = append(
			it.items,
			wrapper)

		return it
	}

	errWrapper := errorwrapper.NewUsingWrapper(
		defaultSkipInternal,
		wrapper,
		refValues...)

	it.items = append(
		it.items,
		errWrapper)

	return it
}

func (it *Collection) AddRefTwo(
	variation errtype.Variation,
	varName1 string, val1 interface{},
	varName2 string, val2 interface{},
	messages ...string,
) *Collection {
	if len(messages) == 0 {
		return it
	}

	var err error

	msgCompiled := strings.Join(messages, errorwrapper.MessagesJoiner)

	err = errors.New(msgCompiled)

	errWrapper := errorwrapper.NewRef2(
		defaultSkipInternal,
		variation,
		err,
		varName1, val1,
		varName2, val2)

	it.items = append(
		it.items,
		errWrapper)

	return it
}

// AddRefOne if messages empty then nothing added
func (it *Collection) AddRefOne(
	variation errtype.Variation,
	var1 string, val1 interface{},
) *Collection {
	errWrapper := errorwrapper.NewOnlyRefs(
		defaultSkipInternal,
		variation,
		ref.Value{
			Value:    val1,
			Variable: var1,
		})

	it.items = append(
		it.items,
		errWrapper)

	return it
}

// AddUsingMessages
//
// Same category multiple errorwrapper.New gets created for each message.
func (it *Collection) AddUsingMessages(
	variation errtype.Variation,
	messages ...string,
) *Collection {
	if messages == nil {
		return it
	}

	for _, message := range messages {
		errMessagePtr := errorwrapper.NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			variation,
			message)

		if errMessagePtr.IsEmpty() {
			continue
		}

		it.items = append(
			it.items,
			errMessagePtr,
		)
	}

	return it
}

// AddUsingMessagesUsingStackSkip
//
// Same category multiple errorwrapper.New gets created for each message.
func (it *Collection) AddUsingMessagesUsingStackSkip(
	skipStartStackIndex int,
	variation errtype.Variation,
	messages ...string,
) *Collection {
	if messages == nil {
		return it
	}

	for _, message := range messages {
		errMessagePtr := errorwrapper.NewMsgDisplayErrorNoReference(
			skipStartStackIndex+defaultSkipInternal,
			variation,
			message)

		if errMessagePtr.IsEmpty() {
			continue
		}

		it.items = append(
			it.items,
			errMessagePtr,
		)
	}

	return it
}

// Adds
//
//  for each variation errorwrapper.New will be created and stored.
func (it *Collection) Adds(
	variations ...errtype.Variation,
) *Collection {
	if variations == nil {
		return it
	}

	for _, v := range variations {
		it.items = append(
			it.items,
			errnew.Type.DefaultUsingStackSkip(
				defaultSkipInternal,
				v),
		)
	}

	return it
}

// AddsUsingStackSkip
//
//  for each variation errorwrapper.New
//  will be created and stored.
func (it *Collection) AddsUsingStackSkip(
	stackStartIndex int,
	variations ...errtype.Variation,
) *Collection {
	if variations == nil {
		return it
	}

	for _, v := range variations {
		it.items = append(
			it.items,
			errnew.Type.UsingStackSkip(
				stackStartIndex+defaultSkipInternal,
				v),
		)
	}

	return it
}

func (it *Collection) HasError() bool {
	return it != nil && it.items != nil && len(it.items) > 0
}

func (it *Collection) IsEmpty() bool {
	return it == nil ||
		len(it.items) == 0
}

func (it *Collection) Length() int {
	if it == nil || it.items == nil {
		return 0
	}

	return len(it.items)
}

func (it *Collection) ToString(
	isIncludeStakeTraces,
	isIncludeHeader bool,
) string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	lines := it.ToStrings(
		isIncludeStakeTraces,
		isIncludeHeader)

	return strings.Join(
		lines,
		constants.NewLineUnix)
}

func (it *Collection) ToStrings(
	isIncludeStakeTraces,
	isIncludeHeader bool,
) []string {
	if it.IsEmpty() {
		return []string{}
	}

	if isIncludeHeader {
		return it.Strings(isIncludeStakeTraces)
	}

	return it.StringsIf(isIncludeStakeTraces)
}

func (it *Collection) Strings(isIncludeStakeTraces bool) []string {
	if it == nil {
		return []string{}
	}

	length := it.Length()
	stringsCollection := make([]string, length+2)
	stringsCollection[coreindexes.First] =
		fmt.Sprintf(
			messages.ErrorWrappersCollectionHeaderFormat,
			length)

	if isIncludeStakeTraces {
		for i, v := range it.items {
			stringsCollection[i+2] = constants.Hyphen +
				constants.Space +
				v.FullStringWithTraces()
		}

		return stringsCollection
	}

	for i, v := range it.items {
		stringsCollection[i+2] = constants.Hyphen +
			constants.Space +
			v.FullString()
	}

	return stringsCollection
}

func (it Collection) String() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return strings.Join(
		it.Strings(false),
		constants.NewLineUnix)
}

func (it Collection) StringIf(isIncludeTraces bool) string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return strings.Join(
		it.Strings(isIncludeTraces),
		constants.NewLineUnix)
}

func (it *Collection) StringStackTracesWithoutHeader() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return strings.Join(
		it.FullStringsWithTraces(),
		constants.NewLineUnix)
}

func (it *Collection) DisplayStringWithTraces() string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	return strings.Join(
		it.FullStringsWithTraces(),
		DisplayJoiner)
}

func (it *Collection) DisplayStringWithLimitTraces(limit int) string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	if limit <= constants.TakeAllMinusOne {
		return it.DisplayStringWithTraces()
	}

	return strings.Join(
		it.FullStringsWithLimitTraces(limit),
		DisplayJoiner)
}

func (it *Collection) LogDisplayStringWithLimitTraces(limit int) {
	if it.IsEmpty() {
		return
	}

	display := it.DisplayStringWithLimitTraces(limit)

	fmt.Print(display)
}

func (it *Collection) StringWithoutHeader() string {
	if it == nil {
		return constants.EmptyString
	}

	return strings.Join(
		it.FullStrings(),
		constants.NewLineUnix)
}

func (it *Collection) List() []*errorwrapper.Wrapper {
	return it.items
}

func (it *Collection) Items() []*errorwrapper.Wrapper {
	return it.items
}

func (it *Collection) StringsWithoutHeader() []string {
	return it.FullStrings()
}

func (it *Collection) LinesIf(
	isIncludeReferences bool,
) []string {
	if isIncludeReferences {
		return it.FullStrings()
	}

	return it.StringsWithoutReferencePlusHeader()
}

func (it *Collection) StringsWithoutReferencePlusHeader() []string {
	if it.IsEmpty() {
		return []string{}
	}

	list := make([]string, it.Length())

	for i, wrapper := range it.items {
		list[i] = wrapper.FullStringWithoutReferences()
	}

	return list
}

func (it *Collection) StringsIf(isIncludeStakeTraces bool) []string {
	if it.IsEmpty() {
		return []string{}
	}

	if isIncludeStakeTraces {
		return it.FullStringsWithTraces()
	}

	return it.FullStrings()
}

func (it *Collection) FullStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	list := make([]string, it.Length())

	for i, wrapper := range it.items {
		list[i] = wrapper.FullString()
	}

	return list
}

func (it *Collection) FullStringsWithTraces() []string {
	if it.IsEmpty() {
		return []string{}
	}

	list := make([]string, it.Length())

	for i, wrapper := range it.items {
		list[i] = wrapper.FullStringWithTraces()
	}

	return list
}

func (it *Collection) FullStringsWithLimitTraces(limit int) []string {
	if it.IsEmpty() {
		return []string{}
	}

	list := make([]string, it.Length())

	for i, wrapper := range it.items {
		list[i] = wrapper.FullStringWithLimitTraces(limit)
	}

	return list
}

func (it *Collection) StringsUsingFmt(fmtFunc errorwrapper.StringFmtFunc) []string {
	if it.IsEmpty() {
		return []string{}
	}

	list := make([]string, it.Length())

	for i, wrapper := range it.items {
		list[i] = fmtFunc(wrapper)
	}

	return list
}

func (it *Collection) StringJoinUsingFmt(
	joiner string,
	fmtFunc errorwrapper.StringFmtFunc,
) string {
	if it.IsEmpty() {
		return constants.EmptyString
	}

	list := it.StringsUsingFmt(fmtFunc)

	return strings.Join(list, joiner)
}

func (it *Collection) Errors() []error {
	if it.IsEmpty() {
		return []error{}
	}

	list := make([]error, it.Length())

	for i, wrapper := range it.items {
		list[i] = wrapper.Error()
	}

	return list
}

func (it *Collection) CompiledErrors() []error {
	if it.IsEmpty() {
		return []error{}
	}

	list := make([]error, it.Length())

	for i, wrapper := range it.items {
		list[i] = wrapper.CompiledError()
	}

	return list
}

func (it *Collection) CompiledErrorsWithStackTraces() []error {
	if it.IsEmpty() {
		return []error{}
	}

	list := make([]error, it.Length())

	for i, wrapper := range it.items {
		list[i] = wrapper.CompiledErrorWithStackTraces()
	}

	return list
}

func (it *Collection) GetTypeVariantStructs() []errtype.VariantStructure {
	if it.IsEmpty() {
		return []errtype.VariantStructure{}
	}

	list := make([]errtype.VariantStructure, it.Length())

	for i, wrapper := range it.items {
		list[i] = wrapper.GetTypeVariantStruct()
	}

	return list
}

func (it *Collection) ErrorTypes() []errtype.Variation {
	list := make([]errtype.Variation, it.Length())

	for i, wrapper := range it.items {
		list[i] = wrapper.Type()
	}

	return list
}

func (it *Collection) IsSuccess() bool {
	return it.IsEmpty()
}

func (it *Collection) IsValid() bool {
	return it.IsEmpty()
}

func (it *Collection) IsFailed() bool {
	return it.HasError()
}

func (it *Collection) ItemsNonPtr() *[]errorwrapper.Wrapper {
	list := make([]errorwrapper.Wrapper, it.Length())

	for i, wrapper := range it.items {
		list[i] = *wrapper
	}

	return &list
}

func (it *Collection) GetAsError() error {
	if it.IsEmpty() {
		return nil
	}

	return errors.New(it.DisplayStringWithTraces())
}

func (it *Collection) GetAsErrorWrapper() errorwrapper.Wrapper {
	return *it.GetAsErrorWrapperPtr()
}

func (it *Collection) GetAsErrorWrapperUsingMessageType(
	errType errtype.Variation,
	message string,
) *errorwrapper.Wrapper {
	if it.IsEmpty() {
		return nil
	}

	finalMessage := finalMergedErrorPrefix +
		message + constants.DefaultLine +
		it.DisplayStringWithTraces()

	return errnew.Messages.ManyUsingStackSkip(
		codestack.Skip1,
		errType,
		finalMessage)
}

func (it *Collection) GetAsErrorWrapperUsingType(
	errType errtype.Variation,
) *errorwrapper.Wrapper {
	if it.IsEmpty() {
		return nil
	}

	message := finalMergedErrorPrefix +
		it.DisplayStringWithTraces()

	return errnew.Messages.ManyUsingStackSkip(
		codestack.Skip1,
		errType,
		message)
}

func (it *Collection) GetAsErrorWrapperPtr() *errorwrapper.Wrapper {
	if it.IsEmpty() {
		return nil
	}

	return it.GetAsErrorWrapperUsingType(errtype.CompileErrors)
}

// Log Skip if no error. Or else log prints.
func (it *Collection) Log() {
	if it.IsEmpty() {
		return
	}

	log.Print(it.String())
}

func (it *Collection) LogWithTraces() {
	if it.IsEmpty() {
		return
	}

	log.Print(it.DisplayStringWithTraces())
}

// LogFatal Skip if no error. Or else log fatal.
func (it *Collection) LogFatal() {
	if it.IsEmpty() {
		return
	}

	log.Fatal(it.DisplayStringWithTraces())
}

// HandleError Skip if no error. Or else panics.
func (it *Collection) HandleError() {
	if it.IsEmpty() {
		return
	}

	panic(it.DisplayStringWithTraces())
}

// HandleWithMsg Skip if no error.
func (it *Collection) HandleWithMsg(msg string) {
	if it.IsEmpty() {
		return
	}

	panic(msg + constants.NewLine + it.String())
}

func (it *Collection) MutexCollection() *MutexCollection {
	if it == nil {
		return MutexEmpty()
	}

	return &MutexCollection{
		internalCollection: it,
	}
}

// Clear
//
// Resets collection
func (it *Collection) Clear() {
	if it.IsEmpty() {
		return
	}

	if it.items == nil {
		return
	}

	tempItems := it.items
	clearFunc := func() {
		for i := range tempItems {
			tempItems[i].Dispose()
			tempItems[i] = nil
		}
	}

	go clearFunc()
	it.items = []*errorwrapper.Wrapper{}
}

// Dispose
//
//  After dispose nothing will work,
//  everything be removed from memory.
func (it *Collection) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.items = nil
}

func (it Collection) JsonModel() []*errorwrapper.Wrapper {
	if it.IsEmpty() {
		return nil
	}

	return it.items
}

func (it Collection) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it Collection) MarshalJSON() ([]byte, error) {
	return serializer(it.JsonModel())
}

func (it *Collection) UnmarshalJSON(data []byte) error {
	var errorWrappers []*errorwrapper.Wrapper
	err := usingBytesDeserializer(
		data, &errorWrappers)

	if err == nil {
		it.items = errorWrappers
	}

	return err
}

func (it Collection) Json() corejson.Result {
	return corejson.New(it)
}

func (it *Collection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(*it)
}

//goland:noinspection GoLinterLocal
func (it *Collection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Collection, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return it, defaulterr.UnmarshallingFailedDueToNilOrEmpty
	}

	err := json.Unmarshal(jsonResult.Bytes, &it)

	if err != nil {
		return it, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *Collection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Collection {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *Collection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *Collection) ValidationErrUsingTextValidator(
	validator *corevalidator.TextValidator,
	params *corevalidator.ValidatorParamsBase,
) *errorwrapper.Wrapper {
	err := validator.VerifyDetailError(
		params,
		it.ToString(false, false))

	if err != nil {
		return errorwrapper.NewUsingError(
			defaultSkipInternal,
			errtype.ValidationMismatch,
			err)
	}

	return nil
}

func (it *Collection) ValidationErrUsingSliceValidator(
	sliceValidator *corevalidator.SliceValidator,
	params *corevalidator.ValidatorParamsBase,
) *errorwrapper.Wrapper {
	sliceValidator.SetActual(
		it.FullStrings())

	err := sliceValidator.AllVerifyError(
		params,
	)

	if err != nil {
		return errorwrapper.NewUsingError(
			defaultSkipInternal,
			errtype.ValidationMismatch,
			err)
	}

	return nil
}

func (it *Collection) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Collection) AsJsoner() corejson.Jsoner {
	return it
}

func (it *Collection) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *Collection) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}

func (it *Collection) AsBasicSlicerContractsBinder() coreinterface.BasicSlicerContractsBinder {
	return it
}

func (it *Collection) AsDynamicLinq() coreinterface.DynamicLinq {
	return it
}

func (it *Collection) CollectCmdOnce(cmdOnce *errcmd.CmdOnce) *Collection {
	if cmdOnce == nil {
		return it
	}

	it.AddWrapperPtr(cmdOnce.CompiledErrorWrapper())

	return it
}

func (it *Collection) CollectCmdOnceCollection(
	cmdOnceCollection *errcmd.CmdOnceCollection,
) *Collection {
	if cmdOnceCollection == nil {
		return it
	}

	it.AddWrappers(cmdOnceCollection.ErrorWrappers()...)

	return it
}

func (it Collection) AsBaseErrorWrapperCollectionDefiner() errcoreinf.BaseErrorWrapperCollectionDefiner {
	return &it
}

func (it Collection) ToPtr() *Collection {
	return &it
}

func (it Collection) Deserialize(toPtr interface{}) error {
	return it.JsonPtr().Deserialize(toPtr)
}

func (it *Collection) AsBasicSliceContractsBinder() coreinterface.BasicSlicerContractsBinder {
	return it
}
