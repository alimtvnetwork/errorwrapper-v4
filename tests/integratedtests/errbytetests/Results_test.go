package errbytetests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errbyte"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Errbyte_Results_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errbyte.Results
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
		r := &errbyte.Results{Values: []byte{}, ErrorWrapper: nil}
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasAnyItem(), ShouldBeFalse)
		So(r.Length(), ShouldEqual, 0)
	})

	Convey("Non-empty slice without error", t, func() {
		r := &errbyte.Results{Values: []byte{65, 66}, ErrorWrapper: nil}
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
		r := &errbyte.Results{Values: []byte{65, 66}, ErrorWrapper: nil}
		r.Clear()
		So(r.Length(), ShouldEqual, 0)
	})

	Convey("Dispose nils values and wrapper", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "x")
		r := &errbyte.Results{Values: []byte{65, 66}, ErrorWrapper: w}
		r.Dispose()
		So(r.Values, ShouldBeNil)
	})

	Convey("Value with error", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := &errbyte.Results{Values: []byte{65, 66}, ErrorWrapper: w}
		So(r.HasError(), ShouldBeTrue)
		So(r.IsEmptyError(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.ErrorWrapperInf(), ShouldNotBeNil)
	})
}

func Test_Errbyte_Results_JSON(t *testing.T) {
	Convey("Json round-trip", t, func() {
		r := &errbyte.Results{Values: []byte{65, 66}, ErrorWrapper: nil}
		j := r.Json()
		So(j.HasError(), ShouldBeFalse)
		So(r.JsonPtr(), ShouldNotBeNil)
	})
}
