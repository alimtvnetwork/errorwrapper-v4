//go:build js || wasip1

package errcmdportable

// Detect returns NoProcessRunner for WASM / WASI edge targets.
func Detect() Runner {
	return NoProcessRunner{}
}
