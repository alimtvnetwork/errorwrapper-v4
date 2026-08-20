package errdefer

import (
	"os"

	"github.com/alimtvnetwork/errorwrapper-v4"
)

func CloseFileHandlerFunc(
	location string,
	existingErrorWrapper *errorwrapper.Wrapper,
	osFile *os.File,
	handlerFunc func(errorWrapper *errorwrapper.Wrapper),
) {
	finalError := CloseFile(
		location,
		existingErrorWrapper,
		osFile)

	handlerFunc(finalError)
}
