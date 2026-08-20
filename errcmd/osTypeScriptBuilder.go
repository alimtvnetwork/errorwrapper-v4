package errcmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"sync"

	"github.com/alimtvnetwork/core-v9/codestack"
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/core-v9/coredata"
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/coredata/corestr"
	"github.com/alimtvnetwork/core-v9/errcore"
	osmixtype "github.com/alimtvnetwork/enum-v10/osdetect"
	"github.com/alimtvnetwork/enum-v10/scripttype"
	"github.com/alimtvnetwork/errorwrapper-v4"
	"github.com/alimtvnetwork/errorwrapper-v4/errtype"
	"github.com/alimtvnetwork/errorwrapper-v4/ref"
)

type osTypeScriptBuilder struct {
	taskName, description, url string
	itemsMap                   map[osmixtype.Variant]*scriptOnceBuilder
}

func (it *osTypeScriptBuilder) AddNewBuilder(
	osMixType osmixtype.Variant,
	scriptType scripttype.Variant,
) OsMixTypeScriptBuilder {
	builder := New.
		ScriptBuilder.
		Default(scriptType).(*scriptOnceBuilder)

	it.itemsMap[osMixType] = builder

	return it
}

func (it *osTypeScriptBuilder) AddBuilder(
	osMixType osmixtype.Variant,
	builder ScriptOnceBuilder,
) OsMixTypeScriptBuilder {
	if builder == nil {
		return it
	}

	casted, isSuccess := builder.(*scriptOnceBuilder)

	if isSuccess {
		it.itemsMap[osMixType] = casted
	}

	return it
}

func (it *osTypeScriptBuilder) HasError() bool {
	if it.IsEmpty() {
		return false
	}

	_, hasAnyErr := it.AllCompiledErrorWrappersMapOnly()

	return hasAnyErr
}

func (it *osTypeScriptBuilder) HasAnyIssues() bool {
	return it.HasError()
}

func (it *osTypeScriptBuilder) ValidationError() error {
	return it.Error()
}

func (it *osTypeScriptBuilder) Error() error {
	allErrorsMap, hasAnyError := it.AllCompiledErrorWrappersMapOnly()

	if !hasAnyError {
		return nil
	}

	return it.compiledMappedErrorWrappersToError(
		allErrorsMap)
}

func (it *osTypeScriptBuilder) compiledMappedErrorWrappersToError(
	errorWrappersMapped map[osmixtype.Variant]*errorwrapper.Wrapper,
) error {
	if len(errorWrappersMapped) == 0 {
		return nil
	}

	rawErrCollection := errcore.RawErrCollection{}

	for mixType, wrapper := range errorWrappersMapped {
		rawErrCollection.Fmt(
			"os-type: %s, cmd err: %s",
			mixType.NameValue(),
			wrapper.CompiledError(),
		)
	}

	return errorwrapper.NewMsgDisplayErrorNoReference(
		codestack.Skip1,
		errtype.FailedProcess,
		rawErrCollection.ErrorString()).
		CompiledErrorWithStackTraces()
}

func (it *osTypeScriptBuilder) IsValid() bool {
	return !it.HasError()
}

func (it *osTypeScriptBuilder) IsSuccess() bool {
	return !it.HasError()
}

func (it *osTypeScriptBuilder) IsFailed() bool {
	return it.HasError()
}

func (it *osTypeScriptBuilder) TaskName() string {
	return it.taskName
}

func (it *osTypeScriptBuilder) SetTaskName(taskName string) OsMixTypeScriptBuilder {
	it.taskName = taskName

	return it
}

func (it *osTypeScriptBuilder) Description() string {
	return it.description
}

func (it *osTypeScriptBuilder) SetDescription(description string) OsMixTypeScriptBuilder {
	it.description = description

	return it
}

func (it *osTypeScriptBuilder) Url() string {
	return it.url
}

func (it *osTypeScriptBuilder) SetUrl(url string) OsMixTypeScriptBuilder {
	it.url = url

	return it
}

func (it *osTypeScriptBuilder) ExistingOsMixTypes() []osmixtype.Variant {
	if it == nil {
		return nil
	}

	slice := make([]osmixtype.Variant, it.Length())

	index := 0
	for variant := range it.itemsMap {
		slice[index] = variant
		index++
	}

	return slice
}

func (it *osTypeScriptBuilder) SetCurrentEnv(
	osMixTypes ...osmixtype.Variant,
) *errorwrapper.Wrapper {
	invalids := corestr.New.SimpleSlice.Empty()

	for _, mixType := range osMixTypes {
		builder, isValid := it.GetWithStat(mixType)

		if !isValid {
			invalids.Add(mixType.NameValue())

			continue
		}

		builder.SetCurrentEnv()
	}

	if invalids.IsEmpty() {
		return nil
	}

	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		codestack.Skip1,
		errtype.CategoryMissing,
		"some of the os-mix-types not found",
		ref.Value{
			Variable: "Not found os types in builder map",
			Value:    invalids.JoinCsv(),
		})
}

func (it *osTypeScriptBuilder) SetCurrentEnvToAll() OsMixTypeScriptBuilder {
	for mixType, builder := range it.itemsMap {
		builder.SetCurrentEnv()
		it.itemsMap[mixType] = builder
	}

	return it
}

func (it *osTypeScriptBuilder) SetCurrentEnvPlus(
	osMixType osmixtype.Variant,
	envValues ...corestr.KeyValuePair,
) *errorwrapper.Wrapper {
	builder, isValid := it.GetWithStat(osMixType)

	if isValid {
		builder.SetCurrentEnvPlus(envValues...)

		return nil
	}

	// has err
	return errorwrapper.NewMsgDisplayErrorReferencesPtr(
		codestack.Skip1,
		errtype.NotDefined,
		osMixType.NameValue()+" not present nor valid in the map.",
		ref.Value{
			Variable: "existing os types",
			Value:    it.ExistingOsMixTypes(),
		})
}

func (it *osTypeScriptBuilder) SetCurrentEnvPlusIf(
	isCondition bool,
	osMixType osmixtype.Variant,
	envValues ...corestr.KeyValuePair,
) *errorwrapper.Wrapper {
	if !isCondition {
		return nil
	}

	return it.SetCurrentEnvPlus(
		osMixType, envValues...)
}

func (it *osTypeScriptBuilder) SetCurrentEnvPlusMapAll(
	envMap map[string]string,
) OsMixTypeScriptBuilder {
	for mixType, builder := range it.itemsMap {
		builder.SetCurrentEnvPlusMap(envMap)
		it.itemsMap[mixType] = builder
	}

	return it
}

func (it *osTypeScriptBuilder) EnvVars(osMixType osmixtype.Variant) *corestr.KeyValueCollection {
	builder, isValid := it.GetBuilder(osMixType)

	if isValid {
		return builder.EnvVars()
	}

	return nil
}

func (it *osTypeScriptBuilder) HasAnyEnvVarsOnAll(
	osMixTypes ...osmixtype.Variant,
) bool {
	if it.IsEmpty() {
		return len(osMixTypes) == 0
	}

	for _, mixType := range osMixTypes {
		builder, isValid := it.GetWithStat(mixType)

		if !isValid || !builder.HasAnyEnvVars() {
			return false
		}
	}

	return true
}

func (it *osTypeScriptBuilder) IsCurrentEnvSetOnAll(
	osMixTypes ...osmixtype.Variant,
) bool {
	if it.IsEmpty() {
		return len(osMixTypes) == 0
	}

	for _, mixType := range osMixTypes {
		builder, isValid := it.GetWithStat(mixType)

		if !isValid || !builder.IsCurrentEnvSet() {
			return false
		}
	}

	return true
}

func (it *osTypeScriptBuilder) AddProcessArgs(
	process string,
	args ...string,
) OsMixTypeScriptBuilder {
	for _, builder := range it.List() {
		if builder == nil {
			continue
		}

		builder.ProcessArgs(
			process,
			args...)
	}

	return it
}

func (it *osTypeScriptBuilder) AddProcessArgsBy(
	osMixType osmixtype.Variant,
	process string,
	args ...string,
) OsMixTypeScriptBuilder {
	builder := it.GetBy(osMixType)

	if builder == nil {
		return it
	}

	builder.ProcessArgs(process, args...)

	return it
}

func (it *osTypeScriptBuilder) AddArgsByIf(
	isAdd bool,
	osMixType osmixtype.Variant,
	args ...string,
) OsMixTypeScriptBuilder {
	if !isAdd {
		return it
	}

	return it.AddArgsBy(osMixType, args...)
}

func (it *osTypeScriptBuilder) AddProcessArgsByIf(
	isAdd bool,
	osMixType osmixtype.Variant,
	process string,
	args ...string,
) OsMixTypeScriptBuilder {
	if !isAdd {
		return it
	}

	return it.AddProcessArgsBy(
		osMixType,
		process,
		args...)
}

func (it *osTypeScriptBuilder) AddArgsBy(
	osMixType osmixtype.Variant,
	args ...string,
) OsMixTypeScriptBuilder {
	builder := it.GetBy(osMixType)

	if builder == nil {
		return it
	}

	builder.Args(args...)

	return it
}

func (it *osTypeScriptBuilder) AddArgs(args ...string) OsMixTypeScriptBuilder {
	for _, builder := range it.List() {
		if builder == nil {
			continue
		}

		builder.Args(args...)
	}

	return it
}

func (it *osTypeScriptBuilder) AddArgsIf(
	isAdd bool,
	args ...string,
) OsMixTypeScriptBuilder {
	if !isAdd {
		return it
	}

	return it.AddArgs(args...)
}

func (it osTypeScriptBuilder) Unix() ScriptOnceBuilder {
	return it.itemsMap[osmixtype.Unix]
}

func (it osTypeScriptBuilder) Windows() ScriptOnceBuilder {
	return it.itemsMap[osmixtype.Windows]
}

func (it osTypeScriptBuilder) Linux() ScriptOnceBuilder {
	return it.itemsMap[osmixtype.Linux]
}

func (it osTypeScriptBuilder) MacOs() ScriptOnceBuilder {
	return it.itemsMap[osmixtype.MacOs]
}

func (it osTypeScriptBuilder) Ubuntu() ScriptOnceBuilder {
	return it.itemsMap[osmixtype.Ubuntu]
}

func (it osTypeScriptBuilder) CentOs() ScriptOnceBuilder {
	return it.itemsMap[osmixtype.Centos]
}

func (it *osTypeScriptBuilder) GetBy(
	osMixType osmixtype.Variant,
) ScriptOnceBuilder {
	if it == nil {
		return nil
	}

	return it.itemsMap[osMixType]
}

// GetWithStat
//
//	isDefined : found and not null
func (it *osTypeScriptBuilder) GetWithStat(
	osMixType osmixtype.Variant,
) (scriptOnceBuilder ScriptOnceBuilder, isDefined bool) {
	if it == nil {
		return nil, false
	}

	builder, has := it.itemsMap[osMixType]

	return builder, has && builder != nil
}

func (it *osTypeScriptBuilder) HasBuilder(
	mixType osmixtype.Variant,
) bool {
	if it.IsEmpty() {
		return false
	}

	_, has := it.itemsMap[mixType]

	return has
}

// IsValidBuilder
//
//	 return it != nil &&
//			it.scriptType.IsValid() &&
//			!it.HasError() &&
//			it.scriptLines.HasAnyItem()
func (it *osTypeScriptBuilder) IsValidBuilder(
	mixType osmixtype.Variant,
) bool {
	if it.IsEmpty() {
		return false
	}

	builder, has := it.itemsMap[mixType]

	return has &&
		builder != nil &&
		builder.IsDefined()
}

// IsInvalidBuilder
//
//	invert of IsValidBuilder
func (it *osTypeScriptBuilder) IsInvalidBuilder(
	mixType osmixtype.Variant,
) bool {
	return !it.IsValidBuilder(mixType)
}

// IsAllBuildersValid
//
//	represents that all builder exist
//	with lines and script type is valid
func (it *osTypeScriptBuilder) IsAllBuildersValid(
	mixTypes ...osmixtype.Variant,
) bool {
	if it.IsEmpty() {
		return len(mixTypes) == 0
	}

	for _, mixType := range mixTypes {
		if it.IsInvalidBuilder(mixType) {
			return false
		}
	}

	return true
}

func (it *osTypeScriptBuilder) IsAnyInvalidBuilders(
	mixTypes ...osmixtype.Variant,
) bool {
	if it.IsEmpty() {
		return len(mixTypes) == 0
	}

	for _, mixType := range mixTypes {
		if it.IsInvalidBuilder(mixType) {
			return true
		}
	}

	return false
}

func (it *osTypeScriptBuilder) IsBuilderMissing(
	mixType osmixtype.Variant,
) bool {
	if it.IsEmpty() {
		return true
	}

	_, has := it.itemsMap[mixType]

	return !has
}

func (it *osTypeScriptBuilder) IsBuilderMissingOrInvalid(
	mixType osmixtype.Variant,
) bool {
	if it.IsEmpty() {
		return true
	}

	builder, has := it.itemsMap[mixType]

	return !(has && builder != nil && builder.IsValid())
}

func (it *osTypeScriptBuilder) HasBuilderScriptLines(
	mixType osmixtype.Variant,
) bool {
	if it.IsEmpty() {
		return true
	}

	builder, has := it.itemsMap[mixType]

	return has && builder != nil && builder.HasAnyItem()
}

func (it *osTypeScriptBuilder) GetBuilder(
	mixType osmixtype.Variant,
) (scriptBuilder ScriptOnceBuilder, hasBuilder bool) {
	if it.IsEmpty() {
		return nil, false
	}

	scriptBuilder, hasBuilder = it.itemsMap[mixType]

	return scriptBuilder, hasBuilder
}

func (it *osTypeScriptBuilder) HasAllBuilder(
	mixTypes ...osmixtype.Variant,
) bool {
	if it.IsEmpty() {
		return len(mixTypes) == 0
	}

	for _, mixType := range mixTypes {
		_, has := it.itemsMap[mixType]

		if !has {
			return false
		}
	}

	return true
}

func (it *osTypeScriptBuilder) HasAnyBuilder(
	mixTypes ...osmixtype.Variant,
) bool {
	if it.IsEmpty() {
		return len(mixTypes) == 0
	}

	for _, mixType := range mixTypes {
		_, has := it.itemsMap[mixType]

		if has {
			return true
		}
	}

	return false
}

func (it *osTypeScriptBuilder) HasUnix() bool {
	return it.HasBuilder(osmixtype.Unix)
}

func (it *osTypeScriptBuilder) HasWindows() bool {
	return it.HasBuilder(osmixtype.Windows)
}

func (it *osTypeScriptBuilder) HasLinux() bool {
	return it.HasBuilder(osmixtype.Linux)
}

func (it *osTypeScriptBuilder) IsEmptyUnix() bool {
	return it.IsBuilderMissing(osmixtype.Unix)
}

func (it *osTypeScriptBuilder) IsEmptyWindows() bool {
	return it.IsBuilderMissing(osmixtype.Windows)
}

func (it *osTypeScriptBuilder) IsEmptyLinux() bool {
	return it.IsBuilderMissing(osmixtype.Linux)
}

func (it *osTypeScriptBuilder) IsEmptyMacOs() bool {
	return it.IsBuilderMissing(osmixtype.MacOs)
}

func (it *osTypeScriptBuilder) IsEmptyBy(osMixType osmixtype.Variant) bool {
	builder := it.GetBy(osMixType)

	return builder.IsEmpty()
}

func (it *osTypeScriptBuilder) HasAnyItemBy(osMixType osmixtype.Variant) bool {
	builder := it.GetBy(osMixType)

	return builder.IsEmpty()
}

func (it *osTypeScriptBuilder) LoopInvokeFunc(processorFunc ScriptBuilderWithTypeProcessorFunc) {
	for osType, builder := range it.ListMap() {
		processorFunc(osType, builder)
	}
}

func (it *osTypeScriptBuilder) Filter(
	filter FilterScriptBuilderFunc,
) (resultMap map[osmixtype.Variant]ScriptOnceBuilder) {
	for osType, builder := range it.ListMap() {
		isTake, isBreak := filter(osType, builder)

		if isTake {
			resultMap[osType] = builder
		}

		if isBreak {
			return resultMap
		}
	}

	return resultMap
}

func (it *osTypeScriptBuilder) FilterOsMixTypes(
	filter FilterScriptBuilderFunc,
) (osMixTypes []osmixtype.Variant) {
	for osType, builder := range it.ListMap() {
		isTake, isBreak := filter(osType, builder)

		if isTake {
			osMixTypes = append(
				osMixTypes,
				osType)
		}

		if isBreak {
			return osMixTypes
		}
	}

	return osMixTypes
}

func (it *osTypeScriptBuilder) GetBuildersMapByOsMixTypes(
	osMixTypes []osmixtype.Variant,
) (resultMap map[osmixtype.Variant]ScriptOnceBuilder) {
	allItemsMap := it.ListMap()

	for _, osType := range osMixTypes {
		scriptBuilder, has := allItemsMap[osType]

		if has {
			resultMap[osType] = scriptBuilder
		}
	}

	return resultMap
}

func (it osTypeScriptBuilder) UnixBuild() *CmdOnce {
	return it.Unix().Build()
}

func (it osTypeScriptBuilder) WindowsBuild() *CmdOnce {
	return it.Windows().Build()
}

func (it osTypeScriptBuilder) LinuxBuild() *CmdOnce {
	return it.Linux().Build()
}

func (it osTypeScriptBuilder) MacOsBuild() *CmdOnce {
	return it.MacOs().Build()
}

func (it osTypeScriptBuilder) UnixBuildCmd() *exec.Cmd {
	return it.Unix().BuildCmd()
}

func (it osTypeScriptBuilder) WindowsBuildCmd() *exec.Cmd {
	return it.Windows().BuildCmd()
}

func (it osTypeScriptBuilder) LinuxBuildCmd() *exec.Cmd {
	return it.Linux().BuildCmd()
}

func (it osTypeScriptBuilder) MacOsBuildCmd() *exec.Cmd {
	return it.MacOs().BuildCmd()
}

func (it *osTypeScriptBuilder) List() []ScriptOnceBuilder {
	if it == nil {
		return nil
	}

	slice := make([]ScriptOnceBuilder, constants.Capacity4)

	index := -1
	for _, builder := range it.itemsMap {
		index++

		slice[index] = builder
	}

	return slice
}

func (it *osTypeScriptBuilder) ListMap() map[osmixtype.Variant]ScriptOnceBuilder {
	if it == nil {
		return nil
	}

	newMap := make(
		map[osmixtype.Variant]ScriptOnceBuilder,
		it.Length())

	for osMixType, builder := range it.itemsMap {
		newMap[osMixType] = builder
	}

	return newMap
}

func (it *osTypeScriptBuilder) IsNull() bool {
	return it == nil
}

func (it osTypeScriptBuilder) AvailableBuilders() []ScriptOnceBuilder {
	slice := make([]ScriptOnceBuilder, 0, constants.Capacity4)

	index := -1
	for _, builder := range it.itemsMap {
		index++

		if builder == nil || builder.IsEmpty() {
			continue
		}

		slice[index] = builder
	}

	return slice
}

func (it osTypeScriptBuilder) BuildersMap() map[osmixtype.Variant]ScriptOnceBuilder {
	return it.ListMap()
}

func (it osTypeScriptBuilder) AvailableBuildersMap() map[osmixtype.Variant]ScriptOnceBuilder {
	newMap := make(
		map[osmixtype.Variant]ScriptOnceBuilder,
		it.Length())

	for osMixType, builder := range it.itemsMap {
		if builder == nil || builder.IsEmpty() {
			continue
		}

		newMap[osMixType] = builder
	}

	return newMap
}

func (it osTypeScriptBuilder) ExecutionResultMap() map[osmixtype.Variant]*Result {
	availableMap := it.AvailableBuildersMap()
	resultsMap := make(
		map[osmixtype.Variant]*Result,
		len(availableMap))

	for osMixType, builder := range availableMap {
		resultsMap[osMixType] = builder.Result()
	}

	return resultsMap
}

func (it osTypeScriptBuilder) BuildMap() map[osmixtype.Variant]*CmdOnce {
	availableMap := it.AvailableBuildersMap()
	resultsMap := make(
		map[osmixtype.Variant]*CmdOnce,
		len(availableMap))

	for osMixType, builder := range availableMap {
		resultsMap[osMixType] = builder.Build()
	}

	return resultsMap
}

func (it osTypeScriptBuilder) BuildMapOnly(
	osMixes ...osmixtype.Variant,
) map[osmixtype.Variant]*CmdOnce {
	availableMap := it.AvailableBuildersMap()
	resultsMap := make(map[osmixtype.Variant]*CmdOnce, len(osMixes))

	for _, osMixType := range osMixes {
		builder, has := availableMap[osMixType]

		if !has {
			continue
		}

		resultsMap[osMixType] = builder.Build()
	}

	return resultsMap
}

func (it osTypeScriptBuilder) ResultMapOnly(
	osMixes ...osmixtype.Variant,
) map[osmixtype.Variant]*Result {
	availableMap := it.AvailableBuildersMap()
	resultsMap := make(map[osmixtype.Variant]*Result, len(osMixes))

	for _, osMixType := range osMixes {
		builder, has := availableMap[osMixType]

		if !has {
			continue
		}

		resultsMap[osMixType] = builder.Result()
	}

	return resultsMap
}

func (it osTypeScriptBuilder) BuildCmdMapOnly(
	osMixes ...osmixtype.Variant,
) map[osmixtype.Variant]*exec.Cmd {
	availableMap := it.AvailableBuildersMap()
	resultsMap := make(map[osmixtype.Variant]*exec.Cmd, len(osMixes))

	for _, osMixType := range osMixes {
		builder, has := availableMap[osMixType]

		if !has {
			continue
		}

		resultsMap[osMixType] = builder.BuildCmd()
	}

	return resultsMap
}

func (it osTypeScriptBuilder) OutputMapOnly(
	osMixes ...osmixtype.Variant,
) map[osmixtype.Variant]string {
	availableMap := it.AvailableBuildersMap()
	resultsMap := make(map[osmixtype.Variant]string, len(osMixes))

	for _, osMixType := range osMixes {
		builder, has := availableMap[osMixType]

		if !has {
			continue
		}

		resultsMap[osMixType] = builder.Output()
	}

	return resultsMap
}

func (it osTypeScriptBuilder) OutputLinesSimpleSliceMapOnly(
	osMixes ...osmixtype.Variant,
) map[osmixtype.Variant]*corestr.SimpleSlice {
	availableMap := it.AvailableBuildersMap()
	resultsMap := make(map[osmixtype.Variant]*corestr.SimpleSlice, len(osMixes))

	for _, osMixType := range osMixes {
		builder, has := availableMap[osMixType]

		if !has {
			continue
		}

		slice := builder.OutputLinesSimpleSlice()
		if slice.IsEmpty() {
			continue
		}

		resultsMap[osMixType] = builder.OutputLinesSimpleSlice()
	}

	return resultsMap
}

func (it osTypeScriptBuilder) DetailedOutputMapOnly(
	osMixes ...osmixtype.Variant,
) map[osmixtype.Variant]string {
	availableMap := it.AvailableBuildersMap()
	resultsMap := make(map[osmixtype.Variant]string, len(osMixes))

	for _, osMixType := range osMixes {
		builder, has := availableMap[osMixType]

		if !has {
			continue
		}

		resultsMap[osMixType] = builder.DetailedOutput()
	}

	return resultsMap
}

func (it osTypeScriptBuilder) BuildersMapOnly(
	osMixes ...osmixtype.Variant,
) map[osmixtype.Variant]ScriptOnceBuilder {
	buildersMap := it.BuildersMap()
	resultsMap := make(map[osmixtype.Variant]ScriptOnceBuilder, len(osMixes))

	for _, osMixType := range osMixes {
		builder, has := buildersMap[osMixType]

		if !has {
			continue
		}

		resultsMap[osMixType] = builder
	}

	return resultsMap
}

func (it osTypeScriptBuilder) CurrentOsTypes() []osmixtype.Variant {
	return osmixtype.CurrentOsMixTypes()
}

func (it osTypeScriptBuilder) CurrentOsTypesMap() map[osmixtype.Variant]bool {
	return osmixtype.CurrentOsTypesMap()
}

func (it osTypeScriptBuilder) CurrentOsScriptBuilder() CurrentOsScriptBuilder {
	return &currentOsScriptBuilder{
		osScriptBuilder:        &it,
		selectedCurrentOsMixes: it.CurrentOsTypes(),
	}
}

func (it osTypeScriptBuilder) CompiledErrorWrappersMapOnly(
	osMixes ...osmixtype.Variant,
) (compiledMap map[osmixtype.Variant]*errorwrapper.Wrapper, hasAnyError bool) {
	buildersMap := it.BuildersMapOnly(osMixes...)
	resultsMap := make(
		map[osmixtype.Variant]*errorwrapper.Wrapper,
		len(osMixes))

	for osMixType, builder := range buildersMap {
		errWp := builder.CompiledErrorWrapper()
		if errWp.IsEmpty() {
			continue
		}

		resultsMap[osMixType] = errWp
	}

	return resultsMap, len(resultsMap) > 0
}

func (it osTypeScriptBuilder) AllCompiledErrorWrappersMapOnly() (
	compiledMap map[osmixtype.Variant]*errorwrapper.Wrapper, hasAnyError bool,
) {
	resultsMap := make(
		map[osmixtype.Variant]*errorwrapper.Wrapper,
		it.Length())

	for osMixType, builder := range it.itemsMap {
		errWp := builder.CompiledErrorWrapper()
		if errWp.IsEmpty() {
			continue
		}

		resultsMap[osMixType] = errWp
	}

	return resultsMap, len(resultsMap) > 0
}

func (it osTypeScriptBuilder) BuildCmdMap() map[osmixtype.Variant]*exec.Cmd {
	availableMap := it.AvailableBuildersMap()
	resultsMap := make(map[osmixtype.Variant]*exec.Cmd, len(availableMap))

	for osMixType, builder := range availableMap {
		resultsMap[osMixType] = builder.BuildCmd()
	}

	return resultsMap
}

func (it osTypeScriptBuilder) OutputMap() map[osmixtype.Variant]string {
	availableMap := it.AvailableBuildersMap()
	resultsMap := make(map[osmixtype.Variant]string, len(availableMap))

	for osMixType, builder := range availableMap {
		resultsMap[osMixType] = builder.Output()
	}

	return resultsMap
}

func (it osTypeScriptBuilder) OutputBytesMap() map[osmixtype.Variant][]byte {
	availableMap := it.AvailableBuildersMap()
	resultsMap := make(map[osmixtype.Variant][]byte, len(availableMap))

	for osMixType, builder := range availableMap {
		resultsMap[osMixType] = builder.OutputBytes()
	}

	return resultsMap
}

func (it osTypeScriptBuilder) CombinedOutputMap() map[osmixtype.Variant]coredata.BytesError {
	availableMap := it.AvailableBuildersMap()
	resultsMap := make(map[osmixtype.Variant]coredata.BytesError, len(availableMap))

	for osMixType, builder := range availableMap {
		allBytes, err := builder.CombinedOutput()

		resultsMap[osMixType] = coredata.BytesError{
			Bytes: allBytes,
			Error: err,
		}
	}

	return resultsMap
}

func (it osTypeScriptBuilder) AsyncExecutionResultMap() map[osmixtype.Variant]*Result {
	availableMap := it.AvailableBuildersMap()
	resultsMap := make(map[osmixtype.Variant]*Result, len(availableMap))
	wg := sync.WaitGroup{}
	wg.Add(len(availableMap))

	asyncFunc := func(
		osMixType osmixtype.Variant,
	) {
		resultsMap[osMixType] = availableMap[osMixType].Result()

		wg.Done()
	}

	for osMixType := range availableMap {
		go asyncFunc(osMixType)
	}

	wg.Wait()

	return resultsMap
}

// StandardOutputCmd
//
//	Set all builders input output to standard os input output
//
//	os.Stdin, os.Stdout, os.Stderr
func (it *osTypeScriptBuilder) StandardOutputCmd() OsMixTypeScriptBuilder {
	for _, builder := range it.List() {
		builder.StandardOutputCmd()
	}

	return it
}

func (it osTypeScriptBuilder) AsyncStart(stdIn, stdOut, stderr *bytes.Buffer) (
	map[osmixtype.Variant]error, *sync.WaitGroup,
) {
	availableMap := it.AvailableBuildersMap()
	resultsMap := make(map[osmixtype.Variant]error, len(availableMap))
	wg := sync.WaitGroup{}
	wg.Add(len(availableMap))

	asyncFunc := func(
		osMixType osmixtype.Variant,
	) {
		builder, has := availableMap[osMixType]

		if !has {
			wg.Done()

			return
		}

		builder.SetInputOutput(stdIn, stdOut, stderr)
		err := builder.BuildCmd().Run()
		resultsMap[osMixType] = err

		wg.Done()
	}

	for osMixType := range availableMap {
		go asyncFunc(osMixType)
	}

	return resultsMap, &wg
}

func (it osTypeScriptBuilder) OutputMapMust() (resultMap map[osmixtype.Variant]string) {
	for osMixType, builder := range it.AvailableBuildersMap() {
		result := builder.Result()
		result.HandleError()
		resultMap[osMixType] = result.OutputString()
	}

	return resultMap
}

func (it osTypeScriptBuilder) TrimmedOutputMap() (resultMap map[osmixtype.Variant]string) {
	for osMixType, builder := range it.AvailableBuildersMap() {
		resultMap[osMixType] = builder.TrimmedOutput()
	}

	return resultMap
}

func (it osTypeScriptBuilder) OutputLinesMap() (resultMap map[osmixtype.Variant][]string) {
	for osMixType, builder := range it.AvailableBuildersMap() {
		resultMap[osMixType] = builder.OutputLines()
	}

	return resultMap
}

func (it osTypeScriptBuilder) ErrorOutputMap() (resultMap map[osmixtype.Variant]string) {
	for osMixType, builder := range it.AvailableBuildersMap() {
		resultMap[osMixType] = builder.ErrorOutput()
	}

	return resultMap
}

func (it osTypeScriptBuilder) TrimmedOutputLinesMap() (resultMap map[osmixtype.Variant][]string) {
	for osMixType, builder := range it.AvailableBuildersMap() {
		resultMap[osMixType] = builder.OutputLines()
	}

	return resultMap
}

func (it osTypeScriptBuilder) BuildClearMap() (resultMap map[osmixtype.Variant]*CmdOnce) {
	for osMixType, builder := range it.AvailableBuildersMap() {
		resultMap[osMixType] = builder.BuildClear()
	}

	return resultMap
}

func (it osTypeScriptBuilder) BuildDisposeMap() (resultMap map[osmixtype.Variant]*CmdOnce) {
	for osMixType, builder := range it.AvailableBuildersMap() {
		resultMap[osMixType] = builder.BuildDispose()
	}

	return resultMap
}

// HasAnyItem
//
// Only counts and checks available builders, O(N), expensive
func (it osTypeScriptBuilder) HasAnyItem() bool {
	return len(it.AvailableBuilders()) > 0
}

// IsEmpty
//
// Only counts and checks available builders, O(N), expensive
func (it osTypeScriptBuilder) IsEmpty() bool {
	return len(it.AvailableBuilders()) == 0
}

func (it osTypeScriptBuilder) Log() {
	buildersMap := it.ListMap()

	for osMixType, builder := range buildersMap {
		fmt.Println(osMixType.Name() + constants.SpaceColonSpace)
		builder.Log()
	}
}

func (it osTypeScriptBuilder) LogWithTraces() {
	buildersMap := it.ListMap()

	for osMixType, builder := range buildersMap {
		fmt.Println(osMixType.Name() + constants.SpaceColonSpace)
		builder.LogWithTraces()
	}
}

func (it osTypeScriptBuilder) LogFatal() {
	buildersMap := it.ListMap()

	for osMixType, builder := range buildersMap {
		fmt.Println(osMixType.Name() + constants.SpaceColonSpace)
		builder.LogFatal()
	}
}

func (it osTypeScriptBuilder) LogFatalWithTraces() {
	buildersMap := it.ListMap()

	for osMixType, builder := range buildersMap {
		fmt.Println(osMixType.Name() + constants.SpaceColonSpace)
		builder.LogFatalWithTraces()
	}
}

func (it osTypeScriptBuilder) LogIf(isLog bool) {
	if !isLog {
		return
	}

	it.Log()
}

func (it *osTypeScriptBuilder) Length() int {
	if it == nil {
		return 0
	}

	return len(it.itemsMap)
}

// ValidLength
//
// Only counts and checks available builders, O(N), expensive
func (it *osTypeScriptBuilder) ValidLength() int {
	if it == nil {
		return 0
	}

	return len(it.AvailableBuilders())
}

func (it osTypeScriptBuilder) String() string {
	newMap := it.StringsMap()

	return corejson.Serialize.ToString(newMap)
}

func (it osTypeScriptBuilder) StringsMap() (resultMap map[osmixtype.Variant][]string) {
	for osMixType, builder := range it.AvailableBuildersMap() {
		resultMap[osMixType] = builder.OutputLines()
	}

	return resultMap
}

func (it osTypeScriptBuilder) AsOsMixTypeScriptBuilder() OsMixTypeScriptBuilder {
	return &it
}
