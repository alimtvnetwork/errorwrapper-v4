package refstests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/alimtvnetwork/core-v9/coredata/corejson"

	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
	"github.com/alimtvnetwork/errorwrapper-v3/refs"
)

func Test_MoreCoverage2_Refs_RemainingConstructors(t *testing.T) {
	Convey("constructor helpers not yet covered", t, func() {
		So(refs.NewClone(), ShouldBeNil)
		So(refs.NewClone(ref.New("k", 1)).Count(), ShouldEqual, 1)

		So(refs.NewUsingMany().Count(), ShouldEqual, 0)
		many := refs.NewUsingMany(
			[]ref.Value{ref.New("a", 1)},
			nil,
			[]ref.Value{ref.New("b", 2), ref.New("c", 3)},
		)
		So(many.Count(), ShouldEqual, 3)

		ex := refs.New(2)
		ex.Add("x", 1)
		plus := refs.NewExistingCollectionPlusAddition(ex, ref.New("y", 2), ref.New("z", 3))
		So(plus.Count(), ShouldEqual, 3)
		So(refs.NewExistingCollectionPlusAddition(nil, ref.New("y", 2)).Count(), ShouldEqual, 1)

		plus2 := refs.NewExistingPlusAddition(
			[]ref.Value{ref.New("a", 1)},
			ref.New("b", 2),
		)
		So(plus2.Count(), ShouldEqual, 2)

		merged := refs.MergeReferences(
			[]ref.Value{ref.New("a", 1)},
			ref.New("b", 2),
		)
		So(merged, ShouldHaveLength, 2)
		So(refs.MergeReferences(nil), ShouldHaveLength, 0)

		pre := refs.PrependReferences(false, nil, ref.New("p", 1))
		So(pre, ShouldHaveLength, 1)
		pre2 := refs.PrependReferences(false, []ref.Value{ref.New("a", 1)})
		So(pre2, ShouldHaveLength, 1)
		pre3 := refs.PrependReferences(true, []ref.Value{ref.New("a", 1)}, ref.New("p", 1))
		So(pre3, ShouldHaveLength, 2)

		So(refs.NewWithItem(2, "k", 1).Count(), ShouldEqual, 1)
		So(refs.NewDirectItem("k", 1).Count(), ShouldEqual, 1)

		emp := refs.Empty()
		So(emp.IsEmpty(), ShouldBeTrue)

		dm := refs.NewDataModel(nil)
		So(dm, ShouldNotBeNil)
		c := refs.New(1)
		c.Add("a", 1)
		dm2 := refs.NewDataModel(c)
		So(dm2.Refs, ShouldHaveLength, 1)
		fromDm := refs.NewFromDataModelPtr(dm2)
		So(fromDm.Count(), ShouldEqual, 1)
		emptyDm := refs.NewFromDataModelPtr(&refs.CollectionDataModel{})
		So(emptyDm.IsEmpty(), ShouldBeTrue)

		So(refs.NewUsingReferencers(ref.NewPtr("k", 1).AsReferencer()).Count(), ShouldEqual, 1)
		So(refs.NewUsingBasicErrWrap(nil).IsEmpty(), ShouldBeTrue)
	})
}

func Test_MoreCoverage2_Refs_Helpers(t *testing.T) {
	Convey("standalone helpers", t, func() {
		So(refs.AllIndividualItemsCount(nil), ShouldEqual, 0)
		a := refs.New(1)
		a.Add("k", 1)
		b := refs.New(1)
		b.Add("k2", 2)
		So(refs.AllIndividualItemsCount(a, b, nil), ShouldEqual, 2)

		So(refs.LengthOfEachItems(nil), ShouldEqual, 0)
		So(refs.LengthOfEachItems([][]ref.Value{
			{ref.New("a", 1), ref.New("b", 2)},
			{ref.New("c", 3)},
		}), ShouldEqual, 3)

		v1 := ref.New("a", 1)
		v2 := ref.New("b", 2)
		slice1 := []*ref.Value{v1.ToPtr(), v2.ToPtr()}
		slice2 := []*ref.Value{v1.ToPtr()}
		manyPtr := []*[]*ref.Value{&slice1, nil, &slice2}
		So(refs.LengthOfEachItemsPtr(&manyPtr), ShouldEqual, 3)

		So(refs.CompileAnyItemsToCsvStringDefault(), ShouldEqual, "")
		So(refs.CompileAnyItemsToCsvStringDefault("a", 1, "b"), ShouldNotBeEmpty)
	})
}

func Test_MoreCoverage2_Refs_QuickReference(t *testing.T) {
	Convey("QuickReference family", t, func() {
		qr := refs.NewQuickReference(errtype.Generic, "x", 1)
		So(qr.CompileLine(), ShouldNotBeEmpty)

		qr2 := refs.NewQuickReferenceStrings(errtype.Generic, "x", "y")
		So(qr2.CompileLine(), ShouldNotBeEmpty)

		So(refs.QuickCompileStrings(), ShouldHaveLength, 0)
		So(refs.QuickCompileStrings(qr, qr2), ShouldHaveLength, 2)

		So(refs.QuickCompileString("|"), ShouldEqual, "")
		So(refs.QuickCompileString("|", qr, qr2), ShouldNotBeEmpty)

		So(refs.QuickCompileStringDefaultEachLine(), ShouldEqual, "")
		So(refs.QuickCompileStringDefaultEachLine(qr), ShouldNotBeEmpty)

		So(refs.QuickCompileStringDefaultInLine(), ShouldEqual, "")
		So(refs.QuickCompileStringDefaultInLine(qr), ShouldNotBeEmpty)
	})
}

func Test_MoreCoverage2_Refs_JsonRoundTrip(t *testing.T) {
	Convey("Collection json round-trip and parse-inject", t, func() {
		c := refs.New(2)
		c.Add("k", "v")
		c.Add("n", 1)

		raw, mErr := c.MarshalJSON()
		So(mErr, ShouldBeNil)
		So(raw, ShouldNotBeEmpty)

		jr := &corejson.Result{Bytes: raw}

		var into refs.Collection
		_, _ = into.ParseInjectUsingJson(jr)

		// nil / empty paths
		var into2 refs.Collection
		_, err2 := into2.ParseInjectUsingJson(nil)
		So(err2, ShouldNotBeNil)
		_, err3 := into2.ParseInjectUsingJson(&corejson.Result{})
		So(err3, ShouldNotBeNil)

		var into3 refs.Collection
		mustPanic := func() { into3.ParseInjectUsingJsonMust(nil) }
		So(mustPanic, ShouldPanic)

		var into4 refs.Collection
		So(into4.ParseInjectUsingJsonMust(jr).Count(), ShouldEqual, 2)

		var into5 refs.Collection
		So(into5.JsonParseSelfInject(jr), ShouldBeNil)

		var into6 refs.Collection
		So(into6.UnmarshalJSON(raw), ShouldBeNil)
		So(into6.Count(), ShouldEqual, 2)
		So(into6.UnmarshalJSON([]byte("not-json")), ShouldNotBeNil)
	})
}

func Test_MoreCoverage2_Refs_ReflectAndEquality(t *testing.T) {
	Convey("ReflectSetTo and Equality edges", t, func() {
		src := refs.New(1)
		src.Add("k", "v")
		var dst refs.Collection
		err := src.ReflectSetTo(&dst)
		So(err, ShouldBeNil)

		var nilC *refs.Collection
		So(nilC.IsEqual(nil), ShouldBeTrue)
		So(nilC.IsEqual(refs.New(0)), ShouldBeFalse)

		empty1 := refs.New(0)
		empty2 := refs.New(0)
		So(empty1.IsEqual(empty2), ShouldBeTrue)

		left := refs.New(1)
		left.Add("k", 1)
		right := refs.New(2)
		right.Add("k", 1)
		right.Add("k2", 2)
		So(left.IsEqual(right), ShouldBeFalse)

		diff := refs.New(1)
		diff.Add("z", 9)
		So(left.IsEqual(diff), ShouldBeFalse)
	})
}
