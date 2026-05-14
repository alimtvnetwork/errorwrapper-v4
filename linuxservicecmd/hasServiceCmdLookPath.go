package linuxservicecmd

import (
	"os/exec"

	"github.com/alimtvnetwork/core-v9/cmdconsts"
)

func hasServiceCmdLookPath() bool {
	_, err := exec.LookPath(cmdconsts.Service)

	return err == nil
}
