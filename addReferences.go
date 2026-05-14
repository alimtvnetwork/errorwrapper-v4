package errorwrapper

import (
	"gitlab.com/evatix-go/errorwrapper/ref"
	"gitlab.com/evatix-go/errorwrapper/refs"
)

func addReferences(
	references []ref.Value,
	clonedNew *Wrapper,
) *Wrapper {
	refLength := len(references)
	if refLength == 0 {
		return clonedNew
	}

	if clonedNew.references == nil {
		clonedNew.references = refs.New(refLength)
	}

	if clonedNew.references != nil {
		clonedNew.references.Adds(references...)
	}

	return clonedNew
}
