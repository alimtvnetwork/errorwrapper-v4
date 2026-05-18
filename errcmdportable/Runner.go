// Package errcmdportable provides a runtime-portable façade over errcmd
// for environments that cannot spawn OS processes (Cloudflare Workers,
// browser/WASM, sandboxed serverless runtimes, etc.).
//
// The façade does not import errcmd directly so it can be compiled and
// embedded into edge targets without pulling `os/exec` transitively.
// Callers should use Detect() which auto-wires the osadapter on native
// OS builds and falls back to NoProcessRunner on edge targets.
package errcmdportable

import (
	"errors"
	"runtime"
)

// Capability reports whether the current runtime can spawn OS processes.
type Capability int

const (
	// CapabilityUnknown means the runtime was not probed yet.
	CapabilityUnknown Capability = iota
	// CapabilityOsExec means real `os/exec`-backed execution is available.
	CapabilityOsExec
	// CapabilityNoProcess means process spawning is unavailable
	// (Workers, browser WASM, sandboxed FaaS). Callers should treat
	// commands as a typed `ErrNotSupported`.
	CapabilityNoProcess
)

// ErrNotSupported is returned by Runner.Run on runtimes that cannot
// spawn OS processes. It is `errors.Is`-comparable.
var ErrNotSupported = errors.New("errcmdportable: command execution not supported on this runtime")

// Result is the minimal portable surface mirroring the subset of
// `errcmd.Result` that edge-safe callers need. Intentionally narrow:
// no `*exec.Cmd`, no `*os.Process`, no buffer pointers.
type Result struct {
	ExitCode int
	Stdout   string
	Stderr   string
	Err      error
}

// HasError reports whether the result carries a non-nil error.
func (r Result) HasError() bool { return r.Err != nil }

// IsNotSupported reports whether the failure is the portable
// no-process sentinel (vs. a real command failure).
func (r Result) IsNotSupported() bool {
	return r.Err != nil && errors.Is(r.Err, ErrNotSupported)
}

// Runner is the portable execution interface. Production code wires an
// `os/exec`-backed implementation; edge builds wire NoProcessRunner.
type Runner interface {
	Capability() Capability
	Run(name string, args ...string) Result
}

// NoProcessRunner is the safe default for Worker / WASM / sandbox
// targets. Every Run() returns ErrNotSupported with ExitCode = -1.
type NoProcessRunner struct{}

func (NoProcessRunner) Capability() Capability { return CapabilityNoProcess }

func (NoProcessRunner) Run(name string, args ...string) Result {
	_ = name
	_ = args
	return Result{
		ExitCode: -1,
		Err:      ErrNotSupported,
	}
}

// Detect returns the conservative best-guess Runner for the current
// build target. Cloudflare Workers / js+wasm / wasip1 → NoProcessRunner.
// Native OS targets fall back to NoProcessRunner here too; consumers
// that want real execution must explicitly wire the osadapter Runner
// (kept out of this package so this file stays import-clean for edge
// bundlers that refuse to resolve `os/exec`).
func Detect() Runner {
	switch runtime.GOOS {
	case "js", "wasip1":
		return NoProcessRunner{}
	}
	// Default safe: callers opt-in to OS execution explicitly.
	return NoProcessRunner{}
}
