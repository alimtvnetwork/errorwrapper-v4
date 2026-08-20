package errdefer

import (
	"os"

	"github.com/alimtvnetwork/errorwrapper-v4"
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
