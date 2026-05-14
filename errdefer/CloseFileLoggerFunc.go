package errdefer

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func CloseFileLoggerFunc(
	location string,
	existingErrorWrapper *errorwrapper.Wrapper,
	osFile *os.File,
	loggerFunc func(errorWrapper *errorwrapper.Wrapper),
) {
	finalError := CloseFile(
		location,
		existingErrorWrapper,
		osFile)

	loggerFunc(finalError)
}
