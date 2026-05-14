package errcmd

import "strings"

// getScriptOfChangeDirPlusScripts cd path && \ \n scripts...
func getScriptOfChangeDirPlusScripts(changeDirPath string, scriptsLines ...string) string {
	compiledNewSlice := append(
		[]string{changeDirScriptLine(changeDirPath)},
		scriptsLines...)

	return strings.Join(compiledNewSlice, SingleLineScriptsJoiner)
}
