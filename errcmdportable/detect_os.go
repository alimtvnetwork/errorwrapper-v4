//go:build !js && !wasip1

package errcmdportable

import (
	"bytes"
	"os/exec"
)

// Detect returns an os/exec-backed Runner for native OS targets.
// For WASM / WASI edge builds the default file (detect_default.go)
// is compiled instead and returns NoProcessRunner.
//
// Implementation note: the os/exec-backed runner is inlined here
// (rather than delegating to the osadapter subpackage) to avoid an
// import cycle — osadapter depends on errcmdportable for the Runner
// interface and Result type. The build tag above ensures js / wasip1
// builds still drop os/exec from the bundle.
func Detect() Runner {
	return autoOsRunner{}
}

type autoOsRunner struct{}

func (autoOsRunner) Capability() Capability { return CapabilityOsExec }

func (autoOsRunner) Run(name string, args ...string) Result {
	cmd := exec.Command(name, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	exitCode := 0
	if exitErr, ok := err.(*exec.ExitError); ok {
		exitCode = exitErr.ExitCode()
	} else if err != nil {
		exitCode = -1
	}

	return Result{
		ExitCode: exitCode,
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		Err:      err,
	}
}
