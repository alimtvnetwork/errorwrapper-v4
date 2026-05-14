package main

import (
	"fmt"

	"github.com/alimtvnetwork/enum-v10/scripttype"
	"github.com/alimtvnetwork/errorwrapper-v3/errcmd"
)

func ScriptTest01() {
	cmdOnce := errcmd.New.Script.ArgsDefault(scripttype.Cmd, "dir /w")
	rs := cmdOnce.RunOnce()

	lines := rs.CompiledTrimmedOutput()

	fmt.Println(lines)

	bAll, _ := cmdOnce.Cmd.CombinedOutput()
	fmt.Println(bAll)

}
