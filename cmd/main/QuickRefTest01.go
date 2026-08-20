package main

import (
	"fmt"

	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/refs"
)

func QuickRefTest01() {
	line := refs.QuickCompileStringDefaultEachLine(
		refs.NewQuickReference(errtype.DbUpdateFailed, "RecordKey", "key-5"),
		refs.NewQuickReference(errtype.KeyNotFound, "Key-1"),
		refs.NewQuickReference(errtype.KeyMissing, "Key-2"))

	fmt.Println(line)
}
