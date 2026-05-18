package erranytests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errany"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrAny_Result_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errany.Result
		So(r.IsAnyNull(), ShouldBeTrue)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
		So(r.IsEmptyError(), ShouldBeTrue)
		So(r.IsValid(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.IsFailed(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.HasSafeItems(), ShouldBeFalse)
		So(r.Str(), ShouldEqual, "")
		So(r.Bool(), ShouldBeFalse)
		So(r.Int(), ShouldEqual, 0)
		So(r.Byte(), ShouldEqual, 0)
		So(r.Float32(), ShouldEqual, float32(0))
		So(r.Float64(), ShouldEqual, float64(0))
		So(r.SafeString(), ShouldEqual, "")
		So(func() { r.Dispose() }, ShouldNotPanic)
	})

	Convey("String value without error", t, func() {
		r := &errany.Result{Value: "hello", ErrorWrapper: nil}
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasError(), ShouldBeFalse)
		So(r.Str(), ShouldEqual, "hello")
		So(r.String(), ShouldEqual, "hello")
		So(r.SafeString(), ShouldEqual, "hello")
		So(r.IsValid(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.IsFailed(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeTrue)
	})

	Convey("Bool value", t, func() {
		r := &errany.Result{Value: true, ErrorWrapper: nil}
		So(r.Bool(), ShouldBeTrue)
		So(r.Int(), ShouldEqual, 0)
	})

	Convey("Int value", t, func() {
		r := &errany.Result{Value: 42, ErrorWrapper: nil}
		So(r.Int(), ShouldEqual, 42)
		So(r.Float32(), ShouldEqual, float32(0))
		So(r.Float64(), ShouldEqual, float64(0))
	})

	Convey("Byte value", t, func() {
		r := &errany.Result{Value: byte(65), ErrorWrapper: nil}
		So(r.Byte(), ShouldEqual, byte(65))
	})

	Convey("Float32 value", t, func() {
		r := &errany.Result{Value: float32(3.14), ErrorWrapper: nil}
		So(r.Float32(), ShouldEqual, float32(3.14))
	})

	Convey("Float64 value", t, func() {
		r := &errany.Result{Value: float64(2.718), ErrorWrapper: nil}
		So(r.Float64(), ShouldEqual, float64(2.718))
	})

	Convey("Zero numeric value is empty", t, func() {
		r := &errany.Result{Value: 0, ErrorWrapper: nil}
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("Value with error", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "bad")
		r := &errany.Result{Value: "x", ErrorWrapper: w}
		So(r.HasError(), ShouldBeTrue)
		So(r.IsEmptyError(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.ErrorWrapperInf(), ShouldNotBeNil)
	})

	Convey("Dispose clears value", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "x")
		r := &errany.Result{Value: "test", ErrorWrapper: w}
		r.Dispose()
		So(r.Value, ShouldBeNil)
	})
}

func Test_ErrAny_Result_JSON(t *testing.T) {
	Convey("Json round-trip", t, func() {
		r := &errany.Result{Value: "abc", ErrorWrapper: nil}
		j := r.Json()
		So(j.HasError(), ShouldBeFalse)
		So(r.JsonPtr(), ShouldNotBeNil)
	})
}
