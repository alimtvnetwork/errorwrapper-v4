package errfloat64

type ResultWithApplicable2 struct {
	Result2
	IsApplicable bool
}

func (it *ResultWithApplicable2) IsAnyNull() bool {
	return it == nil
}
