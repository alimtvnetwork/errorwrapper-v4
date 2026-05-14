package errdefer

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
)

func CloseFileHandlePanic(
	location string,
	existingErrorWrapper *errorwrapper.Wrapper,
	osFile *os.File,
) {
	finalError := CloseFile(
		location,
		existingErrorWrapper,
		osFile)

	finalError.HandleError()
}
