package errtype

func (it Variation) Is(variant Variation) bool {
	return it == variant
}

func (it Variation) IsNoError() bool {
	return it == NoError
}

func (it Variation) HasError() bool {
	return it != NoError
}
