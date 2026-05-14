package errcmd

func Conditional(
	isTrue bool,
	trueCmdOne, falseCmdOnce *CmdOnce,
) *CmdOnce {
	if isTrue {
		return trueCmdOne
	}

	return falseCmdOnce
}
