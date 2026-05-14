package linuxservicecmd

import "os/exec"

// IsNamePathExist
//
// Refers to environment, binary, or
// look path for the given name.
func IsNamePathExist(name string) bool {
	_, err := exec.LookPath(name)

	return err == nil
}
