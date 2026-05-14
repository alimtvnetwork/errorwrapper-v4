package errcmd

func ClearDispose(cmdOnceRan *CmdOnce) {
	cmdOnceRan.Dispose()
	cmdOnceRan = nil
}
