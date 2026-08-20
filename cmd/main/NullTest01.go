package main

import (
	"fmt"

	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
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
