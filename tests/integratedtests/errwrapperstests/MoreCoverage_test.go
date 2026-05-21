package errwrapperstests

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)


func newPopulated() *errwrappers.Collection {
	c := errwrappers.Empty()
	c.AddError(errors.New("alpha"))
	c.AddWrapperPtr(errnew.Messages.Single(errtype.InvalidInput, "beta"))
	c.AddError(errors.New("gamma"))
	return c
}

func Test_Constructors(t *testing.T) {
	Convey("Capacity constructors all return non-nil collections", t, func() {
		So(errwrappers.New(2), ShouldNotBeNil)
		So(errwrappers.NewCap1(), ShouldNotBeNil)
		So(errwrappers.NewCap2(), ShouldNotBeNil)
		So(errwrappers.NewCap3(), ShouldNotBeNil)
		So(errwrappers.NewCap4(), ShouldNotBeNil)
		So(errwrappers.NewEmpty(), ShouldNotBeNil)
		So(errwrappers.EmptyCollection(), ShouldNotBeNil)
	})

	Convey("NewUsingErrors / NewUsingErrorsPtr seed errors", t, func() {
		errs := []error{errors.New("a"), errors.New("b")}
		c := errwrappers.NewUsingErrors(errs...)
		So(c.Count(), ShouldEqual, 2)
		c2 := errwrappers.NewUsingErrorsPtr(&errs)
		So(c2.Count(), ShouldEqual, 2)
	})

	Convey("NewWithType / NewWithMessage / NewWithError variants", t, func() {
		So(errwrappers.NewWithType(errtype.NotFound).Count(), ShouldEqual, 1)
		So(errwrappers.NewWithTypeUsingStackSkip(0, errtype.NotFound).Count(), ShouldEqual, 1)
		So(errwrappers.NewWithMessage(errtype.NotFound, "x").Count(), ShouldEqual, 1)
		So(errwrappers.NewWithMessageUsingStackSkip(0, errtype.NotFound, "y").Count(), ShouldEqual, 1)
		So(errwrappers.NewWithError(2, errtype.Generic, errors.New("err")).Count(), ShouldEqual, 1)
		So(errwrappers.NewWithErrorUsingStackSkip(0, errtype.Generic, errors.New("err")).Count(), ShouldEqual, 1)
		So(errwrappers.NewWithOnlyError(errors.New("e")).Count(), ShouldEqual, 1)
		So(errwrappers.NewWithOnlyCapError(2, errors.New("e")).Count(), ShouldEqual, 1)
	})

	Convey("NewWithItem with capacity + type", t, func() {
		c := errwrappers.NewWithItem(2, errtype.NotFound)
		So(c, ShouldNotBeNil)
	})


	Convey("NewUsingErrorWrappers + clone + ptr", t, func() {
		w1 := errnew.Messages.Single(errtype.NotFound, "w1")
		w2 := errnew.Messages.Single(errtype.NotFound, "w2")
		So(errwrappers.NewUsingErrorWrappers(w1, w2).Count(), ShouldEqual, 2)
		So(errwrappers.NewUsingErrorWrappersClone([]*errorwrapper.Wrapper{w1, w2}).Count(), ShouldEqual, 2)
		wps := []*errwrappers.Collection{newPopulated(), newPopulated()}
		So(errwrappers.NewUsingCollections(wps...).Count(), ShouldBeGreaterThan, 0)
		So(errwrappers.NewUsingCollectionsPtr(&wps).Count(), ShouldBeGreaterThan, 0)
	})
}

func Test_Collection_Reads(t *testing.T) {
	c := newPopulated()

	Convey("Reader methods on a populated collection", t, func() {
		So(c.HasAnyError(), ShouldBeTrue)
		So(c.HasAnyIssues(), ShouldBeTrue)
		So(c.HasAnyItem(), ShouldBeTrue)
		So(c.IsInvalid(), ShouldBeTrue)
		So(c.IsDefined(), ShouldBeTrue)
		So(c.IsNull(), ShouldBeFalse)
		So(c.IsAnyNull(), ShouldBeFalse)
		So(c.Count(), ShouldEqual, 3)
		So(c.LastIndex(), ShouldEqual, 2)
		So(c.HasIndex(0), ShouldBeTrue)
		So(c.HasIndex(99), ShouldBeFalse)
		So(c.Compile(), ShouldNotBeBlank)
		So(c.FullString(), ShouldContainSubstring, "alpha")
		So(c.FullStringWithoutReferences(), ShouldNotBeBlank)
		So(c.FullStringSplitByNewLine(), ShouldNotBeEmpty)
		So(c.ErrorString(), ShouldNotBeBlank)
		So(c.CompiledError(), ShouldNotBeNil)
		So(c.FullStringWithTraces(), ShouldNotBeBlank)
		So(c.FullStringWithTracesIf(true), ShouldNotBeBlank)
		So(c.AllReferences(), ShouldNotBeNil)
		So(c.ReferencesCompiledString(), ShouldNotBeNil)
		So(c.CompiledErrorWithStackTraces(), ShouldNotBeNil)
		So(c.CompiledStackTracesString(), ShouldNotBeBlank)
		So(c.CompiledJsonErrorWithStackTraces(), ShouldNotBeNil)
		So(c.CompiledJsonStringWithStackTraces(), ShouldNotBeBlank)
		So(c.IsCollectionType(), ShouldBeTrue)
		So(c.Value(), ShouldNotBeNil)
	})

	Convey("First/Last accessors", t, func() {
		So(c.First(), ShouldNotBeNil)
		So(c.Last(), ShouldNotBeNil)
		So(c.FirstOrDefault(), ShouldNotBeNil)
		So(c.LastOrDefault(), ShouldNotBeNil)
		So(c.FirstOrDefaultError(), ShouldNotBeNil)
		So(c.LastOrDefaultError(), ShouldNotBeNil)
		So(c.FirstOrDefaultCompiledError(), ShouldNotBeNil)
		So(c.LastOrDefaultCompiledError(), ShouldNotBeNil)
		So(c.FirstOrDefaultFullMessage(), ShouldNotBeBlank)
		So(c.LastOrDefaultFullMessage(), ShouldNotBeBlank)
		So(c.FirstDynamic(), ShouldNotBeNil)
		So(c.LastDynamic(), ShouldNotBeNil)
		So(c.FirstOrDefaultDynamic(), ShouldNotBeNil)
		So(c.LastOrDefaultDynamic(), ShouldNotBeNil)
	})

	Convey("Slice helpers", t, func() {
		So(c.Skip(1).Count(), ShouldEqual, 2)
		So(c.Take(2).Count(), ShouldEqual, 2)
		So(c.TakeFromTo(0, 1).Count(), ShouldBeGreaterThan, 0)
		So(c.SkipDynamic(1), ShouldNotBeNil)
		So(c.TakeDynamic(2), ShouldNotBeNil)
		So(c.LimitDynamic(2), ShouldNotBeNil)
	})

	Convey("Stack trace + basic-wrapper accessors", t, func() {
		So(c.StackTraces(), ShouldNotBeNil)
		So(c.NewDefaultStackTraces(), ShouldNotBeNil)
		So(c.NewStackTraces(0), ShouldNotBeNil)
		So(c.NewDefaultStackTracesJsonResult(), ShouldNotBeNil)
		So(c.NewStackTracesJsonResult(0), ShouldNotBeNil)
		So(c.StackTracesJsonResult(), ShouldNotBeNil)
		So(c.AllStackTraces(), ShouldNotBeNil)
		So(c.GetAsBasicWrapper(), ShouldNotBeNil)
		So(c.CompiledToGenericBasicErrWrapper(), ShouldNotBeNil)
	})

	Convey("Serialize roundtrip", t, func() {
		b, err := c.Serialize()
		So(err, ShouldBeNil)
		So(b, ShouldNotBeEmpty)
		So(c.SerializeMust(), ShouldNotBeEmpty)
		b2, err := c.SerializeWithoutTraces()
		So(err, ShouldBeNil)
		So(b2, ShouldNotBeEmpty)
	})
}

func Test_Collection_Mutations(t *testing.T) {
	Convey("AddErrors / AddErrorsPtr / AddErrorChain", t, func() {
		c := errwrappers.Empty()
		c.AddErrors(errors.New("a"), errors.New("b"))
		So(c.Count(), ShouldEqual, 2)

		errs := []error{errors.New("c")}
		c.AddErrorsPtr(&errs)
		So(c.Count(), ShouldEqual, 3)

		c2 := errwrappers.Empty().AddErrorChain(errors.New("chained"))
		So(c2.Count(), ShouldEqual, 1)
	})

	Convey("AddErrorWithMessages / AddFmt family", t, func() {
		c := errwrappers.Empty()
		c.AddErrorWithMessages(errtype.NotFound, errors.New("e"), "msg1", "msg2")
		So(c.Count(), ShouldBeGreaterThan, 0)
		c.AddFmt(errtype.NotFound, "fmt-%s", "x")
		c.AddFmtErr(errtype.NotFound, errors.New("inner"), "fmt-%s", "y")
		c.AddFmtMsg(errtype.NotFound, "label", "fmt-%s", "z")
		c.AddFmtIf(true, errtype.NotFound, "fmt-%s", "if")
		c.AddFmtIf(false, errtype.NotFound, "skipped")
		So(c.HasAnyError(), ShouldBeTrue)
	})

	Convey("Conditional adders", t, func() {
		c := errwrappers.Empty()
		c.ConditionalAddError(false, errors.New("nope"))
		So(c.Count(), ShouldEqual, 0)
		c.ConditionalAddError(true, errors.New("yes"))
		So(c.Count(), ShouldEqual, 1)
		c.AddIf(false, errnew.Messages.Single(errtype.NotFound, "skipped"))
		So(c.Count(), ShouldEqual, 1)
		c.AddIf(true, errnew.Messages.Single(errtype.NotFound, "kept"))
		So(c.Count(), ShouldEqual, 2)
	})

	Convey("AddSlice + AddRawErrCollection", t, func() {
		c := errwrappers.Empty()
		c.AddSlice(errtype.NotFound, "label", "m1", "m2")
		So(c.HasAnyError(), ShouldBeTrue)
	})
}

func Test_StateCounter_All(t *testing.T) {
	Convey("StateCounter helpers cover compare/check methods", t, func() {
		c := errwrappers.Empty()
		sc := errwrappers.NewStateCount(c)
		So(sc.IsSameState(), ShouldBeTrue)
		So(sc.HasChanges(), ShouldBeFalse)
		So(sc.HasChangesCollection(), ShouldBeFalse)
		So(sc.IsSameStateCollection(), ShouldBeTrue)
		So(sc.IsSuccess(), ShouldBeTrue)
		So(sc.IsFailed(), ShouldBeFalse)
		So(sc.IsValid(), ShouldBeTrue)
		So(sc.IsSameStateUsingCount(0), ShouldBeTrue)
		So(sc.StartStateTracking(0), ShouldEqual, 0)
		So(sc.AsCountStateTrackerBinder(), ShouldNotBeNil)
	})
}

func Test_MutexCollection_All(t *testing.T) {
	Convey("MutexCollection constructors + methods", t, func() {
		mc := errwrappers.MutexEmpty()
		_ = mc.IsEmpty()
		_ = mc.HasError()
		_ = mc.Length()
		_ = mc.IsSuccess()
		_ = mc.IsValid()
		_ = mc.IsFailed()
		So(mc.Collection(), ShouldNotBeNil)

		mc.AddWrapperPtr(errnew.Messages.Single(errtype.NotFound, "m1"))
		mc.AddWrappers(errnew.Messages.Single(errtype.NotFound, "m2"))
		_ = mc.HasError()
		_ = mc.Length()
		_ = mc.IsEmpty()
		So(mc.String(), ShouldNotBeBlank)
		So(mc.ToStringLock(true, true), ShouldNotBeBlank)
		So(mc.ToStringsLock(true, true), ShouldNotBeEmpty)
		So(mc.DisplayStringWithTraces(), ShouldNotBeBlank)
		So(mc.DisplayStringWithLimitTracesLock(2), ShouldNotBeBlank)
		So(mc.StringsIfLock(true), ShouldNotBeEmpty)
		So(mc.FullStrings(), ShouldNotBeEmpty)
		So(mc.FullStringsWithTraces(), ShouldNotBeEmpty)
		So(mc.FullStringsWithLimitTracesLock(1), ShouldNotBeEmpty)
		So(mc.Errors(), ShouldNotBeEmpty)
		So(mc.CompiledErrors(), ShouldNotBeEmpty)
		So(mc.GetAsError(), ShouldNotBeNil)
		So(mc.GetAsErrorWrapperPtr(), ShouldNotBeNil)
		So(mc.StateCounter(), ShouldNotBeNil)

		mc.Clear()
		So(mc.Length(), ShouldEqual, 0)
		mc.Dispose()
	})

	Convey("MutexNew + state counter", t, func() {
		mc := errwrappers.MutexNew(2)
		sc := errwrappers.NewMutexStateCount(mc)
		So(sc.IsSuccess(), ShouldBeTrue)
		So(sc.IsValid(), ShouldBeTrue)
		So(sc.IsFailed(), ShouldBeFalse)
		So(sc.HasChanges(0), ShouldBeFalse)
		So(sc.IsSameState(0), ShouldBeTrue)
		So(sc.HasChangesCollection(), ShouldBeFalse)
		So(sc.IsSameStateCollection(), ShouldBeTrue)
		So(sc.StartStateTracking(0), ShouldEqual, 0)
	})
}

func Test_Deserialize(t *testing.T) {
	c := newPopulated()
	bytes, _ := c.Serialize()

	Convey("Deserialize.UsingBytes returns a non-nil collection", t, func() {
		result, err := errwrappers.Deserialize.UsingBytes(bytes)
		So(err, ShouldBeNil)
		So(result, ShouldNotBeNil)
	})

	Convey("Deserialize.UsingString roundtrips", t, func() {
		result, err := errwrappers.Deserialize.UsingString(string(bytes))
		So(err, ShouldBeNil)
		So(result, ShouldNotBeNil)
	})
}
