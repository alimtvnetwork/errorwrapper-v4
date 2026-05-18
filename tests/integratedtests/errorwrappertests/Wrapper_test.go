package errorwrappertests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

// Test_EmptyPtr_HasError — EmptyPtr() returns nil sentinel; nil receivers
// must still report HasError=false and IsEmpty=true.
func Test_EmptyPtr_HasError(t *testing.T) {
	Convey("EmptyPtr() returns nil sentinel that reports empty/no-error", t, func() {
		w := errorwrapper.EmptyPtr()
		So(w, ShouldBeNil)
		So(w.HasError(), ShouldBeFalse)
		So(w.IsEmpty(), ShouldBeTrue)
	})
}

// Test_NewPtr_TypeOnly — type-only wrapper still empty (no message/ref).
func Test_NewPtr_TypeOnly(t *testing.T) {
	Convey("NewPtr with type but no message carries the variation", t, func() {
		w := errorwrapper.NewPtr(errtype.InvalidInput)
		So(w, ShouldNotBeNil)
		So(w.Type(), ShouldEqual, errtype.InvalidInput)
	})
}

// Test_HasError_WithMessage — Messages.Single produces HasError true.
func Test_HasError_WithMessage(t *testing.T) {
	Convey("A wrapper built with a message reports HasError true", t, func() {
		w := errnew.Messages.Single(errtype.InvalidInput, "boom")
		So(w.HasError(), ShouldBeTrue)
		So(w.IsEmpty(), ShouldBeFalse)
		So(w.Error(), ShouldNotBeNil)
		So(w.FullString(), ShouldContainSubstring, "boom")
	})
}

// Test_ConcatNew_Message — chaining additional messages preserves type & content.
func Test_ConcatNew_Message(t *testing.T) {
	Convey("ConcatNew().Message appends to an existing wrapper", t, func() {
		base := errnew.Messages.Single(errtype.MappingFailed, "first")
		chained := base.ConcatNew().Message("second")
		So(chained, ShouldNotBeNil)
		So(chained.HasError(), ShouldBeTrue)
		full := chained.FullString()
		So(full, ShouldContainSubstring, "first")
		So(full, ShouldContainSubstring, "second")
	})
}

// Test_ClonePtr — clone yields independent wrapper with same data.
func Test_ClonePtr(t *testing.T) {
	Convey("ClonePtr returns an independent copy carrying the same type+message", t, func() {
		base := errnew.Messages.Single(errtype.NotFound, "missing")
		clone := base.ClonePtr()
		So(clone, ShouldNotBeNil)
		So(clone, ShouldNotPointTo, base)
		So(clone.Type(), ShouldEqual, base.Type())
		So(clone.FullString(), ShouldContainSubstring, "missing")
	})
}

// Test_CompiledError — CompiledError returns standard error for non-empty wrapper.
func Test_CompiledError(t *testing.T) {
	Convey("CompiledError returns nil for empty, non-nil for error wrapper", t, func() {
		So(errorwrapper.EmptyPtr().CompiledError(), ShouldBeNil)
		w := errnew.Messages.Single(errtype.InvalidInput, "x")
		So(w.CompiledError(), ShouldNotBeNil)
	})
}
