package errdefer

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
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
