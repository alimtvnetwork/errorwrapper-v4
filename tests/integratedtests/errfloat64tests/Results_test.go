package errfloat64tests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errfloat64"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Errfloat64_Results_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errfloat64.Results
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
		r := &errfloat64.Results{Values: []float64{}, ErrorWrapper: nil}
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasAnyItem(), ShouldBeFalse)
		So(r.Length(), ShouldEqual, 0)
	})

	Convey("Non-empty slice without error", t, func() {
		r := &errfloat64.Results{Values: []float64{1.1, 2.2}, ErrorWrapper: nil}
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasAnyItem(), ShouldBeTrue)
		So(r.Length(), ShouldBeGreaterThan, 0)
		So(r.HasError(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeFalse)
		So(r.SafeValues(), ShouldNotBeEmpty)
		So(r.String(), ShouldNotBeBlank)
	})

	Convey("Clear resets values", t, func() {
		r := &errfloat64.Results{Values: []float64{1.1, 2.2}, ErrorWrapper: nil}
		r.Clear()
		So(r.Length(), ShouldEqual, 0)
	})

	Convey("Dispose nils values and wrapper", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "x")
		r := &errfloat64.Results{Values: []float64{1.1, 2.2}, ErrorWrapper: w}
		r.Dispose()
		So(r.Values, ShouldBeNil)
	})

	Convey("Value with error", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "bad")
		r := &errfloat64.Results{Values: []float64{1.1, 2.2}, ErrorWrapper: w}
		So(r.HasError(), ShouldBeTrue)
		So(r.IsEmptyError(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.ErrorWrapperInf(), ShouldNotBeNil)
	})
}

func Test_Errfloat64_Results_JSON(t *testing.T) {
	Convey("Json round-trip", t, func() {
		r := &errfloat64.Results{Values: []float64{1.1, 2.2}, ErrorWrapper: nil}
		j := r.Json()
		So(j.HasError(), ShouldBeFalse)
		So(r.JsonPtr(), ShouldNotBeNil)
	})
}
