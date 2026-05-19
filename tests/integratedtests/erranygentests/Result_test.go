package erranygentests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/erranygen"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_GenericResult_Basics(t *testing.T) {
	Convey("Result[string] without wrapper, with isZero predicate", t, func() {
		r := erranygen.NewResult("hello", nil, func(s string) bool { return s == "" })

		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.HasError(), ShouldBeFalse)
		So(r.IsEmptyError(), ShouldBeTrue)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.IsFailed(), ShouldBeFalse)
		So(r.IsValid(), ShouldBeTrue)
		So(r.IsInvalid(), ShouldBeFalse)
	})

	Convey("Result[int] with isZero treats 0 as empty", t, func() {
		r := erranygen.NewResult(0, nil, func(v int) bool { return v == 0 })
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeFalse)
	})

	Convey("Without isZero predicate non-nil receiver is never empty", t, func() {
		r := erranygen.NewResult(0, nil, nil)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeTrue)
	})

	Convey("Nil receiver is empty / any-null", t, func() {
		var r *erranygen.Result[string]
		So(r.IsAnyNull(), ShouldBeTrue)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
		So(r.IsEmptyError(), ShouldBeTrue)
	})

	Convey("Wrapper-bearing result reports error", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad input")
		r := erranygen.NewResult("v", w, nil)
		So(r.HasError(), ShouldBeTrue)
		So(r.IsEmptyError(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.IsFailed(), ShouldBeTrue)
		So(r.ErrorWrapperInf(), ShouldNotBeNil)
	})
}

func Test_GenericResult_JSONAndDispose(t *testing.T) {
	Convey("Json / JsonPtr round-trip", t, func() {
		r := erranygen.NewResult("abc", nil, func(s string) bool { return s == "" })
		j := r.Json()
		So(j.HasError(), ShouldBeFalse)
		So(r.JsonPtr(), ShouldNotBeNil)

		var dst erranygen.Result[string]
		jp := r.JsonPtr()
		err := dst.JsonParseSelfInject(jp)
		So(err, ShouldBeNil)
		So(dst.Value, ShouldEqual, "abc")
	})

	Convey("Dispose clears value and wrapper", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "x")
		r := erranygen.NewResult("v", w, nil)
		r.Dispose()
		So(r.Value, ShouldEqual, "")
		So(r.HasError(), ShouldBeFalse)
	})

	Convey("Dispose on nil is a no-op", t, func() {
		var r *erranygen.Result[int]
		So(func() { r.Dispose() }, ShouldNotPanic)
	})
}
