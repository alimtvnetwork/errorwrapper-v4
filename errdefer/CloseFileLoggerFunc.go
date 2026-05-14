package errdefer

import (
	"os"

	"github.com/alimtvnetwork/errorwrapper-v3"
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
