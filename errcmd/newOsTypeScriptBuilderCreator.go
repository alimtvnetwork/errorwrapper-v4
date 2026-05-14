package errcmd

import (
	"github.com/alimtvnetwork/core-v9/constants"
	"github.com/alimtvnetwork/enum-v10/osmixtype"
	"github.com/alimtvnetwork/enum-v10/scripttype"
)

type newOsTypeScriptBuilderCreator struct{}

func (it newOsTypeScriptBuilderCreator) Default(
	taskName, desc, url string,
	capacity int,
) OsMixTypeScriptBuilder {
	return &osTypeScriptBuilder{
		taskName:    taskName,
		description: desc,
		url:         url,
		itemsMap: make(
			map[osmixtype.Variant]*scriptOnceBuilder,
			capacity),
	}
}

func (it newOsTypeScriptBuilderCreator) EmptyWithNames(
	taskName, desc, url string,
) OsMixTypeScriptBuilder {
	return &osTypeScriptBuilder{
		taskName:    taskName,
		description: desc,
		url:         url,
		itemsMap:    map[osmixtype.Variant]*scriptOnceBuilder{},
	}
}

func (it newOsTypeScriptBuilderCreator) Empty() OsMixTypeScriptBuilder {
	return &osTypeScriptBuilder{
		itemsMap: map[osmixtype.Variant]*scriptOnceBuilder{},
	}
}

func (it newOsTypeScriptBuilderCreator) Create(
	taskName, desc, url string,
	osMixType osmixtype.Variant,
	scriptType scripttype.Variant,
) OsMixTypeScriptBuilder {
	osBuildersMap := &osTypeScriptBuilder{
		taskName:    taskName,
		description: desc,
		url:         url,
		itemsMap: make(
			map[osmixtype.Variant]*scriptOnceBuilder,
			constants.Capacity5),
	}

	return osBuildersMap.
		AddNewBuilder(osMixType, scriptType)
}

func (it newOsTypeScriptBuilderCreator) UnixUbuntu(
	taskName, desc, url string,
	unix, ubuntu ScriptOnceBuilder,
) OsMixTypeScriptBuilder {
	osBuildersMap := &osTypeScriptBuilder{
		taskName:    taskName,
		description: desc,
		url:         url,
		itemsMap: make(
			map[osmixtype.Variant]*scriptOnceBuilder,
			constants.Capacity5),
	}

	if unix != nil {
		osBuildersMap.AddBuilder(osmixtype.Unix, unix)
	}

	if ubuntu != nil {
		osBuildersMap.AddBuilder(osmixtype.Ubuntu, ubuntu)
	}

	return osBuildersMap
}

func (it newOsTypeScriptBuilderCreator) WindowsLinux(
	taskName, desc, url string,
	windows, linux ScriptOnceBuilder,
) OsMixTypeScriptBuilder {
	osBuildersMap := &osTypeScriptBuilder{
		taskName:    taskName,
		description: desc,
		url:         url,
		itemsMap: make(
			map[osmixtype.Variant]*scriptOnceBuilder,
			constants.Capacity5),
	}

	if windows != nil {
		osBuildersMap.AddBuilder(osmixtype.Windows, windows)
	}

	if linux != nil {
		osBuildersMap.AddBuilder(osmixtype.Linux, linux)
	}

	return osBuildersMap
}

func (it newOsTypeScriptBuilderCreator) WindowsLinuxMacOs(
	taskName, desc, url string,
	windows, linux, macOs ScriptOnceBuilder,
) OsMixTypeScriptBuilder {
	windowsLinuxMap := it.WindowsLinux(
		taskName,
		desc,
		url,
		windows,
		linux,
	)

	return windowsLinuxMap.AddBuilder(osmixtype.MacOs, macOs)
}

func (it newOsTypeScriptBuilderCreator) WindowsUbuntu(
	taskName, desc, url string,
	windows, ubuntu ScriptOnceBuilder,
) OsMixTypeScriptBuilder {
	windowsLinuxMap := it.WindowsLinux(
		taskName,
		desc,
		url,
		windows,
		nil,
	)

	return windowsLinuxMap.AddBuilder(osmixtype.Ubuntu, ubuntu)
}
