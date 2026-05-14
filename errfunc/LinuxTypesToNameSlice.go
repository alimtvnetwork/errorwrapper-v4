package errfunc

import (
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/enum/linuxtype"
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
