package linuxservicecmd

import (
	"os/exec"

	"github.com/alimtvnetwork/core-v9/cmdconsts"
)

func hasSystemctlCmdLookPath() bool {
	_, err := exec.LookPath(cmdconsts.SystemCtl)

	return err == nil
}
