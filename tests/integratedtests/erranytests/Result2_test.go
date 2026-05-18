package erranytests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errany"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrAny_Result2_Basics(t *testing.T) {
	Convey("Result2 holds two values", t, func() {
		r := &errany.Result2{
			Result: errany.Result{Value: "first", ErrorWrapper: nil},
			Value2: "second",
		}
		So(r.Value, ShouldEqual, "first")
		So(r.Value2, ShouldEqual, "second")
		So(r.IsAnyNull(), ShouldBeFalse)
	})
}

func Test_ErrAny_ResultWithApplicable_Basics(t *testing.T) {
	Convey("ResultWithApplicable", t, func() {
		r := &errany.ResultWithApplicable{
			Result:       errany.Result{Value: "test", ErrorWrapper: nil},
			IsApplicable: true,
		}
		So(r.Value, ShouldEqual, "test")
		So(r.IsApplicable, ShouldBeTrue)
		So(r.IsAnyNull(), ShouldBeFalse)
	})
}
