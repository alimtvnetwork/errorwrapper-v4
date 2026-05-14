package errbool

type ResultWithApplicable struct {
	Result
	IsApplicable bool
}

func (resultWithApplicable *ResultWithApplicable) IsAllTruePlusApplicableNoError() bool {
	return resultWithApplicable.IsApplicable &&
		resultWithApplicable.Value &&
		resultWithApplicable.ErrorWrapper.IsEmpty()
}
