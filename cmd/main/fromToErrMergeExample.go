package main

import (
	"errors"
	"fmt"

	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
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
