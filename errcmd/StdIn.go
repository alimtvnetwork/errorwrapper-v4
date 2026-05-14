package errcmd

import (
	"bytes"

	"gitlab.com/evatix-go/core/constants"
)

type StdIn struct {
	OutBuf, ErrBuf *bytes.Buffer
}

func NewStdIn(outBufSize, errBuffSize int) *StdIn {
	var outBuf, errBuf *bytes.Buffer
	if outBufSize > 0 {
		outBuf = &bytes.Buffer{}
		outBuf.Grow(outBufSize)
	}

	if errBuffSize > 0 {
		errBuf = &bytes.Buffer{}
		errBuf.Grow(errBuffSize)
	}

	return &StdIn{
		OutBuf: outBuf,
		ErrBuf: errBuf,
	}
}

func (it *StdIn) HasStdOut() bool {
	return it != nil && it.OutBuf != nil
}

func (it *StdIn) StdOutString() string {
	if it.HasStdOut() {
		return it.OutBuf.String()
	}

	return constants.EmptyString
}

func (it *StdIn) HasStdErr() bool {
	return it != nil && it.ErrBuf != nil
}

func (it *StdIn) StdErrString() string {
	if it.HasStdErr() {
		return it.ErrBuf.String()
	}

	return constants.EmptyString
}
