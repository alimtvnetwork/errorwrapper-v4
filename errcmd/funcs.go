package errcmd

import "github.com/alimtvnetwork/enum-v10/osmixtype"

type (
	ScriptBuilderWithTypeProcessorFunc func(variant osmixtype.Variant, builder ScriptOnceBuilder)
	FilterScriptBuilderFunc            func(variant osmixtype.Variant, builder ScriptOnceBuilder) (isTake, isBreak bool)
)
