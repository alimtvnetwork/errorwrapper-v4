package errjsontests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errjson"

	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrJson_Constructors_New(t *testing.T) {
	Convey("New.Result.Empty", t, func() {
		r := errjson.New.Result.Empty()
		So(r, ShouldNotBeNil)
		So(r.IsAnyNull(), ShouldBeTrue)
	})

	Convey("New.Result.Item", t, func() {
		jr := corejson.NewPtr([]byte(`{"x":1}`))
		r := errjson.New.Result.Item(jr)
		So(r, ShouldNotBeNil)
		So(r.IsAnyNull(), ShouldBeFalse)
	})

	Convey("New.Result.Error", t, func() {
		r := errjson.New.Result.Error(errtype.InvalidValidate, errors.New("bad"))
		So(r, ShouldNotBeNil)
		So(r.HasError(), ShouldBeTrue)
	})

	Convey("New.Result.Any", t, func() {
		r := errjson.New.Result.Any(`{"ok":true}`)
		So(r, ShouldNotBeNil)
	})

	Convey("New.Result.Bytes", t, func() {
		r := errjson.New.Result.Bytes([]byte(`{"ok":true}`))
		So(r, ShouldNotBeNil)
	})

	Convey("New.Result.Marshal", t, func() {
		r := errjson.New.Result.Marshal(map[string]bool{"ok": true})
		So(r, ShouldNotBeNil)
	})
}

func Test_ErrJson_Constructors_Empty(t *testing.T) {
	Convey("Empty.Result", t, func() {
		r := errjson.Empty.Result()
		So(r, ShouldNotBeNil)
		So(r.IsAnyNull(), ShouldBeTrue)
	})

	Convey("Empty.ResultsCollection", t, func() {
		r := errjson.Empty.ResultsCollection()
		So(r, ShouldNotBeNil)
	})
}
