package errorwrapper

func NewFromDataModel(
	model *WrapperDataModel,
) *Wrapper {
	if model == nil {
		return EmptyPtr()
	}

	wrapper := &Wrapper{}

	return transpileModelToWrapper(
		model,
		wrapper)
}
