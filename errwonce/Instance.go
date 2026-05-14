package errwonce

import (
	"encoding/json"
	"errors"

	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/converters"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type Instance struct {
	innerData       *errorwrapper.Wrapper
	initializerFunc func() *errorwrapper.Wrapper
	isInitialized   bool
}

func New(initializerFunc func() *errorwrapper.Wrapper) Instance {
	return Instance{
		initializerFunc: initializerFunc,
	}
}

func NewPtr(initializerFunc func() *errorwrapper.Wrapper) *Instance {
	return &Instance{
		initializerFunc: initializerFunc,
	}
}

func NewPtrUsingErrFunc(errType errtype.Variation, initializerFunc func() error) *Instance {
	newFunc := func() *errorwrapper.Wrapper {
		err := initializerFunc()

		if err != nil {
			return errnew.Type.Error(errType, err)
		}

		return nil
	}

	return NewPtr(newFunc)
}

func (it Instance) MarshalJSON() ([]byte, error) {
	if it.IsNullOrEmpty() {
		return json.Marshal("")
	}

	return json.Marshal(it.Value().Error())
}

func (it *Instance) UnmarshalJSON(data []byte) error {
	it.isInitialized = true
	var errWrapper errorwrapper.Wrapper

	err := json.Unmarshal(data, &errWrapper)
	it.innerData = &errWrapper

	return err
}

func (it Instance) HasError() bool {
	return !it.IsNullOrEmpty()
}

func (it Instance) IsNull() bool {
	return it.Value() == nil
}

func (it Instance) IsNullOrEmpty() bool {
	err := it.Value()

	return err == nil || err.IsEmpty()
}

func (it *Instance) Message() string {
	if it.IsNull() {
		return constants.EmptyString
	}

	return it.Value().ErrorString()
}

func (it *Instance) IsMessageEqual(msg string) bool {
	if it.IsNull() {
		return false
	}

	return it.Message() == msg
}

// HandleError with panic if error exist or else skip
//
// Skip if no error type (NoError).
func (it *Instance) HandleError() {
	if it.IsNullOrEmpty() {
		return
	}

	panic(it.Value())
}

// HandleErrorWith by concatenating message and then panic if error exist or else skip
//
// Skip if no error type (NoError).
func (it *Instance) HandleErrorWith(messages ...string) {
	if it.IsNullOrEmpty() {
		return
	}

	panic(it.ConcatNewString(messages...))
}

func (it *Instance) ConcatNewString(messages ...string) string {
	additionalMessages :=
		converters.StringsToCsv(
			false,
			messages...,
		)

	if it.IsNullOrEmpty() {
		return additionalMessages
	}

	return it.String() +
		constants.NewLineUnix +
		additionalMessages
}

func (it *Instance) ConcatNewWrapper(
	startSkipStackIndex int,
	messages ...string,
) *errorwrapper.Wrapper {
	return it.
		Value().
		ConcatNew().
		MessagesUsingStackSkip(
			startSkipStackIndex+codestack.Skip1,
			messages...)
}

func (it *Instance) ConcatNewWrapperUsingError(
	startSkipStackIndex int,
	err error,
) *errorwrapper.Wrapper {
	return it.
		Value().
		ConcatNew().
		ErrorUsingStackSkip(
			startSkipStackIndex+codestack.Skip1,
			err)
}

func (it *Instance) ConcatNewWrapperUsingAnother(
	startSkipStackIndex int,
	errWrapper *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	return it.
		Value().
		ConcatNew().
		WrapperUsingStackSkip(
			startSkipStackIndex+codestack.Skip1,
			errWrapper)
}

func (it *Instance) ConcatNewErrors(
	startSkipStackIndex int,
	errs ...error,
) *errorwrapper.Wrapper {
	return it.
		Value().
		ConcatNew().
		ErrorsUsingStackSkip(
			startSkipStackIndex+codestack.Skip1,
			errs...)
}

func (it *Instance) ConcatNew(messages ...string) error {
	return errors.New(it.ConcatNewString(messages...))
}

func (it *Instance) Value() *errorwrapper.Wrapper {
	if it.isInitialized {
		return it.innerData
	}

	it.innerData = it.initializerFunc()
	it.isInitialized = true

	return it.innerData
}

func (it Instance) IsEmpty() bool {
	return it.Value().IsEmpty()
}

func (it Instance) IsSuccess() bool {
	return it.Value().IsEmpty()
}

func (it Instance) IsFailed() bool {
	return it.Value().HasError()
}

func (it Instance) HasReferences() bool {
	return it.Value().HasReferences()
}

func (it Instance) FullString() string {
	return it.Value().FullString()
}

func (it Instance) FullStringWithTraces() string {
	return it.Value().FullStringWithTraces()
}

func (it Instance) Json() corejson.Result {
	return it.Value().Json()
}

func (it Instance) JsonPtr() *corejson.Result {
	return it.Value().JsonPtr()
}

func (it Instance) Serialize() ([]byte, error) {
	return it.Value().Serialize()
}

func (it Instance) SerializeMust() []byte {
	return it.Value().SerializeMust()
}

func (it Instance) ErrorWrapper() *errorwrapper.Wrapper {
	return it.Value()
}

func (it Instance) String() string {
	return it.Value().FullString()
}
