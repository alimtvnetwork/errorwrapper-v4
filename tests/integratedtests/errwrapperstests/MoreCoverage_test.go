package errwrapperstests

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
)

// Test_Constructors — exercises the family of New* constructors.
func Test_Constructors(t *testing.T) {
	Convey("New variants produce non-nil empty collections", t, func() {
		So(errwrappers.New(0), ShouldNotBeNil)
		So(errwrappers.NewCap1(), ShouldNotBeNil)
		So(errwrappers.NewCap2(), ShouldNotBeNil)
		So(errwrappers.NewCap3(), ShouldNotBeNil)
		So(errwrappers.NewCap4(), ShouldNotBeNil)
		So(errwrappers.Empty().IsEmpty(), ShouldBeTrue)
		So(errwrappers.NewEmpty().IsEmpty(), ShouldBeTrue)
		So(errwrappers.EmptyCollection().IsEmpty(), ShouldBeTrue)
	})

	Convey("NewUsingErrors and NewUsingErrorsPtr add multiple errors", t, func() {
		c := errwrappers.NewUsingErrors(errors.New("a"), errors.New("b"))
		So(c.Count(), ShouldEqual, 2)

		errs := []error{errors.New("c")}
		c2 := errwrappers.NewUsingErrorsPtr(&errs)
		So(c2.Count(), ShouldEqual, 1)

		So(errwrappers.NewUsingErrors().IsEmpty(), ShouldBeTrue)
		So(errwrappers.NewUsingErrorsPtr(nil).IsEmpty(), ShouldBeTrue)
	})

	Convey("NewUsingErrorWrappers variants", t, func() {
		w := errnew.Messages.Single(errtype.InvalidInput, "x")
		c := errwrappers.NewUsingErrorWrappers(w)
		So(c.Count(), ShouldEqual, 1)

		c2 := errwrappers.NewUsingErrorWrappersClone([]*errorwrapper.Wrapper{w})
		So(c2.Count(), ShouldEqual, 1)

		So(errwrappers.NewUsingErrorWrappers().IsEmpty(), ShouldBeTrue)
	})

	Convey("NewWithType / NewWithMessage / NewWithError", t, func() {
		So(errwrappers.NewWithType(errtype.InvalidInput).Count(), ShouldEqual, 1)
		So(errwrappers.NewWithMessage(errtype.InvalidInput, "m").Count(), ShouldEqual, 1)
		So(errwrappers.NewWithError(1, errtype.InvalidInput, errors.New("e")).Count(), ShouldEqual, 1)
		So(errwrappers.NewWithError(1, errtype.InvalidInput, nil), ShouldBeNil)
		So(errwrappers.NewWithOnlyError(errors.New("z")).Count(), ShouldEqual, 1)
		So(errwrappers.NewWithOnlyCapError(2, errors.New("z")).Count(), ShouldEqual, 1)
		So(errwrappers.NewWithItem(2, errtype.InvalidInput).Count(), ShouldEqual, 1)
	})

	Convey("NewUsingCollections merges multiple collections", t, func() {
		a := errwrappers.NewWithMessage(errtype.InvalidInput, "1")
		b := errwrappers.NewWithMessage(errtype.InvalidInput, "2")
		merged := errwrappers.NewUsingCollections(a, b)
		So(merged.Count(), ShouldEqual, 2)
		So(errwrappers.NewUsingCollections().IsEmpty(), ShouldBeTrue)
	})

	Convey("Stack-skip variants accept positive index", t, func() {
		So(errwrappers.NewWithTypeUsingStackSkip(0, errtype.InvalidInput).Count(), ShouldEqual, 1)
		So(errwrappers.NewWithMessageUsingStackSkip(0, errtype.InvalidInput, "m").Count(), ShouldEqual, 1)
		So(errwrappers.NewWithErrorUsingStackSkip(0, errtype.InvalidInput, errors.New("e")).Count(), ShouldEqual, 1)
		So(errwrappers.NewWithErrorUsingStackSkip(0, errtype.InvalidInput, nil), ShouldBeNil)
	})
}

// Test_Mutex_Constructors — Mutex wrappers around Collection.
func Test_Mutex_Constructors(t *testing.T) {
	Convey("MutexEmpty/MutexNew are empty and usable", t, func() {
		m := errwrappers.MutexEmpty()
		So(m, ShouldNotBeNil)
		So(m.IsEmpty(), ShouldBeTrue)
		So(m.Length(), ShouldEqual, 0)
		So(m.IsSuccess(), ShouldBeTrue)
		So(m.IsValid(), ShouldBeTrue)
		So(m.IsFailed(), ShouldBeFalse)

		w := errnew.Messages.Single(errtype.InvalidInput, "boom")
		m.AddWrapperPtr(w)
		So(m.Length(), ShouldEqual, 1)
		So(m.IsFailed(), ShouldBeTrue)
		So(m.GetAsError(), ShouldNotBeNil)
		So(m.GetAsErrorWrapperPtr(), ShouldNotBeNil)

		So(m.String(), ShouldContainSubstring, "boom")
		So(m.ToStringLock(false, true), ShouldContainSubstring, "boom")
		So(len(m.ToStringsLock(false, true)), ShouldBeGreaterThan, 0)
		So(len(m.FullStrings()), ShouldBeGreaterThan, 0)
		So(len(m.FullStringsWithTraces()), ShouldBeGreaterThan, 0)
		So(len(m.FullStringsWithLimitTracesLock(2)), ShouldBeGreaterThan, 0)
		So(len(m.StringsIfLock(false)), ShouldBeGreaterThan, 0)
		So(len(m.Errors()), ShouldBeGreaterThan, 0)
		So(len(m.CompiledErrors()), ShouldBeGreaterThan, 0)
		So(m.DisplayStringWithTraces(), ShouldNotBeBlank)
		So(m.DisplayStringWithLimitTracesLock(2), ShouldNotBeBlank)

		m2 := errwrappers.MutexNew(3)
		So(m2, ShouldNotBeNil)
		So(m2.Collection(), ShouldNotBeNil)
		m2.AddWrappers(w)
		So(m2.Length(), ShouldEqual, 1)

		m.Clear()
		So(m.IsEmpty(), ShouldBeTrue)
		m.Dispose()
	})
}

// Test_StateCounter — counter tracks length changes.
func Test_StateCounter(t *testing.T) {
	Convey("StateCounter detects changes", t, func() {
		c := errwrappers.Empty()
		counter := errwrappers.NewStateCount(c)
		So(counter.IsSameState(), ShouldBeTrue)
		So(counter.HasChanges(), ShouldBeFalse)
		So(counter.IsSuccess(), ShouldBeTrue)
		So(counter.IsFailed(), ShouldBeFalse)
		So(counter.IsValid(), ShouldBeTrue)

		c.AddError(errors.New("x"))
		So(counter.HasChanges(), ShouldBeTrue)
		So(counter.HasChangesCollection(), ShouldBeTrue)
		So(counter.IsSameStateCollection(), ShouldBeFalse)
		So(counter.IsSameStateUsingCount(0), ShouldBeTrue)
		So(counter.AsCountStateTrackerBinder(), ShouldNotBeNil)

		lg := errwrappers.NewStateCountUsingLengthGetter(c)
		So(lg.IsSameState(), ShouldBeTrue)
	})

	Convey("MutexCollectionStateCounter detects changes", t, func() {
		m := errwrappers.MutexEmpty()
		counter := errwrappers.NewMutexStateCount(m)
		So(counter.IsSameState(0), ShouldBeTrue)
		So(counter.HasChanges(0), ShouldBeFalse)
		So(counter.IsSuccess(), ShouldBeTrue)
		So(counter.IsValid(), ShouldBeTrue)
		So(counter.IsFailed(), ShouldBeFalse)

		m.AddWrapperPtr(errnew.Messages.Single(errtype.InvalidInput, "x"))
		So(counter.HasChangesCollection(), ShouldBeTrue)
		So(counter.IsSameStateCollection(), ShouldBeFalse)
		counter.StartStateTracking(42)
		So(counter.Start, ShouldEqual, 42)
	})
}

// Test_Collection_Adds — various Add* methods.
func Test_Collection_Adds(t *testing.T) {
	Convey("Add helpers extend the collection", t, func() {
		c := errwrappers.Empty()

		So(c.Add(errtype.InvalidInput).Count(), ShouldEqual, 1)
		So(c.AddUsingMsg(errtype.InvalidInput, "m").Count(), ShouldEqual, 2)
		So(c.AddUsingMsg(errtype.InvalidInput, "").Count(), ShouldEqual, 2) // empty msg ignored
		c.AddPathIssueMessages(errtype.InvalidInput, "/path", "m1", "m2")
		So(c.Count(), ShouldEqual, 3)
		c.AddPathIssue(errtype.InvalidInput, errors.New("boom"), "/p")
		So(c.Count(), ShouldEqual, 4)
		c.AddPathIssue(errtype.InvalidInput, nil, "/p")
		So(c.Count(), ShouldEqual, 4)
		c.AddOnlyRefs(errtype.InvalidInput, ref.New("k", "v"))
		So(c.Count(), ShouldEqual, 5)
		c.AddOnlyRefs(errtype.InvalidInput)
		So(c.Count(), ShouldEqual, 5)
		c.AddRef1Msg(errtype.InvalidInput, "m", "k", 1)
		So(c.Count(), ShouldEqual, 6)
		c.AddRef2Msg(errtype.InvalidInput, "m", "k1", 1, "k2", 2)
		So(c.Count(), ShouldEqual, 7)
		c.AddExpectation(errtype.InvalidInput, "title", "expected", "actual")
		So(c.Count(), ShouldEqual, 8)
		c.WasExpecting(errtype.InvalidInput, "title", "e", "a")
		So(c.Count(), ShouldEqual, 9)

		c.AddErrors(errors.New("a"), errors.New("b"))
		So(c.Count(), ShouldEqual, 11)

		errsPtr := []error{errors.New("c")}
		c.AddErrorsPtr(&errsPtr)
		So(c.Count(), ShouldEqual, 12)

		c.AddFmtMsg(errtype.InvalidInput, "msg %s", "v")
		So(c.Count(), ShouldEqual, 13)
		c.AddFmtErr(errtype.InvalidInput, errors.New("e"), "ctx %d", 1)
		So(c.Count(), ShouldEqual, 14)
	})
}

// Test_Collection_Readers — accessor / formatter methods.
func Test_Collection_Readers(t *testing.T) {
	Convey("Reader methods work on populated collection", t, func() {
		c := errwrappers.Empty()
		c.AddUsingMsg(errtype.InvalidInput, "one")
		c.AddUsingMsg(errtype.InvalidInput, "two")
		c.AddUsingMsg(errtype.InvalidInput, "three")

		So(c.HasAnyError(), ShouldBeTrue)
		So(c.HasAnyIssues(), ShouldBeTrue)
		So(c.HasAnyItem(), ShouldBeTrue)
		So(c.HasError(), ShouldBeTrue)
		So(c.IsEmpty(), ShouldBeFalse)
		So(c.IsDefined(), ShouldBeTrue)
		So(c.IsInvalid(), ShouldBeTrue)
		So(c.IsSuccess(), ShouldBeFalse)
		So(c.IsValid(), ShouldBeFalse)
		So(c.IsFailed(), ShouldBeTrue)
		So(c.IsNull(), ShouldBeFalse)
		So(c.IsAnyNull(), ShouldBeFalse)
		So(c.Length(), ShouldEqual, 3)
		So(c.LastIndex(), ShouldEqual, 2)
		So(c.HasIndex(0), ShouldBeTrue)
		So(c.HasIndex(99), ShouldBeFalse)

		So(c.First(), ShouldNotBeNil)
		So(c.Last(), ShouldNotBeNil)
		So(c.FirstOrDefault(), ShouldNotBeNil)
		So(c.LastOrDefault(), ShouldNotBeNil)
		So(c.FirstDynamic(), ShouldNotBeNil)
		So(c.LastDynamic(), ShouldNotBeNil)
		So(c.FirstOrDefaultDynamic(), ShouldNotBeNil)
		So(c.LastOrDefaultDynamic(), ShouldNotBeNil)
		So(c.FirstOrDefaultError(), ShouldNotBeNil)
		So(c.LastOrDefaultError(), ShouldNotBeNil)
		So(c.FirstOrDefaultCompiledError(), ShouldNotBeNil)
		So(c.LastOrDefaultCompiledError(), ShouldNotBeNil)
		So(c.FirstOrDefaultFullMessage(), ShouldNotBeBlank)
		So(c.LastOrDefaultFullMessage(), ShouldNotBeBlank)

		So(c.Count(), ShouldEqual, 3)
		So(c.Skip(1).Count(), ShouldEqual, 2)
		So(c.Take(2).Count(), ShouldEqual, 2)
		So(c.TakeFromTo(0, 2).Count(), ShouldEqual, 2)
		So(c.SkipDynamic(1), ShouldNotBeNil)
		So(c.TakeDynamic(2), ShouldNotBeNil)
		So(c.LimitDynamic(1), ShouldNotBeNil)

		So(c.String(), ShouldContainSubstring, "one")
		So(c.StringIf(false), ShouldContainSubstring, "one")
		So(c.FullString(), ShouldContainSubstring, "one")
		So(c.FullStringWithTraces(), ShouldContainSubstring, "one")
		So(c.FullStringWithTracesIf(true), ShouldContainSubstring, "one")
		So(c.FullStringWithoutReferences(), ShouldContainSubstring, "one")
		So(c.StringWithoutHeader(), ShouldContainSubstring, "one")
		So(c.DisplayStringWithTraces(), ShouldNotBeBlank)
		So(c.DisplayStringWithLimitTraces(1), ShouldNotBeBlank)
		So(c.ToString(false, true), ShouldContainSubstring, "one")
		So(c.Compile(), ShouldNotBeBlank)
		So(c.ErrorString(), ShouldNotBeBlank)
		So(c.CompiledError(), ShouldNotBeNil)
		So(c.CompiledErrorWithStackTraces(), ShouldNotBeNil)
		So(c.CompiledStackTracesString(), ShouldNotBeBlank)
		So(c.Value(), ShouldNotBeNil)

		So(len(c.FullStrings()), ShouldEqual, 3)
		So(len(c.FullStringsWithTraces()), ShouldEqual, 3)
		So(len(c.FullStringsWithLimitTraces(1)), ShouldEqual, 3)
		So(len(c.Strings(false)), ShouldEqual, 3)
		So(len(c.StringsIf(false)), ShouldEqual, 3)
		So(len(c.StringsWithoutHeader()), ShouldEqual, 3)
		So(len(c.StringsWithoutReferencePlusHeader()), ShouldEqual, 3)
		So(len(c.ToStrings(false, false)), ShouldEqual, 3)
		So(len(c.FullStringSplitByNewLine()), ShouldBeGreaterThan, 0)
		So(len(c.LinesWithoutTraces()), ShouldBeGreaterThan, 0)
		So(len(c.Errors()), ShouldEqual, 3)
		So(len(c.CompiledErrors()), ShouldEqual, 3)
		So(len(c.CompiledErrorsWithStackTraces()), ShouldEqual, 3)
		So(len(c.GetTypeVariantStructs()), ShouldEqual, 3)
		So(len(c.ErrorTypes()), ShouldEqual, 3)
		So(len(c.List()), ShouldEqual, 3)
		So(len(c.Items()), ShouldEqual, 3)
		So(c.ItemsNonPtr(), ShouldNotBeNil)
		So(c.AllReferences(), ShouldNotBeNil)

		So(c.GetAsError(), ShouldNotBeNil)
		So(c.GetAsErrorWrapperPtr(), ShouldNotBeNil)
		So(c.GetAsErrorWrapperUsingType(errtype.InvalidInput), ShouldNotBeNil)

		So(c.AsJsoner(), ShouldNotBeNil)
		So(c.AsJsonMarshaller(), ShouldNotBeNil)
		So(c.AsJsonContractsBinder(), ShouldNotBeNil)
		So(c.AsJsonParseSelfInjector(), ShouldNotBeNil)
		So(c.AsBasicSliceContractsBinder(), ShouldNotBeNil)
		So(c.AsBasicSlicerContractsBinder(), ShouldNotBeNil)
		So(c.AsDynamicLinq(), ShouldNotBeNil)
		So(c.JsonModel(), ShouldNotBeNil)
		So(c.JsonModelAny(), ShouldNotBeNil)

		bytes, err := c.Serialize()
		So(err, ShouldBeNil)
		So(bytes, ShouldNotBeEmpty)
		So(c.SerializeMust(), ShouldNotBeEmpty)
		bytesNoTraces, err := c.SerializeWithoutTraces()
		So(err, ShouldBeNil)
		So(bytesNoTraces, ShouldNotBeEmpty)
		So(c.JsonPtr(), ShouldNotBeNil)
		So(c.JsonResultWithoutTraces(), ShouldNotBeNil)

		mj, err := c.MarshalJSON()
		So(err, ShouldBeNil)
		So(mj, ShouldNotBeEmpty)

		So(c.MutexCollection(), ShouldNotBeNil)
		So(c.ToPtr(), ShouldNotBeNil)
		So(c.StateCounter(), ShouldNotBeNil)
		So(c.StateTracker(), ShouldNotBeNil)
		So(c.AsBaseErrorWrapperCollectionDefiner(), ShouldNotBeNil)

		c.Clear()
		So(c.IsEmpty(), ShouldBeTrue)
	})
}

// Test_Collection_Empty_Defaults — empty collection returns safe zero values.
func Test_Collection_Empty_Defaults(t *testing.T) {
	Convey("Empty collection getters return safe zero values", t, func() {
		c := errwrappers.Empty()
		So(c.HasAnyError(), ShouldBeFalse)
		So(c.IsSuccess(), ShouldBeTrue)
		So(c.IsValid(), ShouldBeTrue)
		So(c.IsFailed(), ShouldBeFalse)
		So(c.GetAsError(), ShouldBeNil)
		So(c.FirstOrDefault(), ShouldNotBeNil) // returns sentinel empty wrapper
		So(c.FirstOrDefaultError(), ShouldBeNil)
		So(c.LastOrDefaultError(), ShouldBeNil)
		So(c.LastIndex(), ShouldEqual, -1)
	})
}

// Test_Concat — ConcatNew variants.
func Test_Concat(t *testing.T) {
	Convey("ConcatNew/ConcatNewClone produce merged collections", t, func() {
		a := errwrappers.NewWithMessage(errtype.InvalidInput, "a")
		b := errwrappers.NewWithMessage(errtype.InvalidInput, "b")
		merged := a.ConcatNew(b)
		So(merged.Count(), ShouldEqual, 2)
		clone := a.ConcatNewClone(b)
		So(clone.Count(), ShouldEqual, 2)
	})
}

// Test_AddIf — conditional add.
func Test_AddIf(t *testing.T) {
	Convey("AddIf only adds when condition is true", t, func() {
		c := errwrappers.Empty()
		c.AddIf(true, errnew.Messages.Single(errtype.InvalidInput, "x"))
		So(c.Count(), ShouldEqual, 1)
		c.AddIf(false, errnew.Messages.Single(errtype.InvalidInput, "y"))
		So(c.Count(), ShouldEqual, 1)
	})
}

// Test_Append — Append alias.
func Test_Append(t *testing.T) {
	Convey("Append adds a wrapper", t, func() {
		c := errwrappers.Empty()
		c.Append(errnew.Messages.Single(errtype.InvalidInput, "x"))
		So(c.Count(), ShouldEqual, 1)
	})
}
