package errtype

import "github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"

func NewUsingTyper(basicErrTyper errcoreinf.BasicErrorTyper) Variation {
	errTypeVal := basicErrTyper.Value()

	return Variation(errTypeVal)
}
