//go:build !js && !wasip1

package errcmdportable

import "github.com/alimtvnetwork/errorwrapper-v3/errcmdportable/osadapter"

// Detect returns an os/exec-backed Runner for native OS targets.
// For WASM / WASI edge builds the default file (detect_default.go)
// is compiled instead and returns NoProcessRunner.
func Detect() Runner {
	return osadapter.New()
}
