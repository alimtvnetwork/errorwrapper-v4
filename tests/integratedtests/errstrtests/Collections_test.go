package errstrtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/errorwrapper-v4/errdata/errstr"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_ErrStr_Collection_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var c *errstr.Collection
		So(c.IsEmpty(), ShouldBeTrue)
		So(c.HasError(), ShouldBeFalse)
		So(c.HasSafeItems(), ShouldBeFalse)
		So(c.IsSuccess(), ShouldBeFalse)
		So(c.IsValid(), ShouldBeFalse)
		So(c.IsFailed(), ShouldBeFalse)
	})

	Convey("Empty collection without error", t, func() {
		c := &errstr.Collection{
			Collection:   corestr.Empty.Collection(),
			ErrorWrapper: nil,
		}
		So(c.IsEmpty(), ShouldBeTrue)
		So(c.HasError(), ShouldBeFalse)
		So(c.HasSafeItems(), ShouldBeFalse)
		So(c.IsSuccess(), ShouldBeFalse)
		So(c.IsValid(), ShouldBeFalse)
		So(c.IsFailed(), ShouldBeFalse)
	})

	Convey("Non-empty without error", t, func() {
		c := &errstr.Collection{
			Collection:   corestr.New.Collection.Strings([]string{"a", "b"}),
			ErrorWrapper: nil,
		}
		So(c.IsEmpty(), ShouldBeFalse)
		So(c.HasError(), ShouldBeFalse)
		So(c.HasSafeItems(), ShouldBeTrue)
		So(c.IsSuccess(), ShouldBeTrue)
		So(c.IsValid(), ShouldBeTrue)
		So(c.IsFailed(), ShouldBeFalse)
	})

	Convey("With error wrapper", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		c := &errstr.Collection{
			Collection:   corestr.New.Collection.Strings([]string{"a"}),
			ErrorWrapper: w,
		}
		So(c.HasError(), ShouldBeTrue)
		So(c.IsEmptyError(), ShouldBeTrue)
		So(c.HasSafeItems(), ShouldBeFalse)
		So(c.IsSuccess(), ShouldBeFalse)
		So(c.IsFailed(), ShouldBeTrue)
	})
}

func Test_ErrStr_LinkedList_Basics(t *testing.T) {
	Convey("NewLinkedList", t, func() {
		ll := errstr.NewLinkedList()
		So(ll, ShouldNotBeNil)
		So(ll.LinkedList, ShouldNotBeNil)
	})

	Convey("NewLinkedListUsingItemsError", t, func() {
		ll := errstr.NewLinkedListUsingItemsError(
			errtype.InvalidValidate, errors.New("bad"),
			[]string{"a", "b"})
		So(ll, ShouldNotBeNil)
		So(ll.LinkedList, ShouldNotBeNil)
		So(ll.ErrorWrapper.HasError(), ShouldBeTrue)
	})

	Convey("NewLinkedListUsingItemsErrorWrapper", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		items := []string{"a"}
		ll := errstr.NewLinkedListUsingItemsErrorWrapper(&items, w)
		So(ll, ShouldNotBeNil)
		So(ll.ErrorWrapper.HasError(), ShouldBeTrue)
	})

	Convey("EmptyLinkedListUsingError", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		ll := errstr.EmptyLinkedListUsingError(w)
		So(ll, ShouldNotBeNil)
		So(ll.ErrorWrapper.HasError(), ShouldBeTrue)
	})

	Convey("EmptyLinkedList", t, func() {
		ll := errstr.EmptyLinkedList()
		So(ll, ShouldNotBeNil)
	})

	Convey("NewLinkedListUsingPtrItemsError", t, func() {
		ptrItems := []*string{}
		ll := errstr.NewLinkedListUsingPtrItemsError(
			ptrItems, errors.New("bad"), errtype.InvalidValidate)
		So(ll, ShouldNotBeNil)
		So(ll.ErrorWrapper.HasError(), ShouldBeTrue)
	})
}

func Test_ErrStr_SimpleStringOnce_Basics(t *testing.T) {
	Convey("Nil receiver", t, func() {
		var s *errstr.SimpleStringOnce
		So(s.IsEmpty(), ShouldBeTrue)
		So(s.HasError(), ShouldBeFalse)
		So(s.IsEmptyError(), ShouldBeTrue)
		So(s.HasSafeItems(), ShouldBeFalse)
		So(s.IsValid(), ShouldBeTrue)
		So(s.IsSuccess(), ShouldBeFalse)
		So(s.IsFailed(), ShouldBeTrue)
		So(s.HasIssuesOrEmpty(), ShouldBeTrue)
		So(s.Int(), ShouldEqual, 0)
		So(s.Byte(), ShouldEqual, 0)
		So(s.Bool(), ShouldBeFalse)
		So(s.SafeString(), ShouldEqual, "")
		So(func() { s.Dispose() }, ShouldNotPanic)
		So(func() { s.ValidValue() }, ShouldNotPanic)
		So(s.SplitLines(), ShouldResemble, []string{})
	})

	Convey("Empty SimpleStringOnce", t, func() {
		s := &errstr.SimpleStringOnce{
			Value:        corestr.Empty.SimpleStringOnce(),
			ErrorWrapper: nil,
		}
		So(s.IsEmpty(), ShouldBeTrue)
		So(s.HasIssuesOrEmpty(), ShouldBeTrue)
	})

	Convey("Initialized value without error", t, func() {
		s := &errstr.SimpleStringOnce{
			Value:        corestr.New.SimpleStringOnce.Create("hello", true),
			ErrorWrapper: nil,
		}
		So(s.IsEmpty(), ShouldBeFalse)
		So(s.IsEmptyOrWhitespace(), ShouldBeFalse)
		So(s.HasError(), ShouldBeFalse)
		So(s.HasIssuesOrEmpty(), ShouldBeFalse)
		So(s.HasSafeItems(), ShouldBeTrue)
		So(s.IsSuccess(), ShouldBeTrue)
		So(s.IsFailed(), ShouldBeFalse)
		So(s.String(), ShouldEqual, "hello")
		So(s.SafeString(), ShouldEqual, "hello")
		So(s.Bool(), ShouldBeFalse)
		So(s.Int(), ShouldEqual, 0)
		So(s.IsEqual("hello"), ShouldBeTrue)
		So(s.IsEqualIgnoreCase("HELLO"), ShouldBeTrue)
	})

	Convey("Numeric string", t, func() {
		s := &errstr.SimpleStringOnce{
			Value:        corestr.New.SimpleStringOnce.Create("42", true),
			ErrorWrapper: nil,
		}
		So(s.Int(), ShouldEqual, 42)
		So(s.Byte(), ShouldEqual, 42)
	})

	Convey("SplitLines", t, func() {
		s := &errstr.SimpleStringOnce{
			Value:        corestr.New.SimpleStringOnce.Create("a\nb", true),
			ErrorWrapper: nil,
		}
		So(s.SplitLines(), ShouldResemble, []string{"a", "b"})
	})

	Convey("Value with error", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "bad")
		s := &errstr.SimpleStringOnce{
			Value:        corestr.New.SimpleStringOnce.Create("x", true),
			ErrorWrapper: w,
		}
		So(s.HasError(), ShouldBeTrue)
		So(s.IsEmptyError(), ShouldBeFalse)
		So(s.IsSuccess(), ShouldBeFalse)
		So(s.HasIssuesOrEmpty(), ShouldBeTrue)
		So(s.ErrorWrapperInf(), ShouldNotBeNil)
	})

	Convey("Dispose", t, func() {
		w := errnew.Message.New(errtype.InvalidValidate, "x")
		s := &errstr.SimpleStringOnce{
			Value:        corestr.New.SimpleStringOnce.Create("test", true),
			ErrorWrapper: w,
		}
		So(func() { s.Dispose() }, ShouldNotPanic)
	})
}
