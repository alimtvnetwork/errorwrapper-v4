package main

import (
	"fmt"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/enum/scripttype"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
)

func ScriptTest02() {
	cmdOnce := errcmd.New.Script.ArgsDefault(scripttype.Cmd, "dir /w")

	// cmdOnce.RunOnce()

	cmd2 := cmdOnce.CmdCloneWithoutStd()

	// fmt.Println(cmd2.Run())
	bAll, err := cmd2.CombinedOutput()

	fmt.Println(string(bAll))
	fmt.Println(err)

	lines, _ := cmdOnce.CmdCloneCompiledOutputTrimStringLines()

	fmt.Println(strings.Join(lines, constants.NewLineUnix))
	fmt.Println(err)

	result := errcmd.New.Script.ArgsDefaultResult(scripttype.Cmd, "dir /w")
	// time.Sleep(10*time.Second)

	fmt.Println(strings.Join(result.CompiledTrimmedOutputLines(), constants.NewLineUnix))
}
