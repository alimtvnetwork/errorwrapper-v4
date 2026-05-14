package main

import (
	"fmt"

	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

func QuickRefTest02() {
	line := errtype.RedisKeyNotFound.ReferencesCsvError(
		"",
		"key-1", "key-2")

	fmt.Println(line)
}
