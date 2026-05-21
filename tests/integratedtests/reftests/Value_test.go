package reftests

import (
	"testing"

	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Ref_Value_Constructors(t *testing.T) {
	Convey("New / NewPtr / ToPtr / ToNonPtr round-trips", t, func() {
		v := ref.New("name", "alice")
		So(v.VarName(), ShouldEqual, "name")
		So(v.VariableName(), ShouldEqual, "name")
		So(v.ValueDynamic(), ShouldEqual, "alice")

		p := v.ToPtr()
		So(p, ShouldNotBeNil)
		So(p.KeyName(), ShouldEqual, "name")
		So(p.ValueAny(), ShouldEqual, "alice")

		So(ref.NewPtr("k", 42).KeyName(), ShouldEqual, "k")

		again := p.ToNonPtr()
		So(again.VarName(), ShouldEqual, "name")
	})

	Convey("NewUsingReferencer with nil returns zero value", t, func() {
		v := ref.NewUsingReferencer(nil)
		So(v.VarName(), ShouldEqual, "")
	})

	Convey("NewUsingKeyAnyVal with nil returns zero value", t, func() {
		v := ref.NewUsingKeyAnyVal(nil)
		So(v.VarName(), ShouldEqual, "")
	})

	Convey("NewUsingReferencer with a Value referencer copies fields", t, func() {
		src := ref.New("k", "v")
		v := ref.NewUsingReferencer(src.AsReferencer())
		So(v.VarName(), ShouldEqual, "k")
		So(v.ValueDynamic(), ShouldEqual, "v")
	})

	Convey("NewUsingKeyAnyVal with a Value definer copies fields", t, func() {
		src := ref.New("k", 7)
		v := ref.NewUsingKeyAnyVal(src.ToPtr())
		So(v.VarName(), ShouldEqual, "k")
		So(v.ValueDynamic(), ShouldEqual, 7)
	})

	Convey("NewFromDataModelPtr handles nil and populated", t, func() {
		empty := ref.NewFromDataModelPtr(nil)
		So(empty, ShouldNotBeNil)
		So(empty.VarName(), ShouldEqual, "")

		model := &ref.ValueDataModel{VariableName: "k", ValueString: "v"}
		got := ref.NewFromDataModelPtr(model)
		So(got.VarName(), ShouldEqual, "k")
		So(got.ValueDynamic(), ShouldEqual, "v")
	})
}

func Test_Ref_Value_Equality(t *testing.T) {
	Convey("IsVariableNameEqual", t, func() {
		v := ref.New("k", 1)
		So(v.IsVariableNameEqual("k"), ShouldBeTrue)
		So(v.IsVariableNameEqual("other"), ShouldBeFalse)
	})

	Convey("IsAnyValueEqual handles nil receiver and nil-both", t, func() {
		var nilV *ref.Value
		So(nilV.IsAnyValueEqual(nil), ShouldBeTrue)
		So(nilV.IsAnyValueEqual(1), ShouldBeFalse)

		v := ref.New("k", 1)
		So(v.ToPtr().IsAnyValueEqual(nil), ShouldBeFalse)
		So(v.ToPtr().IsAnyValueEqual(1), ShouldBeTrue)
		So(v.ToPtr().IsAnyValueEqual(2), ShouldBeFalse)

		// DeepEqual path
		s := ref.New("s", []int{1, 2})
		So(s.ToPtr().IsAnyValueEqual([]int{1, 2}), ShouldBeTrue)
	})

	Convey("IsEqual / IsEqualPtr", t, func() {
		a := ref.New("k", 1)
		b := ref.New("k", 1)
		So(a.IsEqual(b), ShouldBeTrue)
		So(a.ToPtr().IsEqualPtr(b.ToPtr()), ShouldBeTrue)

		var nilV *ref.Value
		So(nilV.IsEqualPtr(nil), ShouldBeTrue)
	})

	Convey("IsEqualReferencer / IsEqualKeyAnyValueDefiner", t, func() {
		a := ref.New("k", 1)
		b := ref.New("k", 1)
		So(a.ToPtr().IsEqualReferencer(b.AsReferencer()), ShouldBeTrue)
		So(a.ToPtr().IsEqualReferencer(nil), ShouldBeFalse)

		var nilV *ref.Value
		So(nilV.IsEqualReferencer(nil), ShouldBeTrue)
		So(nilV.IsEqualKeyAnyValueDefiner(nil), ShouldBeTrue)
		So(a.ToPtr().IsEqualKeyAnyValueDefiner(nil), ShouldBeFalse)
		So(a.ToPtr().IsEqualKeyAnyValueDefiner(b.ToPtr()), ShouldBeTrue)
	})
}

func Test_Ref_Value_StringsAndClone(t *testing.T) {
	Convey("String / FullString / StringWithoutType / Compile", t, func() {
		v := ref.New("count", 42)
		full := v.FullString()
		So(full, ShouldContainSubstring, "count")
		So(v.String(), ShouldEqual, full)
		So(v.Compile(), ShouldEqual, full)

		So(v.StringWithoutType(), ShouldContainSubstring, "count")
		// once-initialized cache path
		So(v.StringWithoutType(), ShouldContainSubstring, "count")
		So(v.FullString(), ShouldEqual, full)
	})

	Convey("ValueString / VariableValueString / VariableValueDynamic", t, func() {
		v := ref.New("k", 7)
		So(v.ValueString(), ShouldContainSubstring, "7")

		name, val := v.VariableValueString()
		So(name, ShouldEqual, "k")
		So(val, ShouldContainSubstring, "7")

		name2, dyn := v.VariableValueDynamic()
		So(name2, ShouldEqual, "k")
		So(dyn, ShouldEqual, 7)
	})

	Convey("Clone / ClonePtr", t, func() {
		v := ref.New("k", 7)
		c := v.Clone()
		So(c.VarName(), ShouldEqual, "k")
		So(c.ValueDynamic(), ShouldNotBeNil)

		cp := v.ToPtr().ClonePtr()
		So(cp, ShouldNotBeNil)
		So(cp.VarName(), ShouldEqual, "k")

		var nilV *ref.Value
		So(nilV.ClonePtr(), ShouldBeNil)
	})
}

func Test_Ref_Value_DataModelAndJson(t *testing.T) {
	Convey("ToDataModel / JsonModel / JsonModelAny", t, func() {
		v := ref.New("k", "v").ToPtr()
		dm := v.ToDataModel()
		So(dm.VariableName, ShouldEqual, "k")

		jm := v.JsonModel()
		So(jm.VariableName, ShouldEqual, "k")

		any := ref.New("k", "v").JsonModelAny()
		So(any, ShouldNotBeNil)
	})

	Convey("Serialize / SerializeMust round-trip via UnmarshalJSON", t, func() {
		v := ref.New("k", "v")
		raw, err := v.Serialize()
		So(err, ShouldBeNil)
		So(len(raw), ShouldBeGreaterThan, 0)

		raw2 := v.SerializeMust()
		So(len(raw2), ShouldBeGreaterThan, 0)

		var dst ref.Value
		err = dst.UnmarshalJSON(raw)
		So(err, ShouldBeNil)
		So(dst.VarName(), ShouldEqual, "k")
	})

	Convey("MarshalJSON produces parsable bytes", t, func() {
		v := ref.New("k", "v")
		raw, err := v.MarshalJSON()
		So(err, ShouldBeNil)
		So(len(raw), ShouldBeGreaterThan, 0)
	})

	Convey("Json / JsonPtr / JsonParseSelfInject round-trip", t, func() {
		v := ref.New("k", "v")
		jp := v.JsonPtr()
		So(jp, ShouldNotBeNil)

		var dst ref.Value
		err := dst.JsonParseSelfInject(jp)
		So(err, ShouldBeNil)
		So(dst.VarName(), ShouldEqual, "k")

		j := v.Json()
		So(j.HasError(), ShouldBeFalse)
	})

	Convey("ParseInjectUsingJson handles nil/empty", t, func() {
		var dst ref.Value
		_, err := dst.ParseInjectUsingJson(nil)
		So(err, ShouldNotBeNil)

		empty := &corejson.Result{}
		_, err = dst.ParseInjectUsingJson(empty)
		So(err, ShouldNotBeNil)
	})

	Convey("ParseInjectUsingJsonMust panics on bad payload", t, func() {
		var dst ref.Value
		So(func() { dst.ParseInjectUsingJsonMust(nil) }, ShouldPanic)
	})

	Convey("ParseInjectUsingJsonMust succeeds on valid payload", t, func() {
		src := ref.New("k", "v")
		jp := src.JsonPtr()
		var dst ref.Value
		So(func() { dst.ParseInjectUsingJsonMust(jp) }, ShouldNotPanic)
		So(dst.VarName(), ShouldEqual, "k")
	})

	Convey("Json contracts binders are non-nil", t, func() {
		v := ref.New("k", "v")
		So(v.AsJsonContractsBinder(), ShouldNotBeNil)
		So(v.AsJsoner(), ShouldNotBeNil)
		So(v.AsJsonParseSelfInjector(), ShouldNotBeNil)
		So(v.AsJsonMarshaller(), ShouldNotBeNil)
		So(v.AsReferencer(), ShouldNotBeNil)
		So(v.AsKeyAnyValueDefinerBinder(), ShouldNotBeNil)
	})
}

func Test_Ref_ValueDataModel_NewDataModel(t *testing.T) {
	Convey("NewDataModel with nil returns zero", t, func() {
		dm := ref.NewDataModel(nil)
		So(dm.VariableName, ShouldEqual, "")
		So(dm.ValueString, ShouldEqual, "")
	})

	Convey("NewDataModel populates from Value", t, func() {
		v := ref.New("k", 9).ToPtr()
		dm := ref.NewDataModel(v)
		So(dm.VariableName, ShouldEqual, "k")
		So(dm.ValueString, ShouldContainSubstring, "9")
	})
}

func Test_Ref_Value_ReflectSetTo(t *testing.T) {
	Convey("ReflectSetTo copies into another *Value", t, func() {
		src := ref.New("k", "v")
		var dst ref.Value
		err := src.ToPtr().ReflectSetTo(&dst)
		So(err, ShouldBeNil)
	})
}
