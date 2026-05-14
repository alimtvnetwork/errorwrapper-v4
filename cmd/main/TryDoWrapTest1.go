package main

import (
	"fmt"

	"gitlab.com/evatix-go/errorwrapper/trydo"
)

func TryDoWrapTest1() {
	exception := trydo.WrapPanic(func() {
		panic(stackTraces1Test().DisplayStringWithLimitTraces(5))
	})

	fmt.Println(exception)
}
