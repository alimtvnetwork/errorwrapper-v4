package main

import (
	"fmt"

	"github.com/alimtvnetwork/errorwrapper-v4/trydo"
)

func GetAsSingleErrorTest1() {
	something := []byte{}

	exception := trydo.WrapPanic(func() {
		panic(stackTraces1Test().GetAsErrorWrapperPtr())
	})

	fmt.Println(something)
	fmt.Println(exception)
}
