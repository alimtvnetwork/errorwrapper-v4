package errjsontests

import (
	"testing"

	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errjson"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrJson_ResultsCollection_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var r *errjson.ResultsCollection
		So(r.IsAnyNull(), ShouldBeTrue)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
		So(r.Length(), ShouldEqual, 0)
		So(func() { r.Dispose() }, ShouldNotPanic)
	})

	Convey("Empty ResultsCollection", t, func() {
		r := &errjson.ResultsCollection{
			ResultsCollection: corejson.Empty.ResultsCollection(),
			ErrorCollection:   errwrappers.EmptyCollection(),
		}
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeTrue)
		So(r.HasError(), ShouldBeFalse)
		So(r.Length(), ShouldEqual, 0)
	})

	Convey("Non-empty without error", t, func() {
		r := &errjson.ResultsCollection{
			ResultsCollection: corejson.NewResultsCollection.AnyItems(
				[]byte(`{"a":1}`),
				[]byte(`{"b":2}`),
			),
			ErrorCollection: errwrappers.EmptyCollection(),
		}
		So(r.IsAnyNull(), ShouldBeFalse)
		So(r.IsEmpty(), ShouldBeFalse)
		So(r.HasError(), ShouldBeFalse)
		So(r.Length(), ShouldBeGreaterThan, 0)
	})

	Convey("Dispose", t, func() {
		r := &errjson.ResultsCollection{
			ResultsCollection: corejson.NewResultsCollection.AnyItems(
				[]byte(`{"x":1}`),
			),
			ErrorCollection: errwrappers.EmptyCollection(),
		}
		So(func() { r.Dispose() }, ShouldNotPanic)
	})
}
