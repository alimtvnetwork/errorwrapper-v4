package errcmd

import (
	"bytes"
	"os/exec"
	"sync"

	"github.com/alimtvnetwork/core-v9/coredata"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/core-v9/coreinterface"
	"github.com/alimtvnetwork/core-v9/coreinterface/errcoreinf"
	osmixtype "github.com/alimtvnetwork/enum-v10/osdetect"
	"github.com/alimtvnetwork/enum-v10/scripttype"
	"github.com/alimtvnetwork/errorwrapper-v3"
)

type ScriptOnceBuilder interface {
	ScriptBuilderFinalizer

	// CompiledProcessArgs
	//
	//  All ScriptLines will be compiled to
	//  args and script type will be output as process name
	CompiledProcessArgs() (
		processName string,
		args []string,
	)

	SetSecure() ScriptOnceBuilder
	SetPlainOutput() ScriptOnceBuilder
	EnableOutput() ScriptOnceBuilder
	DisableOutput() ScriptOnceBuilder

	SetScriptType(
		scriptType scripttype.Variant,
	) ScriptOnceBuilder

	SetScriptLines(
		scriptLines *corestr.SimpleSlice,
	) ScriptOnceBuilder
	SetStdIn(
		stdIn *bytes.Buffer,
	) ScriptOnceBuilder
	SetStdOut(
		stdOut *bytes.Buffer,
	) ScriptOnceBuilder
	SetStderr(
		stderr *bytes.Buffer,
	) ScriptOnceBuilder
	SetScript(
		scriptType scripttype.Variant,
	) ScriptOnceBuilder
	AllBuffers() (
		stdIn, stdOut, stderr *bytes.Buffer,
	)
	SetInputOutput(
		stdIn, stdOut, stderr *bytes.Buffer,
	) ScriptOnceBuilder
	AppendScriptLines(
		lines ...string,
	) ScriptOnceBuilder
	Args(
		args ...string,
	) ScriptOnceBuilder
	ArgsIf(
		isAdd bool,
		args ...string,
	) ScriptOnceBuilder
	ProcessArgs(
		process string,
		args ...string,
	) ScriptOnceBuilder
	Process(
		process string,
	) ScriptOnceBuilder
	ProcessArgsIf(
		isAdd bool,
		process string, args ...string,
	) ScriptOnceBuilder
	ProcessIf(
		isAdd bool,
		process string,
	) ScriptOnceBuilder

	ScriptLines() corestr.SimpleSlice
	Append(appendItems ...ScriptOnceBuilder) ScriptOnceBuilder
	// AppendCmd
	//
	//  Empty or has issues will be ignored.
	AppendCmd(appendItems ...*exec.Cmd) ScriptOnceBuilder
	// AppendCmdOnce
	//
	//  Empty or has issues will be ignored.
	AppendCmdOnce(appendItems ...*CmdOnce) ScriptOnceBuilder

	IsDefined() bool
	// HasPlainOutput
	//
	// returns true if not secure output and has output
	HasPlainOutput() bool
	// HasInsecureOutput
	//
	// returns true if not secure output
	HasInsecureOutput() bool
	HasStdIn() bool
	HasStdOut() bool
	HasStdErr() bool
	Length() int
	HasAnyItem() bool
	IsEmpty() bool
	IsNull() bool
	IsNotNull() bool
	IsAnyNull() bool
	IsValidWithoutItems() bool
	IsDisposed() bool
	IsFinalized() bool

	SetCurrentEnv() ScriptOnceBuilder
	SetCurrentEnvPlus(
		envValues ...corestr.KeyValuePair,
	) ScriptOnceBuilder
	SetCurrentEnvPlusIf(
		isCondition bool,
		envValues ...corestr.KeyValuePair,
	) ScriptOnceBuilder
	SetCurrentEnvPlusMap(envMap map[string]string) ScriptOnceBuilder
	EnvVars() *corestr.KeyValueCollection
	HasAnyEnvVars() bool
	IsCurrentEnvSet() bool

	Dispose()
	// Clear
	//
	//  clears commands stacks
	Clear()
}

type ScriptBuilderFinalizer interface {
	HasError() bool

	Result() *Result

	CompiledErrorWrapper() *errorwrapper.Wrapper
	ErrorWithTraces() error
	ErrorStringWithTraces() string
	// AsyncStart
	//
	//  Starts cmd in async manner,
	//  have to wait to be waited on
	//
	//  Check out exec.Cmd Start()
	AsyncStart() (*exec.Cmd, error)
	// AsyncStartErrorWrapper
	//
	//  Starts cmd in async manner,
	//  have to wait to be waited on
	//
	//  Check out exec.Cmd Start()
	AsyncStartErrorWrapper() (*exec.Cmd, *errorwrapper.Wrapper)

	CombinedOutput() ([]byte, error)
	CombinedOutputErrorWrapper() ([]byte, *errorwrapper.Wrapper)

	// Output
	//
	//  Gives compiled output
	Output() string

	// OutputMust
	//
	//  Gives compiled output,
	//  on error panics
	OutputMust() string

	// TrimmedOutput
	//
	//  Gives compiled output string without whitespace lines
	TrimmedOutput() string

	// String
	//
	//  Gives compiled output
	String() string

	// Strings
	//
	//  Gives compiled output lines
	Strings() []string
	ErrorOutput() string
	DetailedOutput() string
	OutputLines() []string
	OutputLinesSimpleSlice() *corestr.SimpleSlice
	TrimmedOutputLines() []string

	OutputBytes() []byte
	// StandardOutputCmd
	//
	//  refers to os standard output inputs binding
	//
	//  os.Stdin, os.Stdout, os.Stderr
	StandardOutputCmd() *exec.Cmd

	Build() *CmdOnce
	BuildClear() *CmdOnce
	BuildDispose() *CmdOnce
	BuildCmd() *exec.Cmd
	BuildCmdClear() *exec.Cmd

	NewBuilderWithArgs(args ...string) ScriptOnceBuilder
	Clone() ScriptOnceBuilder

	BuildCmdWithStdIn() (cmd *exec.Cmd, stdInWriter *bytes.Buffer)

	coreinterface.CombinedValidationChecker
	coreinterface.IsSuccessValidator
	errcoreinf.CompiledVoidLogger
}

type ScriptOnceBuilderContractsBinder interface {
	ScriptOnceBuilder
	AsScriptOnceBuilderContractsBinder() ScriptOnceBuilderContractsBinder
}

type CurrentOsScriptBuilder interface {
	CurrentOsTypes() []osmixtype.Variant
	CurrentOsTypesMap() map[osmixtype.Variant]bool
	BuildersMap() map[osmixtype.Variant]ScriptOnceBuilder
	ResultsMap() map[osmixtype.Variant]*Result
	CmdOnceMap() map[osmixtype.Variant]*CmdOnce
	OutputMap() map[osmixtype.Variant]string
	DetailedMap() map[osmixtype.Variant]string
	CmdMap() map[osmixtype.Variant]*exec.Cmd
	OutputLinesSimpleSliceMap() map[osmixtype.Variant]*corestr.SimpleSlice

	// HasAnyItem
	//
	// Only counts and checks available builders, O(N), expensive
	HasAnyItem() bool
	// IsEmpty
	//
	// Only counts and checks available builders, O(N), expensive
	IsEmpty() bool

	coreinterface.CombinedValidationChecker
	coreinterface.IsSuccessValidator
	errcoreinf.CompiledVoidLogger

	Length() int
	AvailableLength() int
}

type OsMixTypeScriptBuilder interface {
	TaskName() string
	SetTaskName(taskName string) OsMixTypeScriptBuilder

	AddNewBuilder(
		osMixType osmixtype.Variant,
		scriptType scripttype.Variant,
	) OsMixTypeScriptBuilder
	AddBuilder(
		osMixType osmixtype.Variant,
		builder ScriptOnceBuilder,
	) OsMixTypeScriptBuilder

	Description() string
	SetDescription(description string) OsMixTypeScriptBuilder

	Url() string
	SetUrl(url string) OsMixTypeScriptBuilder

	ExistingOsMixTypes() []osmixtype.Variant
	SetCurrentEnv(
		osMixTypes ...osmixtype.Variant,
	) *errorwrapper.Wrapper
	SetCurrentEnvToAll() OsMixTypeScriptBuilder
	SetCurrentEnvPlus(
		osMixType osmixtype.Variant,
		envValues ...corestr.KeyValuePair,
	) *errorwrapper.Wrapper
	SetCurrentEnvPlusIf(
		isCondition bool,
		osMixType osmixtype.Variant,
		envValues ...corestr.KeyValuePair,
	) *errorwrapper.Wrapper
	SetCurrentEnvPlusMapAll(envMap map[string]string) OsMixTypeScriptBuilder
	EnvVars(osMixType osmixtype.Variant) *corestr.KeyValueCollection
	HasAnyEnvVarsOnAll(osMixTypes ...osmixtype.Variant) bool
	IsCurrentEnvSetOnAll(osMixTypes ...osmixtype.Variant) bool

	AddProcessArgs(
		process string,
		args ...string,
	) OsMixTypeScriptBuilder

	AddProcessArgsBy(
		osMixType osmixtype.Variant,
		process string,
		args ...string,
	) OsMixTypeScriptBuilder
	AddArgsByIf(
		isAdd bool,
		osMixType osmixtype.Variant,
		args ...string,
	) OsMixTypeScriptBuilder
	AddProcessArgsByIf(
		isAdd bool,
		osMixType osmixtype.Variant,
		process string,
		args ...string,
	) OsMixTypeScriptBuilder
	AddArgsBy(
		osMixType osmixtype.Variant,
		args ...string,
	) OsMixTypeScriptBuilder
	AddArgs(args ...string) OsMixTypeScriptBuilder
	AddArgsIf(
		isAdd bool,
		args ...string,
	) OsMixTypeScriptBuilder

	Unix() ScriptOnceBuilder
	Windows() ScriptOnceBuilder
	Linux() ScriptOnceBuilder
	MacOs() ScriptOnceBuilder
	Ubuntu() ScriptOnceBuilder
	CentOs() ScriptOnceBuilder
	GetBy(
		osMixType osmixtype.Variant,
	) ScriptOnceBuilder

	// GetWithStat
	//
	//  isDefined : found and not null
	GetWithStat(
		osMixType osmixtype.Variant,
	) (scriptOnceBuilder ScriptOnceBuilder, isDefined bool)
	HasBuilder(
		mixType osmixtype.Variant,
	) bool

	// IsValidBuilder
	//
	//	represents that builder present
	//	and script lines present
	//	and script type is valid
	//
	//  return it != nil &&
	//		it.scriptType.IsValid() &&
	//		!it.HasError() &&
	//		it.scriptLines.HasAnyItem()
	IsValidBuilder(
		mixType osmixtype.Variant,
	) bool
	// IsInvalidBuilder
	//
	//  invert of IsValidBuilder
	IsInvalidBuilder(
		mixType osmixtype.Variant,
	) bool

	// IsAllBuildersValid
	//
	//  represents that all builder exist
	//  with lines and script type is valid
	IsAllBuildersValid(
		mixTypes ...osmixtype.Variant,
	) bool
	IsAnyInvalidBuilders(
		mixTypes ...osmixtype.Variant,
	) bool
	IsBuilderMissing(
		mixType osmixtype.Variant,
	) bool
	// IsBuilderMissingOrInvalid
	//
	//  Either builder not found or,
	//  builder has no script lines or
	//  script type is not valid
	IsBuilderMissingOrInvalid(
		mixType osmixtype.Variant,
	) bool
	HasBuilderScriptLines(
		mixType osmixtype.Variant,
	) bool
	GetBuilder(
		mixType osmixtype.Variant,
	) (scriptBuilder ScriptOnceBuilder, hasBuilder bool)
	HasAllBuilder(
		mixTypes ...osmixtype.Variant,
	) bool
	HasAnyBuilder(
		mixTypes ...osmixtype.Variant,
	) bool

	HasUnix() bool
	HasWindows() bool
	HasLinux() bool

	IsEmptyUnix() bool
	IsEmptyWindows() bool
	IsEmptyLinux() bool
	IsEmptyMacOs() bool

	IsEmptyBy(osMixType osmixtype.Variant) bool
	HasAnyItemBy(osMixType osmixtype.Variant) bool

	LoopInvokeFunc(processorFunc ScriptBuilderWithTypeProcessorFunc)

	Filter(
		filter FilterScriptBuilderFunc,
	) (resultMap map[osmixtype.Variant]ScriptOnceBuilder)

	FilterOsMixTypes(
		filter FilterScriptBuilderFunc,
	) (osMixTypes []osmixtype.Variant)
	GetBuildersMapByOsMixTypes(
		osMixTypes []osmixtype.Variant,
	) (resultMap map[osmixtype.Variant]ScriptOnceBuilder)
	UnixBuild() *CmdOnce
	WindowsBuild() *CmdOnce
	LinuxBuild() *CmdOnce
	MacOsBuild() *CmdOnce

	UnixBuildCmd() *exec.Cmd
	WindowsBuildCmd() *exec.Cmd
	LinuxBuildCmd() *exec.Cmd
	MacOsBuildCmd() *exec.Cmd

	List() []ScriptOnceBuilder
	ListMap() map[osmixtype.Variant]ScriptOnceBuilder
	IsNull() bool

	AvailableBuilders() []ScriptOnceBuilder
	BuildersMap() map[osmixtype.Variant]ScriptOnceBuilder
	AvailableBuildersMap() map[osmixtype.Variant]ScriptOnceBuilder
	ExecutionResultMap() map[osmixtype.Variant]*Result
	BuildMap() map[osmixtype.Variant]*CmdOnce

	BuildMapOnly(
		osMixes ...osmixtype.Variant,
	) map[osmixtype.Variant]*CmdOnce
	ResultMapOnly(
		osMixes ...osmixtype.Variant,
	) map[osmixtype.Variant]*Result
	BuildCmdMapOnly(
		osMixes ...osmixtype.Variant,
	) map[osmixtype.Variant]*exec.Cmd
	OutputMapOnly(
		osMixes ...osmixtype.Variant,
	) map[osmixtype.Variant]string
	OutputLinesSimpleSliceMapOnly(
		osMixes ...osmixtype.Variant,
	) map[osmixtype.Variant]*corestr.SimpleSlice
	DetailedOutputMapOnly(
		osMixes ...osmixtype.Variant,
	) map[osmixtype.Variant]string
	BuildersMapOnly(
		osMixes ...osmixtype.Variant,
	) map[osmixtype.Variant]ScriptOnceBuilder

	CurrentOsTypes() []osmixtype.Variant
	CurrentOsTypesMap() map[osmixtype.Variant]bool

	CurrentOsScriptBuilder() CurrentOsScriptBuilder

	CompiledErrorWrappersMapOnly(
		osMixes ...osmixtype.Variant,
	) (compiledMap map[osmixtype.Variant]*errorwrapper.Wrapper, hasAnyError bool)

	AllCompiledErrorWrappersMapOnly() (
		compiledMap map[osmixtype.Variant]*errorwrapper.Wrapper, hasAnyError bool,
	)

	BuildCmdMap() map[osmixtype.Variant]*exec.Cmd
	OutputMap() map[osmixtype.Variant]string

	OutputBytesMap() map[osmixtype.Variant][]byte
	CombinedOutputMap() map[osmixtype.Variant]coredata.BytesError
	AsyncExecutionResultMap() map[osmixtype.Variant]*Result

	StandardOutputCmd() OsMixTypeScriptBuilder
	AsyncStart(stdIn, stdOut, stderr *bytes.Buffer) (
		map[osmixtype.Variant]error, *sync.WaitGroup,
	)
	OutputMapMust() (resultMap map[osmixtype.Variant]string)
	TrimmedOutputMap() (resultMap map[osmixtype.Variant]string)
	OutputLinesMap() (resultMap map[osmixtype.Variant][]string)
	ErrorOutputMap() (resultMap map[osmixtype.Variant]string)
	TrimmedOutputLinesMap() (resultMap map[osmixtype.Variant][]string)
	BuildClearMap() (resultMap map[osmixtype.Variant]*CmdOnce)
	BuildDisposeMap() (resultMap map[osmixtype.Variant]*CmdOnce)
	// HasAnyItem
	//
	// Only counts and checks available builders, O(N), expensive
	HasAnyItem() bool
	// IsEmpty
	//
	// Only counts and checks available builders, O(N), expensive
	IsEmpty() bool

	coreinterface.CombinedValidationChecker
	coreinterface.IsSuccessValidator
	errcoreinf.CompiledVoidLogger

	Length() int
	// ValidLength
	//
	// Only counts and checks available builders, O(N), expensive
	ValidLength() int
	String() string
	StringsMap() (resultMap map[osmixtype.Variant][]string)
}
