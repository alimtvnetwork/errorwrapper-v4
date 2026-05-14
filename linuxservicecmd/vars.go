package linuxservicecmd

import "github.com/alimtvnetwork/errorwrapper-v3/errcmd"

var (
	hasService          = hasServiceCmdLookPath()
	hasSystemCtlService = hasSystemctlCmdLookPath()
	bashArgsCmdCreator  = errcmd.New.BashScript.ArgsDefault
)
