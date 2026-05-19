package errstrtests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errstr"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrStr_Constructors_New(t *testing.T) {
	Convey("New.Result.Empty", t, func() {
		r := errstr.New.Result.Empty()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("New.Result.Item", t, func() {
		r := errstr.New.Result.Item("hello")
		So(r.Value, ShouldEqual, "hello")
	})

	Convey("New.Result.String", t, func() {
		r := errstr.New.Result.String("world")
		So(r.Value, ShouldEqual, "world")
	})

	Convey("New.Result.Error", t, func() {
		r := errstr.New.Result.Error(errtype.InvalidValidate, errnew.Message("bad"))
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.ErrorWrapper", t, func() {
		w := errnew.Message.Type(errtype.InvalidValidate, "bad")
		r := errstr.New.Result.ErrorWrapper(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.Create", t, func() {
		w := errnew.Message.Type(errtype.InvalidValidate, "bad")
		r := errstr.New.Result.Create("x", w)
		So(r.Value, ShouldEqual, "x")
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.ValueOnly", t, func() {
		r := errstr.New.Result.ValueOnly("v")
		So(r.Value, ShouldEqual, "v")
		So(r.HasError(), ShouldBeFalse)
	})
}

func Test_ErrStr_Constructors_Empty(t *testing.T) {
	Convey("Empty.Result", t, func() {
		r := errstr.Empty.Result()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("Empty.Results", t, func() {
		r := errstr.Empty.Results()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("Empty.Result2", t, func() {
		r := errstr.Empty.Result2()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithApplicable", t, func() {
		r := errstr.Empty.ResultWithApplicable()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithApplicable2", t, func() {
		r := errstr.Empty.ResultWithApplicable2()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultsWithErrorCollection", t, func() {
		r := errstr.Empty.ResultsWithErrorCollection()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.Collection", t, func() {
		r := errstr.Empty.Collection()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.LinkedList", t, func() {
		r := errstr.Empty.LinkedList()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.LinkedCollections", t, func() {
		r := errstr.Empty.LinkedCollections()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.Hashset", t, func() {
		r := errstr.Empty.Hashset()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.Hashmap", t, func() {
		r := errstr.Empty.Hashmap()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.SimpleStringOnce", t, func() {
		r := errstr.Empty.SimpleStringOnce()
		So(r, ShouldNotBeNil)
	})
}
