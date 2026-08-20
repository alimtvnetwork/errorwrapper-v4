package errfloat64tests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errfloat64"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrFloat64_Result2_Basics(t *testing.T) {
	Convey("Result2 holds two float64 values", t, func() {
		r := &errfloat64.Result2{
			Result: errfloat64.Result{Value: 1.1, ErrorWrapper: nil},
			Value2: 2.2,
		}
		So(r.Value, ShouldAlmostEqual, 1.1)
		So(r.Value2, ShouldAlmostEqual, 2.2)
		So(r.IsAnyNull(), ShouldBeFalse)
	})
}

func Test_ErrFloat64_ResultWithApplicable_Basics(t *testing.T) {
	Convey("ResultWithApplicable", t, func() {
		r := &errfloat64.ResultWithApplicable{
			Result:       errfloat64.Result{Value: 3.14159, ErrorWrapper: nil},
			IsApplicable: true,
		}
		So(r.Value, ShouldAlmostEqual, 3.14159)
		So(r.IsApplicable, ShouldBeTrue)
		So(r.IsAnyNull(), ShouldBeFalse)
	})
}
