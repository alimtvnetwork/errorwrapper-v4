package errdefertests

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errdefer"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	. "github.com/smartystreets/goconvey/convey"
)

// Test_MoreCoverage_ErrorMessagesUsingFunc covers nil-func and func paths.
func Test_MoreCoverage_ErrorMessagesUsingFunc(t *testing.T) {
	Convey("nil errorFunc returns existing wrapper unchanged", t, func() {
		existing := errnew.Type.Error(errtype.Generic, errors.New("e1"))
		w := errdefer.ErrorMessagesUsingFunc(existing, errtype.Generic, nil, "m")
		So(w, ShouldEqual, existing)
	})
	Convey("errorFunc returning error produces non-empty wrapper", t, func() {
		w := errdefer.ErrorMessagesUsingFunc(nil, errtype.Generic, func() error {
			return errors.New("boom")
		}, "ctx")
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
	})
	Convey("errorFunc returning nil yields empty/merged wrapper", t, func() {
		w := errdefer.ErrorMessagesUsingFunc(nil, errtype.Generic, func() error { return nil })
		So(w == nil || !w.HasError(), ShouldBeTrue)
	})
}

// Test_MoreCoverage_ErrorMessagesUsingCollectionFunc
func Test_MoreCoverage_ErrorMessagesUsingCollectionFunc(t *testing.T) {
	Convey("nil errorFunc returns false without modifying collection", t, func() {
		coll := errwrappers.Empty()
		ok := errdefer.ErrorMessagesUsingCollectionFunc(coll, errtype.Generic, nil, "m")
		So(ok, ShouldBeFalse)
		So(coll.HasAnyItem(), ShouldBeFalse)
	})
	Convey("errorFunc returning error adds to collection", t, func() {
		coll := errwrappers.Empty()
		ok := errdefer.ErrorMessagesUsingCollectionFunc(coll, errtype.Generic,
			func() error { return errors.New("boom") }, "msg")
		So(ok, ShouldBeFalse)
		So(coll.HasAnyItem(), ShouldBeTrue)
	})
	Convey("errorFunc returning nil reports success", t, func() {
		coll := errwrappers.Empty()
		ok := errdefer.ErrorMessagesUsingCollectionFunc(coll, errtype.Generic,
			func() error { return nil })
		So(ok, ShouldBeTrue)
	})
}

// Test_MoreCoverage_ErrorUsingCollection
func Test_MoreCoverage_ErrorUsingCollection(t *testing.T) {
	Convey("adds typed error to collection", t, func() {
		coll := errwrappers.Empty()
		errdefer.ErrorUsingCollection(coll, errtype.Generic, errors.New("x"))
		So(coll.HasAnyItem(), ShouldBeTrue)
	})
}

// Test_MoreCoverage_ErrorWithMessagesUsingCollection
func Test_MoreCoverage_ErrorWithMessagesUsingCollection(t *testing.T) {
	Convey("returns false when error is non-nil and records it", t, func() {
		coll := errwrappers.Empty()
		ok := errdefer.ErrorWithMessagesUsingCollection(coll, errtype.Generic, errors.New("oops"), "m1", "m2")
		So(ok, ShouldBeFalse)
		So(coll.HasAnyItem(), ShouldBeTrue)
	})
	Convey("returns true when error is nil", t, func() {
		coll := errwrappers.Empty()
		ok := errdefer.ErrorWithMessagesUsingCollection(coll, errtype.Generic, nil, "m")
		So(ok, ShouldBeTrue)
	})
}

// Test_MoreCoverage_ErrorWrapperFuncUsingCollection
func Test_MoreCoverage_ErrorWrapperFuncUsingCollection(t *testing.T) {
	Convey("non-empty wrapper from func is collected and reports false", t, func() {
		coll := errwrappers.Empty()
		ok := errdefer.ErrorWrapperFuncUsingCollection(coll, func() *errorwrapper.Wrapper {
			return errnew.Type.Error(errtype.Generic, errors.New("z"))
		})
		So(ok, ShouldBeFalse)
		So(coll.HasAnyItem(), ShouldBeTrue)
	})
	Convey("empty wrapper from func reports success", t, func() {
		coll := errwrappers.Empty()
		ok := errdefer.ErrorWrapperFuncUsingCollection(coll, func() *errorwrapper.Wrapper {
			return errorwrapper.StaticEmptyPtr
		})
		So(ok, ShouldBeTrue)
	})
}

// Test_MoreCoverage_CloseFile covers nil-file and real-file paths.
func Test_MoreCoverage_CloseFile(t *testing.T) {
	Convey("nil file returns existing wrapper unchanged", t, func() {
		existing := errnew.Type.Error(errtype.Generic, errors.New("e"))
		w := errdefer.CloseFile("loc", existing, nil)
		So(w, ShouldEqual, existing)
	})
	Convey("closing a real file twice surfaces a closing error", t, func() {
		dir := t.TempDir()
		path := filepath.Join(dir, "f.txt")
		f, err := os.Create(path)
		So(err, ShouldBeNil)
		So(f.Close(), ShouldBeNil)
		w := errdefer.CloseFile(path, nil, f) // second close → error
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
	})
	Convey("closing an open file succeeds", t, func() {
		dir := t.TempDir()
		path := filepath.Join(dir, "g.txt")
		f, err := os.Create(path)
		So(err, ShouldBeNil)
		w := errdefer.CloseFile(path, nil, f)
		So(w == nil || !w.HasError(), ShouldBeTrue)
	})
}

// Test_MoreCoverage_CloseFileUsingErrorCollection
func Test_MoreCoverage_CloseFileUsingErrorCollection(t *testing.T) {
	Convey("nil file returns false default and leaves collection empty", t, func() {
		coll := errwrappers.Empty()
		ok := errdefer.CloseFileUsingErrorCollection("loc", coll, nil)
		So(ok, ShouldBeFalse)
		So(coll.HasAnyItem(), ShouldBeFalse)
	})
	Convey("closing an open file reports success", t, func() {
		dir := t.TempDir()
		path := filepath.Join(dir, "h.txt")
		f, err := os.Create(path)
		So(err, ShouldBeNil)
		coll := errwrappers.Empty()
		ok := errdefer.CloseFileUsingErrorCollection(path, coll, f)
		So(ok, ShouldBeTrue)
	})
	Convey("double-close records a failure", t, func() {
		dir := t.TempDir()
		path := filepath.Join(dir, "i.txt")
		f, err := os.Create(path)
		So(err, ShouldBeNil)
		So(f.Close(), ShouldBeNil)
		coll := errwrappers.Empty()
		ok := errdefer.CloseFileUsingErrorCollection(path, coll, f)
		So(ok, ShouldBeFalse)
		So(coll.HasAnyItem(), ShouldBeTrue)
	})
}

// Test_MoreCoverage_CloseFileHandlerFunc / LoggerFunc / HandlePanic
func Test_MoreCoverage_CloseFileHandlerVariants(t *testing.T) {
	Convey("HandlerFunc invoked with closing-error wrapper on double close", t, func() {
		dir := t.TempDir()
		path := filepath.Join(dir, "h1.txt")
		f, err := os.Create(path)
		So(err, ShouldBeNil)
		So(f.Close(), ShouldBeNil)
		var got *errorwrapper.Wrapper
		errdefer.CloseFileHandlerFunc(path, nil, f, func(w *errorwrapper.Wrapper) { got = w })
		So(got, ShouldNotBeNil)
		So(got.HasError(), ShouldBeTrue)
	})
	Convey("LoggerFunc invoked when closing a fresh file", t, func() {
		dir := t.TempDir()
		path := filepath.Join(dir, "h2.txt")
		f, err := os.Create(path)
		So(err, ShouldBeNil)
		called := false
		errdefer.CloseFileLoggerFunc(path, nil, f, func(_ *errorwrapper.Wrapper) { called = true })
		So(called, ShouldBeTrue)
	})
	Convey("HandlePanic returns silently for nil file", t, func() {
		So(func() { errdefer.CloseFileHandlePanic("loc", nil, nil) }, ShouldNotPanic)
	})
	Convey("HandlePanic returns silently when close succeeds", t, func() {
		dir := t.TempDir()
		path := filepath.Join(dir, "h3.txt")
		f, err := os.Create(path)
		So(err, ShouldBeNil)
		So(func() { errdefer.CloseFileHandlePanic(path, nil, f) }, ShouldNotPanic)
	})
}
