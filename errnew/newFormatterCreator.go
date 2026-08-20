package errnew

import (
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core-v9/converters"
	"github.com/alimtvnetwork/core-v9/isany"
	"github.com/alimtvnetwork/core-v9/simplewrap"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
)

type newFormatterCreator struct{}

func (it newFormatterCreator) If(
	isCreate bool,
	errType errtype.Variation,
	format string,
	v ...interface{},
) *errorwrapper.Wrapper {
	if !isCreate {
		return nil
	}

	if !isCreate || format == "" && len(v) == 0 {
		return nil
	}

	compiled := fmt.Sprintf(
		format,
		v...)

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		compiled,
	)
}

func (it newFormatterCreator) MessageRefs(
	errType errtype.Variation,
	message string,
	references ...interface{},
) *errorwrapper.Wrapper {
	if message == "" && len(references) == 0 {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		message,
		ref.Value{
			Variable: "Reference(s)",
			Value:    references,
		})
}

func (it newFormatterCreator) ErrorRefs(
	errType errtype.Variation,
	err error,
	references ...interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		err.Error(),
		ref.Value{
			Variable: "Reference(s)",
			Value:    references,
		})
}

func (it newFormatterCreator) MessageRef(
	errType errtype.Variation,
	message string,
	reference interface{},
) *errorwrapper.Wrapper {
	if message == "" && isany.Null(reference) {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		defaultSkipInternal,
		errType,
		message,
		ref.Value{
			Variable: "Reference",
			Value:    reference,
		})
}

func (it newFormatterCreator) Default(
	errType errtype.Variation,
	format string,
	v ...interface{},
) *errorwrapper.Wrapper {
	if format == "" && len(v) == 0 {
		return nil
	}

	compiled := fmt.Sprintf(
		format,
		v...)

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		compiled,
	)
}

func (it newFormatterCreator) Format(
	errType errtype.Variation,
	format string,
	v ...interface{},
) *errorwrapper.Wrapper {
	if format == "" && len(v) == 0 {
		return nil
	}

	compiled := fmt.Sprintf(
		format,
		v...)

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		compiled,
	)
}

func (it newFormatterCreator) Fmt(
	errType errtype.Variation,
	format string,
	v ...interface{},
) *errorwrapper.Wrapper {
	if format == "" && len(v) == 0 {
		return nil
	}

	compiled := fmt.Sprintf(
		format,
		v...)

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		compiled,
	)
}

// Error
//
//	err + format given compiled
//
//	returns null on empty format or empty error
func (it newFormatterCreator) Error(
	errType errtype.Variation,
	err error,
	format string,
	v ...interface{},
) *errorwrapper.Wrapper {
	if err == nil || format == "" && len(v) == 0 {
		return nil
	}

	compiled := fmt.Sprintf(
		format,
		v...)

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		err.Error()+compiled,
	)
}

// Message
//
//	Message + format given compiled
//
//	returns null on empty format or message
func (it newFormatterCreator) Message(
	errType errtype.Variation,
	message string,
	format string,
	v ...interface{},
) *errorwrapper.Wrapper {
	if message == "" || format == "" && len(v) == 0 {
		return nil
	}

	compiled := fmt.Sprintf(
		format,
		v...)

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		message+compiled,
	)
}

// MessageError
//
//	message + err + format given compiled
//
//	returns null on empty format or empty error
func (it newFormatterCreator) MessageError(
	errType errtype.Variation,
	message string,
	err error,
	format string,
	v ...interface{},
) *errorwrapper.Wrapper {
	if err == nil || format == "" && len(v) == 0 {
		return nil
	}

	compiled := fmt.Sprintf(
		format,
		v...)

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		message+err.Error()+compiled,
	)
}

func (it newFormatterCreator) UsingMapOptions(
	isCurlyMap bool,
	errType errtype.Variation,
	format string,
	replacerMap map[string]interface{},
) *errorwrapper.Wrapper {
	if format == "" {
		return nil
	}

	if isCurlyMap {
		return it.CurlyFormatUsingMap(
			errType,
			format,
			replacerMap)
	}

	return it.FormatUsingMap(
		errType,
		format,
		replacerMap)
}

func (it newFormatterCreator) ErrorFormatUsingMap(
	errType errtype.Variation,
	err error,
	format string,
	replacerMap map[string]interface{},
) *errorwrapper.Wrapper {
	if err == nil {
		return nil
	}

	return it.MsgFormatUsingMap(
		errType,
		err.Error(),
		format,
		replacerMap)
}

// MsgFormatUsingMapIf
//
//	Put format : something {user} done {action}
//	map[] => {user} => name, {action} => whatever
func (it newFormatterCreator) MsgFormatUsingMapIf(
	isCreate bool,
	errType errtype.Variation,
	msg string,
	format string,
	replacerMap map[string]interface{},
) *errorwrapper.Wrapper {
	if !isCreate {
		return nil
	}

	return it.MsgFormatUsingMap(
		errType,
		msg,
		format,
		replacerMap)
}

// MsgFormatUsingMap
//
//	Put format : something {user} done {action}
//	map[] => {user} => name, {action} => whatever
func (it newFormatterCreator) MsgFormatUsingMap(
	errType errtype.Variation,
	msg string,
	format string,
	replacerMap map[string]interface{},
) *errorwrapper.Wrapper {
	if msg == "" {
		return nil
	}

	for search, replace := range replacerMap {
		format = strings.ReplaceAll(
			format,
			search,
			converters.AnyTo.ToValueString(replace))
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		msg+" "+format,
	)
}

// FormatUsingMap
//
//	Put format : something {user} done {action}
//	map[] => {user} => name, {action} => whatever
func (it newFormatterCreator) FormatUsingMap(
	errType errtype.Variation,
	format string,
	replacerMap map[string]interface{},
) *errorwrapper.Wrapper {
	if format == "" {
		return nil
	}

	for search, replace := range replacerMap {
		format = strings.ReplaceAll(
			format,
			search,
			converters.AnyTo.ToValueString(replace))
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		format,
	)
}

// CurlyFormatUsingMap
//
//	Put format : something {user} done {action}
//	map[] => user => name, action => whatever
//
//	In FormatUsingMap user have to use curly brace by self.
//	Here it will auto put curly braces
func (it newFormatterCreator) CurlyFormatUsingMap(
	errType errtype.Variation,
	format string,
	replacerMap map[string]interface{},
) *errorwrapper.Wrapper {
	if format == "" {
		return nil
	}

	for search, replace := range replacerMap {
		format = strings.ReplaceAll(
			format,
			simplewrap.CurlyWrap(search),
			converters.AnyTo.ToValueString(replace))
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		format,
	)
}

// CurlyMsgFormatUsingMap
//
//	Put format : something {user} done {action}
//	map[] => user => name, action => whatever
//
//	In FormatUsingMap user have to use curly brace by self.
//	Here it will auto put curly braces
func (it newFormatterCreator) CurlyMsgFormatUsingMap(
	errType errtype.Variation,
	message string,
	format string,
	replacerMap map[string]interface{},
) *errorwrapper.Wrapper {
	if format == "" && message == "" {
		return nil
	}

	for search, replace := range replacerMap {
		format = strings.ReplaceAll(
			format,
			simplewrap.CurlyWrap(search),
			converters.AnyTo.ToValueString(replace))
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		defaultSkipInternal,
		errType,
		message+" "+format,
	)
}
