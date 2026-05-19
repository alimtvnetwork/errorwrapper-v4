package errinttests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errint"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrInt_Constructors_New(t *testing.T) {
	Convey("New.Result.Empty", t, func() {
		r := errint.New.Result.Empty()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("New.Result.Item", t, func() {
		r := errint.New.Result.Item(42)
		So(r.Value, ShouldEqual, 42)
	})

	Convey("New.Result.Int", t, func() {
		r := errint.New.Result.Int(99)
		So(r.Value, ShouldEqual, 99)
	})

	Convey("New.Result.Error", t, func() {
		r := errint.New.Result.Error(errtype.InvalidValidate, errors.New("bad"))
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.ErrorWrapper", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errint.New.Result.ErrorWrapper(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.Create", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errint.New.Result.Create(7, w)
		So(r.Value, ShouldEqual, 7)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.ValueOnly", t, func() {
		r := errint.New.Result.ValueOnly(5)
		So(r.Value, ShouldEqual, 5)
		So(r.HasError(), ShouldBeFalse)
	})
}

func Test_ErrInt_Constructors_Empty(t *testing.T) {
	Convey("Empty.Result", t, func() {
		r := errint.Empty.Result()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("Empty.Results", t, func() {
		r := errint.Empty.Results()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("Empty.Result2", t, func() {
		r := errint.Empty.Result2()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultsWithErrorCollection", t, func() {
		r := errint.Empty.ResultsWithErrorCollection()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithApplicable", t, func() {
		r := errint.Empty.ResultWithApplicable()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithApplicable2", t, func() {
		r := errint.Empty.ResultWithApplicable2()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithError", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errint.Empty.ResultWithError(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("Empty.ResultsWithError", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errint.Empty.ResultsWithError(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("Empty.ResultWithValue", t, func() {
		r := errint.Empty.ResultWithValue(42)
		So(r.Value, ShouldEqual, 42)
	})

	Convey("Empty.ResultsWithValue", t, func() {
		r := errint.Empty.ResultsWithValue([]int{1, 2})
		So(r.Values, ShouldResemble, []int{1, 2})
	})
}
