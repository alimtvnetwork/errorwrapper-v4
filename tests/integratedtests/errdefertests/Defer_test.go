package errdefertests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errdefer"
	"github.com/alimtvnetwork/errorwrapper-v3/errfunc"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	. "github.com/smartystreets/goconvey/convey"
)

// Test_Error_Merge verifies basic defer error merging with nil and non-nil existing wrappers.
func Test_Error_Merge(t *testing.T) {
	Convey("Error with nil existing wrapper returns a non-empty wrapper", t, func() {
		w := errdefer.Error(nil, errtype.Generic, errors.New("deferred failure"))
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
		So(w.Message(), ShouldContainSubstring, "deferred failure")
	})

	Convey("Error with existing wrapper concatenates new error", t, func() {
		existing := errnew.Type.Error(errtype.NotFound, errors.New("first"))
		w := errdefer.Error(existing, errtype.Generic, errors.New("second"))
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
		So(w.Message(), ShouldContainSubstring, "first")
		So(w.Message(), ShouldContainSubstring, "second")
	})

	Convey("Error with nil err returns existing wrapper unchanged", t, func() {
		existing := errnew.Type.Error(errtype.NotFound, errors.New("first"))
		w := errdefer.Error(existing, errtype.Generic, nil)
		So(w, ShouldEqual, existing)
	})
}

// Test_ErrorWithMessages verifies defer error creation with extra messages.
func Test_ErrorWithMessages(t *testing.T) {
	Convey("ErrorWithMessages adds messages to the wrapper", t, func() {
		w := errdefer.ErrorWithMessages(nil, errtype.InvalidInput, errors.New("bad input"), "ctx", "field=name")
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
		msg := w.Message()
		So(msg, ShouldContainSubstring, "bad input")
		So(msg, ShouldContainSubstring, "ctx")
		So(msg, ShouldContainSubstring, "field=name")
	})

	Convey("ErrorWithMessages with nil err returns existing wrapper", t, func() {
		existing := errnew.Type.Error(errtype.NotFound, errors.New("first"))
		w := errdefer.ErrorWithMessages(existing, errtype.Generic, nil, "extra")
		So(w, ShouldEqual, existing)
	})
}

// Test_ErrorUsingFunc verifies the func-based defer helpers.
func Test_ErrorUsingFunc(t *testing.T) {
	Convey("ErrorUsingFunc with nil func returns existing wrapper", t, func() {
		existing := errnew.Type.Error(errtype.NotFound, errors.New("first"))
		w := errdefer.ErrorUsingFunc(existing, errtype.Generic, nil)
		So(w, ShouldEqual, existing)
	})

	Convey("ErrorUsingFunc with non-nil func captures returned error", t, func() {
		w := errdefer.ErrorUsingFunc(nil, errtype.CommandExecution, func() error {
			return errors.New("func error")
		})
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
		So(w.Message(), ShouldContainSubstring, "func error")
	})

	Convey("ErrorMessagesUsingFunc with func + messages builds wrapper", t, func() {
		w := errdefer.ErrorMessagesUsingFunc(nil, errtype.MappingFailed, func() error {
			return errors.New("map fail")
		}, "key", "value")
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
		msg := w.Message()
		So(msg, ShouldContainSubstring, "map fail")
		So(msg, ShouldContainSubstring, "key")
	})
}

// Test_ErrorUsingCollection verifies collection-based defer helpers.
func Test_ErrorUsingCollection(t *testing.T) {
	Convey("ErrorUsingCollection appends an error to the collection", t, func() {
		c := errwrappers.NewEmpty()
		errdefer.ErrorUsingCollection(c, errtype.Generic, errors.New("coll error"))
		So(c.StateCounter().HasChanges(), ShouldBeTrue)
	})

	Convey("ErrorWithMessagesUsingCollection returns true when err is nil", t, func() {
		c := errwrappers.NewEmpty()
		ok := errdefer.ErrorWithMessagesUsingCollection(c, errtype.Generic, nil, "msg")
		So(ok, ShouldBeTrue)
		So(c.StateCounter().HasChanges(), ShouldBeFalse)
	})

	Convey("ErrorWithMessagesUsingCollection returns false when err is non-nil", t, func() {
		c := errwrappers.NewEmpty()
		ok := errdefer.ErrorWithMessagesUsingCollection(c, errtype.Generic, errors.New("fail"), "msg")
		So(ok, ShouldBeFalse)
		So(c.StateCounter().HasChanges(), ShouldBeTrue)
	})

	Convey("ErrorMessagesUsingCollectionFunc with nil func does nothing", t, func() {
		c := errwrappers.NewEmpty()
		ok := errdefer.ErrorMessagesUsingCollectionFunc(c, errtype.Generic, nil, "msg")
		So(ok, ShouldBeFalse)
		So(c.StateCounter().HasChanges(), ShouldBeFalse)
	})

	Convey("ErrorMessagesUsingCollectionFunc with func returns true on success", t, func() {
		c := errwrappers.NewEmpty()
		ok := errdefer.ErrorMessagesUsingCollectionFunc(c, errtype.Generic, func() error {
			return nil
		}, "msg")
		So(ok, ShouldBeTrue)
		So(c.StateCounter().HasChanges(), ShouldBeFalse)
	})

	Convey("ErrorMessagesUsingCollectionFunc with func returns false on error", t, func() {
		c := errwrappers.NewEmpty()
		ok := errdefer.ErrorMessagesUsingCollectionFunc(c, errtype.Generic, func() error {
			return errors.New("boom")
		}, "msg")
		So(ok, ShouldBeFalse)
		So(c.StateCounter().HasChanges(), ShouldBeTrue)
	})
}

// Test_ErrorWrapperFunc verifies wrapper-func defer helpers.
func Test_ErrorWrapperFunc(t *testing.T) {
	Convey("ErrorWrapperFunc merges a wrapper produced by a func", t, func() {
		wf := errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.NotFound, errors.New("wf err"))
		})
		w := errdefer.ErrorWrapperFunc(nil, wf)
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
		So(w.Message(), ShouldContainSubstring, "wf err")
	})

	Convey("ErrorWrapperFuncUsingCollection returns true when wrapper is empty", t, func() {
		c := errwrappers.NewEmpty()
		wf := errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errorwrapper.EmptyPtr()
		})
		ok := errdefer.ErrorWrapperFuncUsingCollection(c, wf)
		So(ok, ShouldBeTrue)
		So(c.StateCounter().HasChanges(), ShouldBeFalse)
	})

	Convey("ErrorWrapperFuncUsingCollection returns false when wrapper has error", t, func() {
		c := errwrappers.NewEmpty()
		wf := errfunc.WrapperFunc(func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.Generic, errors.New("wf fail"))
		})
		ok := errdefer.ErrorWrapperFuncUsingCollection(c, wf)
		So(ok, ShouldBeFalse)
		So(c.StateCounter().HasChanges(), ShouldBeTrue)
	})
}

// Test_CloseFile verifies file-close defer helpers (no real OS file needed for nil path).
func Test_CloseFile(t *testing.T) {
	Convey("CloseFile with nil osFile returns existing wrapper", t, func() {
		existing := errnew.Type.Error(errtype.NotFound, errors.New("first"))
		w := errdefer.CloseFile("/tmp/x", existing, nil)
		So(w, ShouldEqual, existing)
	})

	Convey("CloseFileHandlePanic with nil osFile calls HandleError safely", t, func() {
		existing := errnew.Type.Error(errtype.NotFound, errors.New("first"))
		// Should not panic even with nil file
		So(func() {
			errdefer.CloseFileHandlePanic("/tmp/x", existing, nil)
		}, ShouldNotPanic)
	})

	Convey("CloseFileHandlerFunc with nil osFile passes existing to handler", t, func() {
		existing := errnew.Type.Error(errtype.NotFound, errors.New("first"))
		var received *errorwrapper.Wrapper
		errdefer.CloseFileHandlerFunc("/tmp/x", existing, nil, func(w *errorwrapper.Wrapper) {
			received = w
		})
		So(received, ShouldEqual, existing)
	})

	Convey("CloseFileLoggerFunc with nil osFile passes existing to logger", t, func() {
		existing := errnew.Type.Error(errtype.NotFound, errors.New("first"))
		var received *errorwrapper.Wrapper
		errdefer.CloseFileLoggerFunc("/tmp/x", existing, nil, func(w *errorwrapper.Wrapper) {
			received = w
		})
		So(received, ShouldEqual, existing)
	})

	Convey("CloseFileUsingErrorCollection with nil osFile does nothing", t, func() {
		c := errwrappers.NewEmpty()
		ok := errdefer.CloseFileUsingErrorCollection("/tmp/x", c, nil)
		So(ok, ShouldBeFalse)
		So(c.StateCounter().HasChanges(), ShouldBeFalse)
	})
}
