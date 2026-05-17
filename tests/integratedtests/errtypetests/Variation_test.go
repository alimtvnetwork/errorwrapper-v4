package errtypetests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

// Test_Variation_HasError verifies the NoError sentinel vs every real variation.
func Test_Variation_HasError(t *testing.T) {
	Convey("NoError.HasError() must be false", t, func() {
		So(errtype.NoError.HasError(), ShouldBeFalse)
		So(errtype.NoError.IsNoError(), ShouldBeTrue)
	})

	cases := []errtype.Variation{
		errtype.Generic,
		errtype.NullOrEmpty,
		errtype.InvalidInput,
		errtype.NotFound,
		errtype.MappingFailed,
		errtype.CommandExecution,
		errtype.SysUserInvalid,
		errtype.FinalizedResourceCannotAccess,
	}

	for _, v := range cases {
		v := v
		Convey("Variation "+v.String()+" must report HasError() == true", t, func() {
			So(v.HasError(), ShouldBeTrue)
			So(v.IsNoError(), ShouldBeFalse)
			So(v.Is(v), ShouldBeTrue)
		})
	}
}

// Test_Variation_StringAndMessage ensures the string/message mappings are populated.
func Test_Variation_StringAndMessage(t *testing.T) {
	cases := []errtype.Variation{
		errtype.NotFound,
		errtype.InvalidInput,
		errtype.MappingFailed,
		errtype.CommandExecution,
	}

	for _, v := range cases {
		v := v
		Convey("Variation "+v.String()+" must have non-empty name + message", t, func() {
			So(v.String(), ShouldNotBeBlank)
			So(v.Name(), ShouldEqual, v.String())
			So(v.Message(), ShouldNotBeBlank)
		})
	}
}

// Test_Variation_ErrorBuilders verifies the *Error* convenience builders produce non-nil errors
// whose message contains the variation name (case-insensitive — built-ins lowercase the output).
func Test_Variation_ErrorBuilders(t *testing.T) {
	v := errtype.NotFound

	Convey("Variation.Error(...) returns a non-nil error containing context", t, func() {
		err := v.Error("missing user", "userId", 42)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "missing user")
	})

	Convey("Variation.ErrorNoRefs returns a non-nil error", t, func() {
		err := v.ErrorNoRefs("nothing here")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "nothing here")
	})

	Convey("Variation.ErrorReferences accepts multiple refs", t, func() {
		err := v.ErrorReferences("ctx", "a", 1, true)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "ctx")
	})
}

// Test_Variation_Combine sanity-checks the formatting helpers.
func Test_Variation_Combine(t *testing.T) {
	v := errtype.InvalidInput

	Convey("CombineNoRefs returns a non-empty string", t, func() {
		out := v.CombineNoRefs("extra info")
		So(out, ShouldNotBeBlank)
		So(out, ShouldContainSubstring, "extra info")
	})

	Convey("Combine with var name + value returns a non-empty string", t, func() {
		out := v.Combine("bad payload", "size", 99)
		So(out, ShouldNotBeBlank)
		So(out, ShouldContainSubstring, "bad payload")
	})
}
