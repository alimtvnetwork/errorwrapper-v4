package creationtests

import (
	"errors"

	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

var (
	nilErr          error = nil
	passwordCrudErr       = errnew.
			Type.
			Error(errtype.PasswordCrud, errors.New("some password"))
)
