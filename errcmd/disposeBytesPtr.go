package errcmd

func disposeBytesPtr(allBytes *[]byte) {
	if allBytes == nil || *allBytes == nil {
		return
	}

	*allBytes = nil
}
