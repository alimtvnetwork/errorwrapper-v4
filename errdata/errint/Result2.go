package errint

type Result2 struct {
	Result
	Value2 int
}

func (it *Result2) IsAnyNull() bool {
	return it == nil
}
