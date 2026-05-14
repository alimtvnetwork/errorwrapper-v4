package main

import (
	"fmt"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

func NullTest01() {
	var errWp *errorwrapper.Wrapper
	var errWpC *errwrappers.Collection

	err1 := errnew.Null.ManyByChecking(errWpC, errWp)

	fmt.Println(err1)

	err2 := errnew.Null.ManyByChecking(nil, errWp)

	fmt.Println(err2)

	err3 := errnew.Null.ManyWithMessage("Something wrong", nil, errWp)

	fmt.Println(err3)
}
