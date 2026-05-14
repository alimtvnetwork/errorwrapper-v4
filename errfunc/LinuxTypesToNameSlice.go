package errfunc

import (
	"github.com/alimtvnetwork/core-v9/coredata/stringslice"
	"github.com/alimtvnetwork/enum-v10/linuxtype"
)

func LinuxTypesToNameSlice(
	linuxTypes ...linuxtype.Variant,
) []string {
	if linuxTypes == nil {
		return []string{}
	}

	slice := stringslice.Make(
		len(linuxTypes),
		len(linuxTypes))
	for i, linuxType := range linuxTypes {
		slice[i] = linuxType.Name()
	}

	return slice
}
