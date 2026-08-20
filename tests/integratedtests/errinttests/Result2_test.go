package errinttests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errint"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrInt_Result2_Basics(t *testing.T) {
	Convey("Result2 holds two int values", t, func() {
		r := &errint.Result2{
			Result: errint.Result{Value: 1, ErrorWrapper: nil},
			Value2: 2,
		}
		So(r.Value, ShouldEqual, 1)
		So(r.Value2, ShouldEqual, 2)
		So(r.IsAnyNull(), ShouldBeFalse)
	})
}

func Test_ErrInt_ResultWithApplicable_Basics(t *testing.T) {
	Convey("ResultWithApplicable", t, func() {
		r := &errint.ResultWithApplicable{
			Result:       errint.Result{Value: 42, ErrorWrapper: nil},
			IsApplicable: true,
		}
		So(r.Value, ShouldEqual, 42)
		So(r.IsApplicable, ShouldBeTrue)
		So(r.IsAnyNull(), ShouldBeFalse)
	})
}
