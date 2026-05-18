// Package osadapter wires the real `os/exec`-backed Runner for
// errcmdportable. This file is imported only by hosts that can actually
// spawn processes — keeping it in a subpackage means edge bundlers that
// drop unused imports won't pull `os/exec` into the Worker build.
//
// Status: deferred-research draft. Production callers should prefer the
// existing `errcmd` package for full-featured execution; this adapter
// exists so a single piece of business logic can target both edge and
// native runtimes via the `errcmdportable.Runner` interface.
package osadapter

import (
	"bytes"
	"os/exec"

	"github.com/alimtvnetwork/errorwrapper-v3/errcmdportable"
)

// New returns an os/exec-backed Runner.
func New() errcmdportable.Runner { return osRunner{} }

type osRunner struct{}

func (osRunner) Capability() errcmdportable.Capability {
	return errcmdportable.CapabilityOsExec
}

func (osRunner) Run(name string, args ...string) errcmdportable.Result {
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

	return errcmdportable.Result{
		ExitCode: exitCode,
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		Err:      err,
	}
}
