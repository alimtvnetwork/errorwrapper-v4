package errnew

import (
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/ref"
	"gitlab.com/evatix-go/errorwrapper/refs"
)

type newEnumToErrorWrapperCreator struct{}

func (it newEnumToErrorWrapperCreator) Default(
	errType errtype.Variation,
	basicEnumer enuminf.BasicEnumer,
) *errorwrapper.Wrapper {
	if basicEnumer.IsValid() {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		"enum : "+basicEnumer.RangeNamesCsv(),
		ref.Value{
			Variable: "current enum",
			Value:    basicEnumer.NameValue(),
		})
}

func (it newEnumToErrorWrapperCreator) Error(
	errType errtype.Variation,
	err error,
	basicEnumer enuminf.BasicEnumer,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		err.Error()+", enum : "+basicEnumer.RangeNamesCsv(),
		ref.Value{
			Variable: "Current Enum",
			Value:    basicEnumer.NameValue(),
		})
}

func (it newEnumToErrorWrapperCreator) Message(
	errType errtype.Variation,
	msg string,
	basicEnumer enuminf.BasicEnumer,
) *errorwrapper.Wrapper {
	if msg == "" {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		msg+", enum : "+basicEnumer.RangeNamesCsv(),
		ref.Value{
			Variable: "Current Enum",
			Value:    basicEnumer.NameValue(),
		})
}

func (it newEnumToErrorWrapperCreator) ErrorMessage(
	errType errtype.Variation,
	err error,
	msg string,
	basicEnumer enuminf.BasicEnumer,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		err.Error()+" "+msg+" , enum : "+basicEnumer.RangeNamesCsv(),
		ref.Value{
			Variable: "Current Enum",
			Value:    basicEnumer.NameValue(),
		})
}

func (it newEnumToErrorWrapperCreator) Many(
	isOnlyInvalidCollect bool,
	errType errtype.Variation,
	msg string,
	basicEnumers ...enuminf.BasicEnumer,
) *errorwrapper.Wrapper {
	if msg == "" {
		return nil
	}

	refCollection := refs.New(len(basicEnumers))

	for _, enumer := range basicEnumers {
		if isOnlyInvalidCollect && enumer.IsInvalid() {
			refCollection.Add(
				"value-enum-"+enumer.TypeName(),
				enumer.NameValue())
			refCollection.Add(
				"ranges-enum-"+enumer.TypeName(),
				enumer.RangeNamesCsv())

			continue
		}

		refCollection.Add(
			"value-enum-"+enumer.TypeName(),
			enumer.NameValue())
		refCollection.Add(
			"ranges-enum-"+enumer.TypeName(),
			enumer.RangeNamesCsv())
	}

	return errorwrapper.NewMsgDisplayError(
		defaultSkipInternal,
		errType,
		msg,
		refCollection)
}

func (it newEnumToErrorWrapperCreator) SingleErrorManyEnum(
	isOnlyInvalidCollect bool,
	errType errtype.Variation,
	err error,
	basicEnumers ...enuminf.BasicEnumer,
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return it.Many(
		isOnlyInvalidCollect,
		errType,
		err.Error(),
		basicEnumers...)
}
