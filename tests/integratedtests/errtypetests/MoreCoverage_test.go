package errtypetests

import (
	"encoding/json"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

// Test_Variation_Formatting exercises the format/combine/code helpers across many variations.
func Test_Variation_Formatting(t *testing.T) {
	variations := []errtype.Variation{
		errtype.Generic,
		errtype.NotFound,
		errtype.InvalidInput,
		errtype.MappingFailed,
		errtype.CommandExecution,
		errtype.NullOrEmpty,
		errtype.FileNotExist,
		errtype.PermissionFailed,
	}

	for _, v := range variations {
		v := v
		Convey("Variation "+v.String()+" formatting helpers", t, func() {
			So(v.TypeNameCodeMessage(), ShouldNotBeBlank)
			So(v.TypeNameCodeMessageRef("ref-x"), ShouldContainSubstring, "ref-x")
			So(v.CodeWithTypeName(), ShouldContainSubstring, v.Name())
			So(v.CodeTypeNameWithCustomMessage("custom"), ShouldContainSubstring, "custom")
			So(v.ExplicitCodeValueName(), ShouldContainSubstring, v.Name())
			So(v.CodeTypeNameWithReference("ref-line"), ShouldContainSubstring, "ref-line")
			So(v.CodeTypeNameWithReferences("a", "b"), ShouldContainSubstring, v.Name())
			So(v.MessageToRawErrType().String(), ShouldEqual, v.Message())
			So(v.TypeMessage(), ShouldEqual, v.Message())
			So(v.CategoryName(), ShouldEqual, v.Name())
			So(v.TypenameString(), ShouldEqual, v.Name())
			So(v.TypeName(), ShouldNotBeBlank)
		})
	}
}

func Test_Variation_References(t *testing.T) {
	v := errtype.NotFound

	Convey("ReferencesCsv with and without additional message", t, func() {
		So(v.ReferencesCsv("", "a", "b"), ShouldContainSubstring, v.Name())
		So(v.ReferencesCsv("msg", 1, 2), ShouldContainSubstring, "msg")
	})

	Convey("ReferencesLines with and without additional message", t, func() {
		So(v.ReferencesLines("", "x", "y"), ShouldContainSubstring, v.Name())
		So(v.ReferencesLines("ctx", "x"), ShouldContainSubstring, "ctx")
	})

	Convey("ReferencesCsvError and ReferencesLinesError", t, func() {
		So(v.ReferencesCsvError("ctx", "a"), ShouldNotBeNil)
		So(v.ReferencesLinesError("ctx", "a"), ShouldNotBeNil)
	})

	Convey("ShortReferencesCsv returns empty when no refs, non-empty otherwise", t, func() {
		So(v.ShortReferencesCsv(), ShouldBeBlank)
		So(v.ShortReferencesCsvError(), ShouldBeNil)
		So(v.ShortReferencesCsv("a"), ShouldNotBeBlank)
		So(v.ShortReferencesCsvError("a"), ShouldNotBeNil)
	})
}

func Test_Variation_Numeric(t *testing.T) {
	v := errtype.InvalidInput

	Convey("Numeric conversions produce consistent values", t, func() {
		So(v.ValueInt(), ShouldEqual, int(v))
		So(v.ValueInt16(), ShouldEqual, int16(v))
		So(v.ValueInt32(), ShouldEqual, int32(v))
		So(v.ValueUInt(), ShouldEqual, uint(v))
		So(v.ValueUInt16(), ShouldEqual, uint16(v))
		So(v.Value(), ShouldEqual, uint16(v))
		So(v.ValueString(), ShouldNotBeBlank)
		So(v.ToNumberString(), ShouldEqual, v.ValueString())
		
		So(v.MinInt(), ShouldEqual, 0)
		So(v.MaxInt(), ShouldBeGreaterThan, 0)
		So(v.MinValueString(), ShouldNotBeBlank)
		So(v.MaxValueString(), ShouldNotBeBlank)
	})
}

func Test_Variation_NameAndEquality(t *testing.T) {
	v := errtype.NotFound

	Convey("Name + equality helpers", t, func() {
		So(v.IsNameEqual(v.Name()), ShouldBeTrue)
		So(v.IsNameEqual("nope-xyz"), ShouldBeFalse)
		So(v.IsEqualVariant(v), ShouldBeTrue)
		So(v.IsEqualVariant(errtype.Generic), ShouldBeFalse)
		So(v.IsAnyNamesOf(v.Name(), "other"), ShouldBeTrue)
		So(v.IsAnyNamesOf("a", "b"), ShouldBeFalse)
		So(v.NameValue(), ShouldNotBeBlank)
	})
}

func Test_Variation_ValidInvalid(t *testing.T) {
	Convey("NoError is Valid; others are Invalid", t, func() {
		So(errtype.NoError.IsValid(), ShouldBeTrue)
		So(errtype.NoError.IsInvalid(), ShouldBeFalse)
		So(errtype.NoError.IsEmptyError(), ShouldBeTrue)
		So(errtype.NotFound.IsValid(), ShouldBeFalse)
		So(errtype.NotFound.IsInvalid(), ShouldBeTrue)
		So(errtype.NotFound.IsEmptyError(), ShouldBeFalse)
	})
}

func Test_Variation_OnlySupported(t *testing.T) {
	v := errtype.NotFound

	Convey("OnlySupportedErr nil when allowed, error otherwise", t, func() {
		So(v.OnlySupportedErr(v.Name()), ShouldBeNil)
		So(v.OnlySupportedErr("other"), ShouldNotBeNil)
		So(v.OnlySupportedMsgErr("msg", v.Name()), ShouldBeNil)
		So(v.OnlySupportedMsgErr("msg", "other"), ShouldNotBeNil)
	})
}

func Test_Variation_Json(t *testing.T) {
	v := errtype.NotFound

	Convey("JsonModel + Marshal/Unmarshal roundtrip", t, func() {
		model := v.JsonModel()
		So(model.Category, ShouldEqual, v.Name())
		So(model.Id, ShouldEqual, v.ValueUInt16())

		bytes, err := v.MarshalJSON()
		So(err, ShouldBeNil)
		So(len(bytes), ShouldBeGreaterThan, 0)

		serialized, err := v.Serialize()
		So(err, ShouldBeNil)
		So(string(serialized), ShouldEqual, string(bytes))

		mustBytes := v.SerializeMust()
		So(string(mustBytes), ShouldEqual, string(bytes))

		var roundTrip errtype.Variation
		err = json.Unmarshal(bytes, &roundTrip)
		So(err, ShouldBeNil)
		So(roundTrip, ShouldEqual, v)

		any := v.JsonModelAny()
		So(any, ShouldNotBeNil)
	})
}

func Test_Variation_Ranges(t *testing.T) {
	v := errtype.NotFound

	Convey("Range helpers return populated structures", t, func() {
		So(v.IntegerEnumRanges(), ShouldNotBeEmpty)
		So(v.AllNameValues(), ShouldNotBeEmpty)
		So(v.RangeNamesCsv(), ShouldNotBeBlank)
		So(v.RangesDynamicMap(), ShouldNotBeEmpty)
		min, max := v.MinMaxAny()
		So(min, ShouldNotBeNil)
		So(max, ShouldNotBeNil)
	})
}

func Test_Variation_VariantStructure(t *testing.T) {
	v := errtype.NotFound

	Convey("VariantStructure + pointer share name/message", t, func() {
		vs := v.VariantStructure()
		So(vs.Name, ShouldEqual, v.Name())
		vsPtr := v.VariantStructurePtr()
		So(vsPtr, ShouldNotBeNil)
		So(vsPtr.Name, ShouldEqual, v.Name())
		So(vs.CombineNoRefs("msg"), ShouldContainSubstring, "msg")
		So(vs.Combine("msg", "k", "v"), ShouldContainSubstring, "msg")
	})
}

func Test_Variation_Format(t *testing.T) {
	v := errtype.NotFound

	Convey("Format substitutes name and value tokens", t, func() {
		out := v.Format("{name}-{value}")
		So(strings.ToLower(out), ShouldContainSubstring, strings.ToLower(v.Name()))
	})
}

func Test_Variation_Interfaces(t *testing.T) {
	v := errtype.NotFound

	Convey("Interface accessors are non-nil", t, func() {
		So(v.BaseErrorTyper(), ShouldNotBeNil)
		So(v.ErrTypeDetailDefiner(), ShouldNotBeNil)
		So(v.EnumType(), ShouldNotBeNil)
		So(v.ErrorTypeAsBasicEnum(), ShouldNotBeNil)
		So(v.AsBasicErrorTyper(), ShouldNotBeNil)
		So(v.IsErrorTyperEqual(v.BaseErrorTyper()), ShouldBeTrue)
		So(v.IsErrorTyperEqual(nil), ShouldBeFalse)
	})
}

func Test_Variation_Panic(t *testing.T) {
	v := errtype.InvalidInput

	Convey("Panic helpers panic with the formatted message", t, func() {
		So(func() { v.Panic("boom", "k", 1) }, ShouldPanic)
		So(func() { v.PanicNoRefs("boom-no-refs") }, ShouldPanic)
	})
}
