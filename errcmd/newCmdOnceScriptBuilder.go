package errcmd

import (
	"bytes"

	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/core-v9/osconsts"
	"github.com/alimtvnetwork/enum-v10/scripttype"
)

type newCmdOnceScriptBuilder struct{}

// Default
//
//  Has output and insecure (plain text - all errors captured).
func (it newCmdOnceScriptBuilder) Default(
	scriptType scripttype.Variant,
) ScriptOnceBuilder {
	return &scriptOnceBuilder{
		scriptType:  scriptType,
		scriptLines: corestr.New.SimpleSlice.Cap(constants.Capacity5),
		hasOutput:   true,
	}
}

// DefaultBash
//
//  Has output and insecure (plain text - all errors captured).
func (it newCmdOnceScriptBuilder) DefaultBash() ScriptOnceBuilder {
	return it.Default(scripttype.Bash)
}

// DefaultShell
//
//  Has output and insecure (plain text - all errors captured).
func (it newCmdOnceScriptBuilder) DefaultShell() ScriptOnceBuilder {
	return it.Default(scripttype.Shell)
}

// DefaultDependingOnOs
//
//  Has output and insecure (plain text - all errors captured).
//
// Depending on OS:
//  - Windows : returns powershell
//  - Unix    : returns bash
func (it newCmdOnceScriptBuilder) DefaultDependingOnOs() ScriptOnceBuilder {
	scriptType := scripttype.Bash

	if osconsts.IsWindows {
		scriptType = scripttype.Powershell
	}

	return it.Default(scriptType)
}

// DependingOnOs
//
//  Has output and insecure (plain text - all errors captured).
func (it newCmdOnceScriptBuilder) DependingOnOs(
	unixScriptType, windowsScriptType scripttype.Variant,
) ScriptOnceBuilder {
	scriptType := unixScriptType

	if osconsts.IsWindows {
		scriptType = windowsScriptType
	}

	return it.Default(scriptType)
}

// DefaultCmd
//
//  Has output and insecure (plain text - all errors captured).
func (it newCmdOnceScriptBuilder) DefaultCmd() ScriptOnceBuilder {
	return it.Default(scripttype.Cmd)
}

// DefaultPowershell
//
//  Has output and insecure (plain text - all errors captured).
func (it newCmdOnceScriptBuilder) DefaultPowershell() ScriptOnceBuilder {
	return it.Default(scripttype.Powershell)
}

// SecureOutput
//
//  Has output and insecure (plain text - all errors captured).
func (it newCmdOnceScriptBuilder) SecureOutput(
	scriptType scripttype.Variant,
) ScriptOnceBuilder {
	return &scriptOnceBuilder{
		scriptType:  scriptType,
		scriptLines: corestr.New.SimpleSlice.Cap(constants.Capacity5),
		isSecure:    true,
		hasOutput:   true,
	}
}

// NoOutput
//
//  No output and insecure (plain text - all errors captured).
func (it newCmdOnceScriptBuilder) NoOutput(
	scriptType scripttype.Variant,
) ScriptOnceBuilder {
	return &scriptOnceBuilder{
		scriptType:  scriptType,
		scriptLines: corestr.New.SimpleSlice.Cap(constants.Capacity5),
	}
}

// NoOutputSecure
//
//  No output and secure (errors will NOT be displayed or captured).
func (it newCmdOnceScriptBuilder) NoOutputSecure(
	scriptType scripttype.Variant,
) ScriptOnceBuilder {
	return &scriptOnceBuilder{
		scriptType:  scriptType,
		scriptLines: corestr.New.SimpleSlice.Cap(constants.Capacity5),
		isSecure:    true,
		hasOutput:   false,
	}
}

// CustomStdIns
//
//  Has output and insecure (plain text - all errors captured).
func (it newCmdOnceScriptBuilder) CustomStdIns(
	scriptType scripttype.Variant,
	stdIn, stdOut, stderr *bytes.Buffer,
) ScriptOnceBuilder {
	return &scriptOnceBuilder{
		scriptType:  scriptType,
		scriptLines: corestr.New.SimpleSlice.Cap(constants.Capacity5),
		isSecure:    false,
		hasOutput:   true,
		stdIn:       stdIn,
		stdOut:      stdOut,
		stderr:      stderr,
	}
}

func (it newCmdOnceScriptBuilder) ProcessArgs(
	scriptType scripttype.Variant,
	isSecure,
	hasOutput bool,
	process string,
	args ...string,
) ScriptOnceBuilder {
	builder := &scriptOnceBuilder{
		scriptType:  scriptType,
		scriptLines: corestr.New.SimpleSlice.Cap(constants.Capacity5),
		isSecure:    isSecure,
		hasOutput:   hasOutput,
	}

	return builder.ProcessArgs(process, args...)
}

// ProcessArgsDefault
//
//  Has output and insecure (plain text - all errors captured).
func (it newCmdOnceScriptBuilder) ProcessArgsDefault(
	scriptType scripttype.Variant,
	process string,
	args ...string,
) ScriptOnceBuilder {
	builder := &scriptOnceBuilder{
		scriptType:  scriptType,
		scriptLines: corestr.New.SimpleSlice.Cap(constants.Capacity5),
		hasOutput:   true,
	}

	return builder.ProcessArgs(process, args...)
}
