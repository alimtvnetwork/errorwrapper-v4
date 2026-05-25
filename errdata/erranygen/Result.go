// Package erranygen provides a generic, type-parameterized result container
// that captures the type-agnostic surface area shared by every package
// under errdata/ (errstr, errbool, errint, errfloat, errbyte, errjson,
// errany, errfloat64, errcasted).
//
// Status: ADDITIVE proof-of-concept for the Phase 5 generics refactor.
// Existing errdata/* packages are unchanged. This package is safe to
// adopt incrementally — type aliases or full migration of the legacy
// packages is a separate, unblocked decision.
//
// What it covers:
//   - Value + ErrorWrapper pairing
//   - All errorwrapper.ValueWithErrorWrapper interface methods that do
//     NOT depend on T's concrete type (HasError, IsEmptyError, Json,
//     Dispose-of-wrapper, IsAnyNull, ErrorWrapperInf, etc.)
//
// What it does NOT cover (must stay in concrete packages):
//   - Type-specific accessors (Bool, Int, Bytes, SplitLines, ValidValue,
//     IsEqualIgnoreCase, SimpleStringOnce, ...) — those depend on T.
//   - "Is empty value" semantics — what counts as empty depends on T,
//     so callers supply an optional IsZero predicate via NewResult.
package erranygen

import (
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v3"
)

// Result is the generic counterpart of errdata/<type>.Result.
//
//	r := erranygen.NewResult("hello", nil, nil)
//	r := erranygen.NewResult(42, wrapper, func(v int) bool { return v == 0 })
type Result[T any] struct {
	Value        T
	ErrorWrapper *errorwrapper.Wrapper

	// isZero reports whether Value should be considered "empty".
	// When nil, only the wrapper is consulted for emptiness.
	isZero func(T) bool
}

// NewResult constructs a Result. isZero may be nil.
func NewResult[T any](
	value T,
	wrapper *errorwrapper.Wrapper,
	isZero func(T) bool,
) *Result[T] {
	return &Result[T]{
		Value:        value,
		ErrorWrapper: wrapper,
		isZero:       isZero,
	}
}

// --- ValueWithErrorWrapper-style methods (type-agnostic subset) ---

func (it *Result[T]) IsAnyNull() bool { return it == nil }

func (it *Result[T]) HasError() bool {
	return it != nil && it.ErrorWrapper.HasError()
}

func (it *Result[T]) IsEmptyError() bool {
	return it == nil || it.ErrorWrapper.IsEmpty()
}

// IsEmpty reports emptiness based on the caller-supplied isZero
// predicate. When no predicate is provided, only nil-receiver is empty.
func (it *Result[T]) IsEmpty() bool {
	if it == nil {
		return true
	}
	if it.isZero != nil && it.isZero(it.Value) {
		return true
	}
	return false
}

func (it *Result[T]) HasIssuesOrEmpty() bool {
	return it.IsEmpty() || it.HasError()
}

func (it *Result[T]) HasSafeItems() bool {
	return !it.HasIssuesOrEmpty()
}

func (it *Result[T]) IsSuccess() bool { return it.HasSafeItems() }
func (it *Result[T]) IsFailed() bool  { return it.HasIssuesOrEmpty() }
func (it *Result[T]) IsValid() bool   { return it.HasSafeItems() }
func (it *Result[T]) IsInvalid() bool { return it.HasIssuesOrEmpty() }

func (it *Result[T]) Dispose() {
	if it == nil {
		return
	}
	var zero T
	it.Value = zero
	it.ErrorWrapper.Dispose()
}

func (it *Result[T]) ErrorWrapperInf() errorwrapper.ErrWrapper {
	if it == nil {
		return nil
	}
	return it.ErrorWrapper
}

// --- JSON surface ---

func (it Result[T]) Json() corejson.Result    { return corejson.New(it) }
func (it Result[T]) JsonPtr() *corejson.Result { return corejson.NewPtr(it) }
func (it Result[T]) JsonModelAny() interface{} { return it }

func (it *Result[T]) JsonParseSelfInject(jsonResult *corejson.Result) error {
	return jsonResult.Unmarshal(it)
}

func (it *Result[T]) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}
