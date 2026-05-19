package errfloattests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errfloat"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrFloat_Result_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errfloat.Result
		So(r.IsAnyNull(), ShouldBeTrue)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
		So(r.IsEmptyError(), ShouldBeTrue)
		So(r.IsValid(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.IsFailed(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.HasSafeItems(), ShouldBeFalse)
		So(r.Int(), ShouldEqual, 0)
		So(r.Byte(), ShouldEqual, 0)
		So(r.SafeString(), ShouldEqual, "")
		So(func() { r.Dispose() }, ShouldNotPanic)
	})

	Convey("Zero value without error", t, func() {
		r := &errfloat.Result{Value: 0, ErrorWrapper: nil}
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
	})

	Convey("Positive value without error", t, func() {
		r := &errfloat.Result{Value: 3.14, ErrorWrapper: nil}
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasError(), ShouldBeFalse)
		So(r.Int(), ShouldEqual, 3)
		So(r.String(), ShouldNotBeBlank)
		So(r.SafeString(), ShouldNotBeBlank)
		So(r.IsValid(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.IsFailed(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeTrue)
	})

	Convey("Value with error", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := &errfloat.Result{Value: 1.5, ErrorWrapper: w}
		So(r.HasError(), ShouldBeTrue)
		So(r.IsEmptyError(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.ErrorWrapperInf(), ShouldNotBeNil)
	})

	Convey("ValidRange checks", t, func() {
		r := &errfloat.Result{Value: 5.5, ErrorWrapper: nil}
		So(r.IsValidRange(1.0, 10.0), ShouldBeTrue)
		So(r.IsValidRange(6.0, 10.0), ShouldBeFalse)
	})

	Convey("SafeValidRange requires empty error", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := &errfloat.Result{Value: 5.5, ErrorWrapper: w}
		So(r.IsSafeValidRange(1.0, 10.0), ShouldBeFalse)
	})

	Convey("Byte clamps to max uint8", t, func() {
		r := &errfloat.Result{Value: 300.0, ErrorWrapper: nil}
		So(r.Byte(), ShouldEqual, 255)
	})

	Convey("Byte clamps negative to zero", t, func() {
		r := &errfloat.Result{Value: -1.5, ErrorWrapper: nil}
		So(r.Byte(), ShouldEqual, 0)
	})

	Convey("Dispose clears value", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "x")
		r := &errfloat.Result{Value: 1.5, ErrorWrapper: w}
		r.Dispose()
		So(r.Value, ShouldEqual, 0)
	})
}

func Test_ErrFloat_Result_JSON(t *testing.T) {
	Convey("Json round-trip", t, func() {
		r := &errfloat.Result{Value: 2.71, ErrorWrapper: nil}
		j := r.Json()
		So(j.HasError(), ShouldBeFalse)
		So(r.JsonPtr(), ShouldNotBeNil)
	})
}
