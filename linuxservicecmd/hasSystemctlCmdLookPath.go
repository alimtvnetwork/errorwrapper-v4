package linuxservicecmd

import (
	"os/exec"

	"gitlab.com/evatix-go/core/cmdconsts"
)

func hasSystemctlCmdLookPath() bool {
	_, err := exec.LookPath(cmdconsts.SystemCtl)

	return err == nil
}
