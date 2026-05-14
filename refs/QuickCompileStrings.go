package refs

import "gitlab.com/evatix-go/core/coredata/stringslice"

func QuickCompileStrings(
	quickReferences ...QuickReference,
) (lines []string) {
	if len(quickReferences) == 0 {
		return []string{}
	}

	slice := stringslice.MakeLen(len(quickReferences))

	for i, quickRef := range quickReferences {
		slice[i] = quickRef.CompileLine()
	}

	return slice
}
