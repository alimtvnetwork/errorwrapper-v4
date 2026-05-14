package errcmd

import "bytes"

func BufferClone(buffer *bytes.Buffer) *bytes.Buffer {
	if buffer == nil {
		return buffer
	}

	if buffer.Cap() > 0 {
		newBuff := &bytes.Buffer{}
		newBuff.Grow(buffer.Cap())

		newBuff.Write(buffer.Bytes())

		return newBuff
	}

	return &bytes.Buffer{}
}
