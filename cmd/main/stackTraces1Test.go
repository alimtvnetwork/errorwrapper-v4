package main

import (
	"errors"
	"fmt"

	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

func stackTraces1Test() *errwrappers.Collection {
	fmt.Println("----------------------")
	err1 := errnew.Type.Error(
		errtype.SnapshotFailed,
		errors.New("something wrong"))

	err2 := errnew.Type.ErrorUsingStackSkip(
		codestack.Skip2,
		errtype.SnapshotFailed,
		errors.New("something wrong"))

	stackTraces := codestack.NewStacksDefault(codestack.SkipNone, constants.Capacity2)

	err3 := errorwrapper.NewMsgDisplayErrorUsingStackTracesPtr(
		errtype.DbMigration,
		"dbmigrate failed",
		stackTraces,
	)

	// fmt.Println(err1.String())
	// fmt.Println(err1.FullStringWithTraces())

	errCollection := errwrappers.New(10)
	errCollection.AddWrapperPtr(err3)
	errCollection.AddWrapperPtr(err2)
	errCollection.AddWrapperPtr(err1)
	errCollection.Adds(
		errtype.EmptyCollection,
		errtype.Delete,
		errtype.FileDelete)
	// fmt.Println("----------------------")
	// fmt.Println(errCollection.DisplayStringWithTraces())

	return errCollection
}
