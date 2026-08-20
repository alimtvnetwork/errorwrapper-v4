package main

import (
	"fmt"

	"github.com/alimtvnetwork/errorwrapper-v4/trydo"
)

func TryDoWrapTest1() {
	exception := trydo.WrapPanic(func() {
		panic(stackTraces1Test().DisplayStringWithLimitTraces(5))
	})

	fmt.Println(exception)
}
