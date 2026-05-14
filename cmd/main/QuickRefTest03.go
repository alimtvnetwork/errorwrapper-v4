package main

import (
	"fmt"

	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

func QuickRefTest03() {
	line := errtype.RedisKeyNotFound.ReferencesLines(
		"something wrong",
		errtype.KeyNotFound.ShortReferencesCsv("key-1"),
		errtype.DbRecordNotFound.ShortReferencesCsv("db-record-1"))

	fmt.Println(line)

	empty := errwrappers.Empty()

	fmt.Println(empty)
}
