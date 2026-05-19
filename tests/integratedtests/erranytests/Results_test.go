package erranytests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errany"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrAny_Results_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errany.Results
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
		r := &errany.Results{Values: []interface{}{}, ErrorWrapper: nil}
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasAnyItem(), ShouldBeFalse)
		So(r.Length(), ShouldEqual, 0)
	})

	Convey("Non-empty slice without error", t, func() {
		r := &errany.Results{Values: []interface{}{"a", 1, true}, ErrorWrapper: nil}
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasAnyItem(), ShouldBeTrue)
		So(r.Length(), ShouldEqual, 3)
		So(r.HasError(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeFalse)
		So(r.SafeValues(), ShouldResemble, []interface{}{"a", 1, true})
		So(r.String(), ShouldNotBeBlank)
	})

	Convey("Clear resets values", t, func() {
		r := &errany.Results{Values: []interface{}{"a"}, ErrorWrapper: nil}
		r.Clear()
		So(r.Length(), ShouldEqual, 0)
	})

	Convey("Dispose nils values", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "x")
		r := &errany.Results{Values: []interface{}{"a"}, ErrorWrapper: w}
		r.Dispose()
		So(r.Values, ShouldBeNil)
	})
}

func Test_ErrAny_ResultsWithErrorCollection_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errany.ResultsWithErrorCollection
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
		r := &errany.ResultsWithErrorCollection{
			Values:        []interface{}{"a", 1},
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
		So(r.SafeValues(), ShouldResemble, []interface{}{"a", 1})
	})

	Convey("Clear resets values", t, func() {
		r := &errany.ResultsWithErrorCollection{
			Values:        []interface{}{"a"},
			ErrorWrappers: errwrappers.Empty(),
		}
		r.Clear()
		So(r.Length(), ShouldEqual, 0)
	})

	Convey("Dispose nils values", t, func() {
		r := &errany.ResultsWithErrorCollection{
			Values:        []interface{}{"a"},
			ErrorWrappers: errwrappers.Empty(),
		}
		r.Dispose()
		So(r.Values, ShouldBeNil)
	})
}
