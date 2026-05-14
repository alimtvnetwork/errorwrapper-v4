package creationtests

import (
	"errors"

	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

var (
	nilErr          error = nil
	passwordCrudErr       = errnew.
			Type.
			Error(errtype.PasswordCrud, errors.New("some password"))
)
