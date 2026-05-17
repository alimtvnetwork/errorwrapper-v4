package errwrapperstests

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
)

// Test_Empty_NoError — empty collection reports no error.
func Test_Empty_NoError(t *testing.T) {
	Convey("Empty() collection has no error and length 0", t, func() {
		c := errwrappers.Empty()
		So(c, ShouldNotBeNil)
		So(c.HasAnyError(), ShouldBeFalse)
		So(c.IsEmpty(), ShouldBeTrue)
		So(c.Count(), ShouldEqual, 0)
	})
}

// Test_AddError_ReportsError — AddError flips HasAnyError to true.
func Test_AddError_ReportsError(t *testing.T) {
	Convey("AddError increments count and HasAnyError becomes true", t, func() {
		c := errwrappers.Empty()
		c.AddError(errors.New("boom"))
		So(c.HasAnyError(), ShouldBeTrue)
		So(c.IsEmpty(), ShouldBeFalse)
		So(c.Count(), ShouldEqual, 1)
		So(c.FullString(), ShouldContainSubstring, "boom")
	})
}

// Test_AddWrapperPtr — wrapper pushed in keeps its type and message.
func Test_AddWrapperPtr(t *testing.T) {
	Convey("AddWrapperPtr adds an existing *Wrapper to the collection", t, func() {
		c := errwrappers.Empty()
		w := errnew.Messages.Single(errtype.InvalidInput, "bad-input")
		c.AddWrapperPtr(w)
		So(c.HasAnyError(), ShouldBeTrue)
		So(c.Count(), ShouldEqual, 1)
		So(c.FullString(), ShouldContainSubstring, "bad-input")
	})
}

// Test_NewWithMessage_Constructor — collection seeded with a message.
func Test_NewWithMessage_Constructor(t *testing.T) {
	Convey("NewWithMessage seeds a collection with a single typed message", t, func() {
		c := errwrappers.NewWithMessage(errtype.NotFound, "missing-row")
		So(c, ShouldNotBeNil)
		So(c.HasAnyError(), ShouldBeTrue)
		So(c.Count(), ShouldEqual, 1)
		So(c.FullString(), ShouldContainSubstring, "missing-row")
	})
}

// Test_StateCounter_HasChanges — StateCounter tracks collection size changes.
func Test_StateCounter_HasChanges(t *testing.T) {
	Convey("StateCounter detects new errors added after start", t, func() {
		c := errwrappers.Empty()
		sc := errwrappers.NewStateCount(c)
		So(sc.HasChanges(), ShouldBeFalse)
		c.AddError(errors.New("late-failure"))
		So(sc.HasChanges(), ShouldBeTrue)
		So(sc.IsSameState(), ShouldBeFalse)
	})
}

// Test_MutexEmpty — mutex variant constructs without panic.
func Test_MutexEmpty(t *testing.T) {
	Convey("MutexEmpty returns a usable empty MutexCollection", t, func() {
		mc := errwrappers.MutexEmpty()
		So(mc, ShouldNotBeNil)
	})
}
