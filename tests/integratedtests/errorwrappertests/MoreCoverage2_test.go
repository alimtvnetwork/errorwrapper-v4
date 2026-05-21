package errorwrappertests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
	. "github.com/smartystreets/goconvey/convey"
)

func mkWrap(msg string) *errorwrapper.Wrapper {
	return errnew.Type.Error(errtype.Generic, errors.New(msg))
}

func mkWrapWithRef(msg string) *errorwrapper.Wrapper {
	return errnew.Ref.ErrorOne(errtype.Generic, errors.New(msg), "k", "v")
}

func Test_MoreCoverage2_MergeAndCollect(t *testing.T) {
	Convey("MergeNewMessage on an error wrapper returns a non-empty combiner", t, func() {
		w := mkWrap("base")
		got := w.MergeNewMessage("extra")
		So(got, ShouldNotBeNil)
		So(got.HasError(), ShouldBeTrue)
	})

	Convey("MergeNewErrInf with nil right returns receiver", t, func() {
		w := mkWrap("base")
		got := w.MergeNewErrInf(nil)
		So(got, ShouldNotBeNil)
		So(got.HasError(), ShouldBeTrue)
	})

	Convey("MergeNewErrInf with non-empty right returns merged", t, func() {
		w := mkWrap("base")
		right := mkWrap("extra")
		got := w.MergeNewErrInf(right)
		So(got, ShouldNotBeNil)
		So(got.HasError(), ShouldBeTrue)
	})

	Convey("IsCollect with nil/empty returns false", t, func() {
		w := mkWrap("base")
		So(w.IsCollect(nil), ShouldBeFalse)
		So(w.IsCollect(errorwrapper.EmptyPtr()), ShouldBeFalse)
	})

	Convey("IsCollect with a wrapper appends and returns true", t, func() {
		w := mkWrap("base")
		other := mkWrap("other")
		So(w.IsCollect(other), ShouldBeTrue)
	})

	Convey("IsCollect with a Collection wraps and returns true", t, func() {
		w := mkWrap("base")
		c := errwrappers.Empty()
		c.AddTypeError(errtype.Generic, errors.New("coll"))
		So(w.IsCollect(c), ShouldBeTrue)
	})

	Convey("IsCollectedAny / IsEmptyAll with no args", t, func() {
		w := mkWrap("base")
		So(w.IsCollectedAny(), ShouldBeFalse)
		So(w.IsEmptyAll(), ShouldBeTrue)
	})

	Convey("IsCollectedAny with real wrappers returns true", t, func() {
		w := mkWrap("base")
		So(w.IsCollectedAny(mkWrap("a"), mkWrap("b")), ShouldBeTrue)
	})

	Convey("IsCollectOn respects guard flag", t, func() {
		w := mkWrap("base")
		So(w.IsCollectOn(false, mkWrap("x")), ShouldBeFalse)
		So(w.IsCollectOn(true, nil), ShouldBeFalse)
		So(w.IsCollectOn(true, mkWrap("x")), ShouldBeTrue)
	})
}

func Test_MoreCoverage2_References(t *testing.T) {
	Convey("ReferencesList on plain wrapper is nil", t, func() {
		w := mkWrap("base")
		So(w.ReferencesList(), ShouldBeNil)
	})

	Convey("References / HasReferences with refs-bearing wrapper", t, func() {
		w := errnew.RefOne.Error(errtype.Generic, errors.New("x"), "k", "v")
		So(w.References(), ShouldNotBeNil)
		So(w.HasReferences(), ShouldBeTrue)
		So(w.ReferencesList(), ShouldNotBeNil)
	})

	Convey("CloneReferences returns nil for empty wrapper", t, func() {
		So(errorwrapper.EmptyPtr().CloneReferences(), ShouldBeNil)
	})

	Convey("CloneReferences on refs-bearing wrapper is non-nil", t, func() {
		w := errnew.RefOne.Error(errtype.Generic, errors.New("x"), "k", "v")
		So(w.CloneReferences(), ShouldNotBeNil)
	})

	Convey("MergeNewReferences on empty receiver builds from additions", t, func() {
		got := errorwrapper.EmptyPtr().MergeNewReferences(ref.New("a", 1))
		So(got, ShouldNotBeNil)
		So(got.HasAnyItem(), ShouldBeTrue)
	})

	Convey("MergeNewReferences on existing references appends", t, func() {
		w := errnew.RefOne.Error(errtype.Generic, errors.New("x"), "k", "v")
		got := w.MergeNewReferences(ref.New("a", 1))
		So(got, ShouldNotBeNil)
		So(got.HasAnyItem(), ShouldBeTrue)
	})

	Convey("ReferencesCollection / ReferencesCompiledString are safe", t, func() {
		w := errnew.RefOne.Error(errtype.Generic, errors.New("x"), "k", "v")
		So(w.ReferencesCollection(), ShouldNotBeNil)
		So(w.ReferencesCompiledString(), ShouldNotBeBlank)
	})
}

func Test_MoreCoverage2_StackAndTypeNames(t *testing.T) {
	Convey("StackTraces helpers return strings/json results", t, func() {
		w := mkWrap("st")
		So(w.StackTraces(), ShouldNotBeBlank)
		So(w.NewStackTraces(0), ShouldNotBeBlank)
		So(w.NewDefaultStackTraces(), ShouldNotBeBlank)

		jr := w.StackTracesJsonResult()
		So(jr, ShouldNotBeNil)

		So(w.NewStackTracesJsonResult(0), ShouldNotBeNil)
		So(w.NewDefaultStackTracesJsonResult(), ShouldNotBeNil)
		So(w.StackTracesLimit(2), ShouldNotBeNil)
		So(w.FullStringWithLimitTraces(2), ShouldNotBeBlank)
	})

	Convey("TypeNameWithCustomMessage / TypeCodeNameString / TypeNameCodeMessage", t, func() {
		w := mkWrap("t")
		So(w.TypeNameWithCustomMessage("custom"), ShouldNotBeBlank)
		So(w.TypeCodeNameString(), ShouldNotBeBlank)
		So(w.TypeNameCodeMessage(), ShouldNotBeBlank)
		So(w.TypeString(), ShouldNotBeBlank)
		So(w.RawErrorTypeValue(), ShouldBeGreaterThanOrEqualTo, 0)
		So(w.RawErrorTypeName(), ShouldNotBeBlank)
	})

	Convey("IsDefined / HasAnyIssues mirror HasError", t, func() {
		w := mkWrap("d")
		So(w.IsDefined(), ShouldBeTrue)
		So(w.HasAnyIssues(), ShouldBeTrue)

		empty := errorwrapper.EmptyPtr()
		So(empty.IsDefined(), ShouldBeFalse)
		So(empty.HasAnyIssues(), ShouldBeFalse)
	})
}

func Test_MoreCoverage2_NonPtrPtrAndJson(t *testing.T) {
	Convey("NonPtr / Ptr round-trip", t, func() {
		w := mkWrap("rt")
		val := w.NonPtr()
		p := val.Ptr()
		So(p, ShouldNotBeNil)
		So(p.HasError(), ShouldBeTrue)
	})

	Convey("ParseInjectUsingJson handles nil/empty", t, func() {
		var w errorwrapper.Wrapper
		_, err := w.ParseInjectUsingJson(nil)
		So(err, ShouldNotBeNil)

		_, err = w.ParseInjectUsingJson(&corejson.Result{})
		So(err, ShouldNotBeNil)
	})

	Convey("ParseInjectUsingJsonMust panics on nil payload", t, func() {
		var w errorwrapper.Wrapper
		So(func() { w.ParseInjectUsingJsonMust(nil) }, ShouldPanic)
	})

	Convey("Json round-trip via JsonParseSelfInject", t, func() {
		src := mkWrap("round")
		jp := src.JsonPtr()
		So(jp, ShouldNotBeNil)

		var dst errorwrapper.Wrapper
		err := dst.JsonParseSelfInject(jp)
		So(err, ShouldBeNil)
	})

	Convey("GetTypeVariantStruct / ErrorTypeAsBasicErrorTyper", t, func() {
		w := mkWrap("v")
		So(w.GetTypeVariantStruct().Name, ShouldNotBeBlank)
		So(w.ErrorTypeAsBasicErrorTyper(), ShouldNotBeNil)
	})
}

func Test_MoreCoverage2_GetAsBasicWrapperUsingTyper(t *testing.T) {
	Convey("Empty wrapper returns nil interface", t, func() {
		out := errorwrapper.EmptyPtr().GetAsBasicWrapperUsingTyper(errtype.Generic)
		So(out, ShouldBeNil)
	})

	Convey("Non-empty wrapper returns a basic wrapper", t, func() {
		w := mkWrap("basic")
		out := w.GetAsBasicWrapperUsingTyper(errtype.NotFound)
		So(out, ShouldNotBeNil)
		So(out.HasError(), ShouldBeTrue)
	})
}
