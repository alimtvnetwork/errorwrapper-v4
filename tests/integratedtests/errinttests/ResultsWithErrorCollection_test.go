package errinttests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errint"

	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrInt_ResultsWithErrorCollection_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errint.ResultsWithErrorCollection
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
		r := errint.NewResultsWithErrorCollection(
			[]int{1, 2},
			errwrappers.Empty(),
		)
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasAnyItem(), ShouldBeTrue)
		So(r.Length(), ShouldEqual, 2)
		So(r.HasError(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeFalse)
		So(r.SafeValues(), ShouldResemble, []int{1, 2})
	})

	Convey("EmptyResultsWithErrorCollection", t, func() {
		r := errint.EmptyResultsWithErrorCollection()
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
	})

	Convey("NewResultsWithErrorCollectionUsingErrorCollection", t, func() {
		r := errint.NewResultsWithErrorCollectionUsingErrorCollection(
			errwrappers.Empty(),
		)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
	})

	Convey("NewResultsWithErrorCollectionUsingTypeMessage", t, func() {
		r := errint.NewResultsWithErrorCollectionUsingTypeMessage(
			errtype.InvalidValidate, "bad")
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("NewResultsWithErrorCollectionUsingTypeError", t, func() {
		r := errint.NewResultsWithErrorCollectionUsingTypeError(
			errtype.InvalidValidate, errors.New("bad"))
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("NewResultsWithErrorCollectionUsingType", t, func() {
		r := errint.NewResultsWithErrorCollectionUsingType(
			errtype.InvalidValidate)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("Clear resets values", t, func() {
		r := errint.NewResultsWithErrorCollection(
			[]int{1},
			errwrappers.Empty(),
		)
		r.Clear()
		So(r.Length(), ShouldEqual, 0)
	})

	Convey("Dispose nils values", t, func() {
		r := errint.NewResultsWithErrorCollection(
			[]int{1},
			errwrappers.Empty(),
		)
		r.Dispose()
		So(r.Values, ShouldBeNil)
	})
}
