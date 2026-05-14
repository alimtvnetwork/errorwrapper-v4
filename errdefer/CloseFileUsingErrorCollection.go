package errdefer

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

func CloseFileUsingErrorCollection(
	location string,
	errorCollection *errwrappers.Collection,
	osFile *os.File,
) (isSuccess bool) {
	if osFile == nil {
		return
	}

	closerErr := osFile.Close()
	closingErrorWrapper := errnew.Path.Error(
		errtype.FileClosing,
		closerErr,
		location)

	errorCollection.AddWrapperPtr(closingErrorWrapper)

	return closingErrorWrapper.IsEmpty()
}
