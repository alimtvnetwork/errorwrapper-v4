package errany

type Result2 struct {
	Result
	Value2 interface{}
}

func (it *Result2) IsAnyNull() bool {
	return it == nil
}
