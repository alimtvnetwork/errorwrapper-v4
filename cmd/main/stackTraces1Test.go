package main

import (
	"errors"
	"fmt"

	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/errwrappers"
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

	stackTraces := codestack.New.StackTrace.Default(codestack.SkipNone, constants.Capacity2)

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
