package creationtests

import (
	"errors"

	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
)

var (
	nilErr          error = nil
	passwordCrudErr       = errnew.
			Type.
			Error(errtype.PasswordCrud, errors.New("some password"))
)
