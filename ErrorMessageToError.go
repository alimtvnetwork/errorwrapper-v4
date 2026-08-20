package errorwrapper

import (
	"errors"

	"github.com/alimtvnetwork/errorwrapper-v4/internal/consts"
)

// ErrorMessageToError final error is nil if err is nil
func ErrorMessageToError(err error, msg string) error {
	if err == nil {
		return nil
	}

	finalErrMsg := err.Error() +
		consts.DefaultErrorLineSeparator +
		msg

	return errors.New(finalErrMsg)
}
