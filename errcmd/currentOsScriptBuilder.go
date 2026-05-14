package errcmd

import (
	"fmt"
	"os/exec"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/enum/osmixtype"
)

type currentOsScriptBuilder struct {
	osScriptBuilder        *osTypeScriptBuilder
	selectedCurrentOsMixes []osmixtype.Variant
}

func (it currentOsScriptBuilder) AvailableLength() int {
	if it.osScriptBuilder.IsEmpty() {
		return 0
	}

	count := 0

	for _, mixType := range it.selectedCurrentOsMixes {
		if it.osScriptBuilder.HasBuilder(mixType) {
			count++
		}
	}

	return count
}

func (it currentOsScriptBuilder) CurrentOsTypesMap() map[osmixtype.Variant]bool {
	return osmixtype.CurrentOsMixTypesMap()
}

// IsEmpty
//
// Only counts and checks available builders, O(N), expensive
func (it currentOsScriptBuilder) IsEmpty() bool {
	return len(it.selectedCurrentOsMixes) == 0 &&
		it.AvailableLength() == 0
}

// HasError
//
//  Very expensive runs all operations
func (it currentOsScriptBuilder) HasError() bool {
	if it.IsEmpty() {
		return false
	}

	_, hasAnyErr := it.
		osScriptBuilder.
		CompiledErrorWrappersMapOnly(
			it.selectedCurrentOsMixes...)

	return hasAnyErr
}

// HasAnyItem
//
// Only counts and checks available builders, O(N), expensive
func (it currentOsScriptBuilder) HasAnyItem() bool {
	return it.AvailableLength() > 0
}

func (it currentOsScriptBuilder) HasAnyIssues() bool {
	return it.HasError()
}

func (it currentOsScriptBuilder) ValidationError() error {
	return it.Error()
}

// Error
//
//  Very expensive operation
func (it currentOsScriptBuilder) Error() error {
	errorWrappersMapped, hasAnyErr := it.
		osScriptBuilder.
		CompiledErrorWrappersMapOnly(
			it.selectedCurrentOsMixes...)

	if !hasAnyErr {
		return nil
	}

	return it.
		osScriptBuilder.
		compiledMappedErrorWrappersToError(
			errorWrappersMapped)
}

func (it currentOsScriptBuilder) IsValid() bool {
	return !it.HasError()
}

func (it currentOsScriptBuilder) IsSuccess() bool {
	return !it.HasError()
}

func (it currentOsScriptBuilder) IsFailed() bool {
	return it.HasError()
}

func (it currentOsScriptBuilder) Log() {
	for _, osMixType := range it.selectedCurrentOsMixes {
		fmt.Println(osMixType.Name() + constants.SpaceColonSpace)
		builder, isValid := it.osScriptBuilder.
			GetWithStat(osMixType)

		if isValid {
			builder.Log()
		} else {
			// invalid
			fmt.Println("Empty or invalid.")
		}
	}
}

func (it currentOsScriptBuilder) LogWithTraces() {
	for _, osMixType := range it.selectedCurrentOsMixes {
		fmt.Println(osMixType.Name() + constants.SpaceColonSpace)
		builder, isValid := it.osScriptBuilder.
			GetWithStat(osMixType)

		if isValid {
			builder.LogWithTraces()
		} else {
			// invalid
			fmt.Println("Empty or invalid.")
		}
	}
}

func (it currentOsScriptBuilder) LogFatal() {
	for _, osMixType := range it.selectedCurrentOsMixes {
		fmt.Println(osMixType.Name() + constants.SpaceColonSpace)
		builder, isValid := it.osScriptBuilder.
			GetWithStat(osMixType)

		if isValid {
			builder.LogFatal()
		} else {
			// invalid
			fmt.Println("Empty or invalid.")
		}
	}
}

func (it currentOsScriptBuilder) LogFatalWithTraces() {
	for _, osMixType := range it.selectedCurrentOsMixes {
		fmt.Println(osMixType.Name() + constants.SpaceColonSpace)
		builder, isValid := it.osScriptBuilder.
			GetWithStat(osMixType)

		if isValid {
			builder.LogFatalWithTraces()
		} else {
			// invalid
			fmt.Println("Empty or invalid.")
		}
	}
}

func (it currentOsScriptBuilder) LogIf(isLog bool) {
	if !isLog {
		return
	}

	it.Log()
}

func (it currentOsScriptBuilder) Length() int {
	return len(it.selectedCurrentOsMixes)
}

func (it currentOsScriptBuilder) CurrentOsTypes() []osmixtype.Variant {
	return it.selectedCurrentOsMixes
}

func (it currentOsScriptBuilder) BuildersMap() map[osmixtype.Variant]ScriptOnceBuilder {
	currentOsTypes := it.CurrentOsTypes()

	return it.osScriptBuilder.BuildersMapOnly(currentOsTypes...)
}

func (it currentOsScriptBuilder) ResultsMap() map[osmixtype.Variant]*Result {
	currentOsTypes := it.CurrentOsTypes()

	return it.osScriptBuilder.ResultMapOnly(currentOsTypes...)
}

func (it currentOsScriptBuilder) CmdOnceMap() map[osmixtype.Variant]*CmdOnce {
	currentOsTypes := it.CurrentOsTypes()

	return it.osScriptBuilder.BuildMapOnly(currentOsTypes...)
}

func (it currentOsScriptBuilder) OutputMap() map[osmixtype.Variant]string {
	currentOsTypes := it.CurrentOsTypes()

	return it.osScriptBuilder.OutputMapOnly(currentOsTypes...)
}

func (it currentOsScriptBuilder) DetailedMap() map[osmixtype.Variant]string {
	currentOsTypes := it.CurrentOsTypes()

	return it.osScriptBuilder.DetailedOutputMapOnly(currentOsTypes...)
}

func (it currentOsScriptBuilder) CmdMap() map[osmixtype.Variant]*exec.Cmd {
	currentOsTypes := it.CurrentOsTypes()

	return it.osScriptBuilder.BuildCmdMapOnly(currentOsTypes...)
}

func (it currentOsScriptBuilder) OutputLinesSimpleSliceMap() map[osmixtype.Variant]*corestr.SimpleSlice {
	currentOsTypes := it.CurrentOsTypes()

	return it.osScriptBuilder.OutputLinesSimpleSliceMapOnly(currentOsTypes...)
}

func (it currentOsScriptBuilder) AsCurrentOsScriptBuilder() CurrentOsScriptBuilder {
	return &it
}
