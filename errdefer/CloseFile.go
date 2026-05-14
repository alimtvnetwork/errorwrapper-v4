package errdefer

import (
	"os"

	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

func CloseFile(
	location string,
	existingErrorWrapper *errorwrapper.Wrapper, // could be nil
	osFile *os.File,
) *errorwrapper.Wrapper {
	if osFile == nil {
		return existingErrorWrapper
	}

	closerErr := osFile.Close()
	closingErrorWrapper := errnew.Path.Error(
		errtype.FileClosing,
		closerErr,
		location)

	return mergeErrorWrapper(
		existingErrorWrapper,
		closingErrorWrapper)
}
