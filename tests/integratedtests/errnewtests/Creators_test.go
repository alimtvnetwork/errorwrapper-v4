package errnewtests

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
)

// Test_Messages_Single — most-used creator.
func Test_Messages_Single(t *testing.T) {
	Convey("Messages.Single builds a wrapper with the given type + message", t, func() {
		w := errnew.Messages.Single(errtype.InvalidInput, "bad payload")
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
		So(w.Type(), ShouldEqual, errtype.InvalidInput)
		So(w.FullString(), ShouldContainSubstring, "bad payload")
	})
}

// Test_Messages_Many — joins messages.
func Test_Messages_Many(t *testing.T) {
	Convey("Messages.Many concatenates all messages into the wrapper", t, func() {
		w := errnew.Messages.Many(errtype.MappingFailed, "first", "second", "third")
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
		So(w.Type(), ShouldEqual, errtype.MappingFailed)
		full := w.FullString()
		So(full, ShouldContainSubstring, "first")
		So(full, ShouldContainSubstring, "second")
		So(full, ShouldContainSubstring, "third")
	})
}

// Test_Type_Default — pure type-only constructor.
func Test_Type_Default(t *testing.T) {
	Convey("Type.Default returns a wrapper carrying just the variation", t, func() {
		w := errnew.Type.Default(errtype.NotFound)
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
		So(w.Type(), ShouldEqual, errtype.NotFound)
	})
}

// Test_NotFound_Family — convenience creators.
func Test_NotFound_Family(t *testing.T) {
	Convey("NotFound.Type produces a NotFound-typed wrapper", t, func() {
		w := errnew.NotFound.Type()
		So(w, ShouldNotBeNil)
		So(w.Type(), ShouldEqual, errtype.NotFound)
	})

	Convey("NotFound.Error returns nil for a nil error and a wrapper otherwise", t, func() {
		So(errnew.NotFound.Error(nil), ShouldBeNil)

		w := errnew.NotFound.Error(errors.New("user 42 missing"))
		So(w, ShouldNotBeNil)
		So(w.Type(), ShouldEqual, errtype.NotFound)
		So(w.FullString(), ShouldContainSubstring, "user 42 missing")
	})

	Convey("NotFound.Reference accepts an arbitrary reference value", t, func() {
		w := errnew.NotFound.Reference("missing-thing")
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
		So(w.Type(), ShouldEqual, errtype.NotFound)
	})
}

// Test_Error_NoType — wraps a raw error with no specific category.
func Test_Error_NoType(t *testing.T) {
	Convey("Error.NoType wraps a raw error and preserves its text", t, func() {
		raw := errors.New("disk full")
		w := errnew.Error.NoType(raw)
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
		So(w.FullString(), ShouldContainSubstring, "disk full")
	})

	Convey("Error.NoType returns nil for nil input", t, func() {
		So(errnew.Error.NoType(nil), ShouldBeNil)
	})
}

// Test_Ref_WithReference — verifies reference-attached creators round-trip.
func Test_Ref_WithReference(t *testing.T) {
	Convey("Messages.WithRef attaches a ref.Value to the wrapper", t, func() {
		r := ref.New("userId", "42")
		w := errnew.Messages.WithRef(errtype.InvalidId, r, "bad user id")
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
		So(w.Type(), ShouldEqual, errtype.InvalidId)
		full := w.FullString()
		So(full, ShouldContainSubstring, "bad user id")
		So(full, ShouldContainSubstring, "42")
	})
}

// Test_PreBaked_Vars — the package-level pre-constructed wrappers.
func Test_PreBaked_Vars(t *testing.T) {
	Convey("Pre-baked variables are non-nil and carry the right variation", t, func() {
		So(errnew.MappingFailed, ShouldNotBeNil)
		So(errnew.MappingFailed.Type(), ShouldEqual, errtype.MappingFailed)

		So(errnew.InvalidInput, ShouldNotBeNil)
		So(errnew.InvalidInput.Type(), ShouldEqual, errtype.InvalidInput)

		So(errnew.EmptyString, ShouldNotBeNil)
		So(errnew.EmptyString.Type(), ShouldEqual, errtype.EmptyString)

		So(errnew.Unexpected, ShouldNotBeNil)
		So(errnew.Unexpected.Type(), ShouldEqual, errtype.Unexpected)

		So(errnew.FinalizedResourceCannotAccess, ShouldNotBeNil)
		So(errnew.FinalizedResourceCannotAccess.Type(), ShouldEqual, errtype.FinalizedResourceCannotAccess)
	})
}

// Test_Messages_ErrorWithMany — folds an error + extra messages.
func Test_Messages_ErrorWithMany(t *testing.T) {
	Convey("Messages.ErrorWithMany returns nil when err is nil", t, func() {
		So(errnew.Messages.ErrorWithMany(errtype.IO, nil, "ctx"), ShouldBeNil)
	})

	Convey("Messages.ErrorWithMany combines err text with extra messages", t, func() {
		w := errnew.Messages.ErrorWithMany(
			errtype.IO,
			errors.New("permission denied"),
			"while reading",
			"/etc/config")
		So(w, ShouldNotBeNil)
		So(w.HasError(), ShouldBeTrue)
		So(w.Type(), ShouldEqual, errtype.IO)
		full := w.FullString()
		So(full, ShouldContainSubstring, "permission denied")
		So(full, ShouldContainSubstring, "while reading")
		So(full, ShouldContainSubstring, "/etc/config")
	})
}
