package refstests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
	"github.com/alimtvnetwork/errorwrapper-v3/refs"
)

func Test_Ref_Conversions(t *testing.T) {
	v := ref.New("k", "v")

	Convey("Value getters", t, func() {
		So(v.VarName(), ShouldEqual, "k")
		So(v.VariableName(), ShouldEqual, "k")
		So(v.ValueDynamic(), ShouldEqual, "v")
		So(v.ValueString(), ShouldNotBeBlank)
		So(v.String(), ShouldContainSubstring, "k")
		So(v.FullString(), ShouldContainSubstring, "k")
		So(v.StringWithoutType(), ShouldContainSubstring, "k")
		vn, val := v.VariableValueString()
		So(vn, ShouldEqual, "k")
		So(val, ShouldNotBeBlank)
		vn2, val2 := v.VariableValueDynamic()
		So(vn2, ShouldEqual, "k")
		So(val2, ShouldEqual, "v")
	})

	Convey("Pointer-receiver helpers", t, func() {
		p := v.ToPtr()
		So(p, ShouldNotBeNil)
		So(p.KeyName(), ShouldEqual, "k")
		So(p.ValueAny(), ShouldEqual, "v")
		So(p.IsVariableNameEqual("k"), ShouldBeTrue)
		So(p.IsVariableNameEqual("z"), ShouldBeFalse)
		So(p.IsAnyValueEqual("v"), ShouldBeTrue)
		So(p.IsAnyValueEqual("x"), ShouldBeFalse)
	})

	Convey("Clone helpers", t, func() {
		c := v.Clone()
		So(c.Variable, ShouldEqual, v.Variable)
		cp := v.ToPtr().ClonePtr()
		So(cp, ShouldNotBeNil)
		So(cp.Variable, ShouldEqual, v.Variable)
	})

	Convey("Serialize + json", t, func() {
		b, err := v.Serialize()
		So(err, ShouldBeNil)
		So(b, ShouldNotBeEmpty)
		So(v.SerializeMust(), ShouldNotBeEmpty)
		jb, jerr := v.MarshalJSON()
		So(jerr, ShouldBeNil)
		So(jb, ShouldNotBeEmpty)
		So(v.Json(), ShouldNotBeNil)
		So(v.JsonPtr(), ShouldNotBeNil)
		So(v.JsonModelAny(), ShouldNotBeNil)
	})

	Convey("DataModel round-trip", t, func() {
		p := v.ToPtr()
		dm := p.ToDataModel()
		So(dm.VariableName, ShouldEqual, "k")
		from := ref.NewFromDataModelPtr(&dm)
		So(from.Variable, ShouldEqual, "k")
		So(ref.NewFromDataModelPtr(nil), ShouldNotBeNil)
	})

	Convey("NewPtr + IsEqualPtr", t, func() {
		a := ref.NewPtr("k", "v")
		b := ref.NewPtr("k", "v")
		So(a.IsEqualPtr(b), ShouldBeTrue)
	})

	Convey("Interface adapters", t, func() {
		So(v.AsJsoner(), ShouldNotBeNil)
		So(v.AsJsonMarshaller(), ShouldNotBeNil)
		So(v.AsJsonContractsBinder(), ShouldNotBeNil)
		So(v.AsJsonParseSelfInjector(), ShouldNotBeNil)
		So(v.AsReferencer(), ShouldNotBeNil)
		So(v.AsKeyAnyValueDefinerBinder(), ShouldNotBeNil)
	})
}

func Test_Refs_Constructors(t *testing.T) {
	Convey("Empty + New variants", t, func() {
		e := refs.Empty()
		So(e.IsEmpty(), ShouldBeTrue)
		So(refs.EmptyPtr().IsEmpty(), ShouldBeTrue)
		So(refs.New(4).Count(), ShouldEqual, 0)
		So(refs.New2().Count(), ShouldEqual, 0)
		So(refs.New4().Count(), ShouldEqual, 0)
		So(refs.NewWithItem(2, "k", "v").Count(), ShouldEqual, 1)
		So(refs.NewDirectItem("k", "v").Count(), ShouldEqual, 1)
	})

	Convey("NewUsingValues + NewUsingMap + NewUsingMany", t, func() {
		c := refs.NewUsingValues(ref.New("a", 1), ref.New("b", 2))
		So(c, ShouldNotBeNil)
		So(c.Count(), ShouldEqual, 2)

		m := refs.NewUsingMap(map[string]interface{}{"k": "v"})
		So(m.Count(), ShouldEqual, 1)

		manyC := refs.NewUsingMany([]ref.Value{ref.New("a", 1)}, []ref.Value{ref.New("b", 2)})
		So(manyC.Count(), ShouldEqual, 2)
	})

	Convey("NewUsingValues nil + NewUsingMap empty", t, func() {
		So(refs.NewUsingValues(), ShouldBeNil)
		So(refs.NewUsingMap(nil).IsEmpty(), ShouldBeTrue)
	})
}

func Test_Refs_Collection_Mutations(t *testing.T) {
	Convey("Add + Adds + AddsIf + AddMap", t, func() {
		c := refs.EmptyPtr()
		c.Add("a", 1)
		c.Adds(ref.New("b", 2), ref.New("c", 3))
		c.AddsIf(true, ref.New("d", 4))
		c.AddsIf(false, ref.New("skip", 99))
		c.AddMap(map[string]interface{}{"e": 5})
		So(c.Count(), ShouldEqual, 5)
	})

	Convey("AddCollection + AddCollectionCloned + ConcatNew", t, func() {
		c1 := refs.NewDirectItem("a", 1)
		c2 := refs.NewDirectItem("b", 2)
		c1.AddCollection(c2)
		So(c1.Count(), ShouldEqual, 2)

		c3 := refs.NewDirectItem("c", 3)
		merged := c1.ConcatNew(false, c3)
		So(merged.Count(), ShouldBeGreaterThanOrEqualTo, 3)
	})

	Convey("AddVarVal + AddReferencer", t, func() {
		c := refs.EmptyPtr()
		c.AddVarVal("a", 1)
		other := ref.New("b", 2)
		c.AddReferencer(&other)
		So(c.Count(), ShouldEqual, 2)
	})

	Convey("Lists + Length + Strings + Items", t, func() {
		c := refs.NewWithItem(2, "k", "v")
		c.Add("k2", "v2")
		So(c.Length(), ShouldEqual, 2)
		So(c.Items(), ShouldHaveLength, 2)
		So(c.List(), ShouldHaveLength, 2)
		So(c.Collection(), ShouldHaveLength, 2)
		So(c.Strings(), ShouldHaveLength, 2)
		So(c.HasAnyItem(), ShouldBeTrue)
		So(c.IsNull(), ShouldBeFalse)
	})

	Convey("Maps", t, func() {
		c := refs.EmptyPtr().Add("a", 1).Add("b", "x")
		ms := c.MapStringString()
		So(ms["a"], ShouldNotBeBlank)
		So(c.MapStringAny()["b"], ShouldEqual, "x")
	})

	Convey("Compile + String + Clone + Dispose", t, func() {
		c := refs.NewDirectItem("k", "v")
		So(c.Compile(), ShouldContainSubstring, "k")
		So(c.String(), ShouldContainSubstring, "k")
		So(c.Clone().Count(), ShouldEqual, 1)
		So(c.ClonePtr().Count(), ShouldEqual, 1)

		c2 := refs.NewDirectItem("k", "v")
		c2.Dispose()
	})

	Convey("Serialize + Json", t, func() {
		c := refs.NewDirectItem("k", "v")
		b, err := c.Serialize()
		So(err, ShouldBeNil)
		So(b, ShouldNotBeEmpty)
		So(c.SerializeMust(), ShouldNotBeEmpty)
		mj, err2 := c.MarshalJSON()
		So(err2, ShouldBeNil)
		So(mj, ShouldNotBeEmpty)
		So(c.JsonModel(), ShouldNotBeNil)
		So(c.JsonModelAny(), ShouldNotBeNil)
		So(c.Json(), ShouldNotBeNil)
		So(c.JsonPtr(), ShouldNotBeNil)
	})

	Convey("IsEqual self + nil", t, func() {
		a := refs.NewDirectItem("k", "v")
		b := refs.NewDirectItem("k", "v")
		So(a.IsEqual(b), ShouldBeTrue)
		var n *refs.Collection
		So(n.IsEqual(nil), ShouldBeTrue)
	})

	Convey("Interface adapters", t, func() {
		c := refs.NewDirectItem("k", "v")
		So(c.AsJsoner(), ShouldNotBeNil)
		So(c.AsJsonMarshaller(), ShouldNotBeNil)
		So(c.AsJsonContractsBinder(), ShouldNotBeNil)
		So(c.AsJsonParseSelfInjector(), ShouldNotBeNil)
	})
}

func Test_Refs_TopLevelHelpers(t *testing.T) {
	Convey("AllIndividualItemsCount", t, func() {
		a := refs.NewDirectItem("a", 1)
		b := refs.NewDirectItem("b", 2)
		So(refs.AllIndividualItemsCount(a, b), ShouldEqual, 2)
		So(refs.AllIndividualItemsCount(nil), ShouldEqual, 0)
	})

	Convey("LengthOfEachItems", t, func() {
		manyC := [][]ref.Value{
			{ref.New("a", 1)},
			{ref.New("b", 2), ref.New("c", 3)},
		}
		So(refs.LengthOfEachItems(manyC), ShouldEqual, 3)
	})

	Convey("MergeReferences + PrependReferences", t, func() {
		existing := []ref.Value{ref.New("a", 1)}
		merged := refs.MergeReferences(existing, ref.New("b", 2))
		So(merged, ShouldHaveLength, 2)

		prepended := refs.PrependReferences(false, existing, ref.New("z", 9))
		So(prepended, ShouldHaveLength, 2)
	})

	Convey("CompileAnyItemsToCsvStringDefault", t, func() {
		s := refs.CompileAnyItemsToCsvStringDefault("a", 1, true)
		So(s, ShouldNotBeBlank)
	})

	Convey("QuickReference + QuickCompileString", t, func() {
		q := refs.NewQuickReference(errtype.InvalidInput, "a", 1)
		So(q.CompileLine(), ShouldNotBeBlank)
		line := refs.QuickCompileString(",", q)
		So(line, ShouldNotBeBlank)
	})
}
