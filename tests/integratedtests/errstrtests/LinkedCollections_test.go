package errstrtests

import (
	"testing"

	"github.com/alimtvnetwork/errorwrapper-v3/errdata/errstr"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_LinkedCollections_Constructors(t *testing.T) {
	Convey("NewLinkedCollections returns empty collections with nil wrapper", t, func() {
		lc := errstr.NewLinkedCollections()
		So(lc, ShouldNotBeNil)
		So(lc.LinkedCollections, ShouldNotBeNil)
		So(lc.ErrorWrapper, ShouldBeNil)
	})

	Convey("NewLinkedCollectionsUsingItemsError builds collections with items and error", t, func() {
		items := []string{"a", "b"}
		lc := errstr.NewLinkedCollectionsUsingItemsError(errtype.EmptyCollection, nil, items)
		So(lc, ShouldNotBeNil)
		So(lc.LinkedCollections, ShouldNotBeNil)
		So(lc.ErrorWrapper, ShouldNotBeNil)
		So(lc.ErrorWrapper.HasError(), ShouldBeTrue)
	})

	Convey("NewLinkedCollectionsUsingItemsErrorWrapper uses provided wrapper", t, func() {
		items := []string{"x"}
		w := errnew.Message.Type(errtype.Generic, "test")
		lc := errstr.NewLinkedCollectionsUsingItemsErrorWrapper(&items, w)
		So(lc, ShouldNotBeNil)
		So(lc.LinkedCollections, ShouldNotBeNil)
		So(lc.ErrorWrapper, ShouldEqual, w)
	})

	Convey("EmptyLinkedCollectionsUsingError returns empty collections with wrapper", t, func() {
		w := errnew.Message.Type(errtype.Generic, "empty")
		lc := errstr.EmptyLinkedCollectionsUsingError(w)
		So(lc, ShouldNotBeNil)
		So(lc.LinkedCollections, ShouldNotBeNil)
		So(lc.ErrorWrapper, ShouldEqual, w)
	})

	Convey("EmptyLinkedCollections returns empty collections with nil wrapper", t, func() {
		lc := errstr.EmptyLinkedCollections()
		So(lc, ShouldNotBeNil)
		So(lc.LinkedCollections, ShouldNotBeNil)
		So(lc.ErrorWrapper, ShouldBeNil)
	})

	Convey("NewLinkedCollectionsUsingPtrItemsError builds from pointer items", t, func() {
		s1 := "p1"
		s2 := "p2"
		ptrItems := []*string{&s1, &s2}
		lc := errstr.NewLinkedCollectionsUsingPtrItemsError(ptrItems, nil, errtype.EmptySlice)
		So(lc, ShouldNotBeNil)
		So(lc.LinkedCollections, ShouldNotBeNil)
		So(lc.ErrorWrapper, ShouldNotBeNil)
	})
}
