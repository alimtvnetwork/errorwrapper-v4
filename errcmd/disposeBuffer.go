package errcmd

import "bytes"

func disposeBuffer(buff *bytes.Buffer) {
	if buff == nil {
		return
	}

	buff.Reset()
	*buff = bytes.Buffer{}
}
