package errfloat

type ResultWithApplicable struct {
	Result
	IsApplicable bool
}

func (it *ResultWithApplicable) IsAnyNull() bool {
	return it == nil
}
