package errcmd

import (
	"github.com/alimtvnetwork/enum-v10/scripttype"
	"github.com/alimtvnetwork/errorwrapper-v4"
)

type newCmdOnceTypedScriptsCreator struct {
	scriptType scripttype.Variant
}

func (it newCmdOnceTypedScriptsCreator) LinesSecure(
	scriptLines ...string,
) *CmdOnce {
	return New.Script.Lines(
		true,
		true,
		it.scriptType,
		scriptLines...)
}

func (it newCmdOnceTypedScriptsCreator) Line(
	scriptLine string,
) *CmdOnce {
	return New.Script.Lines(
		true,
		false,
		it.scriptType,
		scriptLine)
}

func (it newCmdOnceTypedScriptsCreator) Lines(
	scriptLines ...string,
) *CmdOnce {
	return New.Script.Lines(
		true,
		false,
		it.scriptType,
		scriptLines...)
}

func (it newCmdOnceTypedScriptsCreator) LinesDetailedOutput(
	scriptLines ...string,
) string {
	return New.Script.Lines(
		true,
		false,
		it.scriptType,
		scriptLines...).
		CompiledResult().
		DetailedOutput()
}

func (it newCmdOnceTypedScriptsCreator) ArgsResult(
	args ...string,
) *Result {
	return New.Script.ArgsDefaultResult(
		it.scriptType,
		args...)
}

func (it newCmdOnceTypedScriptsCreator) ArgsErr(
	args ...string,
) *errorwrapper.Wrapper {
	return it.ArgsResult(
		args...).
		CompiledFullErrorWrapper()
}

func (it newCmdOnceTypedScriptsCreator) Args(
	hasOutput,
	hasSecureOutput bool,
	args ...string,
) *CmdOnce {
	return New.Script.Args(
		hasOutput,
		hasSecureOutput,
		it.scriptType,
		args...)
}

func (it newCmdOnceTypedScriptsCreator) ArgsDefault(
	args ...string,
) *CmdOnce {
	return New.Script.ArgsDefault(
		it.scriptType,
		args...)
}

func (it newCmdOnceTypedScriptsCreator) ArgsSecure(
	args ...string,
) *CmdOnce {
	return New.Script.Args(
		true,
		true,
		it.scriptType,
		args...)
}

func (it newCmdOnceTypedScriptsCreator) ArgsNoOutput(
	args ...string,
) *CmdOnce {
	return New.Script.Args(
		false,
		false,
		it.scriptType,
		args...)
}

func (it newCmdOnceTypedScriptsCreator) ArgsNoOutputSecure(
	args ...string,
) *CmdOnce {
	return New.Script.Args(
		false,
		true,
		it.scriptType,
		args...)
}

func (it newCmdOnceTypedScriptsCreator) LinesResult(
	scriptLines ...string,
) *Result {
	cmdOnce := New.Script.Lines(
		true,
		false,
		it.scriptType,
		scriptLines...)

	result := cmdOnce.CompiledResult()
	go cmdOnce.DisposeWithoutResult()
	cmdOnce = nil

	return result
}

func (it newCmdOnceTypedScriptsCreator) Error(
	scriptLines ...string,
) *errorwrapper.Wrapper {
	return it.
		LinesResult(scriptLines...).
		CompiledFullErrorWrapper()
}

func (it newCmdOnceTypedScriptsCreator) ErrorLines(
	scriptLines ...string,
) []string {
	return it.
		LinesResult(scriptLines...).
		CompiledErrorLines()
}

func (it newCmdOnceTypedScriptsCreator) ArgsErrorLines(
	args ...string,
) []string {
	return New.Script.Args(
		false,
		false,
		it.scriptType,
		args...).
		CompiledErrorLines()
}

func (it newCmdOnceTypedScriptsCreator) ArgsTrimmedErrorLines(
	args ...string,
) []string {
	return New.Script.Args(
		false,
		false,
		it.scriptType,
		args...).
		CompiledTrimmedErrorLines()
}

func (it newCmdOnceTypedScriptsCreator) ArgsTrimmedOutputLines(
	args ...string,
) []string {
	return New.Script.Args(
		false,
		false,
		it.scriptType,
		args...).
		CompiledTrimmedOutputLines()
}

func (it newCmdOnceTypedScriptsCreator) ArgsOutputLines(
	args ...string,
) []string {
	return New.Script.Args(
		true,
		false,
		it.scriptType,
		args...).
		CompiledOutputLines()
}

func (it newCmdOnceTypedScriptsCreator) ChangeDirRunArgs(
	changeDirPath string,
	args ...string,
) *CmdOnce {
	argsToSingle :=
		changeDirScriptLine(changeDirPath) +
			SingleLineScriptsJoiner +
			ArgsJoin(args...)

	return New.Script.LinesDefault(
		it.scriptType,
		argsToSingle)
}

func (it newCmdOnceTypedScriptsCreator) ChangeDirRunLines(
	changeDirPath string,
	scriptLines ...string,
) *CmdOnce {
	scriptLine := getScriptOfChangeDirPlusScripts(
		changeDirPath,
		scriptLines...)

	return New.Script.LinesDefault(
		it.scriptType,
		scriptLine)
}

func (it newCmdOnceTypedScriptsCreator) IsSuccessArgs(
	args ...string,
) bool {
	return New.Script.IsSuccessArgs(
		it.scriptType,
		args...)
}

func (it newCmdOnceTypedScriptsCreator) IsSuccessLines(
	scriptLines ...string,
) bool {
	return New.Script.IsSuccessLines(
		it.scriptType,
		scriptLines...)
}

func (it newCmdOnceTypedScriptsCreator) IsFailedArgs(
	args ...string,
) bool {
	return !New.Script.IsSuccessArgs(
		it.scriptType,
		args...)
}

func (it newCmdOnceTypedScriptsCreator) ExitCode(
	args ...string,
) int {
	return New.Script.ExitCode(
		it.scriptType,
		args...)
}

func (it newCmdOnceTypedScriptsCreator) ExitCodeAsByte(
	args ...string,
) byte {
	return New.Script.ExitCodeByte(
		it.scriptType,
		args...)
}
