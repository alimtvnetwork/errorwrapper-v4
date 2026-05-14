package main

import (
	"errors"
	"fmt"

	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/ref"
)

func fromToErrMergeExample() {
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

	var x *string = nil
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

	fromToErr := errnew.FromTo.Messages(errtype.ConversionFailed, true, x, someValue)
	fmt.Println("-------------- ")
	mergeErr := errnew.Merge.New(fromToErr, errRef2)
	fmt.Println(mergeErr.String())
	fmt.Println(mergeErr.String())
	fmt.Println(mergeErr.CompileString())
	fmt.Println(mergeErr.ErrorString())
}
