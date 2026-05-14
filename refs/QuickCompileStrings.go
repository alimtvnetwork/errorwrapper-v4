package refs

import "github.com/alimtvnetwork/core-v9/coredata/stringslice"

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
