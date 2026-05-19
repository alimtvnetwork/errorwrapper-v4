package errbytetests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errbyte"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrByte_ResultsWithErrorCollection_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errbyte.ResultsWithErrorCollection
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
		r := &errbyte.ResultsWithErrorCollection{
			Values:        []byte{65, 66},
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
		So(r.SafeValues(), ShouldResemble, []byte{65, 66})
	})

	Convey("Clear resets values", t, func() {
		r := &errbyte.ResultsWithErrorCollection{
			Values:        []byte{65},
			ErrorWrappers: errwrappers.Empty(),
		}
		r.Clear()
		So(r.Length(), ShouldEqual, 0)
	})

	Convey("Dispose nils values", t, func() {
		r := &errbyte.ResultsWithErrorCollection{
			Values:        []byte{65},
			ErrorWrappers: errwrappers.Empty(),
		}
		r.Dispose()
		So(r.Values, ShouldBeNil)
	})
}
