package main

import (
	"errors"
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/errorwrapper/ref"
)

func main() {
	// marshallingIssueTest()

	// TryDoWrapTest1()
	// GetAsSingleErrorTest1()

	// stackTraces1Test()
	// TryDoWrapTest1()
	// GetAsSingleErrorTest1()

	// QuickRefTest01()
	// QuickRefTest02()
	// QuickRefTest03()
	// ScriptTest02()
	// NullTest01()

	// ScriptTest01()
	// ScriptTest02()
	// DisposingScriptTest02()

	// testErrNew()
	// fromToErrMergeExample()
	// cmdBuilderTest()
	// cmdBuilderAsyncTest()
	// cmdBuilderAsyncTest2()
	// errCreationSimpleStringOnceTest()
	// serializeDeserializeTest()
	errorCollectionTest01()
}

func errorCollectionTest01() {
	collection := errwrappers.Empty()

	collection.AddError(errors.New("something"))

	collection.AddWrapperPtr(errnew.Ref.ManyWithMsg(
		errtype.AccountExpired,
		"some expired issues 2",
		ref.Value{
			Variable: "something",
			Value:    22,
		}))

	isCollect := collection.IsCollectOn(
		false,
		errnew.Ref.ManyWithMsg(
			errtype.AccountExpired,
			"some expired issues 3",
			ref.Value{
				Variable: "something",
				Value:    22,
			}))

	fmt.Println(isCollect)

	isCollect = collection.IsCollectOn(
		true,
		errnew.Ref.ManyWithMsg(
			errtype.AccountExpired,
			"some expired issues 4",
			ref.Value{
				Variable: "something",
				Value:    22,
			}))

	fmt.Println(isCollect)
	// fmt.Println(collection.FullStringWithTraces())
	fmt.Println(collection.CompiledJsonErrorWithStackTraces())

	jsonResult := collection.Json()

	newErrCollection, errWp := errwrappers.Deserialize.UsingJsonResult(jsonResult.Ptr())

	errWp.HandleError()
	fmt.Println(newErrCollection.DisplayStringWithTraces())

	newErrCollection2, errWp2 := errwrappers.Deserialize.UsingJsonResult(nil)

	errWp2.HandleError()
	fmt.Println("final err 1 ", newErrCollection2.DisplayStringWithTraces())

	newErrCollection3, errWp3 := errwrappers.Deserialize.UsingBytes(nil)

	errWp3.HandleError()
	fmt.Println("final err 2 ", newErrCollection3.DisplayStringWithTraces())

	newErrCollection4, errWp4 := errwrappers.Deserialize.UsingBytes(jsonResult.SafeBytes())

	errWp4.HandleError()
	fmt.Println("final err 3 ", newErrCollection4.DisplayStringWithTraces())

	jsonRs := corejson.NewResult.UsingBytes([]byte("some byte which will fail"))
	newErrCollection5, errWp5 := errwrappers.Deserialize.UsingJsonResult(jsonRs.Ptr())

	errWp5.HandleError()
	fmt.Println("final err 5 ", newErrCollection5.DisplayStringWithTraces())
}

func serializeDeserializeTest() {
	newErr := errnew.NotFound.Invalid(
		"some message", "some ref1")

	json := newErr.JsonPtr()
	fmt.Println("First Err", newErr.CompiledErrorWithStackTraces())
	fmt.Println("First Json Err", json.String())

	newErr3 := errnew.DeserializeTo.JsonResultToWrapper(json)

	fmt.Println("2nd Err", newErr3.CompiledErrorWithStackTraces())
	fmt.Println("2nd Json Err", newErr3.Json().String())

	newErr4 := errnew.DeserializeTo.JsonStringToWrapper(newErr3.Json().JsonString())

	fmt.Println("3rd Err", newErr4.CompiledErrorWithStackTraces())
	fmt.Println("3rd Json Err", newErr4.Json().String())

}
