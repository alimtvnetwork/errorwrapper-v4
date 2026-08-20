package errbytetests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errbyte"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Errbyte_Constructors_New(t *testing.T) {
	Convey("New.Result.Empty", t, func() {
		r := errbyte.New.Result.Empty()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("New.Result.Item", t, func() {
		r := errbyte.New.Result.Item(65)
		So(r.Value, ShouldEqual, 65)
	})

	Convey("New.Result.Error", t, func() {
		r := errbyte.New.Result.Error(errtype.InvalidValidate, errors.New("bad"))
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.ErrorWrapper", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errbyte.New.Result.ErrorWrapper(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.Create", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errbyte.New.Result.Create(65, w)
		So(r.Value, ShouldEqual, 65)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.ValueOnly", t, func() {
		r := errbyte.New.Result.ValueOnly(65)
		So(r.Value, ShouldEqual, 65)
		So(r.HasError(), ShouldBeFalse)
	})
}

func Test_Errbyte_Constructors_Empty(t *testing.T) {
	Convey("Empty.Result", t, func() {
		r := errbyte.Empty.Result()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("Empty.Results", t, func() {
		r := errbyte.Empty.Results()
		So(r, ShouldNotBeNil)
		So(r.IsEmpty(), ShouldBeTrue)
	})

	Convey("Empty.Result2", t, func() {
		r := errbyte.Empty.Result2()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithApplicable", t, func() {
		r := errbyte.Empty.ResultWithApplicable()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithApplicable2", t, func() {
		r := errbyte.Empty.ResultWithApplicable2()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultsWithErrorCollection", t, func() {
		r := errbyte.Empty.ResultsWithErrorCollection()
		So(r, ShouldNotBeNil)
	})

	Convey("Empty.ResultWithError", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errbyte.Empty.ResultWithError(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("Empty.ResultsWithError", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		r := errbyte.Empty.ResultsWithError(w)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("Empty.ResultWithValue", t, func() {
		r := errbyte.Empty.ResultWithValue(65)
		So(r.Value, ShouldEqual, 65)
	})

	Convey("Empty.ResultsWithValue", t, func() {
		r := errbyte.Empty.ResultsWithValue([]byte{65, 66})
		So(r.Values, ShouldResemble, []byte{65, 66})
	})
}
