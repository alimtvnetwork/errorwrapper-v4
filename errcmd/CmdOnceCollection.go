package errcmd

import (
	"strings"

	"github.com/alimtvnetwork/core-v9/conditional"
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata/stringslice"
	"github.com/alimtvnetwork/core-v9/coreinterface"
	"github.com/alimtvnetwork/enum-v10/scripttype"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errnew"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
)

type CmdOnceCollection struct {
	items []*CmdOnce
}

func NewCmdOnceCollection2() *CmdOnceCollection {
	return NewCmdOnceCollection(constants.Capacity2)
}

func NewCmdOnceCollection(capacity int) *CmdOnceCollection {
	items := make([]*CmdOnce, 0, capacity)

	return &CmdOnceCollection{
		items: items,
	}
}

func NewCmdOnceCollectionUsingLinesOfScripts(
	scriptType scripttype.Variant,
	scriptLines ...string,
) *CmdOnceCollection {
	length := len(scriptLines)
	cmdOnceCollection := NewCmdOnceCollection(length)

	if length == 0 {
		return cmdOnceCollection
	}

	return cmdOnceCollection.AddEachScriptAsEachCmdOnce(
		scriptType,
		scriptLines...)
}

func NewCmdOnceCollectionUsingLinesDirect(
	scriptType scripttype.Variant,
	isTrimEmptyLines bool,
	scriptLines ...string,
) *CmdOnceCollection {
	if scriptLines == nil {
		return NewCmdOnceCollection2()
	}

	return NewCmdOnceCollectionUsingLines(
		scriptType,
		isTrimEmptyLines,
		scriptLines...)
}

func BashScriptsLinesCmdOneCollection(
	isTrimEmptyLines bool,
	scriptLines ...string,
) *CmdOnceCollection {
	if scriptLines == nil {
		return NewCmdOnceCollection2()
	}

	return NewCmdOnceCollectionUsingLines(
		scripttype.Bash,
		isTrimEmptyLines,
		scriptLines...)
}

func NewCmdOnceCollectionUsingLines(
	scriptType scripttype.Variant,
	isTrimEmptyLines bool,
	scriptLines ...string,
) *CmdOnceCollection {
	length := len(scriptLines)
	if length == 0 {
		return NewCmdOnceCollection2()
	}

	collection := NewCmdOnceCollection(length)

	scriptLines2 := conditional.IfSliceString(
		isTrimEmptyLines,
		stringslice.NonWhitespace(scriptLines),
		scriptLines)

	return collection.
		AddEachScriptAsEachCmdOnce(
			scriptType,
			scriptLines2...)
}

func NewCmdOnceCollectionUsingLinesPtr(
	scriptType scripttype.Variant,
	isTrimEmptyLines bool,
	scriptLines *[]string,
) *CmdOnceCollection {
	if scriptLines == nil || *scriptLines == nil {
		return NewCmdOnceCollection2()
	}

	return NewCmdOnceCollectionUsingLines(
		scriptType,
		isTrimEmptyLines,
		*scriptLines...)
}

func (it *CmdOnceCollection) AddMany(
	cmdOnceItems ...*CmdOnce,
) *CmdOnceCollection {
	if len(cmdOnceItems) == 0 {
		return it
	}

	for _, item := range cmdOnceItems {
		it.items = append(it.items, item)
	}

	return it
}

func (it *CmdOnceCollection) Add(
	cmdOnce *CmdOnce,
) *CmdOnceCollection {
	if cmdOnce == nil {
		return it
	}

	it.items = append(it.items, cmdOnce)

	return it
}

// AddDefaultScript joins using SingleLineScriptsJoiner whereas args joins using space
func (it *CmdOnceCollection) AddDefaultScript(
	scriptType scripttype.Variant,
	scriptsLines ...string,
) *CmdOnceCollection {
	if len(scriptsLines) == 0 {
		return it
	}

	it.items = append(
		it.items,
		New.Script.LinesDefault(scriptType, scriptsLines...))

	return it
}

// AddDefaultScriptArgs args joins with space and whereas Scripts joins with SingleLineScriptsJoiner
func (it *CmdOnceCollection) AddDefaultScriptArgs(
	scriptType scripttype.Variant,
	args ...string,
) *CmdOnceCollection {
	if len(args) == 0 {
		return it
	}

	it.items = append(
		it.items,
		New.Script.ArgsDefault(
			scriptType,
			args...))

	return it
}

func (it *CmdOnceCollection) AddBashArgsDefault(
	args ...string,
) *CmdOnceCollection {
	if len(args) == 0 {
		return it
	}

	it.items = append(
		it.items,
		New.BashScript.ArgsDefault(args...))

	return it
}

func (it *CmdOnceCollection) AddBashEachScriptAsEachCmdOnce(
	scriptLines ...string,
) *CmdOnceCollection {
	if len(scriptLines) == 0 {
		return it
	}

	for _, bashScriptLine := range scriptLines {
		it.items = append(
			it.items,
			New.BashScript.Lines(bashScriptLine))
	}

	return it
}

func (it *CmdOnceCollection) AddEachScriptAsEachCmdOnce(
	scriptType scripttype.Variant,
	scriptLines ...string,
) *CmdOnceCollection {
	if len(scriptLines) == 0 {
		return it
	}

	for _, bashScriptLine := range scriptLines {
		it.items = append(
			it.items,
			New.Script.ArgsDefault(scriptType, bashScriptLine))
	}

	return it
}

func (it *CmdOnceCollection) AddBashChangeDirArgs(
	changeDir string,
	args ...string,
) *CmdOnceCollection {
	if len(args) == 0 {
		return it
	}

	it.items = append(
		it.items,
		New.BashScript.ChangeDirRunArgs(changeDir, args...))

	return it
}

func (it *CmdOnceCollection) AddBashChangeDirScripts(
	changeDir string,
	scripts ...string,
) *CmdOnceCollection {
	if len(scripts) == 0 {
		return it
	}

	it.items = append(
		it.items,
		New.BashScript.ChangeDirRunLines(changeDir, scripts...))

	return it
}

func (it *CmdOnceCollection) Length() int {
	return len(it.items)
}

func (it *CmdOnceCollection) Count() int {
	return it.Length()
}

func (it *CmdOnceCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *CmdOnceCollection) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *CmdOnceCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *CmdOnceCollection) Items() []*CmdOnce {
	return it.items
}

func (it *CmdOnceCollection) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *CmdOnceCollection) CommandsStrings() []string {
	length := it.Length()
	slice := stringslice.Make(length, length)

	for i, item := range it.items {
		slice[i] = item.GetCommandLineDataDependingOnSecurity()
	}

	return slice
}

func (it *CmdOnceCollection) OutputsStrings() []string {
	length := it.Length()
	slice := stringslice.Make(length, length)

	for i, item := range it.items {
		slice[i] = item.CompiledOutput()
	}

	return slice
}

func (it *CmdOnceCollection) ExecuteUntilErr() *errorwrapper.Wrapper {
	for _, cmdOnce := range it.items {
		if cmdOnce.HasAnyIssues() {
			return cmdOnce.CompiledErrorWrapper()
		}
	}

	return nil
}

func (it *CmdOnceCollection) ExecuteAll() *errorwrapper.Wrapper {
	return it.CompiledOneErrorWrapper()
}

func (it *CmdOnceCollection) ExecuteAllWithoutLazyResults() []*Result {
	length := it.Length()
	if length == 0 {
		return []*Result{}
	}

	results := make(
		[]*Result,
		length)

	for i, cmdOnce := range it.items {
		results[i] = cmdOnce.
			Run()
	}

	return results
}

func (it *CmdOnceCollection) ExecuteAllWithoutLazy() *errorwrapper.Wrapper {
	var slice []string
	length := it.Length()
	if length == 0 {
		return nil
	}

	for _, cmdOnce := range it.items {
		errWrapper := cmdOnce.Run().ErrorWrapper()

		if errWrapper.HasError() {
			slice = append(
				slice,
				errWrapper.FullString())
		}
	}

	if len(slice) == 0 {
		return nil
	}

	return errnew.Messages.Many(
		errtype.CmdOnceRunningIssue,
		slice...)
}

func (it *CmdOnceCollection) Results() []*Result {
	length := it.Length()
	slice := make([]*Result, length)

	for i, item := range it.items {
		slice[i] = item.CompiledResult()
	}

	return slice
}

func (it *CmdOnceCollection) ErrorWrappers() []*errorwrapper.Wrapper {
	var slice []*errorwrapper.Wrapper

	for _, item := range it.items {
		errorWrapper := item.CompiledErrorWrapper()

		if errorWrapper.HasError() {
			slice = append(
				slice,
				item.CompiledErrorWrapper())
		}
	}

	return slice
}

func (it *CmdOnceCollection) CompiledOneErrorWrapper() *errorwrapper.Wrapper {
	sliceError := it.ErrorWrappers()
	length := len(sliceError)
	if length == 0 {
		return nil
	}

	slice := make([]string, length)

	for i, item := range sliceError {
		slice[i] = item.FullString()
	}

	return errnew.Messages.Many(
		errtype.CmdOnceRunningIssue,
		slice...)
}

func (it *CmdOnceCollection) HasAnyIssuesRunning() bool {
	if it.IsEmpty() {
		return false
	}

	for _, item := range it.items {
		if item.HasAnyIssues() {
			return true
		}
	}

	return false
}

func (it *CmdOnceCollection) IsAllSuccess() bool {
	if it.IsEmpty() {
		return true
	}

	for _, item := range it.items {
		if item.HasAnyIssues() {
			return false
		}
	}

	return true
}

func (it *CmdOnceCollection) StringsProcessor(
	processor func(cmdOnce *CmdOnce) (processedLine string, isTake, isBreak bool),
) []string {
	length := it.Length()
	slice := stringslice.MakeDefault(length)

	for _, item := range it.items {
		processedLine, isTake, isBreak := processor(item)

		if isTake {
			slice = append(slice, processedLine)
		}

		if isBreak {
			return slice
		}
	}

	return slice
}

// SuccessStringsOnly returns slice of strings output if cmd has success only
func (it *CmdOnceCollection) SuccessStringsOnly() []string {
	return it.StringsProcessor(
		successCmdOneOutputStringProcessor)
}

// FailedOutputStringsOnly returns slice of strings output if cmd has failed only
func (it *CmdOnceCollection) FailedOutputStringsOnly() []string {
	return it.StringsProcessor(
		failedCmdOneOutputStringProcessor)
}

func (it *CmdOnceCollection) Strings() []string {
	length := it.Length()
	slice := stringslice.Make(length, length)

	for i, item := range it.items {
		slice[i] = item.String()
	}

	return slice
}

func (it *CmdOnceCollection) Clear() *CmdOnceCollection {
	if it == nil {
		return it
	}

	tempItems := it.items
	clearFunc := func() {
		for _, item := range tempItems {
			item.Dispose()
		}

		tempItems = nil
	}

	go clearFunc()

	it.items = nil
	it.items = []*CmdOnce{}

	return it
}

func (it *CmdOnceCollection) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.items = nil
}

func (it *CmdOnceCollection) String() string {
	return strings.Join(
		it.Strings(),
		constants.NewLineUnix)
}

func (it *CmdOnceCollection) Clone() *CmdOnceCollection {
	slice := make([]*CmdOnce, it.Length())

	for i, cmdOnce := range it.items {
		slice[i] = cmdOnce.Clone()
	}

	return &CmdOnceCollection{items: slice}
}

func (it *CmdOnceCollection) AsBasicSliceContractsBinder() coreinterface.BasicSlicerContractsBinder {
	return it
}

func (it *CmdOnceCollection) AsBasicSlicerContractsBinder() coreinterface.BasicSlicerContractsBinder {
	return it
}
