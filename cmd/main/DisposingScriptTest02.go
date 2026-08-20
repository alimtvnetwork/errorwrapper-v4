package main

import (
	"fmt"

	"github.com/alimtvnetwork/enum-v10/scripttype"
	"github.com/alimtvnetwork/errorwrapper-v4/errcmd"
)

func DisposingScriptTest02() {
	cmdOnceCollection := errcmd.NewCmdOnceCollectionUsingLinesOfScripts(
		scripttype.Cmd,
		"dir /w")

	fmt.Println(cmdOnceCollection.Strings())
	cmdOnceCollection.Dispose()
}
