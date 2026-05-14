package linuxservicecmd

import (
	"os/exec"

	"gitlab.com/evatix-go/core/cmdconsts"
)

func hasServiceCmdLookPath() bool {
	_, err := exec.LookPath(cmdconsts.Service)

	return err == nil
}
