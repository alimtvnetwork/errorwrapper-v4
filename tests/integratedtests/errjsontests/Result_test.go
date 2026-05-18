package errjsontests

import (
	"testing"

	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errjson"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrJson_Result_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errjson.Result
		So(r.IsAnyNull(), ShouldBeTrue)
		So(r.IsNull(), ShouldBeTrue)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
		So(r.IsEmptyError(), ShouldBeTrue)
		So(r.HasIssuesOrEmpty(), ShouldBeTrue)
		So(r.HasSafeItems(), ShouldBeFalse)
		So(r.HasAnyItem(), ShouldBeFalse)
		So(r.Length(), ShouldEqual, 0)
		So(r.SafeValues(), ShouldResemble, []byte{})
		So(r.SafeBytes(), ShouldResemble, []byte{})
		So(r.SafeString(), ShouldEqual, "")
		So(r.String(), ShouldEqual, "")
		So(r.JsonString(), ShouldEqual, "")
		So(r.PrettyJsonString(), ShouldEqual, "")
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.IsFailed(), ShouldBeTrue)
		So(func() { r.Dispose() }, ShouldNotPanic)
		So(func() { r.SplitLines() }, ShouldResemble, []string{})
		So(func() { r.SplitLinesSimpleSlice() }, ShouldNotPanic)
		So(func() { r.ValidValue() }, ShouldNotPanic)
		So(func() { r.SimpleStringOnce(true) }, ShouldNotPanic)
	})

	Convey("Empty json result without error", t, func() {
		jr := corejson.NewPtr(nil)
		r := &errjson.Result{Result: jr, ErrorWrapper: nil}
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasAnyItem(), ShouldBeFalse)
	})

	Convey("Json result with bytes without error", t, func() {
		jr := corejson.NewPtr(`{"key":"value"}`)
		r := &errjson.Result{Result: jr, ErrorWrapper: nil}
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasAnyItem(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
		So(r.HasIssuesOrEmpty(), ShouldBeFalse)
		So(r.HasSafeItems(), ShouldBeTrue)
		So(r.IsSuccess(), ShouldBeTrue)
		So(r.IsFailed(), ShouldBeFalse)
		So(r.Length() > 0, ShouldBeTrue)
		So(r.SafeValues(), ShouldNotBeEmpty)
		So(r.SafeString(), ShouldNotBeBlank)
		So(r.String(), ShouldNotBeBlank)
		So(r.JsonString(), ShouldNotBeBlank)
	})

	Convey("IsEqual and IsEqualIgnoreCase", t, func() {
		jr := corejson.NewPtr(`"hello"`)
		r := &errjson.Result{Result: jr, ErrorWrapper: nil}
		So(r.IsEqual("hello"), ShouldBeTrue)
		So(r.IsEqualIgnoreCase("HELLO"), ShouldBeTrue)
	})

	Convey("Wrapper error takes precedence", t, func() {
		jr := corejson.NewPtr(`{"ok":true}`)
		w := errnew.Type.Message(errtype.InvalidValidate, "bad")
		r := &errjson.Result{Result: jr, ErrorWrapper: w}
		So(r.HasError(), ShouldBeTrue)
		So(r.IsEmptyError(), ShouldBeFalse)
		So(r.IsSuccess(), ShouldBeFalse)
		So(r.IsFailed(), ShouldBeTrue)
		So(r.CompiledErrorWrapper(), ShouldNotBeNil)
	})

	Convey("SplitLines", t, func() {
		jr := corejson.NewPtr(`{"a":1}
{"b":2}`)
		r := &errjson.Result{Result: jr, ErrorWrapper: nil}
		lines := r.SplitLines()
		So(len(lines) >= 1, ShouldBeTrue)
	})
}

func Test_ErrJson_Result_JSON(t *testing.T) {
	Convey("Json round-trip", t, func() {
		jr := corejson.NewPtr(`{"x":1}`)
		r := &errjson.Result{Result: jr, ErrorWrapper: nil}
		j := r.Json()
		So(j.HasError(), ShouldBeFalse)
		So(r.JsonPtr(), ShouldNotBeNil)
	})
}
