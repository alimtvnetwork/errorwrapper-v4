package errbytetests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errbyte"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrByte_Result_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errbyte.Result
		So(r.IsAnyNull(), ShouldBeTrue)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
		So(r.IsEmptyError(), ShouldBeTrue)
		So(r.IsMax(), ShouldBeFalse)
		So(r.IsValid(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.IsFailed(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.HasSafeItems(), ShouldBeFalse)
		So(r.Int(), ShouldEqual, 0)
		So(r.Float32(), ShouldEqual, float32(0))
		So(r.Float64(), ShouldEqual, float64(0))
		So(r.Bool(), ShouldBeFalse)
		So(r.SafeString(), ShouldEqual, "")
		So(r.NumberString(), ShouldEqual, "0")
		So(func() { r.Dispose() }, ShouldNotPanic)
	})

	Convey("Zero value without error", t, func() {
		r := &errbyte.Result{Value: 0, ErrorWrapper: nil}
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.Bool(), ShouldBeFalse)
		So(r.IsMax(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
	})

	Convey("Non-zero value without error", t, func() {
		r := &errbyte.Result{Value: 42, ErrorWrapper: nil}
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasError(), ShouldBeFalse)
		So(r.IsMax(), ShouldBeFalse)
		So(r.Bool(), ShouldBeTrue)
		So(r.Int(), ShouldEqual, 42)
		So(r.Float32(), ShouldEqual, float32(42))
		So(r.Float64(), ShouldEqual, float64(42))
		So(r.SafeString(), ShouldEqual, "*")
		So(r.NumberString(), ShouldEqual, "42")
		So(r.String(), ShouldEqual, "*")
		So(r.IsValid(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.IsFailed(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeTrue)
	})

	Convey("Max value", t, func() {
		r := &errbyte.Result{Value: 255, ErrorWrapper: nil}
		So(r.IsMax(), ShouldBeTrue)
	})

	Convey("Value with error", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := &errbyte.Result{Value: 10, ErrorWrapper: w}
		So(r.HasError(), ShouldBeTrue)
		So(r.IsEmptyError(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.ErrorWrapperInf(), ShouldNotBeNil)
	})

	Convey("ValidRange checks", t, func() {
		r := &errbyte.Result{Value: 50, ErrorWrapper: nil}
		So(r.IsValidRange(10, 100), ShouldBeTrue)
		So(r.IsValidRange(60, 100), ShouldBeFalse)
	})

	Convey("SafeValidRange requires empty error", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := &errbyte.Result{Value: 50, ErrorWrapper: w}
		So(r.IsSafeValidRange(10, 100), ShouldBeFalse)
	})

	Convey("Dispose clears value", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "x")
		r := &errbyte.Result{Value: 10, ErrorWrapper: w}
		r.Dispose()
		So(r.Value, ShouldEqual, 0)
	})
}

func Test_ErrByte_Result_JSON(t *testing.T) {
	Convey("Json round-trip", t, func() {
		r := &errbyte.Result{Value: 65, ErrorWrapper: nil}
		j := r.Json()
		So(j.HasError(), ShouldBeFalse)
		So(r.JsonPtr(), ShouldNotBeNil)
	})
}
