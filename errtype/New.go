package errtype

import "gitlab.com/evatix-go/core/coreinterface/errcoreinf"

func NewUsingTyper(basicErrTyper errcoreinf.BasicErrorTyper) Variation {
	errTypeVal := basicErrTyper.Value()

	return Variation(errTypeVal)
}
