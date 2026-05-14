package errcmd

import "github.com/alimtvnetwork/errorwrapper-v3"

type BaseCmdWrapper struct {
	baseBufferStdOutError  *baseBufferStdOutError
	initializeErrorWrapper *errorwrapper.Wrapper
	process                string
	wholeCommand           string // Process + Space + ArgumentsCompiledWithSpace
	argumentsSingleLine    string
	arguments              []string
}

func (it *BaseCmdWrapper) Dispose() {
	if it == nil {
		return
	}

	it.DisposeWithoutErrorWrapper()
	it.initializeErrorWrapper.Dispose()
}

func (it *BaseCmdWrapper) DisposeWithoutErrorWrapper() {
	if it == nil {
		return
	}

	it.baseBufferStdOutError.Dispose()
	disposeStringsPtr(&it.arguments)
	it.process = ""
	it.argumentsSingleLine = ""
	it.wholeCommand = ""
}
