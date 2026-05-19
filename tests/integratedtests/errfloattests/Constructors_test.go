package errfloattests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errfloat"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Errfloat_Constructors_New(t *testing.T) {
	Convey("New.Result.Empty", t, func() {
		r := errfloat.New.Result.Empty()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("New.Result.Item", t, func() {
		r := errfloat.New.Result.Item(3.14)
		So(r.Value, ShouldEqual, 3.14)
	})

	Convey("New.Result.Error", t, func() {
		r := errfloat.New.Result.Error(errtype.InvalidValidate, errnew.Message("bad"))
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.ErrorWrapper", t, func() {
		w := errnew.Message.Type(errtype.InvalidValidate, "bad")
		r := errfloat.New.Result.ErrorWrapper(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.Create", t, func() {
		w := errnew.Message.Type(errtype.InvalidValidate, "bad")
		r := errfloat.New.Result.Create(3.14, w)
		So(r.Value, ShouldEqual, 3.14)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.ValueOnly", t, func() {
		r := errfloat.New.Result.ValueOnly(3.14)
		So(r.Value, ShouldEqual, 3.14)
		So(r.HasError(), ShouldBeFalse)
	})
}

func Test_Errfloat_Constructors_Empty(t *testing.T) {
	Convey("Empty.Result", t, func() {
		r := errfloat.Empty.Result()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("Empty.Results", t, func() {
		r := errfloat.Empty.Results()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("Empty.Result2", t, func() {
		r := errfloat.Empty.Result2()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithApplicable", t, func() {
		r := errfloat.Empty.ResultWithApplicable()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithApplicable2", t, func() {
		r := errfloat.Empty.ResultWithApplicable2()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultsWithErrorCollection", t, func() {
		r := errfloat.Empty.ResultsWithErrorCollection()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithError", t, func() {
		w := errnew.Message.Type(errtype.InvalidValidate, "bad")
		r := errfloat.Empty.ResultWithError(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("Empty.ResultsWithError", t, func() {
		w := errnew.Message.Type(errtype.InvalidValidate, "bad")
		r := errfloat.Empty.ResultsWithError(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("Empty.ResultWithValue", t, func() {
		r := errfloat.Empty.ResultWithValue(3.14)
		So(r.Value, ShouldEqual, 3.14)
	})

	Convey("Empty.ResultsWithValue", t, func() {
		r := errfloat.Empty.ResultsWithValue([]float32{1.1, 2.2})
		So(r.Values, ShouldResemble, []float32{1.1, 2.2})
	})
}
