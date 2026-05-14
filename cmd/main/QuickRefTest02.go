package main

import (
	"fmt"

	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func QuickRefTest02() {
	line := errtype.RedisKeyNotFound.ReferencesCsvError(
		"",
		"key-1", "key-2")

	fmt.Println(line)
}
