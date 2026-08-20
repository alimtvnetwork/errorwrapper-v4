package main

import (
	"fmt"

	"github.com/alimtvnetwork/errorwrapper-v4/errcmd"
)

func cmdBuilderAsyncTest2() {
	cmdBuilder := errcmd.New.ScriptBuilder.DefaultPowershell()
	cmdBuilder.Args(
		"Start-Sleep -Seconds 1.5",
	)

	cmdBuilder.Args("ls")
	cmdBuilder.Args("echo something")
	cmdBuilder.Args("xxwfw omething")
	cmd := cmdBuilder.StandardOutputCmd()
	errcmd.CmdSetOsStandardWriters(cmd)

	err := cmd.Start()
	fmt.Println("started but not ended")
	fmt.Println("err", err)

	err2 := cmd.Wait()
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("Finished cmd")
}
