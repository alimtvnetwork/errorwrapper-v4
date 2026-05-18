package errbytetests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errbyte"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrByte_Result2_Basics(t *testing.T) {
	Convey("Result2 holds two byte values", t, func() {
		r := &errbyte.Result2{
			Result: errbyte.Result{Value: 65, ErrorWrapper: nil},
			Value2: 66,
		}
		So(r.Value, ShouldEqual, 65)
		So(r.Value2, ShouldEqual, 66)
		So(r.IsAnyNull(), ShouldBeFalse)
	})
}

func Test_ErrByte_ResultWithApplicable_Basics(t *testing.T) {
	Convey("ResultWithApplicable", t, func() {
		r := &errbyte.ResultWithApplicable{
			Result:       errbyte.Result{Value: 65, ErrorWrapper: nil},
			IsApplicable: true,
		}
		So(r.Value, ShouldEqual, 65)
		So(r.IsApplicable, ShouldBeTrue)
		So(r.IsAnyNull(), ShouldBeFalse)
	})
}
