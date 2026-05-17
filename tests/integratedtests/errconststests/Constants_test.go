package errconststests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/alimtvnetwork/errorwrapper-v3/errconsts"
)

func TestErrconsts_Values(t *testing.T) {
	Convey("Constants should be non-empty strings", t, func() {
		So(errconsts.RangeWithRangeFormat, ShouldNotBeEmpty)
		So(errconsts.VariantStructStringFormat, ShouldNotBeEmpty)
		So(errconsts.ValueHyphenValueFormat, ShouldNotBeEmpty)
		So(errconsts.CombineMessageNoReferencesAdditionalMessageFormat, ShouldNotBeEmpty)
		So(errconsts.CombineMessageNoReferencesNoAdditionalMessageFormat, ShouldNotBeEmpty)
		So(errconsts.SingleReferenceCompile, ShouldNotBeEmpty)
		So(errconsts.CombineMessageAdditionalMessageFormat, ShouldNotBeEmpty)
		So(errconsts.CombineMessageNoAdditionalMessageFormat, ShouldNotBeEmpty)
		So(errconsts.DoubleStringTogetherFormat, ShouldNotBeEmpty)
		So(errconsts.ErrorCodeWithTypeNameFormat, ShouldNotBeEmpty)
		So(errconsts.ErrorCodeHyphenTypeNameFormat, ShouldNotBeEmpty)
		So(errconsts.DirectoryPath, ShouldNotBeEmpty)
		So(errconsts.FilePath, ShouldNotBeEmpty)
		So(errconsts.ErrorStart, ShouldNotBeEmpty)
		So(errconsts.ReferenceStart, ShouldNotBeEmpty)
		So(errconsts.SpaceParenthesisEnd, ShouldNotBeEmpty)
		So(errconsts.ReferenceWithTypeFormat, ShouldNotBeEmpty)
	})

	Convey("Format constants should contain expected placeholders", t, func() {
		So(errconsts.RangeWithRangeFormat, ShouldContainSubstring, "%+v")
		So(errconsts.VariantStructStringFormat, ShouldContainSubstring, "%s")
		So(errconsts.VariantStructStringFormat, ShouldContainSubstring, "%d")
		So(errconsts.ValueHyphenValueFormat, ShouldContainSubstring, "%s")
		So(errconsts.CombineMessageNoReferencesAdditionalMessageFormat, ShouldContainSubstring, "%s")
		So(errconsts.CombineMessageNoReferencesAdditionalMessageFormat, ShouldContainSubstring, "%d")
		So(errconsts.ErrorCodeWithTypeNameFormat, ShouldContainSubstring, "%d")
		So(errconsts.ErrorCodeWithTypeNameFormat, ShouldContainSubstring, "%s")
	})
}
