package linuxservicecmd

import "gitlab.com/evatix-go/errorwrapper/errcmd"

var (
	hasService          = hasServiceCmdLookPath()
	hasSystemCtlService = hasSystemctlCmdLookPath()
	bashArgsCmdCreator  = errcmd.New.BashScript.ArgsDefault
)
