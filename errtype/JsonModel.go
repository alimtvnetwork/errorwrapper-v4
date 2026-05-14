package errtype

type JsonModel struct {
	Category string
	Id       uint16
}

func (it JsonModel) IsTypeOf(errType Variation) bool {
	return it.Category == errType.Name()
}

func (it JsonModel) IsCategory(name string) bool {
	return it.Category == name
}

func (it JsonModel) Name() string {
	return it.Category
}

func (it JsonModel) Value() uint16 {
	return it.Id
}

func (it JsonModel) IsNoError() bool {
	return it.Id == NoError.ValueUInt16()
}

func (it JsonModel) HasError() bool {
	return it.Id > NoError.ValueUInt16()
}

func (it JsonModel) Type() Variation {
	return Variation(it.Id)
}

func (it JsonModel) String() string {
	return it.Category
}
