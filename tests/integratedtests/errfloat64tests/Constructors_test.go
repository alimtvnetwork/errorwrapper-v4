package errfloat64tests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errfloat64"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Errfloat64_Constructors_New(t *testing.T) {
	Convey("New.Result.Empty", t, func() {
		r := errfloat64.New.Result.Empty()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("New.Result.Item", t, func() {
		r := errfloat64.New.Result.Item(3.14159)
		So(r.Value, ShouldEqual, 3.14159)
	})

	Convey("New.Result.Error", t, func() {
		r := errfloat64.New.Result.Error(errtype.InvalidValidate, errors.New("bad"))
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.ErrorWrapper", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errfloat64.New.Result.ErrorWrapper(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.Create", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errfloat64.New.Result.Create(3.14159, w)
		So(r.Value, ShouldEqual, 3.14159)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.ValueOnly", t, func() {
		r := errfloat64.New.Result.ValueOnly(3.14159)
		So(r.Value, ShouldEqual, 3.14159)
		So(r.HasError(), ShouldBeFalse)
	})
}

func Test_Errfloat64_Constructors_Empty(t *testing.T) {
	Convey("Empty.Result", t, func() {
		r := errfloat64.Empty.Result()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("Empty.Results", t, func() {
		r := errfloat64.Empty.Results()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("Empty.Result2", t, func() {
		r := errfloat64.Empty.Result2()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithApplicable", t, func() {
		r := errfloat64.Empty.ResultWithApplicable()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithApplicable2", t, func() {
		r := errfloat64.Empty.ResultWithApplicable2()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultsWithErrorCollection", t, func() {
		r := errfloat64.Empty.ResultsWithErrorCollection()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithError", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errfloat64.Empty.ResultWithError(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("Empty.ResultsWithError", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errfloat64.Empty.ResultsWithError(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("Empty.ResultWithValue", t, func() {
		r := errfloat64.Empty.ResultWithValue(3.14159)
		So(r.Value, ShouldEqual, 3.14159)
	})

	Convey("Empty.ResultsWithValue", t, func() {
		r := errfloat64.Empty.ResultsWithValue([]float64{1.1, 2.2})
		So(r.Values, ShouldResemble, []float64{1.1, 2.2})
	})
}
