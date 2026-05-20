package errfloat

type Result2 struct {
	Result
	Value2 float64
}

func (it *Result2) IsAnyNull() bool {
	return it == nil
}
