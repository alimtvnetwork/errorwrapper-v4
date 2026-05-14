package errcmd

import osmixtype "github.com/alimtvnetwork/enum-v10/osdetect"

type (
	ScriptBuilderWithTypeProcessorFunc func(variant osmixtype.Variant, builder ScriptOnceBuilder)
	FilterScriptBuilderFunc            func(variant osmixtype.Variant, builder ScriptOnceBuilder) (isTake, isBreak bool)
)
