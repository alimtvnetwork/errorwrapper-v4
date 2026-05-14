package main

import (
	"fmt"

	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/refs"
)

func QuickRefTest01() {
	line := refs.QuickCompileStringDefaultEachLine(
		refs.NewQuickReference(errtype.DbUpdateFailed, "RecordKey", "key-5"),
		refs.NewQuickReference(errtype.KeyNotFound, "Key-1"),
		refs.NewQuickReference(errtype.KeyMissing, "Key-2"))

	fmt.Println(line)
}
