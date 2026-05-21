package errtypetests

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

func Test_Variation_Conversions(t *testing.T) {
	v := errtype.NotFound

	Convey("numeric conversions return matching values", t, func() {
		So(v.Value(), ShouldEqual, v.ValueUInt16())
		So(v.ValueInt(), ShouldEqual, int(v.Value()))
		So(v.ValueUInt(), ShouldEqual, uint(v.Value()))
		So(v.ValueInt16(), ShouldEqual, int16(v.Value()))
		So(v.ValueInt32(), ShouldEqual, int32(v.Value()))
		So(v.ValueString(), ShouldNotBeBlank)
		So(v.ToNumberString(), ShouldNotBeBlank)
	})

	Convey("identity helpers", t, func() {
		So(v.Name(), ShouldEqual, v.String())
		So(v.CategoryName(), ShouldNotBeBlank)
		So(v.TypenameString(), ShouldNotBeBlank)
		So(v.TypeMessage(), ShouldEqual, v.Message())
		So(v.TypeName(), ShouldNotBeBlank)
	})
}

func Test_Variation_Formatters(t *testing.T) {
	v := errtype.InvalidInput

	Convey("formatter helpers produce non-empty strings containing the name", t, func() {
		name := v.Name()
		So(v.TypeNameCodeMessage(), ShouldContainSubstring, name)
		So(v.TypeNameCodeMessageRef("ref-val"), ShouldContainSubstring, "ref-val")
		So(v.CodeWithTypeName(), ShouldContainSubstring, name)
		So(v.CodeTypeNameWithCustomMessage("custom"), ShouldContainSubstring, "custom")
		So(v.CodeTypeNameWithReference("ref-line"), ShouldContainSubstring, "ref-line")
		So(v.CodeTypeNameWithReferences("a", "b"), ShouldNotBeBlank)
		So(v.ExplicitCodeValueName(), ShouldNotBeBlank)
		So(v.RangeNamesCsv(), ShouldNotBeBlank)
	})

	Convey("Reference helpers", t, func() {
		So(v.ReferencesCsv("msg", "a", 1), ShouldContainSubstring, "msg")
		So(v.ReferencesCsv("", "a", 1), ShouldNotBeBlank)
		So(v.ReferencesLines("msg", "a", "b"), ShouldContainSubstring, "msg")
		So(v.ReferencesLines("", "a", "b"), ShouldNotBeBlank)
		So(v.ShortReferencesCsv("msg", "x", 9), ShouldNotBeBlank)

		err1 := v.ReferencesCsvError("ctx", "a", 1)
		So(err1, ShouldNotBeNil)
		err2 := v.ReferencesLinesError("ctx", "a", 1)
		So(err2, ShouldNotBeNil)
		err3 := v.ShortReferencesCsvError("ctx", "a", 1)
		So(err3, ShouldNotBeNil)
	})
}

func Test_Variation_VariantStructure(t *testing.T) {
	v := errtype.NotFound

	Convey("VariantStructure returns populated struct", t, func() {
		vs := v.VariantStructure()
		So(vs.Name, ShouldEqual, v.Name())
		So(vs.Message, ShouldEqual, v.Message())
		So(vs.String(), ShouldNotBeBlank)
		So(vs.TypeNameCodeMessage(), ShouldNotBeBlank)
		So(vs.CodeTypeName(), ShouldContainSubstring, vs.Name)
		So(vs.MessageToRawType(), ShouldNotBeBlank)
	})

	Convey("VariantStructurePtr methods", t, func() {
		vsp := v.VariantStructurePtr()
		So(vsp, ShouldNotBeNil)
		So(vsp.CodeTypeNameWithCustomMessage("hello"), ShouldContainSubstring, "hello")
		So(vsp.CodeTypeNameWithReference("r1"), ShouldContainSubstring, "r1")
		So(vsp.CodeTypeNameWithReferences("a", "b"), ShouldNotBeBlank)
	})

	Convey("VariantStructure combine + references", t, func() {
		vs := v.VariantStructure()
		So(vs.CombineNoRefs(""), ShouldNotBeBlank)
		So(vs.CombineNoRefs("extra"), ShouldContainSubstring, "extra")
		So(vs.CombineRefs("", ""), ShouldNotBeBlank)
		So(vs.CombineRefs("extra", ""), ShouldContainSubstring, "extra")
		So(vs.CombineRefs("", "refs"), ShouldContainSubstring, "refs")
		So(vs.CombineRefs("extra", "refs"), ShouldContainSubstring, "refs")
		So(vs.Combine("ctx", "key", "val"), ShouldContainSubstring, "ctx")
		So(vs.Combine("", "key", "val"), ShouldNotBeBlank)
		So(vs.MsgReferenceValues("m", "v"), ShouldContainSubstring, "m")
		So(vs.ReferenceValues("v"), ShouldContainSubstring, "v")
	})

	Convey("VariantStructure Error builders", t, func() {
		vs := v.VariantStructure()
		err := vs.Error("ctx", "key", "val")
		So(err, ShouldNotBeNil)
		So(strings.ToLower(err.Error()), ShouldContainSubstring, "ctx")

		err2 := vs.ErrorNoRefs("ctx2")
		So(err2, ShouldNotBeNil)
		So(strings.ToLower(err2.Error()), ShouldContainSubstring, "ctx2")
	})
}

func Test_Variation_JsonModel(t *testing.T) {
	v := errtype.NotFound

	Convey("JsonModel reflects variation identity", t, func() {
		jm := v.JsonModel()
		So(jm.Category, ShouldEqual, v.Name())
		So(jm.Id, ShouldEqual, v.ValueUInt16())
		So(jm.HasError(), ShouldBeTrue)
		So(jm.IsNoError(), ShouldBeFalse)
		So(jm.IsTypeOf(v), ShouldBeTrue)
		So(jm.IsCategory(v.Name()), ShouldBeTrue)
		So(jm.IsCategory("not-a-real-category"), ShouldBeFalse)
		So(jm.Name(), ShouldEqual, v.Name())
		So(jm.String(), ShouldEqual, v.Name())
		So(jm.Value(), ShouldEqual, v.ValueUInt16())
		So(jm.Type(), ShouldEqual, v)
		So(v.JsonModelAny(), ShouldNotBeNil)
	})

	Convey("NoError JsonModel reports NoError", t, func() {
		jm := errtype.NoError.JsonModel()
		So(jm.HasError(), ShouldBeFalse)
		So(jm.IsNoError(), ShouldBeTrue)
	})
}

func Test_Variation_Serialize(t *testing.T) {
	v := errtype.InvalidInput

	Convey("Serialize returns bytes and no error", t, func() {
		bytes, err := v.Serialize()
		So(err, ShouldBeNil)
		So(bytes, ShouldNotBeEmpty)

		mustBytes := v.SerializeMust()
		So(mustBytes, ShouldNotBeEmpty)

		jsonBytes, err := v.MarshalJSON()
		So(err, ShouldBeNil)
		So(jsonBytes, ShouldNotBeEmpty)
	})
}

func Test_Variation_IsAnyNamesOf(t *testing.T) {
	v := errtype.NotFound

	Convey("matches its own name", t, func() {
		So(v.IsAnyNamesOf(v.Name()), ShouldBeTrue)
		So(v.IsAnyNamesOf("foo", v.Name(), "bar"), ShouldBeTrue)
		So(v.IsAnyNamesOf("foo", "bar"), ShouldBeFalse)
		So(v.IsAnyNamesOf(), ShouldBeFalse)
	})
}

func Test_StringToVariantMap(t *testing.T) {
	Convey("StringToVariantMap returns populated mapping", t, func() {
		m := errtype.StringToVariantMap()
		So(m, ShouldNotBeNil)
		So(len(m), ShouldBeGreaterThan, 0)
		So(m[errtype.NotFound.Name()], ShouldEqual, errtype.NotFound)
	})
}
