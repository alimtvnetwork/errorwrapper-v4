package errfunc

import (
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/coreinterface"
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
