package errstr

import (
	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
)

type newHashmapCreator struct{}

func (it *newHashmapCreator) Empty() *Hashmap {
	return &Hashmap{}
}

func (it *newHashmapCreator) Error(
	errType errtype.Variation,
	err error,
) *Hashmap {
	return &Hashmap{
		ErrorWrapper: errnew.Type.ErrorUsingStackSkip(
			codestack.Skip1,
			errType,
			err),
	}
}

func (it *newHashmapCreator) ErrorWrapper(
	errorWrapper *errorwrapper.Wrapper,
) *Hashmap {
	return &Hashmap{
		ErrorWrapper: errorWrapper,
	}
}

func (it *newHashmapCreator) Create(
	hashmap *corestr.Hashmap,
	errorWrapper *errorwrapper.Wrapper,
) *Hashmap {
	return &Hashmap{
		Hashmap:      hashmap,
		ErrorWrapper: errorWrapper,
	}
}

func (it *newHashmapCreator) Cap(capacity int) *Hashmap {
	return &Hashmap{
		Hashmap: corestr.New.Hashmap.Cap(capacity),
	}
}

// UsingMap having addCapacity <= 0 , keeps the same pointer for the hashset.
func (it *newHashmapCreator) UsingMap(
	addCapacity int,
	items map[string]string,
) *Hashmap {
	return &Hashmap{
		Hashmap: corestr.New.Hashmap.UsingMapOptions(
			false,
			addCapacity,
			items),
		ErrorWrapper: nil,
	}
}

func (it *newHashmapCreator) UsingMapNoClone(
	items map[string]string,
) *Hashmap {
	return &Hashmap{
		Hashmap: corestr.New.Hashmap.UsingMap(
			items),
	}
}

func (it *newHashmapCreator) KeyValuesSeparateCollection(
	keys, values *corestr.Collection,
) *Hashmap {
	return &Hashmap{
		Hashmap: corestr.New.Hashmap.KeyValuesCollection(
			keys, values),
	}
}

// KeyValuesSeparateWithErrorWrapper wrapper nil will point to empty error wrapper
func (it *newHashmapCreator) KeyValuesSeparateWithErrorWrapper(
	keys, values []string,
	wrapper *errorwrapper.Wrapper,
) *Hashmap {
	return &Hashmap{
		Hashmap: corestr.New.Hashmap.KeyValuesStrings(
			keys, values),
		ErrorWrapper: wrapper,
	}
}
