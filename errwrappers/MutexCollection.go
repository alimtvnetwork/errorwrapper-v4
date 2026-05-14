package errwrappers

import (
	"sync"

	"github.com/alimtvnetwork/errorwrapper-v3"
)

type MutexCollection struct {
	internalCollection *Collection
	sync.Mutex
}

// AddWrapperPtr
//
// Skip on empty or nil
func (it *MutexCollection) AddWrapperPtr(
	errorWrapper *errorwrapper.Wrapper,
) *MutexCollection {
	it.Lock()
	defer it.Unlock()

	it.internalCollection.AddWrapperPtr(errorWrapper)

	return it
}

// AddWrappers
//
// Skip on empty or nil
func (it *MutexCollection) AddWrappers(
	errWrappers ...*errorwrapper.Wrapper,
) *MutexCollection {
	it.Lock()
	defer it.Unlock()

	if errWrappers == nil {
		return it
	}

	it.internalCollection.AddWrappersPtr(&errWrappers)

	return it
}

func (it *MutexCollection) HasError() bool {
	it.Lock()
	defer it.Unlock()

	if it == nil {
		return false
	}

	return !it.internalCollection.HasError()
}

func (it *MutexCollection) IsEmpty() bool {
	it.Lock()
	defer it.Unlock()

	if it == nil {
		return true
	}

	return it.internalCollection.IsEmpty()
}

func (it *MutexCollection) Length() int {
	it.Lock()
	defer it.Unlock()

	if it == nil {
		return 0
	}

	return it.internalCollection.Length()
}

func (it *MutexCollection) ToStringLock(
	isIncludeStakeTraces,
	isIncludeHeader bool,
) string {
	it.Lock()
	defer it.Unlock()

	return it.internalCollection.ToString(
		isIncludeStakeTraces,
		isIncludeHeader)
}

func (it *MutexCollection) ToStringsLock(
	isIncludeStakeTraces,
	isIncludeHeader bool,
) []string {
	it.Lock()
	defer it.Unlock()

	return it.internalCollection.ToStrings(isIncludeStakeTraces, isIncludeHeader)
}

func (it MutexCollection) String() string {
	it.Lock()
	defer it.Unlock()

	return it.internalCollection.String()
}

func (it *MutexCollection) DisplayStringWithTraces() string {
	it.Lock()
	defer it.Unlock()

	return it.internalCollection.DisplayStringWithTraces()
}

func (it *MutexCollection) DisplayStringWithLimitTracesLock(limit int) string {
	it.Lock()
	defer it.Unlock()

	return it.internalCollection.DisplayStringWithLimitTraces(limit)
}

func (it *MutexCollection) LogDisplayStringWithLimitTracesLock(limit int) {
	it.Lock()
	defer it.Unlock()

	if it == nil {
		return
	}

	it.internalCollection.LogDisplayStringWithLimitTraces(limit)
}

func (it *MutexCollection) StringsIfLock(isIncludeStakeTraces bool) []string {
	it.Lock()
	defer it.Unlock()

	return it.internalCollection.StringsIf(isIncludeStakeTraces)
}

func (it *MutexCollection) FullStrings() []string {
	it.Lock()
	defer it.Unlock()

	return it.internalCollection.FullStrings()
}

func (it *MutexCollection) FullStringsWithTraces() []string {
	it.Lock()
	defer it.Unlock()

	return it.internalCollection.FullStringsWithTraces()
}

func (it *MutexCollection) FullStringsWithLimitTracesLock(limit int) []string {
	it.Lock()
	defer it.Unlock()

	return it.internalCollection.FullStringsWithLimitTraces(limit)
}

func (it *MutexCollection) StringsUsingFmtLock(
	fmtFunc errorwrapper.StringFmtFunc,
) []string {
	it.Lock()
	defer it.Unlock()

	return it.internalCollection.StringsUsingFmt(fmtFunc)
}

func (it *MutexCollection) StringJoinUsingFmtLock(
	joiner string,
	fmtFunc errorwrapper.StringFmtFunc,
) string {
	it.Lock()
	defer it.Unlock()

	return it.internalCollection.StringJoinUsingFmt(joiner, fmtFunc)
}

func (it *MutexCollection) Errors() []error {
	it.Lock()
	defer it.Unlock()

	return it.internalCollection.Errors()
}

func (it *MutexCollection) CompiledErrors() []error {
	it.Lock()
	defer it.Unlock()

	return it.internalCollection.CompiledErrors()
}

func (it *MutexCollection) IsSuccess() bool {
	it.Lock()
	defer it.Unlock()

	return it.internalCollection.IsEmpty()
}

func (it *MutexCollection) IsValid() bool {
	it.Lock()
	defer it.Unlock()

	return it.internalCollection.IsEmpty()
}

func (it *MutexCollection) IsFailed() bool {
	it.Lock()
	defer it.Unlock()

	return it.internalCollection.HasError()
}

func (it *MutexCollection) GetAsError() error {
	it.Lock()
	defer it.Unlock()

	if it == nil {
		return nil
	}

	return it.internalCollection.GetAsError()
}

func (it *MutexCollection) GetAsErrorWrapperPtr() *errorwrapper.Wrapper {
	it.Lock()
	defer it.Unlock()

	if it == nil {
		return nil
	}

	return it.internalCollection.GetAsErrorWrapperPtr()
}

// Log
//
// Log Skip if no error. Or else log prints.
func (it *MutexCollection) Log() {
	it.Lock()
	defer it.Unlock()

	if it == nil {
		return
	}

	it.internalCollection.Log()
}

func (it *MutexCollection) LogWithTraces() {
	it.Lock()
	defer it.Unlock()

	if it == nil {
		return
	}

	it.internalCollection.LogWithTraces()
}

// LogFatal Skip if no error. Or else log fatal.
func (it *MutexCollection) LogFatal() {
	it.Lock()
	defer it.Unlock()

	if it == nil {
		return
	}

	it.internalCollection.LogFatal()
}

// HandleError
//
// Skip if no error. Or else panics.
func (it *MutexCollection) HandleError() {
	it.Lock()
	defer it.Unlock()

	if it == nil {
		return
	}

	it.internalCollection.HandleError()
}

func (it *MutexCollection) Clear() {
	it.Lock()
	defer it.Unlock()

	if it == nil {
		return
	}

	it.internalCollection.Clear()
}

// Dispose After dispose nothing will work, everything be removed from memory.
func (it *MutexCollection) Dispose() {
	it.Lock()
	defer it.Unlock()

	if it == nil {
		return
	}

	it.internalCollection.Dispose()
	it.internalCollection = nil
}

func (it *MutexCollection) StateCounter() MutexCollectionStateCounter {
	return NewMutexStateCount(it)
}

func (it *MutexCollection) Collection() *Collection {
	it.Lock()
	defer it.Unlock()

	if it == nil {
		return nil
	}

	return it.internalCollection
}
