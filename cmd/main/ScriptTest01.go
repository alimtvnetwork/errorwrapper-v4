package main

import (
	"fmt"

	"gitlab.com/evatix-go/enum/scripttype"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
)

func ScriptTest01() {
	cmdOnce := errcmd.New.Script.ArgsDefault(scripttype.Cmd, "dir /w")
	rs := cmdOnce.RunOnce()

	lines := rs.CompiledTrimmedOutput()

	fmt.Println(lines)

	bAll, _ := cmdOnce.Cmd.CombinedOutput()
	fmt.Println(bAll)

}
