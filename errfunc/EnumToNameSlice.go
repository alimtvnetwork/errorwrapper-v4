package errfunc

import (
	"github.com/alimtvnetwork/core-v9/coredata/stringslice"
	"github.com/alimtvnetwork/core-v9/coreinterface"
)

func EnumToNameSlice(
	enums ...coreinterface.ToNamer,
) []string {
	if enums == nil {
		return []string{}
	}

	slice := stringslice.Make(len(enums), len(enums))
	for i, enum := range enums {
		slice[i] = enum.Name()
	}

	return slice
}
