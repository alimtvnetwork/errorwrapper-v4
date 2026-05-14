package errnew

import (
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
	"github.com/alimtvnetwork/errorwrapper-v3/ref"
	"github.com/alimtvnetwork/errorwrapper-v3/refs"
)

type newMessageWithRefCreator struct{}

func (it newMessageWithRefCreator) References(
	variation errtype.Variation,
	message string,
	references ...interface{},
) *errorwrapper.Wrapper {
	if message == "" && len(references) == 0 {
		return nil
	}

	if len(references) == 0 {
		return errorwrapper.NewMsgDisplayErrorNoReference(
			defaultSkipInternal,
			variation,
			message,
		)
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		variation,
		message,
		ref.Value{
			Variable: "Length",
			Value:    len(references),
		},
		ref.Value{
			Variable: "Reference(s)",
			Value:    references,
		},
	)
}

func (it newMessageWithRefCreator) Default(
	variation errtype.Variation,
	message string,
	reference interface{},
) *errorwrapper.Wrapper {
	if message == "" {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		variation,
		message,
		ref.Value{
			Variable: "Reference",
			Value:    reference,
		})
}

func (it newMessageWithRefCreator) Error(
	variation errtype.Variation,
	err error,
	reference interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		variation,
		err.Error(),
		ref.Value{
			Variable: "Reference",
			Value:    reference,
		})
}

func (it newMessageWithRefCreator) ErrorVarName(
	variation errtype.Variation,
	err error,
	variableName string,
	reference interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		variation,
		err.Error(),
		ref.Value{
			Variable: variableName,
			Value:    reference,
		})
}

func (it newMessageWithRefCreator) ErrorRefs(
	variation errtype.Variation,
	err error,
	references ...interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		variation,
		err.Error(),
		ref.Value{
			Variable: "Reference(s)",
			Value:    references,
		})
}

func (it newMessageWithRefCreator) DefaultVarName(
	variation errtype.Variation,
	message string,
	variableName string,
	valueReference interface{},
) *errorwrapper.Wrapper {
	if message == "" {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		variation,
		message,
		ref.Value{
			Variable: variableName,
			Value:    valueReference,
		})
}

func (it newMessageWithRefCreator) UsingMap(
	variation errtype.Variation,
	message string,
	itemsMap map[string]interface{},
) *errorwrapper.Wrapper {
	if message == "" {
		return nil
	}

	references := refs.NewUsingMap(itemsMap)

	return errorwrapper.NewMsgUsingAllParams(
		defaultSkipInternal,
		variation,
		true,
		message,
		references)
}
