package ref

type ValueDataModel struct {
	VariableName string
	ValueString  string
}

func NewDataModel(value *Value) ValueDataModel {
	if value == nil {
		return ValueDataModel{}
	}

	return ValueDataModel{
		VariableName: value.Variable,
		ValueString:  value.ValueString(),
	}
}
