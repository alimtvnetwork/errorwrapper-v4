package errdefer

import (
	"os"

	"github.com/alimtvnetwork/errorwrapper-v3"
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
