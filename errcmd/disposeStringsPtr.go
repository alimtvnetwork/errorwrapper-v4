package errcmd

func disposeStringsPtr(allBytes *[]string) {
	if allBytes == nil || *allBytes == nil {
		return
	}

	*allBytes = nil
}
