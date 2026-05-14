package errstr

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

type newHashsetCreator struct{}

func (it *newHashsetCreator) Empty() *Hashset {
	return &Hashset{}
}

func (it *newHashsetCreator) Strings(
	items []string,
) *Hashset {
	return &Hashset{
		Hashset: corestr.New.Hashset.Strings(items),
	}
}

func (it *newHashsetCreator) Error(
	errType errtype.Variation,
	err error,
) *Hashset {
	return &Hashset{
		ErrorWrapper: errnew.Type.ErrorUsingStackSkip(
			codestack.Skip1,
			errType,
			err),
	}
}

func (it *newHashsetCreator) ErrorWrapper(
	errorWrapper *errorwrapper.Wrapper,
) *Hashset {
	return &Hashset{
		ErrorWrapper: errorWrapper,
	}
}

func (it *newHashsetCreator) Create(
	hashset *corestr.Hashset,
	errorWrapper *errorwrapper.Wrapper,
) *Hashset {
	return &Hashset{
		Hashset:      hashset,
		ErrorWrapper: errorWrapper,
	}
}

func (it *newHashsetCreator) StringsErrorWrapper(
	items []string,
	errorWrapper *errorwrapper.Wrapper,
) *Hashset {
	return &Hashset{
		Hashset:      corestr.New.Hashset.Strings(items),
		ErrorWrapper: errorWrapper,
	}
}

func (it *newHashsetCreator) ValueOnly(
	items []string,
) *Hashset {
	return &Hashset{
		Hashset: corestr.New.Hashset.Strings(items),
	}
}

func (it *newHashsetCreator) Cap(capacity int) *Hashset {
	return &Hashset{
		Hashset: corestr.New.Hashset.Cap(capacity),
	}
}

func (it *newHashsetCreator) ErrorWrapperWithSpreadStrings(
	wrapper *errorwrapper.Wrapper,
	items ...string,
) *Hashset {
	return &Hashset{
		Hashset: corestr.New.Hashset.Strings(
			items),
		ErrorWrapper: wrapper,
	}
}

// UsingMap having addCapacity <= 0 , keeps the same pointer for the hashset.
func (it *newHashsetCreator) UsingMap(
	addCapacity int,
	items map[string]bool,
) *Hashset {
	return &Hashset{
		Hashset: corestr.New.Hashset.UsingMapOption(
			addCapacity,
			false,
			items,
		),
	}
}

// UsingCollection having addCapacity <= 0 , keeps the same pointer for the hashset.
func (it *newHashsetCreator) UsingCollection(
	collection *corestr.Collection,
) *Hashset {
	return &Hashset{
		Hashset: corestr.New.Hashset.UsingCollection(
			collection),
	}
}
