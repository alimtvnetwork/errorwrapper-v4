package main

import (
	"fmt"

	"gitlab.com/evatix-go/errorwrapper/errcmd"
)

func cmdBuilderAsyncTest() {
	cmdBuilder := errcmd.New.ScriptBuilder.DefaultPowershell()
	cmdBuilder.Args(
		"Start-Sleep -Seconds 1.5",
	)

	cmdBuilder.Args("ls")
	cmd, err := cmdBuilder.AsyncStartErrorWrapper()
	errcmd.CmdSetOsStandardWriters(cmd)

	fmt.Println("started but not ended")
	fmt.Println("err", err)

	cmd.Wait()

	fmt.Println("Finished cmd")
}
