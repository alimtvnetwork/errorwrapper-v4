package erranygentests

import (
	"testing"

	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v4/errdata/erranygen"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_GenericResult_NilReceiver_Paths(t *testing.T) {
	Convey("Nil receiver methods are all safe and consistent", t, func() {
		var r *erranygen.Result[int]

		So(r.IsAnyNull(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
		So(r.IsEmptyError(), ShouldBeTrue)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.HasSafeItems(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.IsFailed(), ShouldBeTrue)
		So(r.IsValid(), ShouldBeFalse)
		So(r.IsInvalid(), ShouldBeTrue)

		_ = r.ErrorWrapperInf()

		So(func() { r.Dispose() }, ShouldNotPanic)
	})
}

func Test_GenericResult_WrapperAndIsZero_Combinations(t *testing.T) {
	Convey("Wrapper present + isZero-true value -> failed & empty", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "oops")
		r := erranygen.NewResult(0, w, func(v int) bool { return v == 0 })

		So(r.HasError(), ShouldBeTrue)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.HasSafeItems(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.IsFailed(), ShouldBeTrue)
		So(r.IsValid(), ShouldBeFalse)
		So(r.IsInvalid(), ShouldBeTrue)
	})

	Convey("Wrapper present + non-zero value still failed", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "oops")
		r := erranygen.NewResult(7, w, func(v int) bool { return v == 0 })

		So(r.HasError(), ShouldBeTrue)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeFalse)
	})

	Convey("No wrapper, no predicate, zero value -> not empty, success", t, func() {
		r := erranygen.NewResult(0, nil, nil)

		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.IsValid(), ShouldBeTrue)
	})

	Convey("ErrorWrapperInf returns underlying wrapper interface", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "x")
		r := erranygen.NewResult("v", w, nil)
		inf := r.ErrorWrapperInf()
		So(inf, ShouldNotBeNil)
		So(inf.HasError(), ShouldBeTrue)
	})
}

func Test_GenericResult_JsonSurface(t *testing.T) {
	Convey("JsonModelAny returns self value", t, func() {
		r := erranygen.NewResult("abc", nil, nil)
		model := r.JsonModelAny()
		So(model, ShouldNotBeNil)
	})

	Convey("AsJsonContractsBinder is non-nil and round-trips", t, func() {
		r := erranygen.NewResult("abc", nil, nil)
		binder := r.AsJsonContractsBinder()
		So(binder, ShouldNotBeNil)

		j := binder.Json()
		So(j.HasError(), ShouldBeFalse)
	})

	Convey("JsonParseSelfInject on bad payload returns error", t, func() {
		bad := &corejson.Result{Bytes: []byte("{not-json")}
		var dst erranygen.Result[string]
		err := dst.JsonParseSelfInject(bad)
		So(err, ShouldNotBeNil)
	})
}
