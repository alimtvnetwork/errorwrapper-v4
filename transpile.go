package errorwrapper

import "errors"

func transpileWrapperToModel(from *Wrapper, to *WrapperDataModel) WrapperDataModel {
	to.IsDisplayableError = from.isDisplayableError
	to.CurrentError = from.ErrorString()
	to.References = from.references
	to.ErrorType = from.errorType
	to.HasError = from.HasError()
	to.StackTraces = from.stackTraces

	return *to
}

func transpileModelToWrapper(from *WrapperDataModel, toWrapper *Wrapper) *Wrapper {
	var err error
	if from.HasError {
		err = errors.New(from.CurrentError)
	}

	toWrapper.isDisplayableError = from.IsDisplayableError
	toWrapper.currentError = err
	toWrapper.references = from.References
	toWrapper.errorType = from.ErrorType
	toWrapper.stackTraces = from.StackTraces

	return toWrapper
}
