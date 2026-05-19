package errfloattests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errfloat"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrFloat_ResultsWithErrorCollection_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errfloat.ResultsWithErrorCollection
		So(r.IsAnyNull(), ShouldBeTrue)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
		So(r.IsEmptyError(), ShouldBeTrue)
		So(r.IsEmptyItems(), ShouldBeTrue)
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

	Convey("Non-empty without error", t, func() {
		r := &errfloat.ResultsWithErrorCollection{
			Values:        []float32{1.1, 2.2},
			ErrorWrappers: errwrappers.Empty(),
		}
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasAnyItem(), ShouldBeTrue)
		So(r.Length(), ShouldEqual, 2)
		So(r.HasError(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeFalse)
		So(r.SafeValues(), ShouldResemble, []float32{1.1, 2.2})
	})

	Convey("Clear resets values", t, func() {
		r := &errfloat.ResultsWithErrorCollection{
			Values:        []float32{1.1},
			ErrorWrappers: errwrappers.Empty(),
		}
		r.Clear()
		So(r.Length(), ShouldEqual, 0)
	})

	Convey("Dispose nils values", t, func() {
		r := &errfloat.ResultsWithErrorCollection{
			Values:        []float32{1.1},
			ErrorWrappers: errwrappers.Empty(),
		}
		r.Dispose()
		So(r.Values, ShouldBeNil)
	})
}
