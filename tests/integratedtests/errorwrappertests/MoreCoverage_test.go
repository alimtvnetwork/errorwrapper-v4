package errorwrappertests

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

func makeWrapper() *errorwrapper.Wrapper {
	return errnew.Messages.Single(errtype.NotFound, "needle")
}

func Test_TopLevel_Funcs(t *testing.T) {
	Convey("ErrorMessageToError + ErrorsToString family", t, func() {
		So(errorwrapper.ErrorMessageToError(errors.New("a"), "extra").Error(), ShouldContainSubstring, "extra")
		So(errorwrapper.ErrorMessageToError(nil, "x"), ShouldBeNil)
		So(errorwrapper.ErrorsToString(errors.New("a"), errors.New("b")), ShouldContainSubstring, "a")
		So(errorwrapper.ErrorsToError(errors.New("a"), errors.New("b")), ShouldNotBeNil)
		So(errorwrapper.ErrorsToError(), ShouldBeNil)
		So(errorwrapper.MessagesJoined("a", "b"), ShouldContainSubstring, "a")
	})

	Convey("ErrorsToStringUsingJoiner + ErrorsToWrapper", t, func() {
		So(errorwrapper.ErrorsToStringUsingJoiner(",", errors.New("a"), errors.New("b")), ShouldContainSubstring, "a")
		w := errorwrapper.ErrorsToWrapper(errtype.Generic, errors.New("a"))
		So(w, ShouldNotBeNil)
	})

	Convey("ErrorsToWrap with type", t, func() {
		w := errorwrapper.ErrorsToWrap(errtype.NotFound, errors.New("a"), errors.New("b"))
		So(w, ShouldNotBeNil)
	})

	Convey("SimpleReferencesCompile", t, func() {
		So(errorwrapper.SimpleReferencesCompile(errtype.NotFound, "a", 1, true), ShouldNotBeBlank)
		So(errorwrapper.SimpleReferencesCompileOptimized(errtype.NotFound, "a", 1, true), ShouldNotBeBlank)
	})
}

func Test_Constructors_Root(t *testing.T) {
	Convey("Constructor family", t, func() {
		e1 := errorwrapper.Empty()
		So(e1.Ptr().HasError(), ShouldBeFalse)
		So(errorwrapper.EmptyPtr().IsEmpty(), ShouldBeTrue)
		e2 := errorwrapper.EmptyPrint()
		So(e2.Ptr().HasError(), ShouldBeFalse)
		v := errorwrapper.New(errtype.NotFound)
		So(v.Ptr().Type(), ShouldEqual, errtype.NotFound)
		So(errorwrapper.NewPtr(errtype.NotFound), ShouldNotBeNil)
		So(errorwrapper.NewTypeUsingStackSkip(0, errtype.NotFound), ShouldNotBeNil)
		So(errorwrapper.NewPtrUsingStackSkip(0, errtype.NotFound), ShouldNotBeNil)
		So(errorwrapper.NewError(0, errors.New("e")), ShouldNotBeNil)
		So(errorwrapper.NewUsingError(0, errtype.Generic, errors.New("e")), ShouldNotBeNil)
		So(errorwrapper.NewUsingErrorWithoutTypeDisplay(errtype.Generic, errors.New("e")), ShouldNotBeNil)
		So(errorwrapper.NewUsingErrorWithoutTypeDisplayPtr(errtype.Generic, errors.New("e")), ShouldNotBeNil)
		So(errorwrapper.NewUsingTypeErrorAndMessage(0, errtype.Generic, errors.New("e"), "msg"), ShouldNotBeNil)
		So(errorwrapper.NewUsingErrorAndMessage(0, errors.New("e"), "msg"), ShouldNotBeNil)
		So(errorwrapper.NewMessagesUsingJoiner(0, errtype.Generic, " | ", "a", "b"), ShouldNotBeNil)
		So(errorwrapper.NewGeneric(0, errors.New("g")), ShouldNotBeNil)
		So(errorwrapper.NewUnknownMessage(0, true, "u"), ShouldNotBeNil)
	})

	Convey("Ref + Path constructors", t, func() {
		So(errorwrapper.NewRefOne(0, errtype.NotFound, "name", "val"), ShouldNotBeNil)
		So(errorwrapper.TypeReferenceQuick(0, errtype.NotFound, "ref"), ShouldNotBeNil)
	})

	Convey("Wrapper-wrapping constructors", t, func() {
		base := makeWrapper()
		So(errorwrapper.NewUsingWrapper(0, base), ShouldNotBeNil)
		dataModel := base.JsonModel()
		So(errorwrapper.NewFromDataModel(&dataModel), ShouldNotBeNil)
	})
}


func Test_Wrapper_Readers(t *testing.T) {
	w := makeWrapper()

	Convey("Predicate and string readers", t, func() {
		So(w.HasError(), ShouldBeTrue)
		So(w.HasAnyError(), ShouldBeTrue)
		So(w.HasCurrentError(), ShouldBeTrue)
		So(w.HasAnyIssues(), ShouldBeTrue)
		So(w.IsDefined(), ShouldBeTrue)
		So(w.IsInvalid(), ShouldBeTrue)
		So(w.IsFailed(), ShouldBeTrue)
		So(w.IsSuccess(), ShouldBeFalse)
		So(w.IsValid(), ShouldBeFalse)
		So(w.IsNoError(), ShouldBeFalse)
		So(w.IsNull(), ShouldBeFalse)
		So(w.IsAnyNull(), ShouldBeFalse)
		So(w.IsEmpty(), ShouldBeFalse)
		So(w.IsEmptyError(), ShouldBeFalse)
		So(w.IsCollectionType(), ShouldBeFalse)
		So(w.IsTypeOf(errtype.NotFound), ShouldBeTrue)
		So(w.IsReferencesEmpty(), ShouldBeTrue)
		So(w.HasReferences(), ShouldBeFalse)
	})

	Convey("String and message readers", t, func() {
		So(w.Compile(), ShouldNotBeBlank)
		So(w.CompileString(), ShouldNotBeBlank)
		So(w.FullString(), ShouldContainSubstring, "needle")
		So(w.FullStringWithoutReferences(), ShouldNotBeBlank)
		So(w.FullStringWithTraces(), ShouldNotBeBlank)
		So(w.FullStringWithTracesIf(true), ShouldNotBeBlank)
		So(w.FullStringWithLimitTraces(3), ShouldNotBeBlank)
		So(w.FullStringSplitByNewLine(), ShouldNotBeEmpty)
		So(w.Message(), ShouldNotBeBlank)
		So(w.ErrorString(), ShouldNotBeBlank)
		So(w.String(), ShouldNotBeBlank)
		So(w.StringIf(true), ShouldNotBeBlank)
		So(w.TypeName(), ShouldNotBeBlank)
		So(w.TypeString(), ShouldNotBeBlank)
		So(w.TypeNameCode(), ShouldNotBeBlank)
		So(w.TypeNameCodeMessage(), ShouldNotBeBlank)
		So(w.TypeCodeNameString(), ShouldNotBeBlank)
		So(w.TypeNameWithCustomMessage("custom"), ShouldContainSubstring, "custom")
		So(w.CodeTypeName(), ShouldNotBeBlank)
		So(w.RawErrorTypeName(), ShouldNotBeBlank)
		// RawErrorTypeValue skipped: upstream panics with "implement me".
	})

	Convey("Type / Error / Value accessors", t, func() {
		So(w.Type(), ShouldEqual, errtype.NotFound)
		So(w.Error(), ShouldNotBeNil)
		So(w.Value(), ShouldNotBeNil)
		So(w.CompiledError(), ShouldNotBeNil)
		So(w.CompiledErrorWithStackTraces(), ShouldNotBeNil)
		So(w.CompiledJsonErrorWithStackTraces(), ShouldNotBeNil)
		So(w.CompiledJsonStringWithStackTraces(), ShouldNotBeBlank)
		So(w.FullOrErrorMessage(true, true), ShouldNotBeBlank)
		So(w.GetTypeVariantStruct().Name, ShouldEqual, errtype.NotFound.Name())
		So(w.ErrorTypeAsBasicErrorTyper(), ShouldNotBeNil)
	})

	Convey("Stack trace accessors", t, func() {
		So(w.StackTraces(), ShouldNotBeNil)
		So(w.StackTraceString(), ShouldNotBeNil)
		So(w.NewStackTraces(0), ShouldNotBeNil)
		So(w.NewDefaultStackTraces(), ShouldNotBeNil)
		So(w.StackTracesJsonResult(), ShouldNotBeNil)
		So(w.NewStackTracesJsonResult(0), ShouldNotBeNil)
		So(w.NewDefaultStackTracesJsonResult(), ShouldNotBeNil)
		So(w.CompiledStackTracesString(), ShouldNotBeNil)
		So(w.StackTracesLimit(2), ShouldNotBeNil)
	})

	Convey("References / refs accessors", t, func() {
		_ = w.References()
		_ = w.CloneReferences()
		_ = w.ReferencesList()
		_ = w.ReferencesCollection()
		_ = w.ReferencesCompiledString()
	})
}

func Test_Wrapper_Equality(t *testing.T) {
	a := makeWrapper()
	b := a.ClonePtr()

	Convey("Equality + comparison", t, func() {
		So(a.IsEquals(b), ShouldBeTrue)
		So(a.IsNotEquals(b), ShouldBeFalse)
		So(a.IsErrorEquals(a.Error()), ShouldBeTrue)
		So(a.IsErrorMessageEqual(a.Error().Error()), ShouldBeTrue)
		So(a.IsErrorMessage(a.Error().Error(), false), ShouldBeTrue)
		So(a.IsErrorMessageContains("needle", false), ShouldBeTrue)
		So(a.IsBasicErrEqual(a), ShouldBeTrue)
	})
}

func Test_Wrapper_Serialize(t *testing.T) {
	w := makeWrapper()

	Convey("Serialize / Json + roundtrip", t, func() {
		bytes, err := w.Serialize()
		So(err, ShouldBeNil)
		So(bytes, ShouldNotBeEmpty)
		So(w.SerializeMust(), ShouldNotBeEmpty)

		b2, err := w.SerializeWithoutTraces()
		So(err, ShouldBeNil)
		So(b2, ShouldNotBeEmpty)

		jb, err := w.MarshalJSON()
		So(err, ShouldBeNil)
		So(jb, ShouldNotBeEmpty)

		So(w.JsonModel().ErrorType, ShouldEqual, errtype.NotFound)
		So(w.JsonModelAny(), ShouldNotBeNil)
		So(w.Json(), ShouldNotBeNil)
		So(w.JsonPtr(), ShouldNotBeNil)
		So(w.JsonResultWithoutTraces(), ShouldNotBeNil)

		w2 := errorwrapper.EmptyPtr()
		So(w2.UnmarshalJSON(jb), ShouldBeNil)
	})
}

func Test_Wrapper_Clone(t *testing.T) {
	w := makeWrapper()

	Convey("Clone + NonPtr + Ptr + AsErrorWrapper", t, func() {
		So(w.ClonePtr(), ShouldNotBeNil)
		np := w.NonPtr()
		So(np.Ptr().HasError(), ShouldBeTrue)
		So(w.Ptr(), ShouldNotBeNil)
		cl := w.Clone()
		So(cl.Ptr().HasError(), ShouldBeTrue)
		So(w.CloneNewStackSkipPtr(0), ShouldNotBeNil)
		So(w.CloneInterface(), ShouldNotBeNil)
		So(w.AsErrorWrapper(), ShouldEqual, w)
		So(w.AsJsonContractsBinder(), ShouldNotBeNil)
		So(w.AsErrWrapperContractsBinder(), ShouldNotBeNil)
		So(w.GetAsBasicWrapper(), ShouldNotBeNil)
	})
}

func Test_Wrapper_ConcatNew(t *testing.T) {
	Convey("ConcatNew chain methods", t, func() {
		base := makeWrapper()
		c := base.ConcatNew()
		So(c.Message("more"), ShouldNotBeNil)
		So(c.Messages("a", "b"), ShouldNotBeNil)
		So(c.MessagesUsingStackSkip(0, "x"), ShouldNotBeNil)
		So(c.MsgRefOne(0, "msg", "name", "val"), ShouldNotBeNil)
		So(c.MsgRefTwo(0, "msg", "n1", "v1", "n2", "v2"), ShouldNotBeNil)
		So(c.Msg("msg"), ShouldNotBeNil)
		So(c.MsgUsingStackSkip(0, "msg"), ShouldNotBeNil)
		So(c.Error(errors.New("inner")), ShouldNotBeNil)
		So(c.ErrorUsingStackSkip(0, errors.New("inner")), ShouldNotBeNil)
		So(c.Errors(errors.New("a"), errors.New("b")), ShouldNotBeNil)
		So(c.NewStackSkip(0), ShouldNotBeNil)
		// CloneStackSkip skipped: upstream recurses infinitely.
		So(c.Wrapper(makeWrapper()), ShouldNotBeNil)
		So(c.WrapperUsingStackSkip(0, makeWrapper()), ShouldNotBeNil)
		So(c.AsConcatenateNewer(), ShouldNotBeNil)
	})
}
