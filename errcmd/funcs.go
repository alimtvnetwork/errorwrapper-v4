package errcmd

import "gitlab.com/evatix-go/enum/osmixtype"

type (
	ScriptBuilderWithTypeProcessorFunc func(variant osmixtype.Variant, builder ScriptOnceBuilder)
	FilterScriptBuilderFunc            func(variant osmixtype.Variant, builder ScriptOnceBuilder) (isTake, isBreak bool)
)
