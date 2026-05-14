package main

import (
	"errors"
	"fmt"

	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
)

func testErrNew() {
	errmsg := errnew.Messages
	errMsg := errmsg.Single(errtype.InvalidData, "invalid data")
	fmt.Println(errMsg)

	pathErr := errnew.
		Path.
		Error(
			errtype.NotFound,
			errors.New("file not found"),
			"/path/something")

	// pathErr.HandleError()
	fmt.Println(pathErr)

	var x *string = nil
	nullErr := errnew.Null.Simple(x)
	fmt.Println(nullErr)

	errRef := errnew.
		Ref.
		ManyWithError(
			errtype.ValidationFailed,
			errors.New("validation failed"),
			ref.Value{
				Variable: "x",
				Value:    x,
			},
		)

	fmt.Println(errRef)

	type some struct {
		name string
		age  int
		addr []string
	}

	someValue := &some{
		name: "Doe",
		age:  42,
		addr: []string{"dhaka", "bangladesh"},
	}

	errRef2 := errnew.
		Ref.
		ManyWithError(
			errtype.ValidationFailed,
			errors.New("validation failed"),
			ref.Value{
				Variable: "x",
				Value:    x,
			},
			ref.Value{
				Variable: "someValue",
				Value:    someValue,
			},
		)

	fmt.Println(errRef2)

	errProcess := errnew.Type.Create(errtype.Process).TypeNameWithCustomMessage("abcd process failed")
	fmt.Println(errProcess)

}
