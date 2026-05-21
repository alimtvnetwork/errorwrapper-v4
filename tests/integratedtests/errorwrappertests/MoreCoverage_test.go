package errorwrappertests

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
)

// Test_ErrorMessageToError — concatenated error or nil.
func Test_ErrorMessageToError(t *testing.T) {
	Convey("nil input yields nil output", t, func() {
		So(errorwrapper.ErrorMessageToError(nil, "msg"), ShouldBeNil)
	})
	Convey("non-nil input concatenates the message", t, func() {
		out := errorwrapper.ErrorMessageToError(errors.New("base"), "extra")
		So(out, ShouldNotBeNil)
		So(out.Error(), ShouldContainSubstring, "base")
		So(out.Error(), ShouldContainSubstring, "extra")
	})
}

// Test_ErrorsToStringUsingJoiner — joins messages, ignores nils.
func Test_ErrorsToStringUsingJoiner(t *testing.T) {
	Convey("Empty input returns empty string", t, func() {
		So(errorwrapper.ErrorsToStringUsingJoiner(", "), ShouldEqual, "")
	})
	Convey("Joins with joiner ignoring nil entries", t, func() {
		out := errorwrapper.ErrorsToStringUsingJoiner(
			", ",
			errors.New("a"),
			nil,
			errors.New("b"),
		)
		So(out, ShouldEqual, "a, b")
	})
}

// Test_ErrorsToWrapper — combines errors into a wrapper.
func Test_ErrorsToWrapper(t *testing.T) {
	Convey("Wraps multiple errors into a single typed wrapper", t, func() {
		w := errorwrapper.ErrorsToWrapper(
			errtype.InvalidInput,
			errors.New("a"),
			errors.New("b"))
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
		So(w.FullString(), ShouldContainSubstring, "a")
		So(w.FullString(), ShouldContainSubstring, "b")
	})
}

// Test_NewFromDataModel — round-trip via WrapperDataModel.
func Test_NewFromDataModel(t *testing.T) {
	Convey("nil model returns empty wrapper", t, func() {
		w := errorwrapper.NewFromDataModel(nil)
		So(w, ShouldNotBeNil)
		So(w.IsEmpty(), ShouldBeTrue)
	})

	Convey("data model with HasError rebuilds the error", t, func() {
		model := &errorwrapper.WrapperDataModel{
			HasError:           true,
			CurrentError:       "rebuilt",
			IsDisplayableError: true,
			ErrorType:          errtype.InvalidInput,
		}
		w := errorwrapper.NewFromDataModel(model)
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
		So(w.ErrorString(), ShouldContainSubstring, "rebuilt")
	})
}

// Test_SimpleReferencesCompile — formats type + references.
func Test_SimpleReferencesCompile(t *testing.T) {
	Convey("Without references returns name/code", t, func() {
		out := errorwrapper.SimpleReferencesCompile(errtype.InvalidInput)
		So(out, ShouldNotBeBlank)
	})
	Convey("With references appends compiled string", t, func() {
		out := errorwrapper.SimpleReferencesCompile(errtype.InvalidInput, "k", 1)
		So(out, ShouldNotBeBlank)
	})
}

// Test_SimpleReferencesCompileOptimized — short variant.
func Test_SimpleReferencesCompileOptimized(t *testing.T) {
	Convey("Empty refs return variant name only", t, func() {
		out := errorwrapper.SimpleReferencesCompileOptimized(errtype.InvalidInput)
		So(out, ShouldNotBeBlank)
	})
	Convey("With refs returns name + compiled refs", t, func() {
		out := errorwrapper.SimpleReferencesCompileOptimized(errtype.InvalidInput, "x", 2)
		So(out, ShouldNotBeBlank)
	})
}

// Test_MessagesJoined — joins with the package joiner.
func Test_MessagesJoined(t *testing.T) {
	Convey("Joins messages with single space", t, func() {
		So(errorwrapper.MessagesJoined("a", "b"), ShouldEqual, "a b")
		So(errorwrapper.MessagesJoined(), ShouldEqual, "")
	})
}

// Test_New_TypeConstructors — basic type-only wrappers.
func Test_New_TypeConstructors(t *testing.T) {
	Convey("New / NewPtr / NewPtrUsingStackSkip / NewTypeUsingStackSkip", t, func() {
		So(errorwrapper.New(errtype.InvalidInput).Type(), ShouldEqual, errtype.InvalidInput)
		So(errorwrapper.NewPtr(errtype.InvalidInput).Type(), ShouldEqual, errtype.InvalidInput)
		So(errorwrapper.NewPtrUsingStackSkip(0, errtype.InvalidInput), ShouldNotBeNil)
		So(errorwrapper.NewTypeUsingStackSkip(0, errtype.InvalidInput), ShouldNotBeNil)
	})
}

// Test_NewMessage_Constructors — message-based constructors.
func Test_NewMessage_Constructors(t *testing.T) {
	Convey("NewMsgDisplayError variants", t, func() {
		So(errorwrapper.NewMsgDisplayError(0, errtype.InvalidInput, true, "m"), ShouldNotBeNil)
		So(errorwrapper.NewMsgDisplayErrorNoReference(0, errtype.InvalidInput, "m"), ShouldNotBeNil)
		So(errorwrapper.NewMessagesUsingJoiner(0, errtype.InvalidInput, " | ", "a", "b"), ShouldNotBeNil)
		So(errorwrapper.NewUnknownMessage(0, true, "u"), ShouldNotBeNil)
		So(errorwrapper.NewGeneric(0, errors.New("g")), ShouldNotBeNil)
	})
}

// Test_NewError_Constructors — error-based constructors.
func Test_NewError_Constructors(t *testing.T) {
	Convey("NewError / NewUsingError variants", t, func() {
		err := errors.New("e")
		So(errorwrapper.NewError(0, errtype.InvalidInput, err), ShouldNotBeNil)
		So(errorwrapper.NewUsingError(0, errtype.InvalidInput, err), ShouldNotBeNil)
		So(errorwrapper.NewUsingErrorWithoutTypeDisplay(0, errtype.InvalidInput, err), ShouldNotBeNil)
		So(errorwrapper.NewUsingErrorWithoutTypeDisplayPtr(0, errtype.InvalidInput, err), ShouldNotBeNil)
		So(errorwrapper.NewUsingTypeErrorAndMessage(0, errtype.InvalidInput, err, "m"), ShouldNotBeNil)
		So(errorwrapper.NewUsingErrorAndMessage(0, errtype.InvalidInput, err, "m"), ShouldNotBeNil)
	})

	Convey("nil error short-circuits to nil", t, func() {
		So(errorwrapper.NewError(0, errtype.InvalidInput, nil), ShouldBeNil)
	})
}

// Test_NewRef_Constructors — reference-based constructors.
func Test_NewRef_Constructors(t *testing.T) {
	Convey("Reference constructors produce non-nil wrappers", t, func() {
		So(errorwrapper.NewRef(0, errtype.InvalidInput, errors.New("e"), "k", 1), ShouldNotBeNil)
		So(errorwrapper.NewRefOne(0, errtype.InvalidInput, "k", 1), ShouldNotBeNil)
		So(errorwrapper.NewRef1Msg(0, errtype.InvalidInput, "m", "k", 1), ShouldNotBeNil)
		So(errorwrapper.NewRef2Msg(0, errtype.InvalidInput, "m", "k1", 1, "k2", 2), ShouldNotBeNil)
		So(errorwrapper.NewOnlyRefs(0, errtype.InvalidInput, ref.New("k", "v")), ShouldNotBeNil)
		So(errorwrapper.NewRefs(0, errtype.InvalidInput, errors.New("e"), ref.New("k", "v")), ShouldNotBeNil)
		So(errorwrapper.NewRefWithMessage(0, errtype.InvalidInput, "m", ref.New("k", "v")), ShouldNotBeNil)
		So(errorwrapper.TypeReferenceQuick(0, errtype.InvalidInput, ref.New("k", "v")), ShouldNotBeNil)
		So(errorwrapper.NewRef2(0, errtype.InvalidInput, errors.New("e"), "k1", 1, "k2", 2), ShouldNotBeNil)
		So(errorwrapper.NewErrorRef1(0, errtype.InvalidInput, errors.New("e"), "k", 1), ShouldNotBeNil)
	})

	Convey("Nil error returns nil for NewRef", t, func() {
		So(errorwrapper.NewRef(0, errtype.InvalidInput, nil, "k", 1), ShouldBeNil)
	})
}

// Test_NewPath_Constructors — path-based constructors.
func Test_NewPath_Constructors(t *testing.T) {
	Convey("Path constructors produce non-nil wrappers", t, func() {
		So(errorwrapper.NewPath(0, errtype.InvalidInput, errors.New("e"), "/p"), ShouldNotBeNil)
		So(errorwrapper.NewPathMsg(0, errtype.InvalidInput, errors.New("e"), "/p", "m"), ShouldNotBeNil)
		So(errorwrapper.NewPathMessages(0, errtype.InvalidInput, "/p", "m1", "m2"), ShouldNotBeNil)
	})
}

// Test_NewUsingWrapper — wraps an existing wrapper.
func Test_NewUsingWrapper(t *testing.T) {
	Convey("nil/empty returns nil", t, func() {
		So(errorwrapper.NewUsingWrapper(0, nil), ShouldBeNil)
		So(errorwrapper.NewUsingWrapper(0, errorwrapper.EmptyPtr()), ShouldBeNil)
	})
	Convey("non-empty wrapper produces a derived wrapper", t, func() {
		base := errnew.Messages.Single(errtype.InvalidInput, "x")
		out := errorwrapper.NewUsingWrapper(0, base, ref.New("k", "v"))
		So(out, ShouldNotBeNil)
		So(out.HasError(), ShouldBeTrue)
	})
}

// Test_Wrapper_Readers — exercise inspector methods.
func Test_Wrapper_Readers(t *testing.T) {
	Convey("Inspector methods report consistent state for a populated wrapper", t, func() {
		w := errnew.Messages.Single(errtype.InvalidInput, "boom")
		So(w.HasError(), ShouldBeTrue)
		So(w.HasAnyError(), ShouldBeTrue)
		So(w.HasAnyIssues(), ShouldBeTrue)
		So(w.HasCurrentError(), ShouldBeTrue)
		So(w.IsEmpty(), ShouldBeFalse)
		So(w.IsEmptyError(), ShouldBeFalse)
		So(w.IsDefined(), ShouldBeTrue)
		So(w.IsInvalid(), ShouldBeTrue)
		So(w.IsSuccess(), ShouldBeFalse)
		So(w.IsValid(), ShouldBeFalse)
		So(w.IsFailed(), ShouldBeTrue)
		So(w.IsNoError(), ShouldBeFalse)
		So(w.IsNull(), ShouldBeFalse)
		So(w.IsAnyNull(), ShouldBeFalse)
		So(w.IsTypeOf(errtype.InvalidInput), ShouldBeTrue)
		So(w.IsTypeOf(errtype.NotFound), ShouldBeFalse)
		So(w.IsCollectionType(), ShouldBeFalse)

		So(w.Type(), ShouldEqual, errtype.InvalidInput)
		So(w.TypeName(), ShouldNotBeBlank)
		So(w.TypeString(), ShouldNotBeBlank)
		So(w.TypeNameCode(), ShouldNotBeBlank)
		So(w.TypeCodeNameString(), ShouldNotBeBlank)
		So(w.TypeNameCodeMessage(), ShouldNotBeBlank)
		So(w.TypeNameWithCustomMessage("ctx"), ShouldNotBeBlank)
		So(w.RawErrorTypeName(), ShouldNotBeBlank)
		So(w.CodeTypeName(), ShouldNotBeBlank)
		So(w.RawErrorTypeValue(), ShouldNotEqual, uint16(0))
		So(w.ErrorTypeAsBasicErrorTyper(), ShouldNotBeNil)
		So(w.GetTypeVariantStruct().Name, ShouldNotBeBlank)

		So(w.Message(), ShouldContainSubstring, "boom")
		So(w.ErrorString(), ShouldContainSubstring, "boom")
		So(w.Error(), ShouldNotBeNil)
		So(w.Value(), ShouldNotBeNil)
		So(w.CompiledError(), ShouldNotBeNil)
		So(w.CompiledErrorWithStackTraces(), ShouldNotBeNil)
		So(w.Compile(), ShouldNotBeBlank)
		So(w.CompileString(), ShouldNotBeBlank)
		So(w.String(), ShouldNotBeBlank)
		So(w.StringIf(false), ShouldNotBeBlank)
		So(w.FullString(), ShouldContainSubstring, "boom")
		So(w.FullStringWithTraces(), ShouldContainSubstring, "boom")
		So(w.FullStringWithLimitTraces(2), ShouldContainSubstring, "boom")
		So(w.FullStringWithTracesIf(true), ShouldContainSubstring, "boom")
		So(w.FullStringWithoutReferences(), ShouldContainSubstring, "boom")
		So(w.FullOrErrorMessage(false), ShouldNotBeBlank)
		So(len(w.FullStringSplitByNewLine()), ShouldBeGreaterThan, 0)

		So(w.StackTraceString(), ShouldNotBeBlank)
		So(w.StackTraces(), ShouldNotBeBlank)
		So(w.NewDefaultStackTraces(), ShouldNotBeBlank)
		So(w.NewStackTraces(0), ShouldNotBeBlank)
		So(w.CompiledStackTracesString(), ShouldNotBeBlank)
		So(w.StackTracesLimit(1), ShouldNotBeNil)
		So(w.StackTracesJsonResult(), ShouldNotBeNil)
		So(w.NewDefaultStackTracesJsonResult(), ShouldNotBeNil)
		So(w.NewStackTracesJsonResult(0), ShouldNotBeNil)
	})
}

// Test_Wrapper_Serialize — serialization helpers.
func Test_Wrapper_Serialize(t *testing.T) {
	Convey("Serialize variants produce non-empty bytes", t, func() {
		w := errnew.Messages.Single(errtype.InvalidInput, "boom")
		b, err := w.Serialize()
		So(err, ShouldBeNil)
		So(b, ShouldNotBeEmpty)
		So(w.SerializeMust(), ShouldNotBeEmpty)
		b2, err := w.SerializeWithoutTraces()
		So(err, ShouldBeNil)
		So(b2, ShouldNotBeEmpty)

		So(w.Json().IsEmpty(), ShouldBeFalse)
		So(w.JsonPtr(), ShouldNotBeNil)
		So(w.JsonResultWithoutTraces(), ShouldNotBeNil)
		So(w.JsonModel(), ShouldNotBeNil)
		So(w.JsonModelAny(), ShouldNotBeNil)

		mj, err := w.MarshalJSON()
		So(err, ShouldBeNil)
		So(mj, ShouldNotBeEmpty)

		So(w.CompiledJsonErrorWithStackTraces(), ShouldNotBeNil)
		So(w.CompiledJsonStringWithStackTraces(), ShouldNotBeBlank)
	})
}

// Test_Wrapper_Compare — equality / contains.
func Test_Wrapper_Compare(t *testing.T) {
	Convey("IsEquals / IsNotEquals / IsErrorEquals / IsErrorMessageEqual / IsErrorMessageContains", t, func() {
		w1 := errnew.Messages.Single(errtype.InvalidInput, "abc")
		w2 := errnew.Messages.Single(errtype.InvalidInput, "abc")
		So(w1.IsEquals(w2), ShouldBeTrue)
		So(w1.IsNotEquals(w2), ShouldBeFalse)
		So(w1.IsErrorMessageEqual("abc"), ShouldBeTrue)
		So(w1.IsErrorMessage("ABC", false), ShouldBeTrue)
		So(w1.IsErrorMessage("ABC", true), ShouldBeFalse)
		So(w1.IsErrorMessageContains("ab", true), ShouldBeTrue)
		So(w1.IsErrorEquals(errors.New("abc")), ShouldBeTrue)
	})
}

// Test_Wrapper_Clone — ClonePtr/Clone/NonPtr/Ptr.
func Test_Wrapper_Clone(t *testing.T) {
	Convey("Clone variants produce independent wrappers", t, func() {
		w := errnew.Messages.Single(errtype.InvalidInput, "x")
		So(w.ClonePtr(), ShouldNotBeNil)
		c := w.Clone()
		So(c.HasError(), ShouldBeTrue)
		So(w.NonPtr().HasError(), ShouldBeTrue)
		So(w.Ptr(), ShouldNotBeNil)
		So(w.CloneNewStackSkipPtr(0), ShouldNotBeNil)
		So(w.AsErrorWrapper(), ShouldNotBeNil)
		So(w.AsJsonContractsBinder(), ShouldNotBeNil)
		So(w.AsErrWrapperContractsBinder(), ShouldNotBeNil)
		So(w.CloneInterface(), ShouldNotBeNil)
	})
}

// Test_Wrapper_References — references and refs collection helpers.
func Test_Wrapper_References(t *testing.T) {
	Convey("References accessors return non-nil for ref-laden wrapper", t, func() {
		w := errorwrapper.NewRefWithMessage(0, errtype.InvalidInput, "m", ref.New("k", "v"))
		So(w.HasReferences(), ShouldBeTrue)
		So(w.IsReferencesEmpty(), ShouldBeFalse)
		So(w.References(), ShouldNotBeNil)
		So(w.CloneReferences(), ShouldNotBeNil)
		So(w.ReferencesCollection(), ShouldNotBeNil)
		So(w.ReferencesCompiledString(), ShouldNotBeBlank)
		So(len(w.ReferencesList()), ShouldBeGreaterThan, 0)
		So(w.MergeNewReferences(ref.New("k2", "v2")), ShouldNotBeNil)
	})
}

// Test_EmptyPtr_Inspectors — empty sentinel reports safe zero state.
func Test_EmptyPtr_Inspectors(t *testing.T) {
	Convey("Empty wrapper inspectors", t, func() {
		w := errorwrapper.EmptyPtr()
		So(w.HasError(), ShouldBeFalse)
		So(w.HasAnyError(), ShouldBeFalse)
		So(w.HasAnyIssues(), ShouldBeFalse)
		So(w.IsEmpty(), ShouldBeTrue)
		So(w.IsEmptyError(), ShouldBeTrue)
		So(w.IsSuccess(), ShouldBeTrue)
		So(w.IsValid(), ShouldBeTrue)
		So(w.IsFailed(), ShouldBeFalse)
		So(w.IsNoError(), ShouldBeTrue)
		So(w.Error(), ShouldBeNil)
		So(w.CompiledError(), ShouldBeNil)
	})
}
