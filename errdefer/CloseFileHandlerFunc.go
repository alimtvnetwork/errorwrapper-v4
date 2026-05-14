package errdefer

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
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
