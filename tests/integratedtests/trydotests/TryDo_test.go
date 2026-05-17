package trydotests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	"github.com/alimtvnetwork/errorwrapper-v3/trydo"
	. "github.com/smartystreets/goconvey/convey"
)

// Test_Block_Do verifies try/catch/finally semantics.
func Test_Block_Do(t *testing.T) {
	Convey("Try runs without catch or finally", t, func() {
		var ran bool
		trydo.Block{
			Try: func() { ran = true },
		}.Do()
		So(ran, ShouldBeTrue)
	})

	Convey("Catch receives a thrown exception", t, func() {
		var caught trydo.Exception
		trydo.Block{
			Try: func() { trydo.Throw(errors.New("boom")) },
			Catch: func(e trydo.Exception) { caught = e },
		}.Do()
		So(caught, ShouldNotBeNil)
		err, ok := caught.(error)
		So(ok, ShouldBeTrue)
		So(err.Error(), ShouldEqual, "boom")
	})

	Convey("Finally runs even when Try panics", t, func() {
		var finallyRan bool
		var caught trydo.Exception
		trydo.Block{
			Try:     func() { trydo.Throw("panic") },
			Catch:   func(e trydo.Exception) { caught = e },
			Finally: func() { finallyRan = true },
		}.Do()
		So(finallyRan, ShouldBeTrue)
		So(caught, ShouldNotBeNil)
	})

	Convey("Finally runs when Try succeeds", t, func() {
		var finallyRan bool
		trydo.Block{
			Try:     func() {},
			Finally: func() { finallyRan = true },
		}.Do()
		So(finallyRan, ShouldBeTrue)
	})

	Convey("No catch re-panics through Block.Do", t, func() {
		So(func() {
			trydo.Block{
				Try: func() { panic("no catch") },
			}.Do()
		}, ShouldPanic)
	})
}

// Test_WrapPanic catches arbitrary panics.
func Test_WrapPanic(t *testing.T) {
	Convey("WrapPanic captures a string panic", t, func() {
		ex := trydo.WrapPanic(func() { panic("string panic") })
		So(ex, ShouldNotBeNil)
		So(ex, ShouldEqual, "string panic")
	})

	Convey("WrapPanic captures an error panic", t, func() {
		ex := trydo.WrapPanic(func() { panic(errors.New("err panic")) })
		So(ex, ShouldNotBeNil)
		err, ok := ex.(error)
		So(ok, ShouldBeTrue)
		So(err.Error(), ShouldEqual, "err panic")
	})

	Convey("WrapPanic returns nil when no panic", t, func() {
		ex := trydo.WrapPanic(func() {})
		So(ex, ShouldBeNil)
	})
}

// Test_GetErrorWrapperWrappedPanic extracts Wrapper from panic.
func Test_GetErrorWrapperWrappedPanic(t *testing.T) {
	Convey("returns the wrapper when a Wrapper is panicked", t, func() {
		w := errnew.Type.Error(errtype.Generic, errors.New("panic wrapper"))
		got := trydo.GetErrorWrapperWrappedPanic(func() { panic(w) })
		So(got, ShouldNotBeNil)
		So(got.HasError(), ShouldBeTrue)
		So(got.Message(), ShouldContainSubstring, "panic wrapper")
	})

	Convey("returns empty wrapper when panic is not a Wrapper", t, func() {
		got := trydo.GetErrorWrapperWrappedPanic(func() { panic("not a wrapper") })
		So(got, ShouldNotBeNil)
		So(got.HasError(), ShouldBeFalse)
	})

	Convey("returns empty wrapper when no panic", t, func() {
		got := trydo.GetErrorWrapperWrappedPanic(func() {})
		So(got, ShouldNotBeNil)
		So(got.HasError(), ShouldBeFalse)
	})
}

// Test_GetErrorWrapperCollectionWrappedPanic extracts Collection from panic.
func Test_GetErrorWrapperCollectionWrappedPanic(t *testing.T) {
	Convey("returns the collection when a Collection is panicked", t, func() {
		c := errwrappers.NewEmpty()
		c.AddTypeError(errtype.Generic, errors.New("coll panic"))
		got := trydo.GetErrorWrapperCollectionWrappedPanic(func() { panic(c) })
		So(got, ShouldNotBeNil)
		So(got.StateCounter().HasChanges(), ShouldBeTrue)
	})

	Convey("returns empty collection when panic is not a Collection", t, func() {
		got := trydo.GetErrorWrapperCollectionWrappedPanic(func() { panic("not a coll") })
		So(got, ShouldNotBeNil)
		So(got.StateCounter().HasChanges(), ShouldBeFalse)
	})

	Convey("returns empty collection when no panic", t, func() {
		got := trydo.GetErrorWrapperCollectionWrappedPanic(func() {})
		So(got, ShouldNotBeNil)
		So(got.StateCounter().HasChanges(), ShouldBeFalse)
	})
}

// Test_WrapPanicToBaseErrorCollection coerces any panic into a BaseErrorOrCollectionWrapper.
func Test_WrapPanicToBaseErrorCollection(t *testing.T) {
	Convey("captures a string panic as an error in the collection", t, func() {
		bec := trydo.WrapPanicToBaseErrorCollection(func() { panic("stringy") })
		So(bec, ShouldNotBeNil)
		So(bec.GetAsErrorWrapperPtr().HasError(), ShouldBeTrue)
	})

	Convey("captures an error panic as an error in the collection", t, func() {
		bec := trydo.WrapPanicToBaseErrorCollection(func() { panic(errors.New("plain err")) })
		So(bec, ShouldNotBeNil)
		So(bec.GetAsErrorWrapperPtr().HasError(), ShouldBeTrue)
	})

	Convey("captures a Wrapper panic directly into the collection", t, func() {
		w := errnew.Type.Error(errtype.NotFound, errors.New("wrapped"))
		bec := trydo.WrapPanicToBaseErrorCollection(func() { panic(w) })
		So(bec, ShouldNotBeNil)
		So(bec.GetAsErrorWrapperPtr().HasError(), ShouldBeTrue)
	})

	Convey("captures a Collection panic directly", t, func() {
		c := errwrappers.NewEmpty()
		c.AddTypeError(errtype.Generic, errors.New("from coll"))
		bec := trydo.WrapPanicToBaseErrorCollection(func() { panic(c) })
		So(bec, ShouldNotBeNil)
		So(bec.GetAsErrorWrapperPtr().HasError(), ShouldBeTrue)
	})

	Convey("returns empty collection when no panic", t, func() {
		bec := trydo.WrapPanicToBaseErrorCollection(func() {})
		So(bec, ShouldNotBeNil)
		So(bec.GetAsErrorWrapperPtr().HasError(), ShouldBeFalse)
	})
}
