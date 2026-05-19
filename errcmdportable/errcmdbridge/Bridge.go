// Package errcmdbridge converts a full-featured `errcmd.Result` into the
// portable `errcmdportable.Result` shape so callers that target both
// native and edge runtimes can share a single result handler.
//
// Kept in its own package to avoid pulling `errcmd` (and its os/exec
// transitive deps) into Worker bundles via errcmdportable.
package errcmdbridge

import (
	"github.com/alimtvnetwork/errorwrapper-v3/errcmd"
	"github.com/alimtvnetwork/errorwrapper-v3/errcmdportable"
)

// FromErrcmdResult downgrades an `*errcmd.Result` to the portable shape.
// Nil input returns a zero-value Result (no error, exit 0).
func FromErrcmdResult(r *errcmd.Result) errcmdportable.Result {
	if r == nil {
		return errcmdportable.Result{}
	}

	stdoutLine, _ := r.AllErrorBytes() // best-effort; ignore compiled wrapper here
	_ = stdoutLine                     // kept for symmetry; stdout/stderr below

	stdout := ""
	stderr := ""
	if base := r.CompiledTrimmedOutput(); base != "" {
		stdout = base
	}
	if base := r.CompiledTrimmedErrorOutput(); base != "" {
		stderr = base
	}

	var portableErr error
	if w := r.ErrorWrapper(); w != nil && w.HasError() {
		portableErr = w.Error()
	}

	return errcmdportable.Result{
		ExitCode: r.ExitCode,
		Stdout:   stdout,
		Stderr:   stderr,
		Err:      portableErr,
	}
}
