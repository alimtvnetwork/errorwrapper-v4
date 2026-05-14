package main

import (
	"fmt"

	"gitlab.com/evatix-go/enum/scripttype"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
)

func DisposingScriptTest02() {
	cmdOnceCollection := errcmd.NewCmdOnceCollectionUsingLinesOfScripts(
		scripttype.Cmd,
		"dir /w")

	fmt.Println(cmdOnceCollection.Strings())
	cmdOnceCollection.Dispose()
}
