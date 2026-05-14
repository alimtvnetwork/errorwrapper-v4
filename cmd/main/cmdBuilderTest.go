package main

import (
	"fmt"

	"github.com/alimtvnetwork/errorwrapper-v3/errcmd"
)

func cmdBuilderTest() {
	cmdBuilder := errcmd.New.ScriptBuilder.DefaultCmd()
	cmdBuilder.Args(
		"dir/w",
	)

	cmdBuilder.ArgsIf(
		true,
		"c: && dir/w",
	)

	lines := cmdBuilder.Output()

	fmt.Println(lines)

	cmdBuilder.Args("ls to give error")

	cmd := cmdBuilder.StandardOutputCmd()

	cmd.Run()

	cmdBuilder.AppendCmd(cmd)
	cmdBuilder.Append(cmdBuilder)
	scriptLines := cmdBuilder.ScriptLines()

	fmt.Println(scriptLines.JoinLine())

	cmd2 := cmdBuilder.BuildCmd()
	combinedAll, err := cmd2.CombinedOutput()

	fmt.Println("combined all :", string(combinedAll))
	fmt.Println("combined err :", err)

	combinedAll2, err2 := cmdBuilder.CombinedOutput()

	fmt.Println("combined all :", string(combinedAll2))
	fmt.Println("combined err :", err2)
	combinedAll3, errWp2 := cmdBuilder.CombinedOutputErrorWrapper()

	fmt.Println("combined all :", string(combinedAll3))
	fmt.Println("combined err :", errWp2.CompiledJsonErrorWithStackTraces())
}
