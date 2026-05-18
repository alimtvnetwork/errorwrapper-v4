package erranytests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errany"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Errany_Constructors_New(t *testing.T) {
	Convey("New.Result.Empty", t, func() {
		r := errany.New.Result.Empty()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("New.Result.Item", t, func() {
		r := errany.New.Result.Item("hello")
		So(r.Value, ShouldEqual, "hello")
	})

	Convey("New.Result.Error", t, func() {
		r := errany.New.Result.Error(errtype.InvalidValidate, errnew.Message("bad"))
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.ErrorWrapper", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "bad")
		r := errany.New.Result.ErrorWrapper(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.Create", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "bad")
		r := errany.New.Result.Create("hello", w)
		So(r.Value, ShouldEqual, "hello")
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.ValueOnly", t, func() {
		r := errany.New.Result.ValueOnly("hello")
		So(r.Value, ShouldEqual, "hello")
		So(r.HasError(), ShouldBeFalse)
	})
}

func Test_Errany_Constructors_Empty(t *testing.T) {
	Convey("Empty.Result", t, func() {
		r := errany.Empty.Result()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("Empty.Results", t, func() {
		r := errany.Empty.Results()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("Empty.Result2", t, func() {
		r := errany.Empty.Result2()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithApplicable", t, func() {
		r := errany.Empty.ResultWithApplicable()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithApplicable2", t, func() {
		r := errany.Empty.ResultWithApplicable2()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultsWithErrorCollection", t, func() {
		r := errany.Empty.ResultsWithErrorCollection()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithError", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "bad")
		r := errany.Empty.ResultWithError(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("Empty.ResultsWithError", t, func() {
		w := errnew.Type.Message(errtype.InvalidValidate, "bad")
		r := errany.Empty.ResultsWithError(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("Empty.ResultWithValue", t, func() {
		r := errany.Empty.ResultWithValue("hello")
		So(r.Value, ShouldEqual, "hello")
	})

	Convey("Empty.ResultsWithValue", t, func() {
		r := errany.Empty.ResultsWithValue([]interface{}{"a", 1})
		So(r.Values, ShouldResemble, []interface{}{"a", 1})
	})
}
