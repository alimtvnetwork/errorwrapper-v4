package errbooltests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errbool"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrBool_Constructors_New(t *testing.T) {
	Convey("New.Result.Empty", t, func() {
		r := errbool.New.Result.Empty()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("New.Result.True", t, func() {
		r := errbool.New.Result.True()
		So(r, ShouldNotBeNil)
		So(r.Value, ShouldBeTrue)
		So(r.IsTrue(), ShouldBeTrue)
	})

	Convey("New.Result.False", t, func() {
		r := errbool.New.Result.False()
		So(r, ShouldNotBeNil)
		So(r.Value, ShouldBeFalse)
		So(r.IsFalse(), ShouldBeTrue)
	})

	Convey("New.Result.Item", t, func() {
		r := errbool.New.Result.Item(true)
		So(r.Value, ShouldBeTrue)
	})

	Convey("New.Result.Bool", t, func() {
		r := errbool.New.Result.Bool(false)
		So(r.Value, ShouldBeFalse)
	})

	Convey("New.Result.TrueWithErr", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errbool.New.Result.TrueWithErr(w)
		So(r.Value, ShouldBeTrue)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.FalseWithErr", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errbool.New.Result.FalseWithErr(w)
		So(r.Value, ShouldBeFalse)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.Error", t, func() {
		r := errbool.New.Result.Error(errtype.InvalidValidate, errors.New("bad"))
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.ErrorWrapper", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errbool.New.Result.ErrorWrapper(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.Create", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errbool.New.Result.Create(true, w)
		So(r.Value, ShouldBeTrue)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.ValueOnly", t, func() {
		r := errbool.New.Result.ValueOnly(true)
		So(r.Value, ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
	})
}

func Test_ErrBool_Constructors_Empty(t *testing.T) {
	Convey("Empty.Result", t, func() {
		r := errbool.Empty.Result()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("Empty.Results", t, func() {
		r := errbool.Empty.Results()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("Empty.Result2", t, func() {
		r := errbool.Empty.Result2()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.Result3", t, func() {
		r := errbool.Empty.Result3()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultsWithErrorCollection", t, func() {
		r := errbool.Empty.ResultsWithErrorCollection()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithApplicable", t, func() {
		r := errbool.Empty.ResultWithApplicable()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithApplicable2", t, func() {
		r := errbool.Empty.ResultWithApplicable2()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithError", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errbool.Empty.ResultWithError(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("Empty.ResultsWithError", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errbool.Empty.ResultsWithError(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("Empty.ResultWithValue", t, func() {
		r := errbool.Empty.ResultWithValue(true)
		So(r.Value, ShouldBeTrue)
	})

	Convey("Empty.ResultsWithValue", t, func() {
		r := errbool.Empty.ResultsWithValue([]bool{true, false})
		So(r.Values, ShouldResemble, []bool{true, false})
	})
}
