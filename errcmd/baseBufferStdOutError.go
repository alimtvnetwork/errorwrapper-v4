package errcmd

import (
	"bytes"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coredata/stringslice"
)

type baseBufferStdOutError struct {
	stdErr, stdOut                                    *bytes.Buffer
	compiledOutputBytes                               []byte
	compiledErrorBytes                                []byte
	compiledOutputString                              corestr.SimpleStringOnce
	compiledErrorString                               corestr.SimpleStringOnce
	compiledOutputLines                               []string
	compiledErrorLines                                []string
	combinedBothErrorOutputLines                      []string
	compiledTrimmedOutputLines                        []string
	compiledTrimmedErrorLines                         []string
	compiledTrimmedOutput, compiledTrimmedErrorOutput corestr.SimpleStringOnce
}

func (it *baseBufferStdOutError) Dispose() {
	if it == nil {
		return
	}

	disposeBuffer(it.stdErr)
	disposeBuffer(it.stdOut)
	disposeBytesPtr(&it.compiledOutputBytes)
	disposeBytesPtr(&it.compiledErrorBytes)
	disposeStringsPtr(&it.compiledOutputLines)
	disposeStringsPtr(&it.compiledErrorLines)
	disposeStringsPtr(&it.combinedBothErrorOutputLines)
	disposeStringsPtr(&it.compiledTrimmedOutputLines)
	disposeStringsPtr(&it.compiledTrimmedErrorLines)
	it.compiledOutputString.Dispose()
	it.compiledErrorString.Dispose()
	it.compiledTrimmedOutput.Dispose()
	it.compiledTrimmedErrorOutput.Dispose()
}

func (it *baseBufferStdOutError) CompiledTrimmedOutput() string {
	if it.compiledTrimmedOutput.IsInitialized() {
		return it.compiledTrimmedOutput.String()
	}

	trimmed := strings.Join(
		it.CompiledTrimmedOutputLines(),
		constants.NewLineUnix)

	return it.
		compiledTrimmedOutput.
		GetPlusSetOnUninitialized(trimmed)
}

func (it *baseBufferStdOutError) CompiledTrimmedErrorOutput() string {
	if it.compiledTrimmedErrorOutput.IsInitialized() {
		return it.compiledTrimmedErrorOutput.String()
	}

	trimmed := strings.Join(
		it.CompiledErrorLines(),
		constants.NewLineUnix)

	return it.
		compiledTrimmedErrorOutput.
		GetPlusSetOnUninitialized(trimmed)
}

func (it *baseBufferStdOutError) CombinedBothErrorOutputLines() []string {
	if it.combinedBothErrorOutputLines != nil {
		return it.combinedBothErrorOutputLines
	}

	it.combinedBothErrorOutputLines = *stringslice.MergeNew(
		it.CompiledErrorLines(),
		it.CompiledOutputLines()...)

	return it.combinedBothErrorOutputLines
}

func (it *baseBufferStdOutError) CompiledErrorLines() []string {
	if it.compiledErrorLines != nil {
		return it.compiledErrorLines
	}

	stringLines := strings.Split(
		it.ErrorString(),
		constants.NewLineUnix)
	it.compiledErrorLines = stringLines

	return stringLines
}

func (it *baseBufferStdOutError) CompiledTrimmedErrorLines() []string {
	if it.compiledTrimmedErrorLines != nil {
		return it.compiledTrimmedErrorLines
	}

	stringLines := it.CompiledErrorLines()
	filteredLines := it.trimmedLines(stringLines)
	it.compiledTrimmedErrorLines = filteredLines

	return filteredLines
}

func (it *baseBufferStdOutError) CompiledOutputLines() []string {
	if it.compiledOutputLines != nil {
		return it.compiledOutputLines
	}

	stringLines := strings.Split(
		it.OutputString(),
		constants.NewLineUnix)
	it.compiledOutputLines = stringLines

	return stringLines
}

func (it *baseBufferStdOutError) CompiledTrimmedOutputLines() []string {
	if it.compiledTrimmedOutputLines != nil {
		return it.compiledTrimmedOutputLines
	}

	stringLines := it.CompiledOutputLines()
	filteredLines := it.trimmedLines(stringLines)
	it.compiledTrimmedOutputLines = filteredLines

	return filteredLines
}

func (it *baseBufferStdOutError) trimmedLines(stringLines []string) []string {
	filteredLines := make([]string, constants.Zero, len(stringLines))
	for _, stringLine := range stringLines {
		trimmed := strings.TrimSpace(stringLine)

		if len(trimmed) > 0 {
			filteredLines = append(filteredLines, trimmed)
		}
	}

	return filteredLines
}

func (it *baseBufferStdOutError) HasOutputBuffer() bool {
	return it != nil && it.stdOut != nil
}

func (it *baseBufferStdOutError) HasErrorBuffer() bool {
	return it != nil && it.stdErr != nil
}

func (it *baseBufferStdOutError) HasOutputBufferData() bool {
	return it.HasOutputBuffer() && it.stdOut.Len() > 0
}

func (it *baseBufferStdOutError) HasErrorBufferData() bool {
	return it.HasErrorBuffer() && it.stdErr.Len() > 0
}

// IsEmptyOutput
//
// Indicates output buffer has NO data.
func (it *baseBufferStdOutError) IsEmptyOutput() bool {
	return it == nil || it.stdOut == nil || it.stdOut.Len() == 0
}

// HasOutputContent
//
// Indicates output buffer has data.
// alias for HasOutputBufferData
func (it *baseBufferStdOutError) HasOutputContent() bool {
	return it.HasErrorBufferData()
}

// HasErrorContent
//
// Indicates error buffer has data.
// alias for HasErrorBufferData
func (it *baseBufferStdOutError) HasErrorContent() bool {
	return it.HasErrorBufferData()
}

// IsEmptyErrorBuffer
//
// Indicates error buffer has NO data.
func (it *baseBufferStdOutError) IsEmptyErrorBuffer() bool {
	return it == nil || it.stdErr == nil || it.stdErr.Len() == 0
}

// SafeBytes
//
// Alias for OutputBytes
func (it *baseBufferStdOutError) SafeBytes() []byte {
	return it.OutputBytes()
}

// SafeErrorBytes
//
// Alias for ErrorBytes
func (it *baseBufferStdOutError) SafeErrorBytes() []byte {
	return it.ErrorBytes()
}

func (it *baseBufferStdOutError) OutputBytes() []byte {
	if it.compiledOutputBytes != nil {
		return it.compiledOutputBytes
	}

	if it.stdOut == nil {
		it.compiledOutputBytes = []byte{}

		return it.compiledOutputBytes
	}

	allBytes := it.stdOut.Bytes()
	it.compiledOutputBytes = allBytes

	return it.compiledOutputBytes
}

func (it *baseBufferStdOutError) ErrorBytes() []byte {
	if it.compiledErrorBytes != nil {
		return it.compiledErrorBytes
	}

	if it.stdErr == nil {
		it.compiledErrorBytes = []byte{}

		return it.compiledErrorBytes
	}

	allBytes := it.stdErr.Bytes()
	it.compiledErrorBytes = allBytes

	return it.compiledErrorBytes
}

func (it *baseBufferStdOutError) OutputString() string {
	if it.compiledOutputString.IsInitialized() {
		return it.compiledOutputString.String()
	}

	if it.HasOutputBufferData() {
		allBytes := it.OutputBytes()
		toString := string(allBytes)

		return it.
			compiledOutputString.
			GetPlusSetOnUninitialized(toString)
	}

	// no output
	return it.
		compiledOutputString.
		GetPlusSetEmptyOnUninitialized()
}

func (it *baseBufferStdOutError) ErrorString() string {
	if it.compiledErrorString.IsInitialized() {
		return it.compiledErrorString.String()
	}

	if it.HasErrorBufferData() {
		allBytes := it.ErrorBytes()
		toString := string(allBytes)

		return it.
			compiledErrorString.
			GetPlusSetOnUninitialized(toString)
	}

	// no error
	return it.
		compiledErrorString.
		GetPlusSetEmptyOnUninitialized()
}
