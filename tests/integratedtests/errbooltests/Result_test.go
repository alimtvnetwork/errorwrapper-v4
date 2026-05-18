package errbooltests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errbool"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrBool_Result_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errbool.Result
		So(r.IsAnyNull(), ShouldBeTrue)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
		So(r.IsEmptyError(), ShouldBeTrue)
		So(r.IsTrue(), ShouldBeFalse)
		So(r.IsFalse(), ShouldBeTrue)
		So(r.IsApplicable(), ShouldBeFalse)
		So(r.IsValid(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.IsFailed(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.HasSafeItems(), ShouldBeFalse)
		So(r.Int(), ShouldEqual, 0)
		So(r.Byte(), ShouldEqual, 0)
		So(r.SafeString(), ShouldEqual, "false")
		So(func() { r.Dispose() }, ShouldNotPanic)
	})

	Convey("True value without error", t, func() {
		r := &errbool.Result{Value: true, ErrorWrapper: nil}
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasError(), ShouldBeFalse)
		So(r.IsEmptyError(), ShouldBeTrue)
		So(r.IsTrue(), ShouldBeTrue)
		So(r.IsFalse(), ShouldBeFalse)
		So(r.IsApplicable(), ShouldBeTrue)
		So(r.IsValid(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.IsFailed(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeTrue)
		So(r.Int(), ShouldEqual, 1)
		So(r.Byte(), ShouldEqual, 1)
		So(r.SafeString(), ShouldEqual, "true")
		So(r.String(), ShouldEqual, "true")
	})

	Convey("False value without error", t, func() {
		r := &errbool.Result{Value: false, ErrorWrapper: nil}
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.IsTrue(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
	})

	Convey("True value with error wrapper", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "bad")
		r := &errbool.Result{Value: true, ErrorWrapper: w}
		So(r.HasError(), ShouldBeTrue)
		So(r.IsEmptyError(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.IsFailed(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.HasSafeItems(), ShouldBeFalse)
		So(r.ErrorWrapperInf(), ShouldNotBeNil)
	})
}

func Test_ErrBool_Result_JSON(t *testing.T) {
	Convey("Json round-trip", t, func() {
		r := &errbool.Result{Value: true, ErrorWrapper: nil}
		j := r.Json()
		So(j.HasError(), ShouldBeFalse)
		So(r.JsonPtr(), ShouldNotBeNil)
	})
}

func Test_ErrBool_Results_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errbool.Results
		So(r.IsAnyNull(), ShouldBeTrue)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
		So(r.IsEmptyError(), ShouldBeTrue)
		So(r.Length(), ShouldEqual, 0)
		So(r.HasAnyItem(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeFalse)
		So(r.IsValid(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.IsFailed(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.SafeString(), ShouldEqual, "")
		So(func() { r.Dispose() }, ShouldNotPanic)
		So(func() { r.Clear() }, ShouldNotPanic)
	})

	Convey("Empty slice", t, func() {
		r := &errbool.Results{Values: []bool{}, ErrorWrapper: nil}
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasAnyItem(), ShouldBeFalse)
		So(r.Length(), ShouldEqual, 0)
	})

	Convey("Non-empty slice without error", t, func() {
		r := &errbool.Results{Values: []bool{true, false}, ErrorWrapper: nil}
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasAnyItem(), ShouldBeTrue)
		So(r.Length(), ShouldEqual, 2)
		So(r.HasError(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeFalse)
		So(r.SafeValues(), ShouldResemble, []bool{true, false})
	})

	Convey("Clear resets values", t, func() {
		r := &errbool.Results{Values: []bool{true, false}, ErrorWrapper: nil}
		r.Clear()
		So(r.Length(), ShouldEqual, 0)
	})

	Convey("Dispose nils values and wrapper", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "x")
		r := &errbool.Results{Values: []bool{true}, ErrorWrapper: w}
		r.Dispose()
		So(r.Values, ShouldBeNil)
	})
}
