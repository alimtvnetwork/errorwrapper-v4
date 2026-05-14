package errbool

type ResultWithApplicable2 struct {
	Result2
	IsApplicable bool
}

func (resultWithApplicable2 *ResultWithApplicable2) IsAllTrue() bool {
	return resultWithApplicable2.Value &&
		resultWithApplicable2.Value2
}

func (resultWithApplicable2 *ResultWithApplicable2) IsAllTruePlusApplicable() bool {
	return resultWithApplicable2.Value &&
		resultWithApplicable2.Value2 &&
		resultWithApplicable2.IsApplicable
}

func (resultWithApplicable2 *ResultWithApplicable2) IsAllTruePlusApplicableNoError() bool {
	return resultWithApplicable2.Value &&
		resultWithApplicable2.Value2 &&
		resultWithApplicable2.IsApplicable &&
		resultWithApplicable2.ErrorWrapper.IsEmpty()
}

func (resultWithApplicable2 *ResultWithApplicable2) IsAnyFalse() bool {
	return !resultWithApplicable2.Value || !resultWithApplicable2.Value2
}
