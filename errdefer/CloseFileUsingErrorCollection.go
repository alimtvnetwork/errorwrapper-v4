package errdefer

import (
	"os"

	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/errwrappers"
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
