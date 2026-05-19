package errinttests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errint"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrInt_Result_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errint.Result
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
		r := &errint.Result{Value: 0, ErrorWrapper: nil}
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
	})

	Convey("Positive value without error", t, func() {
		r := &errint.Result{Value: 42, ErrorWrapper: nil}
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasError(), ShouldBeFalse)
		So(r.Int(), ShouldEqual, 42)
		So(r.Byte(), ShouldEqual, 42)
		So(r.String(), ShouldEqual, "42")
		So(r.SafeString(), ShouldEqual, "42")
		So(r.IsValid(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.IsFailed(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeTrue)
	})

	Convey("Value with error", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := &errint.Result{Value: 10, ErrorWrapper: w}
		So(r.HasError(), ShouldBeTrue)
		So(r.IsEmptyError(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.ErrorWrapperInf(), ShouldNotBeNil)
	})

	Convey("ValidRange checks", t, func() {
		r := &errint.Result{Value: 50, ErrorWrapper: nil}
		So(r.IsValidRange(10, 100), ShouldBeTrue)
		So(r.IsValidRange(60, 100), ShouldBeFalse)
	})

	Convey("SafeValidRange requires empty error", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := &errint.Result{Value: 50, ErrorWrapper: w}
		So(r.IsSafeValidRange(10, 100), ShouldBeFalse)
	})

	Convey("Byte clamps to max uint8", t, func() {
		r := &errint.Result{Value: 300, ErrorWrapper: nil}
		So(r.Byte(), ShouldEqual, 255)
	})

	Convey("Byte clamps negative to zero", t, func() {
		r := &errint.Result{Value: -5, ErrorWrapper: nil}
		So(r.Byte(), ShouldEqual, 0)
	})

	Convey("Dispose clears value", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "x")
		r := &errint.Result{Value: 10, ErrorWrapper: w}
		r.Dispose()
		So(r.Value, ShouldEqual, 0)
	})
}

func Test_ErrInt_Result_JSON(t *testing.T) {
	Convey("Json round-trip", t, func() {
		r := &errint.Result{Value: 99, ErrorWrapper: nil}
		j := r.Json()
		So(j.HasError(), ShouldBeFalse)
		So(r.JsonPtr(), ShouldNotBeNil)
	})
}
