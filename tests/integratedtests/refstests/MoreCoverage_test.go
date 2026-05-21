package refstests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/alimtvnetwork/errorwrapper-v3/ref"
	"github.com/alimtvnetwork/errorwrapper-v3/refs"
)

func Test_MoreCoverage_Ref_Value(t *testing.T) {
	Convey("ref.Value accessors and conversions", t, func() {
		v := ref.New("k", "v")
		So(v.VarName(), ShouldEqual, "k")
		So(v.VariableName(), ShouldEqual, "k")
		So(v.ValueDynamic(), ShouldEqual, "v")
		So(v.ValueString(), ShouldContainSubstring, "v")
		So(v.FullString(), ShouldContainSubstring, "k")
		So(v.StringWithoutType(), ShouldContainSubstring, "k")
		So(v.String(), ShouldNotBeEmpty)
		So(v.Compile(), ShouldNotBeEmpty)

		name, val := v.VariableValueString()
		So(name, ShouldEqual, "k")
		So(val, ShouldContainSubstring, "v")
		n2, d := v.VariableValueDynamic()
		So(n2, ShouldEqual, "k")
		So(d, ShouldEqual, "v")

		ptr := v.ToPtr()
		So(ptr, ShouldNotBeNil)
		So(ptr.KeyName(), ShouldEqual, "k")
		So(ptr.ValueAny(), ShouldEqual, "v")
		So(ptr.IsVariableNameEqual("k"), ShouldBeTrue)
		So(ptr.IsVariableNameEqual("nope"), ShouldBeFalse)
		So(ptr.IsAnyValueEqual("v"), ShouldBeTrue)
		So(ptr.IsAnyValueEqual("other"), ShouldBeFalse)

		clone := v.Clone()
		So(clone.Variable, ShouldEqual, "k")
		cp := v.ToPtr().ClonePtr()
		So(cp, ShouldNotBeNil)

		dm := v.ToPtr().ToDataModel()
		So(dm.VariableName, ShouldEqual, "k")
		jm := v.ToPtr().JsonModel()
		So(jm.VariableName, ShouldEqual, "k")
		So(v.JsonModelAny(), ShouldNotBeNil)

		bytes, err := v.Serialize()
		So(err, ShouldBeNil)
		So(bytes, ShouldNotBeEmpty)
		So(v.SerializeMust(), ShouldNotBeEmpty)
		So(v.Json().IsEmpty(), ShouldBeFalse)
		So(v.JsonPtr(), ShouldNotBeNil)

		So(v.AsJsoner(), ShouldNotBeNil)
		So(v.AsJsonMarshaller(), ShouldNotBeNil)
		So(v.AsJsonContractsBinder(), ShouldNotBeNil)
		So(v.AsJsonParseSelfInjector(), ShouldNotBeNil)
		So(v.AsReferencer(), ShouldNotBeNil)
		So(v.AsKeyAnyValueDefinerBinder(), ShouldNotBeNil)
	})
}

func Test_MoreCoverage_Ref_Constructors(t *testing.T) {
	Convey("ref constructor variants", t, func() {
		So(ref.NewPtr("a", 1), ShouldNotBeNil)
		So(ref.NewUsingReferencer(nil).Variable, ShouldEqual, "")
		So(ref.NewUsingKeyAnyVal(nil).Variable, ShouldEqual, "")
		So(ref.NewFromDataModelPtr(nil), ShouldNotBeNil)
		dm := &ref.ValueDataModel{VariableName: "x", ValueString: "y"}
		got := ref.NewFromDataModelPtr(dm)
		So(got.Variable, ShouldEqual, "x")
	})
}

func Test_MoreCoverage_Ref_Equality(t *testing.T) {
	Convey("ref equality methods", t, func() {
		a := ref.New("k", "v")
		b := ref.New("k", "v")
		So(a.IsEqual(b), ShouldBeTrue)
		So(a.ToPtr().IsEqualPtr(b.ToPtr()), ShouldBeTrue)
		So(a.ToPtr().IsEqualPtr(nil), ShouldBeTrue)
		var nilV *ref.Value
		So(nilV.IsEqualPtr(nil), ShouldBeTrue)
	})
}

func Test_MoreCoverage_Refs_Constructors(t *testing.T) {
	Convey("refs constructor family", t, func() {
		So(refs.New(4), ShouldNotBeNil)
		So(refs.New2(), ShouldNotBeNil)
		So(refs.New4(), ShouldNotBeNil)
		So(refs.NewUsingMap(map[string]interface{}{"a": 1, "b": 2}).Count(), ShouldEqual, 2)
		So(refs.NewUsingMap(nil), ShouldNotBeNil)
		So(refs.NewUsingRefsOrNil(), ShouldBeNil)
		So(refs.NewUsingRefsOrNil(ref.New("k", 1)).Count(), ShouldEqual, 1)
		So(refs.NewUsingValues(ref.New("k", 1)).Count(), ShouldEqual, 1)
		So(refs.NewUsingRefs(true, ref.New("k", 1)).Count(), ShouldEqual, 1)
		So(refs.NewUsingRefs(false), ShouldNotBeNil)
		So(refs.NewUsingReferencers().Count(), ShouldEqual, 0)
		So(refs.NewUsingCollection(false, refs.EmptyPtr()), ShouldNotBeNil)
		So(refs.NewUsingCollection(true, refs.EmptyPtr(), refs.EmptyPtr()), ShouldNotBeNil)
	})
}

func Test_MoreCoverage_Refs_Collection(t *testing.T) {
	Convey("refs.Collection mutation and accessors", t, func() {
		c := refs.New(4)
		c.AddVarVal("a", 1)
		c.Add("b", 2)
		c.Adds(ref.New("c", 3))
		c.AddsIf(true, ref.New("d", 4))
		c.AddsIf(false, ref.New("skip", 0))

		So(c.Count(), ShouldEqual, 4)
		So(c.Length(), ShouldEqual, 4)
		So(c.HasAnyItem(), ShouldBeTrue)
		So(c.IsEmpty(), ShouldBeFalse)
		So(c.IsNull(), ShouldBeFalse)
		So(c.Items(), ShouldHaveLength, 4)
		So(c.List(), ShouldHaveLength, 4)
		So(c.Collection(), ShouldHaveLength, 4)
		So(c.Strings(), ShouldHaveLength, 4)
		So(c.String(), ShouldNotBeEmpty)
		So(c.Compile(), ShouldNotBeEmpty)
		So(c.MapStringString(), ShouldNotBeEmpty)
		So(c.MapStringAny(), ShouldNotBeEmpty)
		So(c.DynamicMap(), ShouldNotBeNil)
		So(c.ReferencesList(), ShouldHaveLength, 4)
		So(c.ReferencerCollection(), ShouldHaveLength, 4)

		cl := c.Clone()
		So(cl.Length(), ShouldEqual, 4)
		clp := c.ClonePtr()
		So(clp.Length(), ShouldEqual, 4)

		bytes, err := c.Serialize()
		So(err, ShouldBeNil)
		So(bytes, ShouldNotBeEmpty)
		So(c.SerializeMust(), ShouldNotBeEmpty)
		So(c.Json().IsEmpty(), ShouldBeFalse)
		So(c.JsonPtr(), ShouldNotBeNil)
		So(c.JsonModel(), ShouldHaveLength, 4)
		So(c.JsonModelAny(), ShouldNotBeNil)

		So(c.AsJsoner(), ShouldNotBeNil)
		So(c.AsJsonMarshaller(), ShouldNotBeNil)
		So(c.AsJsonContractsBinder(), ShouldNotBeNil)
		So(c.AsJsonParseSelfInjector(), ShouldNotBeNil)
		So(c.CloneNewDefiner(), ShouldNotBeNil)

		other := refs.New(2)
		other.Add("x", "y")
		c.AddCollection(other)
		c.AddCollections(other)
		c.AddCollectionCloned(other)
		So(c.ConcatNew(other), ShouldNotBeNil)

		c.AddsPtr(ref.NewPtr("p", 9))
		c.AddsByCloningItems(ref.New("q", 10))
		c.AddsPtrByCloningItems(ref.NewPtr("r", 11))
		c.AddMap(map[string]interface{}{"m": 1})

		c.Dispose()
	})
}

func Test_MoreCoverage_Refs_Equality(t *testing.T) {
	Convey("refs.Collection equality", t, func() {
		a := refs.New(2)
		a.Add("k", "v")
		b := refs.New(2)
		b.Add("k", "v")
		So(a.IsEqual(b), ShouldBeTrue)
	})
}
