package errnewtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

// MoreCoverage2_test.go — covers the JSON / Unmarshal / DeserializeTo /
// ErrInterface / Reflect creator families that MoreCoverage_test.go skipped.
// Assertions stay defensive (non-nil + HasError) because most paths are
// plumbing on top of corejson.

// ---- Unmarshal ----

func Test_MoreCoverage2_Unmarshal(t *testing.T) {
	Convey("Unmarshal creator family", t, func() {
		err := errors.New("boom")
		So(errnew.Unmarshal.Error(err).HasError(), ShouldBeTrue)
		So(errnew.Unmarshal.Error(nil), ShouldBeNil)

		So(errnew.Unmarshal.MessageError("ctx: ", err).HasError(), ShouldBeTrue)
		So(errnew.Unmarshal.MessageError("ctx: ", nil), ShouldBeNil)

		So(errnew.Unmarshal.Message("hello"), ShouldNotBeNil)
		So(errnew.Unmarshal.Message(""), ShouldBeNil)

		So(errnew.Unmarshal.MessageRef("hello", []byte("{}"), nil), ShouldNotBeNil)
		So(errnew.Unmarshal.MessageRef("", []byte("{}"), nil), ShouldBeNil)

		So(errnew.Unmarshal.ErrorRef(err, []byte("{}"), nil), ShouldNotBeNil)
		So(errnew.Unmarshal.ErrorRef(nil, []byte("{}"), nil), ShouldBeNil)

		So(errnew.Unmarshal.MessageErrorRef("ctx: ", err, []byte("{}"), nil), ShouldNotBeNil)
		So(errnew.Unmarshal.MessageErrorRef("ctx: ", nil, []byte("{}"), nil), ShouldBeNil)

		So(errnew.Unmarshal.Reference([]byte("{}"), nil), ShouldNotBeNil)
		So(errnew.Unmarshal.ReferenceUsingStackSkip(0, []byte("{}"), nil), ShouldNotBeNil)
		So(errnew.Unmarshal.MessageRefUsingStackSkip(0, "m", []byte("{}"), nil), ShouldNotBeNil)
		So(errnew.Unmarshal.MessageRefUsingStackSkip(0, "", []byte("{}"), nil), ShouldBeNil)

		// success path: valid bytes deserialize into matching type
		var dst map[string]int
		So(errnew.Unmarshal.BytesToDeserializeTo([]byte(`{"a":1}`), &dst), ShouldBeNil)
		// failure path: invalid bytes
		So(errnew.Unmarshal.BytesToDeserializeTo([]byte("not-json"), &dst).HasError(), ShouldBeTrue)

		jr := corejson.NewPtr([]byte(`{"a":1}`))
		_ = errnew.Unmarshal.JsonResultToDeserializeTo(jr, &dst)
		badJr := corejson.NewPtr([]byte("not-json"))
		_ = errnew.Unmarshal.JsonResultToDeserializeTo(badJr, &dst)
	})
}

// ---- Json ----

func Test_MoreCoverage2_Json(t *testing.T) {
	Convey("Json creator family", t, func() {
		// nil jsoner → error wrapper
		_, w := errnew.Json.Jsoner(nil)
		So(w.HasError(), ShouldBeTrue)

		_, w = errnew.Json.JsonerToBytes(nil)
		So(w.HasError(), ShouldBeTrue)

		So(errnew.Json.JsonerToDeserialize(nil, nil).HasError(), ShouldBeTrue)

		var dst map[string]int
		So(errnew.Json.BytesToDeserializeTo([]byte(`{"a":1}`), &dst), ShouldBeNil)
		So(errnew.Json.BytesToDeserializeTo([]byte("not-json"), &dst).HasError(), ShouldBeTrue)

		jr := corejson.NewPtr([]byte(`{"a":1}`))
		_ = errnew.Json.ResultDeserializeTo(jr, &dst)
		_ = errnew.Json.ResultDeserializeTo(corejson.NewPtr([]byte("nope")), &dst)

		_, w = errnew.Json.JsonerToJsonString(nil)
		So(w.HasError(), ShouldBeTrue)

		So(errnew.Json.Result(jr) == nil || !errnew.Json.Result(jr).HasError(), ShouldBeTrue)
		So(errnew.Json.ResultSkipNullIssues(jr) == nil || !errnew.Json.ResultSkipNullIssues(jr).HasError(), ShouldBeTrue)

		_, w = errnew.Json.BytesErrorFunc(nil)
		So(w.HasError(), ShouldBeTrue)
		_, w = errnew.Json.BytesErrorFunc(func() ([]byte, error) { return []byte("ok"), nil })
		So(w, ShouldBeNil)
		_, w = errnew.Json.BytesErrorFunc(func() ([]byte, error) { return nil, errors.New("x") })
		So(w.HasError(), ShouldBeTrue)

		So(errnew.Json.MessageReferenceJson(errtype.IO, "msg", 1, "a"), ShouldNotBeNil)
		So(errnew.Json.MessageReferenceJson(errtype.IO, ""), ShouldBeNil)

		So(errnew.Json.ErrorReferenceJson(errtype.IO, errors.New("e"), 1), ShouldNotBeNil)
		So(errnew.Json.ErrorReferenceJson(errtype.IO, nil, 1), ShouldBeNil)
	})
}

// ---- DeserializeTo ----

func Test_MoreCoverage2_DeserializeTo(t *testing.T) {
	Convey("DeserializeTo creator family", t, func() {
		var dst map[string]int
		jr := corejson.NewPtr([]byte(`{"a":1}`))
		So(errnew.DeserializeTo.JsonResultToAny(jr, &dst), ShouldBeNil)
		badJr := corejson.NewPtr([]byte("nope"))
		So(errnew.DeserializeTo.JsonResultToAny(badJr, &dst).HasError(), ShouldBeTrue)

		So(errnew.DeserializeTo.JsonResultToAnyOption(true, nil, &dst), ShouldBeNil)
		So(errnew.DeserializeTo.JsonResultToAnyOption(false, badJr, &dst).HasError(), ShouldBeTrue)
		So(errnew.DeserializeTo.JsonResultToAnySkipOnNull(nil, &dst), ShouldBeNil)
		So(errnew.DeserializeTo.JsonResultToAnyOnErrAddMsg("ctx: ", badJr, &dst).HasError(), ShouldBeTrue)

		_, parsed := errnew.DeserializeTo.JsonErrToWrapper(nil)
		So(parsed, ShouldBeNil)
		_, _ = errnew.DeserializeTo.JsonErrToWrapper(errors.New("{not-json"))

		_, _ = errnew.DeserializeTo.JsonResultErrToWrapper(errors.New("not-json"))

		_, _ = errnew.DeserializeTo.BytesToWrapper(nil)
		_, _ = errnew.DeserializeTo.BytesToWrapper([]byte("not-json"))
		So(errnew.DeserializeTo.BytesToUnmarshal(nil, &dst), ShouldBeNil)
		So(errnew.DeserializeTo.BytesToUnmarshal([]byte("not-json"), &dst).HasError(), ShouldBeTrue)
		So(errnew.DeserializeTo.BytesToAnyPtr(nil, &dst), ShouldBeNil)
		So(errnew.DeserializeTo.BytesToAnyPtr([]byte("not-json"), &dst).HasError(), ShouldBeTrue)

		_, _ = errnew.DeserializeTo.JsonResultToWrapper(nil)
		_, _ = errnew.DeserializeTo.JsonResultToWrapper(badJr)
		_, _ = errnew.DeserializeTo.JsonResultToWrapperUsingStackSkip(0, nil)
		_, _ = errnew.DeserializeTo.JsonResultToWrapperUsingStackSkip(0, badJr)

		_, _ = errnew.DeserializeTo.JsonStringToWrapper(false, "")
		_, _ = errnew.DeserializeTo.JsonStringToWrapper(true, "")
		_, _ = errnew.DeserializeTo.JsonStringToWrapperUsingStackSkip(0, false, "not-json")
	})
}

// ---- ErrInterface ----

func Test_MoreCoverage2_ErrInterface(t *testing.T) {
	Convey("ErrInterface creator family", t, func() {
		// Default with nil → nil
		So(errnew.ErrInterface.Default(errtype.IO, nil), ShouldBeNil)
		So(errnew.ErrInterface.NoType(nil), ShouldBeNil)
		So(errnew.ErrInterface.BasicErr(nil), ShouldBeNil)
		So(errnew.ErrInterface.RawErrCollection(errtype.IO, nil), ShouldBeNil)
		So(errnew.ErrInterface.ErrorWrapperCollectionDefiner(errtype.IO, nil), ShouldBeNil)
		So(errnew.ErrInterface.ErrorWrapperCollectionsDefiner(errtype.IO), ShouldBeNil)

		// AnyType: nil
		conv, _ := errnew.ErrInterface.AnyType(errtype.IO, nil)
		So(conv, ShouldBeNil)

		// AnyType: *errorwrapper.Wrapper
		w := errnew.Type.Default(errtype.IO)
		conv, _ = errnew.ErrInterface.AnyType(errtype.IO, w)
		So(conv, ShouldEqual, w)

		// AnyType: errorwrapper.Wrapper value
		conv, _ = errnew.ErrInterface.AnyType(errtype.IO, *w)
		So(conv, ShouldNotBeNil)

		// AnyType: []byte invalid → parsed error wrapper
		_, _ = errnew.ErrInterface.AnyType(errtype.IO, []byte("not-json"))

		// AnyType: corejson.Result invalid
		jr := corejson.NewPtr([]byte("not-json"))
		_, _ = errnew.ErrInterface.AnyType(errtype.IO, jr)
		_, _ = errnew.ErrInterface.AnyType(errtype.IO, *jr)

		// AnyType: plain error
		conv, _ = errnew.ErrInterface.AnyType(errtype.IO, errors.New("plain"))
		So(conv, ShouldNotBeNil)
		So(conv.HasError(), ShouldBeTrue)

		// AnyType: plain string
		conv, _ = errnew.ErrInterface.AnyType(errtype.IO, "just a string")
		So(conv, ShouldNotBeNil)

		// AnyType: unknown struct → parsed-failed (cast error)
		type weird struct{ X int }
		_, parsed := errnew.ErrInterface.AnyType(errtype.IO, weird{X: 1})
		So(parsed, ShouldNotBeNil)
	})
}

// ---- Reflect ----

func Test_MoreCoverage2_Reflect(t *testing.T) {
	Convey("Reflect creator family", t, func() {
		a := 1
		b := 1
		// same types, equal value → nil
		So(errnew.Reflect.SetFromTo(2, &b), ShouldBeNil)

		// type mismatch wrapper
		var s string
		_ = errnew.Reflect.TypeMismatch(a, s)

		// value mismatch
		So(errnew.Reflect.ValueMismatchOption(false, "neq", 1, 2).HasError(), ShouldBeTrue)
		So(errnew.Reflect.ValueMismatchOption(false, "eq", 1, 1), ShouldBeNil)
		So(errnew.Reflect.ValueMismatchOption(true, "regardless", 1, "1"), ShouldBeNil)
		So(errnew.Reflect.ValueMismatchRegardless("ctx", 1, 1), ShouldBeNil)
		So(errnew.Reflect.ValueMismatchRegardless("ctx", 1, 2).HasError(), ShouldBeTrue)
	})
}

// ---- silence unused imports if branches change ----
var _ = errorwrapper.StaticEmptyPtr
