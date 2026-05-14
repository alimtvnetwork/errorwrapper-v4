package main

import (
	"fmt"

	"gitlab.com/evatix-go/errorwrapper/trydo"
)

func GetAsSingleErrorTest1() {
	something := []byte{}

	exception := trydo.WrapPanic(func() {
		panic(stackTraces1Test().GetAsErrorWrapperPtr())
	})

	fmt.Println(something)
	fmt.Println(exception)
}
