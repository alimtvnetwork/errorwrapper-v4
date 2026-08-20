package errfloattests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errfloat"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrFloat_Result2_Basics(t *testing.T) {
	Convey("Result2 holds two float32 values", t, func() {
		r := &errfloat.Result2{
			Result: errfloat.Result{Value: 1.1, ErrorWrapper: nil},
			Value2: 2.2,
		}
		So(r.Value, ShouldAlmostEqual, 1.1)
		So(r.Value2, ShouldAlmostEqual, 2.2)
		So(r.IsAnyNull(), ShouldBeFalse)
	})
}

func Test_ErrFloat_ResultWithApplicable_Basics(t *testing.T) {
	Convey("ResultWithApplicable", t, func() {
		r := &errfloat.ResultWithApplicable{
			Result:       errfloat.Result{Value: 3.14, ErrorWrapper: nil},
			IsApplicable: true,
		}
		So(r.Value, ShouldAlmostEqual, 3.14)
		So(r.IsApplicable, ShouldBeTrue)
		So(r.IsAnyNull(), ShouldBeFalse)
	})
}
