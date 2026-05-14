package errtype

var (
	errTypesMap = map[Variation]VariantStructure{
		Generic: {
			Name:    "Generic",
			Message: "Generic error",
			Variant: Generic,
		},
		NullOrEmpty: {
			Name:    "NullOrEmpty",
			Message: "Null or empty reference.",
			Variant: NullOrEmpty,
		},
		Null: {
			Name:    "Null",
			Message: "Null reference.",
			Variant: Null,
		},
		Empty: {
			Name:    "Empty",
			Message: "Empty value or empty object or empty reference.",
			Variant: Empty,
		},
		Unknown: {
			Name:    "Unknown",
			Message: "Unknown value / result / status / data or un-predictable.",
			Variant: Unknown,
		},
		// Input output error
		IO: {
			Name:    "IO",
			Message: "Input / output error.",
			Variant: IO,
		},
		// Read write error
		ReadWrite: {
			Name:    "ReadWrite",
			Message: "Read or Write error or ReadWrite combined error.",
			Variant: ReadWrite,
		},
		// Data is not as expected
		InvalidInput: {
			Name:    "InvalidInput",
			Message: "Invalid: Input error.",
			Variant: InvalidInput,
		},
		InvalidOutput: {
			Name:    "InvalidOutput",
			Message: "Invalid: Output error.",
			Variant: InvalidOutput,
		},
		InvalidCondition: {
			Name:    "InvalidCondition",
			Message: "Invalid: Condition error.",
			Variant: InvalidCondition,
		},
		EmptyString: {
			Name:    "EmptyString",
			Message: "Empty string error.",
			Variant: EmptyString,
		},
		EmptySlice: {
			Name:    "EmptySlice",
			Message: "Empty slice error.",
			Variant: EmptySlice,
		},
		EmptyArray: {
			Name:    "EmptyArray",
			Message: "Empty array error.",
			Variant: EmptyArray,
		},
		EmptyMap: {
			Name:    "EmptyMap",
			Message: "Empty map error.",
			Variant: EmptyMap,
		},
		EmptyResult: {
			Name:    "EmptyResult",
			Message: "Empty result error.",
			Variant: EmptyResult,
		},
		UnexpectedInteger: {
			Name:    "UnexpectedInteger",
			Message: "Unexpected: integer error.",
			Variant: UnexpectedInteger,
		},
		UnexpectedString: {
			Name:    "UnexpectedString",
			Message: "Unexpected: string error.",
			Variant: UnexpectedString,
		},
		UnexpectedFloat: {
			Name:    "UnexpectedFloat",
			Message: "Unexpected: float error.",
			Variant: UnexpectedFloat,
		},
		UnexpectedValue: {
			Name:    "UnexpectedValue",
			Message: "Unexpected: value error.",
			Variant: UnexpectedValue,
		},
		UnexpectedArray: {
			Name:    "UnexpectedArray",
			Message: "Unexpected: array error.",
			Variant: UnexpectedArray,
		},
		UnexpectedArrayLength: {
			Name:    "UnexpectedArrayLength",
			Message: "Unexpected: array length error.",
			Variant: UnexpectedArrayLength,
		},
		UnexpectedSliceLength: {
			Name:    "UnexpectedSliceLength",
			Message: "Unexpected: slice length error.",
			Variant: UnexpectedSliceLength,
		},
		UnexpectedMapLength: {
			Name:    "UnexpectedMapLength",
			Message: "Unexpected: map length error.",
			Variant: UnexpectedMapLength,
		},
		UnexpectedInterface: {
			Name:    "UnexpectedInterface",
			Message: "Unexpected: interface error.",
			Variant: UnexpectedInterface,
		},
		UnexpectedType: {
			Name:    "UnexpectedType",
			Message: "Unexpected: type error.",
			Variant: UnexpectedType,
		},
		UnexpectedPointer: {
			Name:    "UnexpectedPointer",
			Message: "Unexpected: pointer error.",
			Variant: UnexpectedPointer,
		},
		UnexpectedStatus: {
			Name:    "UnexpectedStatus",
			Message: "Unexpected: status error.",
			Variant: UnexpectedStatus,
		},
		UnexpectedFilePath: {
			Name:    "UnexpectedFilePath",
			Message: "Unexpected: file path error.",
			Variant: UnexpectedFilePath,
		},
		UnexpectedDirectory: {
			Name:    "UnexpectedDirectory",
			Message: "Given directory path is unexpected.",
			Variant: UnexpectedDirectory,
		},
		UnexpectedCondition: {
			Name:    "UnexpectedCondition",
			Message: "Given condition is unexpected.",
			Variant: UnexpectedCondition,
		},
		OutOfRangeValue: {
			Name:    "OutOfRangeValue",
			Message: "Given value is out of range.",
			Variant: OutOfRangeValue,
		},
		OutOfRangeInteger: {
			Name:    "OutOfRangeInteger",
			Message: "Given integer is out of range.",
			Variant: OutOfRangeInteger,
		},
		OutOfRangeString: {
			Name:    "OutOfRangeString",
			Message: "Given string is out of range.",
			Variant: OutOfRangeString,
		},
		OutOfRangeFloat: {
			Name:    "OutOfRangeFloat",
			Message: "Given floating point number is out of range.",
			Variant: OutOfRangeFloat,
		},
		OutOfRangeType: {
			Name:    "OutOfRangeType",
			Message: "Given type format is out of range.",
			Variant: OutOfRangeType,
		},
		OutOfRangeCategory: {
			Name:    "OutOfRangeCategory",
			Message: "Given category is out of range.",
			Variant: OutOfRangeCategory,
		},
		OutOfRangeLength: {
			Name:    "OutOfRangeLength",
			Message: "Given data length is out of range.",
			Variant: OutOfRangeLength,
		},
		OutOfRangeStatus: {
			Name:    "OutOfRangeStatus",
			Message: "Given status value is out of range.",
			Variant: OutOfRangeStatus,
		},
		Warning: {
			Name:    "Warning",
			Message: "Warning",
			Variant: Warning,
		},
		DevelopmentEnvironment: {
			Name:    "DevelopmentEnvironment",
			Message: "Halted in development environment only.",
			Variant: DevelopmentEnvironment,
		},
		ProductionEnvironment: {
			Name:    "ProductionEnvironment",
			Message: "Halted in production environment only.",
			Variant: ProductionEnvironment,
		},
		TestEnvironment: {
			Name:    "TestEnvironment",
			Message: "Halted in test environment only.",
			Variant: TestEnvironment,
		},
		AssertionFailed: {
			Name:    "AssertionFailed",
			Message: "Assertion failed for the given case or case(s).",
			Variant: AssertionFailed,
		},
		EmptyStatus: {
			Name:    "EmptyStatus",
			Message: "Given empty status is unexpected.",
			Variant: EmptyStatus,
		},
		FileNotExist: {
			Name:    "FileNotExist",
			Message: "File does not exist at the path given.",
			Variant: FileNotExist,
		},
		File: {
			Name:    "File",
			Message: "File error, something went wrong with the file.",
			Variant: File,
		},
		FileWrite: {
			Name:    "FileWrite",
			Message: "Failed to write in file.",
			Variant: FileWrite,
		},
		FileAccess: {
			Name:    "FileAccess",
			Message: "Failed to access the file.",
			Variant: FileAccess,
		},
		FileAppend: {
			Name:    "FileAppend",
			Message: "Failed to append text or bytes in the file.",
			Variant: FileAppend,
		},
		DirectoryNotExist: {
			Name:    "DirectoryNotExist",
			Message: "Directory does not exist at the given path.",
			Variant: DirectoryNotExist,
		},
		Directory: {
			Name:    "Directory",
			Message: "Directory has an issue, cannot process the request.",
			Variant: Directory,
		},
		DirectoryAccess: {
			Name:    "DirectoryAccess",
			Message: "Directory access has an issue, cannot process the request.",
			Variant: DirectoryAccess,
		},
		Execution: {
			Name:    "Execution",
			Message: "Execution issue, cannot process the request.",
			Variant: Execution,
		},
		Process: {
			Name:    "Process",
			Message: "Process issue, cannot process the request.",
			Variant: Process,
		},
		LockFailed: {
			Name:    "LockFailed",
			Message: "Lock failed issue, cannot process the request.",
			Variant: LockFailed,
		},
		EmptyContent: {
			Name:    "EmptyContent",
			Message: "Empty content issue, cannot process the request.",
			Variant: EmptyContent,
		},
		EmptyFile: {
			Name:    "EmptyFile",
			Message: "Empty file issue, cannot process the request.",
			Variant: EmptyFile,
		},
		EmptyDirectory: {
			Name:    "EmptyDirectory",
			Message: "Empty directory issue, cannot process the request.",
			Variant: EmptyDirectory,
		},
		NotContainsExpectation: {
			Name:    "NotContainsExpectation",
			Message: "Data is not contained in the collection, cannot process the request. Data should be in the collection.",
			Variant: NotContainsExpectation,
		},
		ContainsExpectation: {
			Name:    "ContainsExpectation",
			Message: "Data is contained in the collection, cannot process the request. Expectation shall not be contained in the collection.",
			Variant: ContainsExpectation,
		},
		EqualExpectation: {
			Name:    "EqualExpectation",
			Message: "Data is equal which is unexpected, cannot process the request. Expectation shall not be equal.",
			Variant: EqualExpectation,
		},
		NotEqualExpectation: {
			Name:    "NotEqualExpectation",
			Message: "Data is not equal which is unexpected, thus cannot process the request. Expectation data should be equal.",
			Variant: NotEqualExpectation,
		},
		LessThanExpectation: {
			Name:    "LessThanExpectation",
			Message: "Data is less than the reference which is unexpected, cannot process the request. Expectation shall not be less than the reference value.",
			Variant: LessThanExpectation,
		},
		LessThanEqualsExpectation: {
			Name:    "LessThanEqualsExpectation",
			Message: "Data is less than or equals to the reference value which is unexpected, cannot process the request. Expectation shall not be less than or equals to the reference value.",
			Variant: LessThanEqualsExpectation,
		},
		GreaterThanExpectation: {
			Name:    "GreaterThanExpectation",
			Message: "Data is greater than the reference value which is unexpected, cannot process the request. Expectation shall not be greater than the reference value.",
			Variant: GreaterThanExpectation,
		},
		GreaterThanEqualsExpectation: {
			Name:    "GreaterThanEqualsExpectation",
			Message: "Data is greater than or equals to the reference value which is unexpected, cannot process the request. Expectation shall not be greater than or equal to the reference value.",
			Variant: GreaterThanEqualsExpectation,
		},
		NotEmptyFile: {
			Name:    "NotEmptyFile",
			Message: "File not empty which is unexpected, cannot process the request. Expected an empty file with 0 byte.",
			Variant: NotEmptyFile,
		},
		Math: {
			Name:    "Math",
			Message: "Math related error, cannot process the request.",
			Variant: Math,
		},
		Calculation: {
			Name:    "Calculation",
			Message: "Calculation related error, cannot process the request.",
			Variant: Calculation,
		},
		Conversion: {
			Name:    "Conversion",
			Message: "Conversion related error, cannot process the request.",
			Variant: Conversion,
		},
		ConversionValueToString: {
			Name:    "ConversionValueToString",
			Message: "Any value to string convert error, cannot process the request.",
			Variant: ConversionValueToString,
		},
		ConversionValueToInteger: {
			Name:    "ConversionValueToInteger",
			Message: "Any value to integer convert error, cannot process the request.",
			Variant: ConversionValueToInteger,
		},
		ConversionValueToFloat: {
			Name:    "ConversionValueToFloat",
			Message: "Any value to float convert error, cannot process the request.",
			Variant: ConversionValueToFloat,
		},
		ConversionValueToArray: {
			Name:    "ConversionValueToArray",
			Message: "Any value to array convert error, cannot process the request.",
			Variant: ConversionValueToArray,
		},
		ConversionValueToSlice: {
			Name:    "ConversionValueToSlice",
			Message: "Any value to slice convert error, cannot process the request.",
			Variant: ConversionValueToSlice,
		},
		ConversionValueToMap: {
			Name:    "ConversionValueToMap",
			Message: "Any value to map convert error, cannot process the request.",
			Variant: ConversionValueToMap,
		},
		ConversionValueToInterface: {
			Name:    "ConversionValueToInterface",
			Message: "Any value to interface or any convert error, cannot process the request.",
			Variant: ConversionValueToInterface,
		},
		ConversionValueToStruct: {
			Name:    "ConversionValueToStruct",
			Message: "Any value to struct convert error, cannot process the request.",
			Variant: ConversionValueToStruct,
		},
		ConversionValueToPointer: {
			Name:    "ConversionValueToPointer",
			Message: "Any value to pointer convert error, cannot process the request.",
			Variant: ConversionValueToPointer,
		},
		MathDivide: {
			Name:    "MathDivide",
			Message: "Arithmetic or math divide error, cannot process the request.",
			Variant: MathDivide,
		},
		Between: {
			Name:    "Between",
			Message: "Values or Range is in between the boundary, cannot process the request. Expectation: value should not be in the range.",
			Variant: Between,
		},
		NotBetween: {
			Name:    "NotBetween",
			Message: "Values or Range is not in between the boundary, cannot process the request. Expectation: value should be in the range.",
			Variant: NotBetween,
		},
		NotSupportParameter: {
			Name:    "NotSupportParameter",
			Message: "Given parameter is not supported.",
			Variant: NotSupportParameter,
		},
		NotSupportParameters: {
			Name:    "NotSupportParameters",
			Message: "Given parameters are not supported.",
			Variant: NotSupportParameters,
		},
		NotSupportFunction: {
			Name:    "NotSupportFunction",
			Message: "Given function is not supported.",
			Variant: NotSupportFunction,
		},
		NotSupportFunctions: {
			Name:    "NotSupportFunctions",
			Message: "Given functions are not supported.",
			Variant: NotSupportFunctions,
		},
		NotSupportInjections: {
			Name:    "NotSupportInjections",
			Message: "Given injections are not supported.",
			Variant: NotSupportInjections,
		},
		NotSupportValue: {
			Name:    "NotSupportValue",
			Message: "Given value or values are not supported.",
			Variant: NotSupportValue,
		},
		NotSupportOperatingSystem: {
			Name:    "NotSupportOperatingSystem",
			Message: "Operating system is not supported.",
			Variant: NotSupportOperatingSystem,
		},
		NotSupportArchitecture: {
			Name:    "NotSupportArchitecture",
			Message: "System architecture is not supported.",
			Variant: NotSupportArchitecture,
		},
		NotSupportOsVersion: {
			Name:    "NotSupportOsVersion",
			Message: "Operating System version is not supported.",
			Variant: NotSupportOsVersion,
		},
		NotSupportGroups: {
			Name:    "NotSupportGroups",
			Message: "Given groups/categories are not supported.",
			Variant: NotSupportGroups,
		},
		NotSupportInWindows: {
			Name:    "NotSupportInWindows",
			Message: "Current request is not supported in Windows Operating system.",
			Variant: NotSupportInWindows,
		},
		NotSupportInLinux: {
			Name:    "NotSupportInLinux",
			Message: "Current request is not supported in Linux Operating system.",
			Variant: NotSupportInLinux,
		},
		NotSupportInMac: {
			Name:    "NotSupportInMac",
			Message: "Current request is not supported in Macintosh (aka. macOs or iOS) Operating system.",
			Variant: NotSupportInMac,
		},
		NotSupportInUnix: {
			Name:    "NotSupportInUnix",
			Message: "Current request is not supported in Unix Based Operating systems (such as : Linux, CentOs, Darwin/MacOs, Ubuntu).",
			Variant: NotSupportInUnix,
		},
		NotSupportInAndroid: {
			Name:    "NotSupportInAndroid",
			Message: "Current request is not supported in Android Based Operating systems.",
			Variant: NotSupportInAndroid,
		},
		NotSupportVersion: {
			Name:    "NotSupportVersion",
			Message: "Current request is not supported in the version given.",
			Variant: NotSupportVersion,
		},
		UnexpectedVersion: {
			Name:    "UnexpectedVersion",
			Message: "Given version is not expected.",
			Variant: UnexpectedVersion,
		},
		NotImplemented: {
			Name:    "NotImplemented",
			Message: "Current functionality is not implemented.",
			Variant: NotImplemented,
		},
		FileOrDirectoryRelatedExecution: {
			Name:    "FileOrDirectoryRelatedExecution",
			Message: "File or directory related request failed, due to an execution error.",
			Variant: FileOrDirectoryRelatedExecution,
		},
		SymbolicLink: {
			Name:    "SymbolicLink",
			Message: "Request for SymbolicLink failed.",
			Variant: SymbolicLink,
		},
		CommandExecution: {
			Name:    "CommandExecution",
			Message: "Command (cmd) execution (exe) failed, thus cannot process the request.",
			Variant: CommandExecution,
		},
		CommandExecutionNotFound: {
			Name:    "CommandExecutionNotFound",
			Message: "Command (cmd) execution (exe) process path is not found, thus cannot process the request.",
			Variant: CommandExecutionNotFound,
		},
		NoError: {
			Name:    "NoError",
			Message: "No error, it is empty error. Everything is alright.",
			Variant: NoError,
		},
		FileInfo: {
			Name:    "FileInfo",
			Message: "File info related error (could be a directory or file), mostly related to specific path file info read issue - missing/permission.",
			Variant: FileInfo,
		},
		InvalidDefaultSwitchCase: {
			Name:    "InvalidDefaultSwitchCase",
			Message: "None of the switch case matched and fallen into default invalid switch case.",
			Variant: InvalidDefaultSwitchCase,
		},
		NotSupportedOption: {
			Name:    "NotSupportedOption",
			Message: "None of the option is supported.",
			Variant: NotSupportedOption,
		},
		CastingFailed: {
			Name:    "CastingFailed",
			Message: "Failed to cast one type to another.",
			Variant: CastingFailed,
		},
		EmptyFilePath: {
			Name:    "EmptyFilePath",
			Message: "Unexpected empty file path provided.",
			Variant: EmptyFilePath,
		},
		UnexpectedAccessToResource: {
			Name:    "UnexpectedAccessToResource",
			Message: "Unexpected access to (a) resource, which is not permitted. One is not suppose to access or make request to access the resource or resources.",
			Variant: UnexpectedAccessToResource,
		},
		UnexpectedAccessRequest: {
			Name:    "UnexpectedAccessRequest",
			Message: "Unexpected access request, which is not permitted. One is not suppose to access or make request to access the resource or resources.",
			Variant: UnexpectedAccessRequest,
		},
		FailedAccessRequest: {
			Name:    "FailedAccessRequest",
			Message: "Failed to access the request, which is not permitted. Request failed due to access issues.",
			Variant: FailedAccessRequest,
		},
		RequestFailed: {
			Name:    "RequestFailed",
			Message: "Request failed to process, can be number of reasons behind failures.",
			Variant: RequestFailed,
		},
		UnexpectedRequest: {
			Name:    "UnexpectedRequest",
			Message: "Request is not expected. One is not supposed to make the request.",
			Variant: UnexpectedRequest,
		},
		UnexpectedTry: {
			Name:    "UnexpectedTry",
			Message: "Request tried is not expected. One is not supposed to make the request.",
			Variant: UnexpectedTry,
		},
		RetryExceed: {
			Name:    "RetryExceed",
			Message: "Number of requests exceeded the retry counts or retry failed for unknown reasons.",
			Variant: RetryExceed,
		},
		UnexpectedRetry: {
			Name:    "UnexpectedRetry",
			Message: "Unexpected retry counts or retry failed for unknown reasons.",
			Variant: UnexpectedRetry,
		},
		FailedRetry: {
			Name:    "FailedRetry",
			Message: "Failed to retry the event. Please try again or check more error logs for further investigations.",
			Variant: FailedRetry,
		},
		FailedProcess: {
			Name:    "FailedProcess",
			Message: "Requested process or task is failed.",
			Variant: FailedProcess,
		},
		CannotGrantPermission: {
			Name:    "CannotGrantPermission",
			Message: "Cannot grant permission for the event or request.",
			Variant: CannotGrantPermission,
		},
		PermissionIssue: {
			Name:    "PermissionIssue",
			Message: "Permission related issue, one doesn't have the access to proceed with the request or task.",
			Variant: PermissionIssue,
		},
		PermissionNotFound: {
			Name:    "PermissionNotFound",
			Message: "Permission not found related issue, one doesn't have the access rights or permission not found for the current operation request.",
			Variant: PermissionNotFound,
		},
		PermissionFailed: {
			Name:    "PermissionFailed",
			Message: "Permission failed to process the request. Doesn't have rights or broken rights issue.",
			Variant: PermissionFailed,
		},
		AccessRequestFailed: {
			Name:    "AccessRequestFailed",
			Message: "Access request failed.",
			Variant: AccessRequestFailed,
		},
		UnAuthorizedAccess: {
			Name:    "UnAuthorizedAccess",
			Message: "One invoked unauthorized access.",
			Variant: UnAuthorizedAccess,
		},
		ResourceFreezeCannotAccess: {
			Name:    "ResourceFreezeCannotAccess",
			Message: "Resource is frozen and cannot access the resource.",
			Variant: ResourceFreezeCannotAccess,
		},
		DisposedResourceCannotAccess: {
			Name:    "DisposedResourceCannotAccess",
			Message: "Disposed resource are removed or set for garbage collection. Moreover, blocked for access.",
			Variant: DisposedResourceCannotAccess,
		},
		FinalizedResourceCannotAccess: {
			Name:    "FinalizedResourceCannotAccess",
			Message: "Finalized resources are blocked for access.",
			Variant: FinalizedResourceCannotAccess,
		},
		FileMovedOut: {
			Name:    "FileMovedOut",
			Message: "File is not there anymore. File moved out.",
			Variant: FileMovedOut,
		},
		PolicyRelated: {
			Name:    "PolicyRelated",
			Message: "Policy related issue.",
			Variant: PolicyRelated,
		},
		PolicyFailed: {
			Name:    "PolicyFailed",
			Message: "Policy failed to proceed.",
			Variant: PolicyFailed,
		},
		PolicyRequestRestricted: {
			Name:    "PolicyRequestRestricted",
			Message: "Policy request is restricted.",
			Variant: PolicyRequestRestricted,
		},
		FatalError: {
			Name:    "FatalError",
			Message: "Fatal errors are dangerous and the last stage thrown. Please consult with appropriate personnel to get the solution.",
			Variant: FatalError,
		},
		WarningError: {
			Name:    "WarningError",
			Message: "It is not an error but warning are given to avoid errors in future.",
			Variant: WarningError,
		},
		SystemDamaged: {
			Name:    "SystemDamaged",
			Message: "System is somehow damaged. Please consult with appropriate personnel to get the solution.",
			Variant: SystemDamaged,
		},
		DiagnosticFailed: {
			Name:    "DiagnosticFailed",
			Message: "Diagnostic failed.",
			Variant: DiagnosticFailed,
		},
		UnexpectedDiagnostic: {
			Name:    "UnexpectedDiagnostic",
			Message: "Unexpected diagnostic results. Please consult with appropriate personnel to get the solution.",
			Variant: UnexpectedDiagnostic,
		},
		AnalysisFailed: {
			Name:    "AnalysisFailed",
			Message: "Analysis failed. Please consult with appropriate personnel to get the solution.",
			Variant: AnalysisFailed,
		},
		RecurringProcessFailed: {
			Name:    "RecurringProcessFailed",
			Message: "Recurring process failed.",
			Variant: RecurringProcessFailed,
		},
		Step1Failed: {
			Name:    "Step1Failed",
			Message: "First step failed.",
			Variant: Step1Failed,
		},
		Step2Failed: {
			Name:    "Step2Failed",
			Message: "Second step failed.",
			Variant: Step2Failed,
		},
		Step3Failed: {
			Name:    "Step3Failed",
			Message: "Third step failed.",
			Variant: Step3Failed,
		},
		Step4Failed: {
			Name:    "Step4Failed",
			Message: "Fourth step failed.",
			Variant: Step4Failed,
		},
		Step5Failed: {
			Name:    "Step5Failed",
			Message: "Fifth step failed.",
			Variant: Step5Failed,
		},
		Step6Failed: {
			Name:    "Step6Failed",
			Message: "Sixth step failed.",
			Variant: Step6Failed,
		},
		Step7Failed: {
			Name:    "Step7Failed",
			Message: "Seventh step failed.",
			Variant: Step7Failed,
		},
		Step8Failed: {
			Name:    "Step8Failed",
			Message: "Eighth step failed.",
			Variant: Step8Failed,
		},
		StepFailed: {
			Name:    "StepFailed",
			Message: "Mentioned step failed.",
			Variant: Step8Failed, // is it correct one?
		},
		NextStepFailed: {
			Name:    "NextStepFailed",
			Message: "Next step failed.",
			Variant: NextStepFailed,
		},
		PreviousStepFailed: {
			Name:    "PreviousStepFailed",
			Message: "Previous step failed.",
			Variant: PreviousStepFailed,
		},
		FinalStepFailed: {
			Name:    "FinalStepFailed",
			Message: "Final step failed.",
			Variant: FinalStepFailed,
		},
		CompletionFailed: {
			Name:    "CompletionFailed",
			Message: "Completion step failed.",
			Variant: CompletionFailed,
		},
		CompileFailed: {
			Name:    "CompileFailed",
			Message: "Compilation failed.",
			Variant: CompileFailed,
		},
		SearchFailed: {
			Name:    "SearchFailed",
			Message: "Searching failed.",
			Variant: SearchFailed,
		},
		QueryFailed: {
			Name:    "QueryFailed",
			Message: "Requested query failed.",
			Variant: QueryFailed,
		},
		DatabaseOperationFailed: {
			Name:    "DatabaseOperationFailed",
			Message: "Current database request operation failed.",
			Variant: DatabaseOperationFailed,
		},
		UIOperationFailed: {
			Name:    "UIOperationFailed",
			Message: "User interface operation failed.",
			Variant: UIOperationFailed,
		},
		SynchronizationFailed: {
			Name:    "SynchronizationFailed",
			Message: "Synchronization mismatches, thus synchronization failed.",
			Variant: SynchronizationFailed,
		},
		ParsingFailed: {
			Name:    "ParsingFailed",
			Message: "Parsing data mismatches, thus parsing failed.",
			Variant: ParsingFailed,
		},
		SyntaxFailed: {
			Name:    "SyntaxFailed",
			Message: "Syntax is not proper, thus failed to process syntax parsing.",
			Variant: SyntaxFailed,
		},
		ExternalToolFailed: {
			Name:    "ExternalToolFailed",
			Message: "External tool running failed or failed to run or didn't get the expected return.",
			Variant: ExternalToolFailed,
		},
		CommandLineFailed: {
			Name:    "CommandLineFailed",
			Message: "Command line tool running failed or failed to run or didn't get the expected return.",
			Variant: CommandLineFailed,
		},
		CommunicationFailed: {
			Name:    "CommunicationFailed",
			Message: "Communication failed to return the expected data.",
			Variant: CommunicationFailed,
		},
		GarbageCollectionFailed: {
			Name:    "GarbageCollectionFailed",
			Message: "Garbage collection failed.",
			Variant: GarbageCollectionFailed,
		},
		CreationRequestFailed: {
			Name:    "CreationRequestFailed",
			Message: "Creation request failed.",
			Variant: CreationRequestFailed,
		},
		ReadFailed: {
			Name:    "ReadFailed",
			Message: "Failed to read.",
			Variant: ReadFailed,
		},
		EditFailed: {
			Name:    "EditFailed",
			Message: "Failed to edit.",
			Variant: EditFailed,
		},
		SaveFailed: {
			Name:    "SaveFailed",
			Message: "Failed to save data.",
			Variant: SaveFailed,
		},
		DeleteFailed: {
			Name:    "DeleteFailed",
			Message: "Failed to delete data.",
			Variant: DeleteFailed,
		},
		ReadRequestFailed: {
			Name:    "ReadRequestFailed",
			Message: "Read request failed, thus no data.",
			Variant: ReadRequestFailed,
		},
		EditRequestFailed: {
			Name:    "EditRequestFailed",
			Message: "Edit request failed, thus saving failed.",
			Variant: EditRequestFailed,
		},
		SaveRequestFailed: {
			Name:    "SaveRequestFailed",
			Message: "Save request failed, thus saving failed.",
			Variant: SaveRequestFailed,
		},
		DeleteRequestFailed: {
			Name:    "DeleteRequestFailed",
			Message: "Delete request failed, thus saving failed.",
			Variant: DeleteRequestFailed,
		},
		CrudOperationFailed: {
			Name:    "CrudOperationFailed",
			Message: "(Create or Read or Update or Delete) request failed, saving failed.",
			Variant: CrudOperationFailed,
		},
		NetworkIssue: {
			Name:    "NetworkIssue",
			Message: "Network related issue.",
			Variant: NetworkIssue,
		},
		NetworkOffline: {
			Name:    "NetworkOffline",
			Message: "Network is offline so no transaction was possible.",
			Variant: NetworkOffline,
		},
		AdapterDisconnected: {
			Name:    "AdapterDisconnected",
			Message: "Adapter device disconnected.",
			Variant: AdapterDisconnected,
		},
		Disconnected: {
			Name:    "Disconnected",
			Message: "Any device or event or process or task is disconnected.",
			Variant: Disconnected,
		},
		Offline: {
			Name:    "Offline",
			Message: "System is offline.",
			Variant: Offline,
		},
		Online: {
			Name:    "Online",
			Message: "System is online issue.",
			Variant: Online,
		},
		Connection: {
			Name:    "Connection",
			Message: "Connection related issues.",
			Variant: Connection,
		},
		Connected: {
			Name:    "Connected",
			Message: "Connected related issues.", // seems improperly phrased?
			Variant: Connected,
		},
		TranspileFailed: {
			Name:    "TranspileFailed",
			Message: "Transpiling data failed.",
			Variant: TranspileFailed,
		},
		NginxParsingFailed: {
			Name:    "NginxParsingFailed",
			Message: "Nginx parsing failed.",
			Variant: NginxParsingFailed,
		},
		ApacheParsingFailed: {
			Name:    "ApacheParsingFailed",
			Message: "Apache parsing failed.",
			Variant: ApacheParsingFailed,
		},
		EndSyntaxWrong: {
			Name:    "EndSyntaxWrong",
			Message: "Ending syntax is wrong.",
			Variant: EndSyntaxWrong,
		},
		BracketsMissing: {
			Name:    "BracketsMissing",
			Message: "Bracket or Brackets are missing.",
			Variant: BracketsMissing,
		},
		WrongSyntax: {
			Name:    "WrongSyntax",
			Message: "Wrong syntax given.",
			Variant: WrongSyntax,
		},
		StartSyntaxWrong: {
			Name:    "StartSyntaxWrong",
			Message: "Starting syntax is not as expected.",
			Variant: StartSyntaxWrong,
		},
		TokenMissing: {
			Name:    "TokenMissing",
			Message: "Token is missing.",
			Variant: TokenMissing,
		},
		Missing: {
			Name:    "Missing",
			Message: "Generic missing",
			Variant: Missing,
		},
		IntegrityBroken: {
			Name:    "IntegrityBroken",
			Message: "Integrity is broken",
			Variant: IntegrityBroken,
		},
		DataIntegrityLost: {
			Name:    "DataIntegrityLost",
			Message: "Data integrity is broken or lost.",
			Variant: DataIntegrityLost,
		},
		MappingFailed: {
			Name:    "MappingFailed",
			Message: "Mapping failed.",
			Variant: MappingFailed,
		},
		CountFailed: {
			Name:    "CountFailed",
			Message: "Count failed.",
			Variant: CountFailed,
		},
		RemoveFailed: {
			Name:    "RemoveFailed",
			Message: "Remove failed.",
			Variant: RemoveFailed,
		},
		NotAsExpected: {
			Name:    "NotAsExpected",
			Message: "Not as expected / unexpected.",
			Variant: NotAsExpected,
		},
		Unexpected: {
			Name:    "Unexpected",
			Message: "Not as expected / unexpected.",
			Variant: Unexpected,
		},
		CorruptedData: {
			Name:    "CorruptedData",
			Message: "Data integrity is not correct and/or corrupted.",
			Variant: CorruptedData,
		},
		NotFound: {
			Name:    "NotFound",
			Message: "Not found.",
			Variant: NotFound,
		},
		Found: {
			Name:    "Found",
			Message: "Found",
			Variant: Found,
		},
		Encoding: {
			Name:    "Encoding",
			Message: "Encoding or serializing failed, cannot process the request further.",
			Variant: Encoding,
		},
		Decoding: {
			Name:    "Decoding",
			Message: "Decoding or deserializing failed, cannot process the request further.",
			Variant: Encoding,
		},
		Marshalling: {
			Name:    "Marshalling",
			Message: "Marshalling or serializing failed, cannot process the request further.",
			Variant: Marshalling,
		},
		Unmarshalling: {
			Name:    "Unmarshalling",
			Message: "Unmarshalling or deserializing failed, cannot process the request further.",
			Variant: Unmarshalling,
		},
		InvalidProcess: {
			Name:    "InvalidProcess",
			Message: "Process is invalid cannot process further.",
			Variant: InvalidProcess,
		},
		NotFoundProcess: {
			Name:    "NotFoundProcess",
			Message: "Process not found, cannot process further.",
			Variant: NotFoundProcess,
		},
		Chmod: {
			Name:    "Chmod",
			Message: "Chmod related issue, probably cannot read or apply to a path or anything.",
			Variant: Chmod,
		},
		ChmodInvalid: {
			Name:    "ChmodInvalid",
			Message: "Chmod invalid, probably cannot read or apply invalid chmod values.",
			Variant: ChmodInvalid,
		},
		ChmodApplyFailed: {
			Name:    "ChmodApplyFailed",
			Message: "Failed to apply chmod to a path.",
			Variant: ChmodApplyFailed,
		},
		ExistingChmodReadFailed: {
			Name:    "ExistingChmodReadFailed",
			Message: "Existing CHMOD read failed from system.",
			Variant: ExistingChmodReadFailed,
		},
		ChownIssue: {
			Name:    "ChownIssue",
			Message: "Chown related issue (probably read/write/apply/mask issue). It could also be related to user, group. Cannot perform the operation properly.",
			Variant: ChownIssue,
		},
		ChownUserOrGroupApplyIssue: {
			Name:    "ChownUserOrGroupApplyIssue",
			Message: "Chown user or group or UserGroup apply issue. Cannot perform the operation properly.",
			Variant: ChownUserOrGroupApplyIssue,
		},
		Permission: {
			Name:    "Permission",
			Message: "Permission related issue, cannot apply permission or cannot read permission or cannot access resource because of permission issue.",
			Variant: Permission,
		},
		Authorization: {
			Name:    "Authorization",
			Message: "Authorization related issue, cannot apply authorization or cannot read permission or cannot access resource because of permission issue.",
			Variant: Authorization,
		},
		AuthorizationFailed: {
			Name:    "AuthorizationFailed",
			Message: "Authorization failed, user is not authorized to access or failed to save the authorization data.",
			Variant: AuthorizationFailed,
		},
		AuthorizationFailedToSave: {
			Name:    "AuthorizationFailedToSave",
			Message: "Authorization data failed to save in the system.",
			Variant: AuthorizationFailedToSave,
		},
		InstructionUndefined: {
			Name:    "InstructionUndefined",
			Message: "Instruction is not as expected it is undefined or nil or empty. Instruction definition doesn't match as expectation.",
			Variant: InstructionUndefined,
		},
		InstructionMustExist: {
			Name:    "InstructionMustExist",
			Message: "Instruction is not as expected it is undefined or nil or empty. Instruction definition must exist.",
			Variant: InstructionMustExist,
		},
		UnexpectedDefinition: {
			Name:    "UnexpectedDefinition",
			Message: "Definition is not as expected. Definition must be defined properly.",
			Variant: UnexpectedDefinition,
		},
		BashScriptIssue: {
			Name:    "BashScriptIssue",
			Message: "Bash script related issue (probably bash script running/reading/writing/async issue).",
			Variant: BashScriptIssue,
		},
		ShellScriptIssue: {
			Name:    "ShellScriptIssue",
			Message: "Shell script related issue (probably shell script running/reading/writing/async issue).",
			Variant: ShellScriptIssue,
		},
		PerlScriptIssue: {
			Name:    "PerlScriptIssue",
			Message: "Perl script related issue (probably Perl script running/reading/writing/async issue).",
			Variant: PerlScriptIssue,
		},
		PythonScriptIssue: {
			Name:    "PythonScriptIssue",
			Message: "Python script related issue (probably Python script running/reading/writing/async issue).",
			Variant: PythonScriptIssue,
		},
		BashScriptRunningProcessFailed: {
			Name:    "BashScriptRunningProcessFailed",
			Message: "Bash script related issue running or process generic error.",
			Variant: BashScriptRunningProcessFailed,
		},
		ShellScriptRunningProcessFailed: {
			Name:    "ShellScriptRunningProcessFailed",
			Message: "Shell script related issue running or process generic error.",
			Variant: ShellScriptRunningProcessFailed,
		},
		ScriptRunningFailed: {
			Name:    "ScriptRunningFailed",
			Message: "Script or process running failed, unexpected error.",
			Variant: ScriptRunningFailed,
		},
		PowershellIssue: {
			Name:    "PowershellIssue",
			Message: "Powershell script or process running failed, unexpected error.",
			Variant: PowershellIssue,
		},
		CmdIssue: {
			Name:    "CmdIssue",
			Message: "Command script or process running failed, unexpected error.",
			Variant: CmdIssue,
		},
		PathStatusCannotRead: {
			Name:    "PathStatusCannotRead",
			Message: "Cannot read the path status properly, unexpected error.",
			Variant: PathStatusCannotRead,
		},
		FileRead: {
			Name:    "FileRead",
			Message: "Cannot read the file or files properly, unexpected error.",
			Variant: FileRead,
		},
		FilesRead: {
			Name:    "FilesRead",
			Message: "Cannot read the files properly, unexpected error.",
			Variant: FileRead,
		},
		WriteFiles: {
			Name:    "WriteFiles",
			Message: "Cannot write multiple files properly (one or more file has writing issue), unexpected error. Could be related to file system permission or authorization issue.",
			Variant: WriteFiles,
		},
		PathIssue: {
			Name:    "PathIssue",
			Message: "Path related unexpected issue (could be directory or file). Could be related to any of the issues (read, modify, access, apply permission).",
			Variant: PathIssue,
		},
		PathNotFound: {
			Name:    "PathNotFound",
			Message: "Give path is not exist or not found in the file system.",
			Variant: PathNotFound,
		},
		PathExists: {
			Name:    "PathExists",
			Message: "Give path is exist or found in the file system which is not expected.",
			Variant: PathExists,
		},
		InvalidPath: {
			Name:    "InvalidPath",
			Message: "Invalid path is provided. Cannot process further!",
			Variant: InvalidPath,
		},
		PathRelatedIssue: {
			Name:    "PathRelatedIssue",
			Message: "Issue is related to path (probable cause can be nil data, permission issue, invalid path etc)!",
			Variant: PathRelatedIssue,
		},
		PathSyntaxIssue: {
			Name:    "PathSyntaxIssue",
			Message: "Provided path syntax is not appropriate to process further!",
			Variant: PathSyntaxIssue,
		},
		PathHasInvalidDataOrNilData: {
			Name:    "PathHasInvalidDataOrNilData",
			Message: "Provided path has invalid data (which doesn't meet the expectation) or has nil data.",
			Variant: PathHasInvalidDataOrNilData,
		},
		RequestResourceNonExist: {
			Name:    "RequestResourceNonExist",
			Message: "Requested resource doesn't exist in the system.",
			Variant: RequestResourceNonExist,
		},
		ResourceMutating: {
			Name:    "ResourceMutating",
			Message: "Resource is mutating bad, programming example! Please try non mutating or read-only resource to proceed with.",
			Variant: ResourceMutating,
		},
		ReadOnlyExpected: {
			Name:    "ReadOnlyExpected",
			Message: "Resource is expected to be readonly, cannot proceed with mutating data.",
			Variant: ReadOnlyExpected,
		},
		Env: {
			Name:    "Env",
			Message: "Environment related issue (path, variable or anything which unexpected in the environment settings).",
			Variant: Env,
		},
		EnvironmentVariable: {
			Name:    "EnvironmentVariable",
			Message: "Environment variable related issue, probably environment is already exist or not found.",
			Variant: EnvironmentVariable,
		},
		EnvironmentVariableExist: {
			Name:    "EnvironmentVariableExist",
			Message: "Environment variable exist which is not expected.",
			Variant: EnvironmentVariableExist,
		},
		EnvironmentVariableNotExist: {
			Name:    "EnvironmentVariableNotExist",
			Message: "Environment variable doesn't exist which is not expected.",
			Variant: EnvironmentVariableNotExist,
		},
		EnvironmentVariableModify: {
			Name:    "EnvironmentVariableModify",
			Message: "Environment variable cannot modify successfully.",
			Variant: EnvironmentVariableModify,
		},
		EnvPath: {
			Name:    "EnvPath",
			Message: "Environment path related issue (probably cannot read or doesn't exist or cannot update).",
			Variant: EnvPath,
		},
		LookupIssue: {
			Name:    "LookupIssue",
			Message: "Look up or search failed, due to some system restriction or process halt.",
			Variant: LookupIssue,
		},
		IntConvertIssue: {
			Name:    "IntConvertIssue",
			Message: "Failed to convert data to integer value.",
			Variant: IntConvertIssue,
		},
		FloatConvertIssue: {
			Name:    "FloatConvertIssue",
			Message: "Failed to convert data to float value.",
			Variant: FloatConvertIssue,
		},
		StringConvertIssue: {
			Name:    "StringConvertIssue",
			Message: "Failed to convert data to string value.",
			Variant: StringConvertIssue,
		},
		StringToIntConvertIssue: {
			Name:    "StringToIntConvertIssue",
			Message: "Failed to convert string to integer value.",
			Variant: StringToIntConvertIssue,
		},
		StringToFloatConvertIssue: {
			Name:    "StringToFloatConvertIssue",
			Message: "Failed to convert string to float value.",
			Variant: StringToFloatConvertIssue,
		},
		StringToExpectedConvertIssue: {
			Name:    "StringToExpectedConvertIssue",
			Message: "Failed to convert string to expected data type.",
			Variant: StringToExpectedConvertIssue,
		},
		InterfaceToExpectedConvertIssue: {
			Name:    "InterfaceToExpectedConvertIssue",
			Message: "Failed to convert interface data / dynamic to expected data type.",
			Variant: InterfaceToExpectedConvertIssue,
		},
		StringToStringConvertIssue: {
			Name:    "StringToStringConvertIssue",
			Message: "Failed to convert string to string (not meeting expectation).",
			Variant: StringToStringConvertIssue,
		},
		ZippingFailed: {
			Name:    "ZippingFailed",
			Message: "Failed to zip file or files.",
			Variant: ZippingFailed,
		},
		CompressingFailed: {
			Name:    "CompressingFailed",
			Message: "Failed to compress file or files.",
			Variant: CompressingFailed,
		},
		CompressingIssue: {
			Name:    "CompressingIssue",
			Message: "Compress related issue (command/process/file-system related issue).",
			Variant: CompressingIssue,
		},
		DecompressingFailed: {
			Name:    "DecompressingFailed",
			Message: "Decompressing failed!",
			Variant: DecompressingFailed,
		},
		DecompressingIssue: {
			Name:    "DecompressingIssue",
			Message: "Decompressing related issue, decompressing failed!",
			Variant: DecompressingIssue,
		},
		InstructionIssue: {
			Name:    "InstructionIssue",
			Message: "Instruction related issue, instruction was not as expected. Failed to proceed further!",
			Variant: InstructionIssue,
		},
		EmptyCollection: {
			Name:    "EmptyCollection",
			Message: "Empty collection which is unexpected!",
			Variant: EmptyCollection,
		},
		DictionaryNotContainsKey: {
			Name:    "DictionaryNotContainsKey",
			Message: "Dictionary or map doesn't contain the key!",
			Variant: DictionaryNotContainsKey,
		},
		KeyNotExist: {
			Name:    "KeyNotExist",
			Message: "Key doesn't exist in the system!",
			Variant: KeyNotExist,
		},
		ChangeDirectoryFailed: {
			Name:    "ChangeDirectoryFailed",
			Message: "Change directory or ch failed to execute!",
			Variant: ChangeDirectoryFailed,
		},
		CmdOnceFailed: {
			Name:    "CmdOnceFailed",
			Message: "Command once script running failed!",
			Variant: CmdOnceFailed,
		},
		CmdOnceRunningIssue: {
			Name:    "CmdOnceRunningIssue",
			Message: "Command once script running failed!",
			Variant: CmdOnceRunningIssue,
		},
		Invalid: {
			Name:    "Invalid",
			Message: "Invalid!",
			Variant: Invalid,
		},
		InvalidRequest: {
			Name:    "InvalidRequest",
			Message: "Invalid request, cannot process it further!",
			Variant: InvalidRequest,
		},
		InvalidData: {
			Name:    "InvalidData",
			Message: "Invalid data, cannot process it further!",
			Variant: InvalidData,
		},
		SystemRelated: {
			Name:    "SystemRelated",
			Message: "System related issue!",
			Variant: SystemRelated,
		},
		OsRelated: {
			Name:    "OsRelated",
			Message: "Operating System related issue!",
			Variant: OsRelated,
		},
		OperatingSystemInternal: {
			Name:    "OperatingSystemInternal",
			Message: "Operating System internal issue!",
			Variant: OperatingSystemInternal,
		},
		SystemUser: {
			Name:    "SystemUser",
			Message: "System user related issue!",
			Variant: SystemUser,
		},
		SysUserInvalid: {
			Name:    "SysUserInvalid",
			Message: "System user invalid!",
			Variant: SysUserInvalid,
		},
		SysGroupInvalid: {
			Name:    "SysGroupInvalid",
			Message: "System group invalid!",
			Variant: SysGroupInvalid,
		},
		CommandLineExecution: {
			Name:    "CommandLineExecution",
			Message: "Command line execution failed!",
			Variant: CommandLineExecution,
		},
		NotFoundData: {
			Name:    "NotFoundData",
			Message: "Expected data not found in the system!",
			Variant: NotFoundData,
		},
		NotFoundObject: {
			Name:    "NotFoundObject",
			Message: "Expected object not found or cannot create!",
			Variant: NotFoundObject,
		},
		NotFoundItems: {
			Name:    "NotFoundItems",
			Message: "Items not found in the system or in the collection!",
			Variant: NotFoundItems,
		},
		NotFoundGroup: {
			Name:    "NotFoundGroup",
			Message: "Group not found in the system or in the collection!",
			Variant: NotFoundGroup,
		},
		NotFoundName: {
			Name:    "NotFoundName",
			Message: "Category not found in the system or in the collection!",
			Variant: NotFoundName,
		},
		InvalidId: {
			Name:    "InvalidId",
			Message: "Invalid id issue!",
			Variant: InvalidId,
		},
		InvalidName: {
			Name:    "InvalidName",
			Message: "Invalid name issue!",
			Variant: InvalidName,
		},
		InvalidIdentifier: {
			Name:    "InvalidIdentifier",
			Message: "Invalid identifier issue!",
			Variant: InvalidIdentifier,
		},
		InvalidDir: {
			Name:    "InvalidDir",
			Message: "Invalid directory issue!",
			Variant: InvalidDir,
		},
		BashScriptRunning: {
			Name:    "BashScriptRunning",
			Message: "Bash script running failed!",
			Variant: BashScriptRunning,
		},
		BashScriptFailed: {
			Name:    "BashScriptFailed",
			Message: "Bash script running failed!",
			Variant: BashScriptFailed,
		},
		ShellScriptFailed: {
			Name:    "ShellScriptFailed",
			Message: "Shell script running failed!",
			Variant: ShellScriptFailed,
		},
		PathInfoFailed: {
			Name:    "PathInfoFailed",
			Message: "Path info failed!",
			Variant: PathInfoFailed,
		},
		StatFailed: {
			Name:    "StatFailed",
			Message: "Stat / Status failed!",
			Variant: StatFailed,
		},
		PathStatFailed: {
			Name:    "PathStatFailed",
			Message: "PathStat failed!",
			Variant: PathStatFailed,
		},
		StatusInvalid: {
			Name:    "StatusInvalid",
			Message: "Invalid status found!",
			Variant: StatusInvalid,
		},
		PathMissingOrInvalid: {
			Name:    "PathMissingOrInvalid",
			Message: "Path missing or invalid!",
			Variant: PathMissingOrInvalid,
		},
		PathMismatch: {
			Name:    "PathMismatch",
			Message: "Path mismatch error, expectation didn't meet!",
			Variant: PathMismatch,
		},
		Mismatch: {
			Name:    "Mismatch",
			Message: "Mismatch!",
			Variant: Mismatch,
		},
		MismatchStatus: {
			Name:    "MismatchStatus",
			Message: "Mismatch status!",
			Variant: MismatchStatus,
		},
		MismatchRequest: {
			Name:    "MismatchRequest",
			Message: "Mismatch request, expectation didn't meet!",
			Variant: MismatchRequest,
		},
		MismatchExpectation: {
			Name:    "MismatchExpectation",
			Message: "Mismatch request, expectation didn't meet!",
			Variant: MismatchExpectation,
		},
		RwxMismatch: {
			Name:    "RwxMismatch",
			Message: "Rwx mismatch!",
			Variant: RwxMismatch,
		},
		ChmodMismatch: {
			Name:    "ChmodMismatch",
			Message: "Chmod mismatch, expectation didn't meet!",
			Variant: ChmodMismatch,
		},
		DataMismatch: {
			Name:    "DataMismatch",
			Message: "Data mismatch, expectation didn't meet!",
			Variant: DataMismatch,
		},
		OutOfSync: {
			Name:    "OutOfSync",
			Message: "Out of sync!",
			Variant: OutOfSync,
		},
		ExpectationFailed: {
			Name:    "ExpectationFailed",
			Message: "Expectation doesn't meet the situation or data!",
			Variant: ExpectationFailed,
		},
		TestFailed: {
			Name:    "TestFailed",
			Message: "Test failed!",
			Variant: TestFailed,
		},
		CreateDirectoryFailed: {
			Name:    "CreateDirectoryFailed",
			Message: "Directory create failed!",
			Variant: CreateDirectoryFailed,
		},
		RenameFailed: {
			Name:    "RenameFailed",
			Message: "Rename failed!",
			Variant: RenameFailed,
		},
		RenameFileFailed: {
			Name:    "RenameFileFailed",
			Message: "Rename file failed!",
			Variant: RenameFileFailed,
		},
		RenamePathFailed: {
			Name:    "RenamePathFailed",
			Message: "Rename path failed!",
			Variant: RenamePathFailed,
		},
		CreatePathFailed: {
			Name:    "CreatePathFailed",
			Message: "Create path failed!",
			Variant: CreatePathFailed,
		},
		DeletePathFailed: {
			Name:    "DeletePathFailed",
			Message: "Delete path failed!",
			Variant: DeletePathFailed,
		},
		ModifyDataFailed: {
			Name:    "ModifyDataFailed",
			Message: "Modify data failed!",
			Variant: ModifyDataFailed,
		},
		MissingPathsOrInvalidPaths: {
			Name:    "MissingPathsOrInvalidPaths",
			Message: "Missing path(s) or invalid path(s)!",
			Variant: MissingPathsOrInvalidPaths,
		},
		Copy: {
			Name:    "Copy",
			Message: "Copy failed!",
			Variant: Copy,
		},
		PathCopy: {
			Name:    "PathCopy",
			Message: "Path copy failed!",
			Variant: PathCopy,
		},
		Move: {
			Name:    "Move",
			Message: "Move failed!",
			Variant: Move,
		},
		PathMove: {
			Name:    "PathMove",
			Message: "Path move failed!",
			Variant: PathMove,
		},
		Append: {
			Name:    "Append",
			Message: "Append failed!",
			Variant: Append,
		},
		Prepend: {
			Name:    "Prepend",
			Message: "Prepend failed!",
			Variant: Prepend,
		},
		PathExpand: {
			Name:    "PathExpand",
			Message: "Path expand, or reading directory failed to expand!",
			Variant: PathExpand,
		},
		PathReadDir: {
			Name:    "PathReadDir",
			Message: "Path read dir, or reading directory failed to expand!",
			Variant: PathReadDir,
		},
		Remove: {
			Name:    "Remove",
			Message: "Remove failed!",
			Variant: Remove,
		},
		Write: {
			Name:    "Write",
			Message: "Write failed!",
			Variant: Write,
		},
		LineIssue: {
			Name:    "LineIssue",
			Message: "Line related failed!",
			Variant: LineIssue,
		},
		LinesIssue: {
			Name:    "LinesIssue",
			Message: "Lines related failed!",
			Variant: LinesIssue,
		},
		CollectionIssue: {
			Name:    "CollectionIssue",
			Message: "Collection related failed!",
			Variant: CollectionIssue,
		},
		ValidationFailed: {
			Name:    "ValidationFailed",
			Message: "Validation failed!",
			Variant: ValidationFailed,
		},
		Delete: {
			Name:    "Delete",
			Message: "Delete failed!",
			Variant: Delete,
		},
		DeletePath: {
			Name:    "DeletePath",
			Message: "Delete path failed!",
			Variant: DeletePath,
		},
		DeletePaths: {
			Name:    "DeletePaths",
			Message: "Delete paths failed!",
			Variant: DeletePaths,
		},
		WriteFailed: {
			Name:    "WriteFailed",
			Message: "Write failed!",
			Variant: WriteFailed,
		},
		CreatePath: {
			Name:    "CreatePath",
			Message: "Create path failed!",
			Variant: CreatePath,
		},
		NewPath: {
			Name:    "NewPath",
			Message: "New path failed!",
			Variant: NewPath,
		},
		Network: {
			Name:    "Network",
			Message: "Network related issue!",
			Variant: Network,
		},
		NetPlan: {
			Name:    "NetPlan",
			Message: "NetPlan related issue!",
			Variant: NetPlan,
		},
		HostName: {
			Name:    "HostName",
			Message: "HostName related issue!",
			Variant: HostName,
		},
		ResolveIssue: {
			Name:    "ResolveIssue",
			Message: "Resolve issue!",
			Variant: ResolveIssue,
		},
		Sync: {
			Name:    "Sync",
			Message: "Synchronization related issue!",
			Variant: Sync,
		},
		ConflictIssue: {
			Name:    "ConflictIssue",
			Message: "Conflict issue, synchronization related issue!",
			Variant: ConflictIssue,
		},
		MergeIssue: {
			Name:    "MergeIssue",
			Message: "Merge issue, synchronization related issue!",
			Variant: MergeIssue,
		},
		NonSync: {
			Name:    "NonSync",
			Message: "Non synchronized related issue!",
			Variant: NonSync,
		},
		Binding: {
			Name:    "Binding",
			Message: "Binding related issue!",
			Variant: Binding,
		},
		MVVMIssue: {
			Name:    "MVVMIssue",
			Message: "MVVM related issue!",
			Variant: MVVMIssue,
		},
		DesignIssue: {
			Name:    "DesignIssue",
			Message: "Design related issue!",
			Variant: DesignIssue,
		},
		QuotationMissing: {
			Name:    "QuotationMissing",
			Message: "Quotation missing, syntax issue!",
			Variant: QuotationMissing,
		},
		RegexIssue: {
			Name:    "RegexIssue",
			Message: "Regex issue!",
			Variant: RegexIssue,
		},
		RegexValidationIssue: {
			Name:    "RegexValidationIssue",
			Message: "Regex validation issue!",
			Variant: RegexValidationIssue,
		},
		CannotApplyInstruction: {
			Name:    "CannotApplyInstruction",
			Message: "Cannot apply instruction!",
			Variant: CannotApplyInstruction,
		},
		InstructionFormatIssue: {
			Name:    "InstructionFormatIssue",
			Message: "Instruction format issue!",
			Variant: InstructionFormatIssue,
		},
		InstructionFormatMissingElementIssue: {
			Name:    "InstructionFormatMissingElementIssue",
			Message: "Instruction format has some important parts missing, failed to execute!",
			Variant: InstructionFormatMissingElementIssue,
		},
		FileSystemInstructionFailed: {
			Name:    "FileSystemInstructionFailed",
			Message: "File system instruction failed to execute!",
			Variant: FileSystemInstructionFailed,
		},
		EthernetInstructionFailed: {
			Name:    "EthernetInstructionFailed",
			Message: "Ethernet instruction failed to execute!",
			Variant: EthernetInstructionFailed,
		},
		CronTabsInstructionFailed: {
			Name:    "CronTabsInstructionFailed",
			Message: "Cron tabs instruction failed to execute!",
			Variant: CronTabsInstructionFailed,
		},
		DatabaseInstructionFailed: {
			Name:    "DatabaseInstructionFailed",
			Message: "Database instruction failed to execute!",
			Variant: DatabaseInstructionFailed,
		},
		MySqlIssue: {
			Name:    "MySqlIssue",
			Message: "MySQL related issue!",
			Variant: MySqlIssue,
		},
		PostgreSqlIssue: {
			Name:    "PostgreSqlIssue",
			Message: "PostgreSQL related issue!",
			Variant: PostgreSqlIssue,
		},
		LightSpeedParsingIssue: {
			Name:    "LightSpeedParsingIssue",
			Message: "Light speed parsing issue!",
			Variant: LightSpeedParsingIssue,
		},
		LightSpeedSyntaxIssue: {
			Name:    "LightSpeedSyntaxIssue",
			Message: "Light speed syntax issue!",
			Variant: LightSpeedSyntaxIssue,
		},
		SemicolonMissing: {
			Name:    "SemicolonMissing",
			Message: "Semicolon(:) symbol missing issue!",
			Variant: SemicolonMissing,
		},
		ColonMissing: {
			Name:    "ColonMissing",
			Message: "Colon(:) symbol missing issue!",
			Variant: ColonMissing,
		},
		EqualMissing: {
			Name:    "EqualMissing",
			Message: "Equal(=) symbol missing issue!",
			Variant: EqualMissing,
		},
		ExpectationMismatch: {
			Name:    "ExpectationMismatch",
			Message: "Expectation mismatch issue!",
			Variant: ExpectationMismatch,
		},
		StringValidationMismatch: {
			Name:    "StringValidationMismatch",
			Message: "String content validation mismatch issue!",
			Variant: StringValidationMismatch,
		},
		BooleanValidationMismatch: {
			Name:    "BooleanValidationMismatch",
			Message: "Boolean validation mismatch!",
			Variant: BooleanValidationMismatch,
		},
		IntegerValidationMismatch: {
			Name:    "IntegerValidationMismatch",
			Message: "Integer validation mismatch!",
			Variant: IntegerValidationMismatch,
		},
		FloatValidationMismatch: {
			Name:    "FloatValidationMismatch",
			Message: "Floating pointer number validation mismatch!",
			Variant: FloatValidationMismatch,
		},
		StatusValidationMismatch: {
			Name:    "StatusValidationMismatch",
			Message: "Status validation mismatch!",
			Variant: StatusValidationMismatch,
		},
		RegexValidationMismatch: {
			Name:    "RegexValidationMismatch",
			Message: "Regex validation mismatch! Regex couldn't find the expected content!",
			Variant: RegexValidationMismatch,
		},
		ClosingError: {
			Name:    "ClosingError",
			Message: "Closing error!",
			Variant: ClosingError,
		},
		FileClosing: {
			Name:    "FileClosing",
			Message: "File closing error!",
			Variant: FileClosing,
		},
		DeferError: {
			Name:    "DeferError",
			Message: "Defer running error!",
			Variant: DeferError,
		},
		MismatchValidation: {
			Name:    "MismatchValidation",
			Message: "Validation mismatch!",
			Variant: MismatchValidation,
		},
		ContentValidationFailed: {
			Name:    "ContentValidationFailed",
			Message: "Content validation mismatch!",
			Variant: ContentValidationFailed,
		},
		CommandLineValidationFailed: {
			Name:    "CommandLineValidationFailed",
			Message: "Command line validation failed!",
			Variant: CommandLineValidationFailed,
		},
		IniParsingFailed: {
			Name:    "IniParsingFailed",
			Message: "INI File parsing failed!",
			Variant: IniParsingFailed,
		},
		PsqlParsingFailed: {
			Name:    "PsqlParsingFailed",
			Message: "PostgreSQL file parsing failed!",
			Variant: PsqlParsingFailed,
		},
		HbaParsingFailed: {
			Name:    "HbaParsingFailed",
			Message: "Hba file parsing failed!",
			Variant: HbaParsingFailed,
		},
		ConfFileParsingFailed: {
			Name:    "ConfFileParsingFailed",
			Message: "Configuration (conf) file parsing failed!",
			Variant: ConfFileParsingFailed,
		},
		RegexCompiledFailed: {
			Name:    "RegexCompiledFailed",
			Message: "Regex compiled failed!",
			Variant: RegexCompiledFailed,
		},
		ChmodReadWriteFailed: {
			Name:    "ChmodReadWriteFailed",
			Message: "Chmod read/write failed!",
			Variant: ChmodReadWriteFailed,
		},
		ChmodValidationFailed: {
			Name:    "ChmodValidationFailed",
			Message: "Chmod validation failed or mismatch!",
			Variant: ChmodValidationFailed,
		},
		Validation: {
			Name:    "Validation",
			Message: "Validation failed or mismatch!",
			Variant: Validation,
		},
		ValidationMismatch: {
			Name:    "ValidationMismatch",
			Message: "Validation failed or mismatch!",
			Variant: ValidationMismatch,
		},
		TypeIssue: {
			Name:    "TypeIssue",
			Message: "Type related issue!",
			Variant: TypeIssue,
		},
		DirIssue: {
			Name:    "DirIssue",
			Message: "Directory related issue!",
			Variant: DirIssue,
		},
		ApplyFailed: {
			Name:    "ApplyFailed",
			Message: "CollectError failed!",
			Variant: ApplyFailed,
		},
		ExecuteFailed: {
			Name:    "ExecuteFailed",
			Message: "Execute (exec) failed!",
			Variant: ExecuteFailed,
		},
		UnixFileSystemIssue: {
			Name:    "UnixFileSystemIssue",
			Message: "Unix file system issue!",
			Variant: UnixFileSystemIssue,
		},
		DatabaseIssue: {
			Name:    "DatabaseIssue",
			Message: "Database issue!",
			Variant: DatabaseIssue,
		},
		DatabaseReadIssue: {
			Name:    "DatabaseReadIssue",
			Message: "Database read issue!",
			Variant: DatabaseReadIssue,
		},
		DatabaseAuthenticationFailed: {
			Name:    "DatabaseAuthenticationFailed",
			Message: "Database authentication failed!",
			Variant: DatabaseAuthenticationFailed,
		},
		DatabaseQueryFailed: {
			Name:    "DatabaseQueryFailed",
			Message: "Database query execute or parsing failed!",
			Variant: DatabaseQueryFailed,
		},
		EqualityCheckerFailed: {
			Name:    "EqualityCheckerFailed",
			Message: "Equality checker failed!",
			Variant: EqualityCheckerFailed,
		},
		ComparisonMismatch: {
			Name:    "ComparisonMismatch",
			Message: "Comparison mismatch issue!",
			Variant: ComparisonMismatch,
		},
		ComparisonFailed: {
			Name:    "ComparisonFailed",
			Message: "Comparison mismatch or comparison failed!",
			Variant: ComparisonFailed,
		},
		ComparisonMethodFailed: {
			Name:    "ComparisonMethodFailed",
			Message: "Comparison method failed!",
			Variant: ComparisonMethodFailed,
		},
		YamlParsingFailed: {
			Name:    "YamlParsingFailed",
			Message: "YAML content parsing failed!",
			Variant: YamlParsingFailed,
		},
		YamlWritingFailed: {
			Name:    "YamlWritingFailed",
			Message: "YAML file content writing failed!",
			Variant: YamlWritingFailed,
		},
		YamlReadFailed: {
			Name:    "YamlReadFailed",
			Message: "YAML file read failed!",
			Variant: YamlReadFailed,
		},
		YamlSyntaxIssue: {
			Name:    "YamlSyntaxIssue",
			Message: "YAML content syntax issue!",
			Variant: YamlSyntaxIssue,
		},
		YamlFileEmpty: {
			Name:    "YamlFileEmpty",
			Message: "YAML file content empty issue!",
			Variant: YamlFileEmpty,
		},
		ConfFileEmpty: {
			Name:    "ConfFileEmpty",
			Message: "Configuration (conf) file content empty issue!",
			Variant: ConfFileEmpty,
		},
		JsonSyntaxIssue: {
			Name:    "JsonSyntaxIssue",
			Message: "Json content syntax issue!",
			Variant: JsonSyntaxIssue,
		},
		IniSyntaxIssue: {
			Name:    "IniSyntaxIssue",
			Message: "INI content syntax issue!",
			Variant: JsonSyntaxIssue,
		},
		ConfSyntaxIssue: {
			Name:    "ConfSyntaxIssue",
			Message: "Configuration (conf) content syntax issue!",
			Variant: ConfSyntaxIssue,
		},
		FileSystemPermissionIssue: {
			Name:    "FileSystemPermissionIssue",
			Message: "File system permission issue!",
			Variant: FileSystemPermissionIssue,
		},
		LinesOrderMismatch: {
			Name:    "LinesOrderMismatch",
			Message: "Lines order mismatch issue!",
			Variant: LinesOrderMismatch,
		},
		LineIndexValidationFailed: {
			Name:    "LineIndexValidationFailed",
			Message: "Lines index validation failed!",
			Variant: LineIndexValidationFailed,
		},
		PathStat: {
			Name:    "PathStat",
			Message: "Path status or state failed!",
			Variant: PathStat,
		},
		NoConditionMatchingAsPerExpectation: {
			Name:    "NoConditionMatchingAsPerExpectation",
			Message: "None of the condition matched! One of the condition must match to finalize the request!",
			Variant: NoConditionMatchingAsPerExpectation,
		},
		AnyOfConditionShouldMatch: {
			Name:    "AnyOfConditionShouldMatch",
			Message: "None of the condition matched! One of the condition must match to finalize the request!",
			Variant: AnyOfConditionShouldMatch,
		},
		AnyOfConditionShouldMatchForCreation: {
			Name:    "AnyOfConditionShouldMatchForCreation",
			Message: "None of the condition matched for creation! One of the condition must match to finalize the creation request!",
			Variant: AnyOfConditionShouldMatchForCreation,
		},
		AnyOfConditionShouldMatchForParsing: {
			Name:    "AnyOfConditionShouldMatchForParsing",
			Message: "None of the condition matched for parsing! One of the condition must match to finalize the parsing request!",
			Variant: AnyOfConditionShouldMatchForParsing,
		},
		AnyOfConditionShouldMatchForSelecting: {
			Name:    "AnyOfConditionShouldMatchForSelecting",
			Message: "None of the condition matched for selecting! One of the condition must match to finalize the selecting request!",
			Variant: AnyOfConditionShouldMatchForSelecting,
		},
		AnyOfConditionShouldMatchForWriting: {
			Name:    "AnyOfConditionShouldMatchForWriting",
			Message: "None of the condition matched for writing! One of the condition must match to finalize the writing request!",
			Variant: AnyOfConditionShouldMatchForWriting,
		},
		AnyOfConditionShouldMatchForReading: {
			Name:    "AnyOfConditionShouldMatchForReading",
			Message: "None of the condition matched for reading! One of the condition must match to finalize the reading request!",
			Variant: AnyOfConditionShouldMatchForReading,
		},
		AnyOfConditionShouldMatchForStatus: {
			Name:    "AnyOfConditionShouldMatchForStatus",
			Message: "None of the condition matched for status case! One of the condition must match to finalize the status request!",
			Variant: AnyOfConditionShouldMatchForStatus,
		},
		AnyOfConditionShouldMatchForValue: {
			Name:    "AnyOfConditionShouldMatchForValue",
			Message: "None of the condition matched for selecting specific value! One of the condition must match to finalize the value selection request!",
			Variant: AnyOfConditionShouldMatchForValue,
		},
		AppendFailed: {
			Name:    "AppendFailed",
			Message: "Append failed!",
			Variant: AppendFailed,
		},
		PrependFailed: {
			Name:    "PrependFailed",
			Message: "Prepend failed!",
			Variant: PrependFailed,
		},
		AddNewFailed: {
			Name:    "AddNewFailed",
			Message: "Add new failed!",
			Variant: AddNewFailed,
		},
		NewFailed: {
			Name:    "NewFailed",
			Message: "New / creation failed!",
			Variant: NewFailed,
		},
		IncorrectData: {
			Name:    "IncorrectData",
			Message: "Incorrect data!",
			Variant: IncorrectData,
		},
		KeyMismatch: {
			Name:    "KeyMismatch",
			Message: "Key mismatch!",
			Variant: KeyMismatch,
		},
		ValueMismatch: {
			Name:    "ValueMismatch",
			Message: "Value mismatch, it is not as expected!",
			Variant: ValueMismatch,
		},
		KeyValueNotFound: {
			Name:    "KeyValueNotFound",
			Message: "Key-Value not found!",
			Variant: KeyValueNotFound,
		},
		KeyNotFound: {
			Name:    "KeyNotFound",
			Message: "Key not found!",
			Variant: KeyNotFound,
		},
		IdNotFound: {
			Name:    "IdNotFound",
			Message: "Id not found!",
			Variant: IdNotFound,
		},
		IdentifierNotFound: {
			Name:    "IdentifierNotFound",
			Message: "Identifier not found!",
			Variant: IdentifierNotFound,
		},
		IdentifierEmpty: {
			Name:    "IdentifierEmpty",
			Message: "Identifier empty!",
			Variant: IdentifierEmpty,
		},
		IdEmpty: {
			Name:    "IdEmpty",
			Message: "Id empty!",
			Variant: IdEmpty,
		},
		KeyEmpty: {
			Name:    "KeyEmpty",
			Message: "Key empty!",
			Variant: KeyEmpty,
		},
		KeyNotFoundInMap: {
			Name:    "KeyNotFoundInMap",
			Message: "Key not found in the map or dictionary!",
			Variant: KeyNotFoundInMap,
		},
		MapNotContainsExpectedKey: {
			Name:    "MapNotContainsExpectedKey",
			Message: "Map or dictionary doesn't contain the expected key or id!",
			Variant: MapNotContainsExpectedKey,
		},
		FileInfoEmptyOrNil: {
			Name:    "FileInfoEmptyOrNil",
			Message: "File info empty or nil!",
			Variant: FileInfoEmptyOrNil,
		},
		Apache: {
			Name:    "Apache",
			Message: "Apache related issue!",
			Variant: Apache,
		},
		CashBin: {
			Name:    "CashBin",
			Message: "Cash Bin related acl issue!",
			Variant: CashBin,
		},
		CronTabs: {
			Name:    "CronTabs",
			Message: "CronTabs related issue!",
			Variant: CronTabs,
		},
		Dovecot: {
			Name:    "Dovecot",
			Message: "Dovecot related issue!",
			Variant: Dovecot,
		},
		Environment: {
			Name:    "Environment",
			Message: "Environment related issue!",
			Variant: Environment,
		},
		Ethernet: {
			Name:    "Ethernet",
			Message: "Ethernet related issue!",
			Variant: Ethernet,
		},
		Exim4: {
			Name:    "Exim4",
			Message: "Exim4 related issue!",
			Variant: Exim4,
		},
		InstallPackage: {
			Name:    "InstallPackage",
			Message: "InstallPackage related issue!",
			Variant: InstallPackage,
		},
		LightSpeed: {
			Name:    "LightSpeed",
			Message: "LightSpeed related issue!",
			Variant: LightSpeed,
		},
		LightSpeedParsing: {
			Name:    "LightSpeedParsing",
			Message: "LightSpeed parsing related issue!",
			Variant: LightSpeedParsing,
		},
		LightSpeedConfiguration: {
			Name:    "LightSpeedConfiguration",
			Message: "LightSpeed configuration related issue!",
			Variant: LightSpeedConfiguration,
		},
		LightSpeedBroken: {
			Name:    "LightSpeedBroken",
			Message: "LightSpeed broken issue!",
			Variant: LightSpeedBroken,
		},
		MySql: {
			Name:    "MySql",
			Message: "MySql related issue!",
			Variant: MySql,
		},
		PostgreSql: {
			Name:    "PostgreSql",
			Message: "PostgreSql related issue!",
			Variant: PostgreSql,
		},
		PowerDns: {
			Name:    "PowerDns",
			Message: "PowerDns related issue!",
			Variant: PowerDns,
		},
		PureFtp: {
			Name:    "PureFtp",
			Message: "PureFtp related issue!",
			Variant: PureFtp,
		},
		Scoping: {
			Name:    "Scoping",
			Message: "Scoping related issue!",
			Variant: Scoping,
		},
		OperatingSystemGeneral: {
			Name:    "OperatingSystemGeneral",
			Message: "General operating system related issue!",
			Variant: OperatingSystemGeneral,
		},
		OperatingSystemRelated: {
			Name:    "OperatingSystemRelated",
			Message: "Operating system related issue!",
			Variant: OperatingSystemRelated,
		},
		SSL: {
			Name:    "SSL",
			Message: "SSL related issue!",
			Variant: SSL,
		},
		SSH: {
			Name:    "SSH",
			Message: "SSH related issue!",
			Variant: SSH,
		},
		PhpFpm: {
			Name:    "PhpFpm",
			Message: "PhpFpm related issue!",
			Variant: PhpFpm,
		},
		PhpAdmin: {
			Name:    "PhpAdmin",
			Message: "PhpAdmin related issue!",
			Variant: PhpAdmin,
		},
		OperatingSystemRepositoryIssue: {
			Name:    "OperatingSystemRepositoryIssue",
			Message: "Operating system repository related issue!",
			Variant: OperatingSystemRepositoryIssue,
		},
		CoreConfigIssue: {
			Name:    "CoreConfigIssue",
			Message: "Core config related issue!",
			Variant: CoreConfigIssue,
		},
		InstructionNotSupportedInOperatingSystem: {
			Name:    "InstructionNotSupportedInOperatingSystem",
			Message: "Instruction is not supported in current operating system!",
			Variant: InstructionNotSupportedInOperatingSystem,
		},
		InstallInstructionNotSupportedInOperatingSystem: {
			Name:    "InstallInstructionNotSupportedInOperatingSystem",
			Message: "Install Package instruction is not supported in current operating system!",
			Variant: InstallInstructionNotSupportedInOperatingSystem,
		},
		DockerNotSupport: {
			Name:    "DockerNotSupport",
			Message: "Docker not supported!",
			Variant: DockerNotSupport,
		},
		MethodNotSupportedInDocker: {
			Name:    "MethodNotSupportedInDocker",
			Message: "Current method not supported in docker operating system!",
			Variant: MethodNotSupportedInDocker,
		},
		MethodNotSupportedInOs: {
			Name:    "MethodNotSupportedInOs",
			Message: "Current method not supported in current operating system!",
			Variant: MethodNotSupportedInOs,
		},
		InvalidOption: {
			Name:    "InvalidOption",
			Message: "Selected option is invalid!",
			Variant: InvalidOption,
		},
		InvalidSelection: {
			Name:    "InvalidSelection",
			Message: "Current selection is invalid!",
			Variant: InvalidSelection,
		},
		InvalidOptionForFunction: {
			Name:    "InvalidOptionForFunction",
			Message: "Selected option is invalid for current function!",
			Variant: InvalidOptionForFunction,
		},
		Defined: {
			Name:    "Defined",
			Message: "Defined but expectation should be NOT defined!",
			Variant: Defined,
		},
		DefinedEmpty: {
			Name:    "DefinedEmpty",
			Message: "Defined but empty expectation should be defined with proper data or record or status or value!",
			Variant: DefinedEmpty,
		},
		NotDefined: {
			Name:    "NotDefined",
			Message: "Not defined or undefined expectation should be defined with proper data or record or status or value!",
			Variant: NotDefined,
		},
		DefinedUninitialized: {
			Name:    "DefinedUninitialized",
			Message: "Defined as uninitialized, value needs to be initialized properly to proceed further!",
			Variant: DefinedUninitialized,
		},
		AlreadyDefined: {
			Name:    "AlreadyDefined",
			Message: "Expectation should not be defined before but it was defined already!",
			Variant: AlreadyDefined,
		},
		CaseMismatch: {
			Name:    "CaseMismatch",
			Message: "Case doesn't match properly!",
			Variant: CaseMismatch,
		},
		NoCaseMatch: {
			Name:    "NoCaseMatch",
			Message: "None of cases matches properly!",
			Variant: NoCaseMatch,
		},
		Mutex: {
			Name:    "Mutex",
			Message: "Mutex issue!",
			Variant: Mutex,
		},
		Lock: {
			Name:    "Lock",
			Message: "Lock issue!",
			Variant: Lock,
		},
		WaitGroupIssue: {
			Name:    "WaitGroupIssue",
			Message: "Wait group issue!",
			Variant: WaitGroupIssue,
		},
		UnAuthorized: {
			Name:    "UnAuthorized",
			Message: "Unauthorized request, cannot proceed further!",
			Variant: UnAuthorized,
		},
		Hash: {
			Name:    "Hash",
			Message: "Hashing issue!",
			Variant: Hash,
		},
		HashCrypto: {
			Name:    "HashCrypto",
			Message: "HashCrypto issue!",
			Variant: HashCrypto,
		},
		Md5: {
			Name:    "Md5",
			Message: "Md5 related hashing issue!",
			Variant: Md5,
		},
		Sha1: {
			Name:    "Sha1",
			Message: "Sha1 related hashing issue!",
			Variant: Sha1,
		},
		Sha256: {
			Name:    "Sha256",
			Message: "Sha256 related hashing issue!",
			Variant: Sha256,
		},
		Sha512: {
			Name:    "Sha512",
			Message: "Sha512 related hashing issue!",
			Variant: Sha512,
		},
		CheckSum: {
			Name:    "CheckSum",
			Message: "CheckSum related hashing issue, mismatch or cannot generate properly!",
			Variant: CheckSum,
		},
		CheckSumMismatch: {
			Name:    "CheckSumMismatch",
			Message: "CheckSum mismatch, failed to verify checksum!",
			Variant: CheckSumMismatch,
		},
		DuplicateIssue: {
			Name:    "DuplicateIssue",
			Message: "Value/Record/Object/Key/File/Path is duplicated, which is unexpected!",
			Variant: DuplicateIssue,
		},
		StringCorrupted: {
			Name:    "StringCorrupted",
			Message: "String value corrupted, which is unexpected!",
			Variant: StringCorrupted,
		},
		DataCorrupted: {
			Name:    "DataCorrupted",
			Message: "Data value corrupted, which is unexpected!",
			Variant: DataCorrupted,
		},
		StatusCorrupted: {
			Name:    "StatusCorrupted",
			Message: "Status corrupted, which is unexpected!",
			Variant: StatusCorrupted,
		},
		CheckSumCorrupted: {
			Name:    "CheckSumCorrupted",
			Message: "Checksum value corrupted, which is unexpected!",
			Variant: CheckSumCorrupted,
		},
		DeferIssue: {
			Name:    "DeferIssue",
			Message: "Defer related issue!",
			Variant: DeferIssue,
		},
		Undefined: {
			Name:    "Undefined",
			Message: "Undefined but should be defined to proceed further!",
			Variant: Undefined,
		},
		UnInitialized: {
			Name:    "UnInitialized",
			Message: "UnInitialized but should be initialized properly to proceed further!",
			Variant: UnInitialized,
		},
		Unmapped: {
			Name:    "Unmapped",
			Message: "Unmapped value or records or data, should be mapped!",
			Variant: Unmapped,
		},
		InvalidInstruction: {
			Name:    "InvalidInstruction",
			Message: "Give instruction is invalid!",
			Variant: InvalidInstruction,
		},
		Default: {
			Name:    "Default",
			Message: "Default is invalid!",
			Variant: Default,
		},
		FileOpen: {
			Name:    "FileOpen",
			Message: "File opening related issue, probably file is invalid or directory no file or permission issue or locked by other process!",
			Variant: FileOpen,
		},
		FileOpenOrRead: {
			Name:    "FileOpenOrRead",
			Message: "File opening or read related issue, probably file is invalid or directory no file or permission issue or locked by other process!",
			Variant: FileOpenOrRead,
		},
		FileCreate: {
			Name:    "FileCreate",
			Message: "File create failed!",
			Variant: FileOpenOrRead,
		},
		FileDelete: {
			Name:    "FileDelete",
			Message: "File delete failed!",
			Variant: FileDelete,
		},
		FileEdit: {
			Name:    "FileEdit",
			Message: "File edit failed!",
			Variant: FileEdit,
		},
		FilePrepend: {
			Name:    "FilePrepend",
			Message: "File prepend failed!",
			Variant: FilePrepend,
		},
		FileCrud: {
			Name:    "FileCrud",
			Message: "File CRUD failed!",
			Variant: FileCrud,
		},
		DatabaseCrud: {
			Name:    "DatabaseCrud",
			Message: "Database CRUD (Create or Read or Update or Delete) failed!",
			Variant: DatabaseCrud,
		},
		RecordCreate: {
			Name:    "RecordCreate",
			Message: "Record create failed!",
			Variant: RecordCreate,
		},
		RecordSave: {
			Name:    "RecordSave",
			Message: "Record save or edit failed!",
			Variant: RecordSave,
		},
		RecordEdit: {
			Name:    "RecordEdit",
			Message: "Record edit failed!",
			Variant: RecordEdit,
		},
		RecordIdMissing: {
			Name:    "RecordIdMissing",
			Message: "Record id or identifier missing!",
			Variant: RecordIdMissing,
		},
		RecordPrimaryKeyMissing: {
			Name:    "RecordPrimaryKeyMissing",
			Message: "Record primary key missing!",
			Variant: RecordPrimaryKeyMissing,
		},
		RecordForeignKeyMissing: {
			Name:    "RecordForeignKeyMissing",
			Message: "Record foreign key missing!",
			Variant: RecordForeignKeyMissing,
		},
		RecordDelete: {
			Name:    "RecordDelete",
			Message: "Record delete failed!",
			Variant: RecordDelete,
		},
		ObjectDelete: {
			Name:    "ObjectDelete",
			Message: "Object delete failed!",
			Variant: ObjectDelete,
		},
		ObjectInsert: {
			Name:    "ObjectInsert",
			Message: "Object insert failed!",
			Variant: ObjectInsert,
		},
		OrmInsert: {
			Name:    "OrmInsert",
			Message: "Orm insert failed!",
			Variant: OrmInsert,
		},
		OrmCreate: {
			Name:    "OrmCreate",
			Message: "Orm create failed!",
			Variant: OrmCreate,
		},
		OrmEdit: {
			Name:    "OrmEdit",
			Message: "Orm edit failed!",
			Variant: OrmEdit,
		},
		OrmSave: {
			Name:    "OrmSave",
			Message: "Orm save or edit failed!",
			Variant: OrmSave,
		},
		OrmDelete: {
			Name:    "OrmDelete",
			Message: "Orm delete failed!",
			Variant: OrmDelete,
		},
		UnChanged: {
			Name:    "UnChanged",
			Message: "Value or Record or Data is unchanged but expectation should be changed or mutated, failed the expectation!",
			Variant: UnChanged,
		},
		FileUnChanged: {
			Name:    "FileUnChanged",
			Message: "File data is unchanged but expectation should be changed or mutated, failed the expectation!",
			Variant: FileUnChanged,
		},
		FileLocked: {
			Name:    "FileLocked",
			Message: "File locked!",
			Variant: FileLocked,
		},
		FilePermissionIssue: {
			Name:    "FilePermissionIssue",
			Message: "File permission related issue!",
			Variant: FilePermissionIssue,
		},
		FileInvalid: {
			Name:    "FileInvalid",
			Message: "File invalid or missing issue!",
			Variant: FileInvalid,
		},
		FileInvalidOrMissing: {
			Name:    "FileInvalidOrMissing",
			Message: "File invalid or missing issue!",
			Variant: FileInvalidOrMissing,
		},
		Open: {
			Name:    "Open",
			Message: "Open failed!",
			Variant: Open,
		},
		FileExpand: {
			Name:    "FileExpand",
			Message: "File expand failed!",
			Variant: FileExpand,
		},
		Seek: {
			Name:    "Seek",
			Message: "Seek failed!",
			Variant: Seek,
		},
		Commit: {
			Name:    "Commit",
			Message: "Commit failed!",
			Variant: Commit,
		},
		FileSync: {
			Name:    "FileSync",
			Message: "File sync failed!",
			Variant: FileSync,
		},
		FileCommit: {
			Name:    "FileCommit",
			Message: "File commit failed!",
			Variant: FileCommit,
		},
		DataCommit: {
			Name:    "DataCommit",
			Message: "Data commit failed!",
			Variant: DataCommit,
		},
		RecordCommit: {
			Name:    "RecordCommit",
			Message: "Record commit failed!",
			Variant: RecordCommit,
		},
		StatusCommit: {
			Name:    "StatusCommit",
			Message: "Status commit failed!",
			Variant: StatusCommit,
		},
		DependencyIssue: {
			Name:    "DependencyIssue",
			Message: "Dependency issue!",
			Variant: DependencyIssue,
		},
		DependencyApiIssue: {
			Name:    "DependencyApiIssue",
			Message: "Dependency api issue!",
			Variant: DependencyApiIssue,
		},
		CircularDependency: {
			Name:    "CircularDependency",
			Message: "Circular dependency issue!",
			Variant: CircularDependency,
		},
		DependencyNotFound: {
			Name:    "DependencyNotFound",
			Message: "Dependency not found or missing!",
			Variant: DependencyNotFound,
		},
		DependencyMissing: {
			Name:    "DependencyMissing",
			Message: "Dependency not found or missing!",
			Variant: DependencyMissing,
		},
		PackageMissing: {
			Name:    "PackageMissing",
			Message: "Package not found or missing!",
			Variant: PackageMissing,
		},
		IdentifierMissing: {
			Name:    "IdentifierMissing",
			Message: "Identifier not found or missing!",
			Variant: IdentifierMissing,
		},
		IdentifierValidationFailed: {
			Name:    "IdentifierValidationFailed",
			Message: "Identifier validation failed!",
			Variant: IdentifierValidationFailed,
		},
		KeyValidationFailed: {
			Name:    "KeyValidationFailed",
			Message: "Key validation failed!",
			Variant: KeyValidationFailed,
		},
		IdValidationFailed: {
			Name:    "IdValidationFailed",
			Message: "Id validation failed!",
			Variant: IdValidationFailed,
		},
		SectionValidationFailed: {
			Name:    "SectionValidationFailed",
			Message: "Section validation failed!",
			Variant: SectionValidationFailed,
		},
		GroupValidationFailed: {
			Name:    "GroupValidationFailed",
			Message: "Group validation failed!",
			Variant: GroupValidationFailed,
		},
		TypeValidationFailed: {
			Name:    "TypeValidationFailed",
			Message: "Type validation failed!",
			Variant: TypeValidationFailed,
		},
		HeaderValidationFailed: {
			Name:    "HeaderValidationFailed",
			Message: "Header validation failed!",
			Variant: HeaderValidationFailed,
		},
		InstructionValidationFailed: {
			Name:    "InstructionValidationFailed",
			Message: "Instruction validation failed!",
			Variant: InstructionValidationFailed,
		},
		LineNotFound: {
			Name:    "LineNotFound",
			Message: "Lines or Line not found or missing!",
			Variant: LineNotFound,
		},
		SectionNotFound: {
			Name:    "SectionNotFound",
			Message: "Section or Sections not found or missing!",
			Variant: SectionNotFound,
		},
		TypeNotFound: {
			Name:    "TypeNotFound",
			Message: "Type not found or missing!",
			Variant: SectionNotFound,
		},
		ItemNotFound: {
			Name:    "ItemNotFound",
			Message: "Item or items not found or missing!",
			Variant: SectionNotFound,
		},
		ItemMissing: {
			Name:    "ItemMissing",
			Message: "Item or items not found or missing!",
			Variant: ItemMissing,
		},
		CategoryNotFound: {
			Name:    "CategoryNotFound",
			Message: "Category or categories not found or missing!",
			Variant: CategoryNotFound,
		},
		HeaderNotFound: {
			Name:    "HeaderNotFound",
			Message: "Header or headers not found or missing!",
			Variant: HeaderNotFound,
		},
		RootNotFound: {
			Name:    "RootNotFound",
			Message: "Root not found or undefined!",
			Variant: RootNotFound,
		},
		MapNotFound: {
			Name:    "MapNotFound",
			Message: "Map not found or undefined!",
			Variant: MapNotFound,
		},
		DataNotFound: {
			Name:    "DataNotFound",
			Message: "Data not found or missing or undefined!",
			Variant: DataNotFound,
		},
		RecordNotFound: {
			Name:    "RecordNotFound",
			Message: "Record not found or missing or undefined!",
			Variant: RecordNotFound,
		},
		RecordMissing: {
			Name:    "RecordMissing",
			Message: "Record not found or missing or undefined!",
			Variant: RecordMissing,
		},
		KeyMissing: {
			Name:    "KeyMissing",
			Message: "Key not found or missing or undefined!",
			Variant: KeyMissing,
		},
		RootMissing: {
			Name:    "RootMissing",
			Message: "Root not found or missing or undefined!",
			Variant: RootMissing,
		},
		SectionMissing: {
			Name:    "SectionMissing",
			Message: "Section not found or missing or undefined!",
			Variant: RootMissing,
		},
		IdMissing: {
			Name:    "IdMissing",
			Message: "Id not found or missing or undefined!",
			Variant: IdMissing,
		},
		HeaderMissing: {
			Name:    "HeaderMissing",
			Message: "Header not found or missing or undefined!",
			Variant: HeaderMissing,
		},
		RequestHeaderMissing: {
			Name:    "RequestHeaderMissing",
			Message: "Request header not found or missing or undefined!",
			Variant: RequestHeaderMissing,
		},
		RequestDefinitionMissing: {
			Name:    "RequestDefinitionMissing",
			Message: "Request header not found or missing or undefined!",
			Variant: RequestDefinitionMissing,
		},
		TypeMissing: {
			Name:    "TypeMissing",
			Message: "Type not found or missing or undefined!",
			Variant: TypeMissing,
		},
		GroupMissing: {
			Name:    "GroupMissing",
			Message: "Group not found or missing or undefined!",
			Variant: GroupMissing,
		},
		DataMissing: {
			Name:    "DataMissing",
			Message: "Data not found or missing or undefined!",
			Variant: DataMissing,
		},
		StatusMissing: {
			Name:    "StatusMissing",
			Message: "Status not found or missing or undefined!",
			Variant: StatusMissing,
		},
		InstructionMissing: {
			Name:    "InstructionMissing",
			Message: "Instruction or Instructions not found or missing or undefined!",
			Variant: InstructionMissing,
		},
		ElementMissing: {
			Name:    "ElementMissing",
			Message: "Element(s) not found or missing or undefined!",
			Variant: ElementMissing,
		},
		ValueOrValuesMissing: {
			Name:    "ValueOrValuesMissing",
			Message: "Value or values not found or missing or undefined!",
			Variant: ValueOrValuesMissing,
		},
		InstructionRootElementMissing: {
			Name:    "InstructionRootElementMissing",
			Message: "Instruction root element not found or missing or undefined!",
			Variant: InstructionRootElementMissing,
		},
		InstructionNotDefined: {
			Name:    "InstructionNotDefined",
			Message: "Instruction not found or missing or undefined!",
			Variant: InstructionNotDefined,
		},
		RecordInsert: {
			Name:    "RecordInsert",
			Message: "Record insert failed!",
			Variant: RecordInsert,
		},
		RecordUndefined: {
			Name:    "RecordUndefined",
			Message: "Record undefined!",
			Variant: RecordUndefined,
		},
		RecordNotInitialized: {
			Name:    "RecordNotInitialized",
			Message: "Record not initialized properly!",
			Variant: RecordNotInitialized,
		},
		NotInitialized: {
			Name:    "NotInitialized",
			Message: "Not initialized properly!",
			Variant: NotInitialized,
		},
		RequestValidationFailed: {
			Name:    "RequestValidationFailed",
			Message: "Request validation failed!",
			Variant: RequestValidationFailed,
		},
		HtmlValidationFailed: {
			Name:    "HtmlValidationFailed",
			Message: "HTML validation failed!",
			Variant: HtmlValidationFailed,
		},
		TagValidationFailed: {
			Name:    "TagValidationFailed",
			Message: "Tag validation failed!",
			Variant: TagValidationFailed,
		},
		JsonValidationFailed: {
			Name:    "JsonValidationFailed",
			Message: "JSON validation failed!",
			Variant: JsonValidationFailed,
		},
		YamlValidationFailed: {
			Name:    "YamlValidationFailed",
			Message: "YAML validation failed!",
			Variant: YamlValidationFailed,
		},
		ConfigValidationFailed: {
			Name:    "ConfigValidationFailed",
			Message: "Configuration validation failed!",
			Variant: ConfigValidationFailed,
		},
		DataUnordered: {
			Name:    "DataUnordered",
			Message: "Data is not in proper order as per expectation!",
			Variant: DataUnordered,
		},
		RecordUnordered: {
			Name:    "RecordUnordered",
			Message: "Record is not in proper order as per expectation!",
			Variant: RecordUnordered,
		},
		ValueOrValuesUnordered: {
			Name:    "ValueOrValuesUnordered",
			Message: "Value or values are not in proper order as per expectation!",
			Variant: ValueOrValuesUnordered,
		},
		ItemsUnordered: {
			Name:    "ItemsUnordered",
			Message: "Item(s) is/(are) not in proper order as per expectation!",
			Variant: ItemsUnordered,
		},
		PowerDnsAddFailed: {
			Name:    "PowerDnsAddFailed",
			Message: "Addition or creation failed for powerdns (record resource (RR) or zone or domain)!",
			Variant: PowerDnsAddFailed,
		},
		PowerDnsEditFailed: {
			Name:    "PowerDnsEditFailed",
			Message: "Edit failed for powerdns (record resource (RR) or zone or domain)!",
			Variant: PowerDnsEditFailed,
		},
		PowerDnsDomainOrZoneMissing: {
			Name:    "PowerDnsDomainOrZoneMissing",
			Message: "PowerDNS Domain or zone missing or not found!",
			Variant: PowerDnsDomainOrZoneMissing,
		},
		PowerDnsUpdateFailed: {
			Name:    "PowerDnsUpdateFailed",
			Message: "PowerDNS updated failed for domain or zone or record resource (RR)!",
			Variant: PowerDnsUpdateFailed,
		},
		PowerDnsRecordResourceIssue: {
			Name:    "PowerDnsRecordResourceIssue",
			Message: "PowerDNS record resource (RR) related issue!",
			Variant: PowerDnsRecordResourceIssue,
		},
		PowerDnsRecordResourceCreateIssue: {
			Name:    "PowerDnsRecordResourceCreateIssue",
			Message: "PowerDNS record resource (RR) create issue!",
			Variant: PowerDnsRecordResourceCreateIssue,
		},
		PowerDnsRecordResourceDeleteIssue: {
			Name:    "PowerDnsRecordResourceDeleteIssue",
			Message: "PowerDNS record resource (RR) delete issue!",
			Variant: PowerDnsRecordResourceDeleteIssue,
		},
		PowerDnsRecordResourceMissing: {
			Name:    "PowerDnsRecordResourceMissing",
			Message: "PowerDNS record resource (RR) missing or not found by the query!",
			Variant: PowerDnsRecordResourceMissing,
		},
		PureFtpPermissionIssue: {
			Name:    "PureFtpPermissionIssue",
			Message: "PureFTP permission related issue!",
			Variant: PureFtpPermissionIssue,
		},
		PureFtpCrudIssue: {
			Name:    "PureFtpCrudIssue",
			Message: "PureFTP CRUD issue!",
			Variant: PureFtpCrudIssue,
		},
		PureFtpUserCreateFailed: {
			Name:    "PureFtpUserCreateFailed",
			Message: "PureFTP user creation failed!",
			Variant: PureFtpUserCreateFailed,
		},
		PureFtpUserMissing: {
			Name:    "PureFtpUserMissing",
			Message: "PureFTP user missing or not found!",
			Variant: PureFtpUserMissing,
		},
		PureFtpUserEditFailed: {
			Name:    "PureFtpUserEditFailed",
			Message: "PureFTP user edit failed!",
			Variant: PureFtpUserEditFailed,
		},
		PureFtpConfigEditFailed: {
			Name:    "PureFtpConfigEditFailed",
			Message: "PureFTP config edit failed!",
			Variant: PureFtpConfigEditFailed,
		},
		PureFtpConfigParsingIssue: {
			Name:    "PureFtpConfigParsingIssue",
			Message: "PureFTP config parsing issue!",
			Variant: PureFtpConfigParsingIssue,
		},
		PureFtpConfigMissing: {
			Name:    "PureFtpConfigMissing",
			Message: "PureFTP config missing issue!",
			Variant: PureFtpConfigMissing,
		},
		LengthMismatch: {
			Name:    "LengthMismatch",
			Message: "Length mismatch!",
			Variant: LengthMismatch,
		},
		Exim4ConfigIssue: {
			Name:    "Exim4ConfigIssue",
			Message: "Exim4 config related issue!",
			Variant: Exim4ConfigIssue,
		},
		Exim4UserIssue: {
			Name:    "Exim4UserIssue",
			Message: "Exim4 user related issue!",
			Variant: Exim4UserIssue,
		},
		Exim4UserCreateIssue: {
			Name:    "Exim4UserCreateIssue",
			Message: "Exim4 user creation issue!",
			Variant: Exim4UserCreateIssue,
		},
		Exim4UserEditIssue: {
			Name:    "Exim4UserEditIssue",
			Message: "Exim4 user edit issue!",
			Variant: Exim4UserEditIssue,
		},
		Exim4UserDeleteIssue: {
			Name:    "Exim4UserDeleteIssue",
			Message: "Exim4 user delete issue!",
			Variant: Exim4UserDeleteIssue,
		},
		Exim4ConfCrudIssue: {
			Name:    "Exim4ConfCrudIssue",
			Message: "Exim4 config CRUD issue!",
			Variant: Exim4ConfCrudIssue,
		},
		Exim4ConfDeleteIssue: {
			Name:    "Exim4ConfDeleteIssue",
			Message: "Exim4 config delete issue!",
			Variant: Exim4ConfDeleteIssue,
		},
		Exim4ConfCreateIssue: {
			Name:    "Exim4ConfCreateIssue",
			Message: "Exim4 config create issue!",
			Variant: Exim4ConfCreateIssue,
		},
		NotValidDirectory: {
			Name:    "NotValidDirectory",
			Message: "Not a valid directory or doesn't exist in the system!",
			Variant: NotValidDirectory,
		},
		NotValidFile: {
			Name:    "NotValidFile",
			Message: "Not a valid file or doesn't exist in the system!",
			Variant: NotValidFile,
		},
		ServiceNotWorkingProperly: {
			Name:    "ServiceNotWorkingProperly",
			Message: "Service not working properly!",
			Variant: ServiceNotWorkingProperly,
		},
		OperatingSystemServiceNotWorkingProperly: {
			Name:    "OperatingSystemServiceNotWorkingProperly",
			Message: "Operating system service not working properly!",
			Variant: OperatingSystemServiceNotWorkingProperly,
		},
		ServiceFailedToReceiveStatus: {
			Name:    "ServiceFailedToReceiveStatus",
			Message: "Service failed to receive status!",
			Variant: ServiceFailedToReceiveStatus,
		},
		ServiceInvalid: {
			Name:    "ServiceInvalid",
			Message: "Service invalid!",
			Variant: ServiceInvalid,
		},
		UnknownService: {
			Name:    "UnknownService",
			Message: "Service is unknown in the operating system!",
			Variant: UnknownService,
		},
		ServiceDead: {
			Name:    "ServiceDead",
			Message: "Service is dead!",
			Variant: ServiceDead,
		},
		ServiceNotRunning: {
			Name:    "ServiceNotRunning",
			Message: "Service is not running in the system!",
			Variant: ServiceNotRunning,
		},
		ServiceDeadButPidExist: {
			Name:    "ServiceDeadButPidExist",
			Message: "Service dead but pid exist!",
			Variant: ServiceDeadButPidExist,
		},
		SystemDOrSystemCtlMissing: {
			Name:    "SystemDOrSystemCtlMissing",
			Message: "systemd or systemctl is not present in the system!",
			Variant: SystemDOrSystemCtlMissing,
		},
		ServiceStatusCodeMismatch: {
			Name:    "ServiceStatusCodeMismatch",
			Message: "service status code or exit code mismatch!",
			Variant: ServiceStatusCodeMismatch,
		},
		ExitCodeMismatch: {
			Name:    "ExitCodeMismatch",
			Message: "Exit code mismatch!",
			Variant: ExitCodeMismatch,
		},
		StatusMismatch: {
			Name:    "StatusMismatch",
			Message: "Status mismatch!",
			Variant: StatusMismatch,
		},
		UpdateFailed: {
			Name:    "UpdateFailed",
			Message: "Update failed!",
			Variant: UpdateFailed,
		},
		DbUpdateFailed: {
			Name:    "DbUpdateFailed",
			Message: "Database update failed!",
			Variant: DbUpdateFailed,
		},
		DbCreateFailed: {
			Name:    "DbCreateFailed",
			Message: "Database creation failed!",
			Variant: DbCreateFailed,
		},
		DbEditFailed: {
			Name:    "DbEditFailed",
			Message: "Database edit failed!",
			Variant: DbEditFailed,
		},
		DbDeleteFailed: {
			Name:    "DbDeleteFailed",
			Message: "Database delete or remove operation failed!",
			Variant: DbDeleteFailed,
		},
		DbCrudFailed: {
			Name:    "DbCrudFailed",
			Message: "Database CRUD - (Create or Read or Update or Delete) operation failed!",
			Variant: DbCrudFailed,
		},
		DbLocked: {
			Name:    "DbLocked",
			Message: "Database locked, operation failed!",
			Variant: DbLocked,
		},
		DbFileMissing: {
			Name:    "DbFileMissing",
			Message: "Database file missing, operation failed!",
			Variant: DbFileMissing,
		},
		RedisCrudFailed: {
			Name:    "RedisCrudFailed",
			Message: "Redis CRUD (Create or Read or Update or Delete) operation failed!",
			Variant: RedisCrudFailed,
		},
		RedisUpdateFailed: {
			Name:    "RedisUpdateFailed",
			Message: "Redis update operation failed!",
			Variant: RedisUpdateFailed,
		},
		RedisEditFailed: {
			Name:    "RedisEditFailed",
			Message: "Redis edit operation failed!",
			Variant: RedisEditFailed,
		},
		RedisDeleteFailed: {
			Name:    "RedisDeleteFailed",
			Message: "Redis delete operation failed!",
			Variant: RedisDeleteFailed,
		},
		RedisCreateFailed: {
			Name:    "RedisCreateFailed",
			Message: "Redis create operation failed!",
			Variant: RedisCreateFailed,
		},
		RedisNullCmd: {
			Name:    "RedisNullCmd",
			Message: "Redis cmd not found thus null!",
			Variant: RedisNullCmd,
		},
		RedisNotAvailable: {
			Name:    "RedisNotAvailable",
			Message: "Redis not available!",
			Variant: RedisNotAvailable,
		},
		RedisConnectionIssue: {
			Name:    "RedisConnectionIssue",
			Message: "Redis connection issue!",
			Variant: RedisConnectionIssue,
		},
		RedisNullValue: {
			Name:    "RedisNullValue",
			Message: "Redis null value!",
			Variant: RedisNullValue,
		},
		RedisAdd: {
			Name:    "RedisAdd",
			Message: "Redis add failed!",
			Variant: RedisAdd,
		},
		RedisList: {
			Name:    "RedisList",
			Message: "Redis list issue or read list issue!",
			Variant: RedisList,
		},
		RedisMap: {
			Name:    "RedisMap",
			Message: "Redis map issue or read map issue!",
			Variant: RedisMap,
		},
		RecordUpdateFailed: {
			Name:    "RecordUpdateFailed",
			Message: "Record update operation failed!",
			Variant: RecordUpdateFailed,
		},
		RecordEditFailed: {
			Name:    "RecordEditFailed",
			Message: "Record edit operation failed!",
			Variant: RecordEditFailed,
		},
		RecordDeleteFailed: {
			Name:    "RecordDeleteFailed",
			Message: "Record delete or remove operation failed!",
			Variant: RecordDeleteFailed,
		},
		RecordCreationFailed: {
			Name:    "RecordCreationFailed",
			Message: "Record create operation failed!",
			Variant: RecordCreationFailed,
		},
		DbRecordNotFound: {
			Name:    "DbRecordNotFound",
			Message: "Db Record not found by specific query or key!",
			Variant: DbRecordNotFound,
		},
		ParentDirCreateFailed: {
			Name:    "ParentDirCreateFailed",
			Message: "Prent dir or base dir creation failed!",
			Variant: ParentDirCreateFailed,
		},
		DbMigration: {
			Name:    "DbMigration",
			Message: "Db migration failed, serious system fault!",
			Variant: DbMigration,
		},
		Migration: {
			Name:    "Migration",
			Message: "Migration issue or failed, serious system fault!",
			Variant: Migration,
		},
		VersionIssue: {
			Name:    "VersionIssue",
			Message: "Version issue or mismatch issue!",
			Variant: VersionIssue,
		},
		ConversionFailed: {
			Name:    "ConversionFailed",
			Message: "Conversion failed!",
			Variant: ConversionFailed,
		},
		ConversionInterfaceToSpecificStruct: {
			Name:    "ConversionInterfaceToSpecificStruct",
			Message: "Conversion from interface to specific struct is failed!",
			Variant: ConversionInterfaceToSpecificStruct,
		},
		StructureIntegrityIssue: {
			Name:    "StructureIntegrityIssue",
			Message: "Structure integrity issue!",
			Variant: StructureIntegrityIssue,
		},
		CachingIntegrityIssue: {
			Name:    "CachingIntegrityIssue",
			Message: "Caching integrity issue!",
			Variant: CachingIntegrityIssue,
		},
		IntegrityIssue: {
			Name:    "IntegrityIssue",
			Message: "Integrity issue!",
			Variant: IntegrityIssue,
		},
		ChownCopyFailed: {
			Name:    "ChownCopyFailed",
			Message: "Chown (userid or group id or both) copy failed!",
			Variant: ChownCopyFailed,
		},
		ChmodCopyFailed: {
			Name:    "ChmodCopyFailed",
			Message: "Chmod or permission copy failed!",
			Variant: ChmodCopyFailed,
		},
		SnapshotFailed: {
			Name:    "SnapshotFailed",
			Message: "Snapshot failed!",
			Variant: SnapshotFailed,
		},
		SnapshotIndexNotFound: {
			Name:    "SnapshotIndexNotFound",
			Message: "Snapshot index not found!",
			Variant: SnapshotIndexNotFound,
		},
		SnapshotNotFound: {
			Name:    "SnapshotNotFound",
			Message: "Snapshot not found!",
			Variant: SnapshotNotFound,
		},
		SnapshotKeyNotFound: {
			Name:    "SnapshotKeyNotFound",
			Message: "Snapshot key not found!",
			Variant: SnapshotKeyNotFound,
		},
		BackupNotFound: {
			Name:    "BackupNotFound",
			Message: "Backup not found!",
			Variant: BackupNotFound,
		},
		BackupKeyNotFound: {
			Name:    "BackupKeyNotFound",
			Message: "Backup key not found!",
			Variant: BackupKeyNotFound,
		},
		BackupDbRecordMissing: {
			Name:    "BackupDbRecordMissing",
			Message: "Backup db record missing!",
			Variant: BackupDbRecordMissing,
		},
		BackupDbRecordNotFound: {
			Name:    "BackupDbRecordNotFound",
			Message: "Backup db record not found!",
			Variant: BackupDbRecordNotFound,
		},
		SnapshotDbRecordMissing: {
			Name:    "SnapshotDbRecordMissing",
			Message: "Backup db record missing!",
			Variant: SnapshotDbRecordMissing,
		},
		SnapshotDbRecordNotFound: {
			Name:    "SnapshotDbRecordNotFound",
			Message: "Snapshot db record not found!",
			Variant: SnapshotDbRecordNotFound,
		},
		DbRecordKeyNotFound: {
			Name:    "DbRecordKeyNotFound",
			Message: "Db record key not found!",
			Variant: DbRecordKeyNotFound,
		},
		IndexInvalid: {
			Name:    "IndexInvalid",
			Message: "Index given invalid!",
			Variant: IndexInvalid,
		},
		IndexIssue: {
			Name:    "IndexIssue",
			Message: "Index issue!",
			Variant: IndexIssue,
		},
		IndexOutOfRange: {
			Name:    "IndexOutOfRange",
			Message: "Index out of range or out of boundary!",
			Variant: IndexOutOfRange,
		},
		CompileErrors: {
			Name:    "CompileErrors",
			Message: "Compiled or merged error of collection!",
			Variant: CompileErrors,
		},
		RedisKeyNotFound: {
			Name:    "RedisKeyNotFound",
			Message: "Redis key not found!",
			Variant: RedisKeyNotFound,
		},
		RedisIssue: {
			Name:    "RedisIssue",
			Message: "Redis generic issue!",
			Variant: RedisIssue,
		},
		RedisCrud: {
			Name:    "RedisCrud",
			Message: "Redis CRUD (create or read or update or delete) related issue!",
			Variant: RedisCrud,
		},
		RedisDataMissing: {
			Name:    "RedisDataMissing",
			Message: "Redis data missing!",
			Variant: RedisDataMissing,
		},
		RedisRelationRecordMissing: {
			Name:    "RedisRelationRecordMissing",
			Message: "Redis related record missing!",
			Variant: RedisRelationRecordMissing,
		},
		RedisNotRunning: {
			Name:    "RedisRelationRecordMissing",
			Message: "Redis related record missing!",
			Variant: RedisRelationRecordMissing,
		},
		RedisIntegrityLost: {
			Name:    "RedisIntegrityLost",
			Message: "Redis integrity lost!",
			Variant: RedisIntegrityLost,
		},
		RedisSeedingDataMissing: {
			Name:    "RedisSeedingDataMissing",
			Message: "Redis seeding data missing!",
			Variant: RedisSeedingDataMissing,
		},
		RedisRetrieveRecordIssue: {
			Name:    "RedisRetrieveRecordIssue",
			Message: "Redis retrieve record issue!",
			Variant: RedisRetrieveRecordIssue,
		},
		RedisUnknownIssue: {
			Name:    "RedisUnknownIssue",
			Message: "Redis unknown issue!",
			Variant: RedisUnknownIssue,
		},
		RedisConnectionLost: {
			Name:    "RedisConnectionLost",
			Message: "Redis connection lost!",
			Variant: RedisConnectionLost,
		},
		RedisKeyAlreadyExist: {
			Name:    "RedisKeyAlreadyExist",
			Message: "Redis key already exist!",
			Variant: RedisKeyAlreadyExist,
		},
		RedisDataStateMissing: {
			Name:    "RedisDataStateMissing",
			Message: "Redis data state missing!",
			Variant: RedisDataStateMissing,
		},
		RedisRecordStatusMissing: {
			Name:    "RedisRecordStatusMissing",
			Message: "Redis record status missing!",
			Variant: RedisRecordStatusMissing,
		},
		RedisWriteFailed: {
			Name:    "RedisWriteFailed",
			Message: "Redis write failed!",
			Variant: RedisWriteFailed,
		},
		RedisReadFailed: {
			Name:    "RedisReadFailed",
			Message: "Redis read failed!",
			Variant: RedisReadFailed,
		},
		RedisRecordMismatch: {
			Name:    "RedisRecordMismatch",
			Message: "Redis record mismatch!",
			Variant: RedisRecordMismatch,
		},
		DbRecordMismatch: {
			Name:    "DbRecordMismatch",
			Message: "Database record mismatch!",
			Variant: DbRecordMismatch,
		},
		Critical: {
			Name:    "Critical",
			Message: "Critical!",
			Variant: Critical,
		},
		CriticalData: {
			Name:    "CriticalData",
			Message: "Critical data!",
			Variant: CriticalData,
		},
		CriticalStatus: {
			Name:    "CriticalStatus",
			Message: "Critical data!",
			Variant: CriticalStatus,
		},
		DomainMissing: {
			Name:    "DomainMissing",
			Message: "Domain missing!",
			Variant: DomainMissing,
		},
		ServerMissing: {
			Name:    "ServerMissing",
			Message: "Server missing!",
			Variant: ServerMissing,
		},
		ServerBlockIssue: {
			Name:    "ServerBlockIssue",
			Message: "Server block issue!",
			Variant: ServerBlockIssue,
		},
		ServerState: {
			Name:    "ServerState",
			Message: "Server state issue!",
			Variant: ServerState,
		},
		ServerStateUnknown: {
			Name:    "ServerStateUnknown",
			Message: "Server state unknown!",
			Variant: ServerStateUnknown,
		},
		ServerStateCorrupted: {
			Name:    "ServerStateCorrupted",
			Message: "Server state corrupted!",
			Variant: ServerStateCorrupted,
		},
		LastKnowStateIssue: {
			Name:    "LastKnowStateIssue",
			Message: "Last know state corrupted or has issues!",
			Variant: LastKnowStateIssue,
		},
		ServerIdCorrupted: {
			Name:    "ServerIdCorrupted",
			Message: "Server id corrupted!",
			Variant: ServerIdCorrupted,
		},
		IdCorrupted: {
			Name:    "IdCorrupted",
			Message: "Id or identifier corrupted!",
			Variant: IdCorrupted,
		},
		ServerUndefined: {
			Name:    "ServerUndefined",
			Message: "Server undefined!",
			Variant: ServerUndefined,
		},
		SiteUndefined: {
			Name:    "SiteUndefined",
			Message: "Site undefined!",
			Variant: SiteUndefined,
		},
		SiteMissing: {
			Name:    "SiteMissing",
			Message: "Site missing!",
			Variant: SiteUndefined,
		},
		SiteIdMissing: {
			Name:    "SiteIdMissing",
			Message: "Site id missing!",
			Variant: SiteIdMissing,
		},
		SubdomainMissing: {
			Name:    "SubdomainMissing",
			Message: "Subdomain missing!",
			Variant: SubdomainMissing,
		},
		SslIssue: {
			Name:    "SslIssue",
			Message: "SSL/TLS related issue!",
			Variant: SslIssue,
		},
		SslMissing: {
			Name:    "SslMissing",
			Message: "SSL missing or ssl file missing!",
			Variant: SslMissing,
		},
		SslInvalid: {
			Name:    "SslInvalid",
			Message: "SSL Invalid!",
			Variant: SslInvalid,
		},
		SslExpired: {
			Name:    "SslExpired",
			Message: "SSL Expired!",
			Variant: SslExpired,
		},
		SslMismatch: {
			Name:    "SslMismatch",
			Message: "SSL mismatch!",
			Variant: SslMismatch,
		},
		SslDryRunFailed: {
			Name:    "SslDryRunFailed",
			Message: "SSL dry run failed, will fail to renew or update!",
			Variant: SslDryRunFailed,
		},
		SslRenewFailed: {
			Name:    "SslRenewFailed",
			Message: "SSL renew failed!",
			Variant: SslRenewFailed,
		},
		Challenge: {
			Name:    "Challenge",
			Message: "Challenge failed!",
			Variant: Challenge,
		},
		ChallengeIssue: {
			Name:    "ChallengeIssue",
			Message: "Challenge issue!",
			Variant: ChallengeIssue,
		},
		SslChallengeIssue: {
			Name:    "SslChallengeIssue",
			Message: "SSL Challenge issue, couldn't verify SSL challenge!",
			Variant: SslChallengeIssue,
		},
		DnsChallengeIssue: {
			Name:    "DnsChallengeIssue",
			Message: "DNS or DNS-01 Challenge issue, couldn't verify it!",
			Variant: DnsChallengeIssue,
		},
		HttpChallengeIssue: {
			Name:    "HttpChallengeIssue",
			Message: "Http or Http01 Challenge issue, couldn't verify it!",
			Variant: HttpChallengeIssue,
		},
		HttpsChallengeIssue: {
			Name:    "HttpsChallengeIssue",
			Message: "Https Challenge issue, couldn't verify it!",
			Variant: HttpsChallengeIssue,
		},
		TlsSniChallengeIssue: {
			Name:    "TlsSniChallengeIssue",
			Message: "TLS-SNI-01 Challenge issue, couldn't verify it!",
			Variant: TlsSniChallengeIssue,
		},
		TlsAlpnChallengeIssue: {
			Name:    "TlsAlpnChallengeIssue",
			Message: "TLS-ALPN-01 Challenge issue, couldn't verify it!",
			Variant: TlsAlpnChallengeIssue,
		},
		CAChallengeIssue: {
			Name:    "CAChallengeIssue",
			Message: "Certificate Authority challenge issue, couldn't verify it!",
			Variant: CAChallengeIssue,
		},
		DomainMismatch: {
			Name:    "DomainMismatch",
			Message: "Domain mismatch!",
			Variant: DomainMismatch,
		},
		SiteContentMismatch: {
			Name:    "SiteContentMismatch",
			Message: "Site content mismatch!",
			Variant: SiteContentMismatch,
		},
		ServerContentMismatch: {
			Name:    "ServerContentMismatch",
			Message: "Server content mismatch!",
			Variant: ServerContentMismatch,
		},
		ContentMismatch: {
			Name:    "ContentMismatch",
			Message: "Content mismatch!",
			Variant: ContentMismatch,
		},
		AlreadyExist: {
			Name:    "AlreadyExist",
			Message: "Already exist!",
			Variant: AlreadyExist,
		},
		RecordAlreadyExist: {
			Name:    "RecordAlreadyExist",
			Message: "Record already exist!",
			Variant: RecordAlreadyExist,
		},
		ContentAlreadyExist: {
			Name:    "ContentAlreadyExist",
			Message: "Content already exist!",
			Variant: ContentAlreadyExist,
		},
		AlreadyPersist: {
			Name:    "AlreadyPersist",
			Message: "Already persist!",
			Variant: AlreadyPersist,
		},
		InvalidUsername: {
			Name:    "InvalidUsername",
			Message: "Invalid username!",
			Variant: InvalidUsername,
		},
		UsernameMissing: {
			Name:    "UsernameMissing",
			Message: "Username missing!",
			Variant: UsernameMissing,
		},
		UsernameEmpty: {
			Name:    "UsernameEmpty",
			Message: "Username empty!",
			Variant: UsernameEmpty,
		},
		UserNotFound: {
			Name:    "UserNotFound",
			Message: "User not found!",
			Variant: UserNotFound,
		},
		UserMismatch: {
			Name:    "UserMismatch",
			Message: "User mismatch!",
			Variant: UserMismatch,
		},
		Banned: {
			Name:    "Banned",
			Message: "Banned!",
			Variant: Banned,
		},
		Disabled: {
			Name:    "Disabled",
			Message: "Disabled!",
			Variant: Disabled,
		},
		AccountDisabled: {
			Name:    "AccountDisabled",
			Message: "Account Disabled!",
			Variant: AccountDisabled,
		},
		AccountSuspended: {
			Name:    "AccountSuspended",
			Message: "Account Suspended!",
			Variant: AccountSuspended,
		},
		Suspended: {
			Name:    "Suspended",
			Message: "Suspended!",
			Variant: Suspended,
		},
		UserSuspended: {
			Name:    "UserSuspended",
			Message: "User Suspended!",
			Variant: UserSuspended,
		},
		RoleSuspended: {
			Name:    "RoleSuspended",
			Message: "Role Suspended!",
			Variant: RoleSuspended,
		},
		ContentDisabled: {
			Name:    "ContentDisabled",
			Message: "Content Disabled!",
			Variant: ContentDisabled,
		},
		ContentBlocked: {
			Name:    "ContentBlocked",
			Message: "Content Blocked!",
			Variant: ContentBlocked,
		},
		Blocked: {
			Name:    "Blocked",
			Message: "Blocked!",
			Variant: Blocked,
		},
		RecordBlocked: {
			Name:    "RecordBlocked",
			Message: "Record Blocked!",
			Variant: RecordBlocked,
		},
		RoleRemoved: {
			Name:    "RoleRemoved",
			Message: "Role Removed!",
			Variant: RoleRemoved,
		},
		RightsRevoked: {
			Name:    "RightsRevoked",
			Message: "Rights Revoked!",
			Variant: RightsRevoked,
		},
		RightsRemoved: {
			Name:    "RightsRemoved",
			Message: "Rights Removed!",
			Variant: RightsRemoved,
		},
		UserBanned: {
			Name:    "UserBanned",
			Message: "User Banned!",
			Variant: UserBanned,
		},
		UserRestricted: {
			Name:    "UserRestricted",
			Message: "User restricted!",
			Variant: UserRestricted,
		},
		UserRemoved: {
			Name:    "UserRemoved",
			Message: "User Removed!",
			Variant: UserRemoved,
		},
		RecordRemoved: {
			Name:    "RecordRemoved",
			Message: "Record Removed!",
			Variant: RecordRemoved,
		},
		AccessRights: {
			Name:    "AccessRights",
			Message: "Access Rights issue!",
			Variant: AccessRights,
		},
		AccessRightsDeny: {
			Name:    "AccessRightsDeny",
			Message: "Access Rights Deny!",
			Variant: AccessRightsDeny,
		},
		UserManagementIssue: {
			Name:    "UserManagementIssue",
			Message: "User Management Issue!",
			Variant: UserManagementIssue,
		},
		SuspiciousBehaviour: {
			Name:    "SuspiciousBehaviour",
			Message: "Suspicious Behaviour!",
			Variant: SuspiciousBehaviour,
		},
		Rights: {
			Name:    "Rights",
			Message: "Rights Issue!",
			Variant: Rights,
		},
		RoleIssue: {
			Name:    "RoleIssue",
			Message: "Role Issue!",
			Variant: RoleIssue,
		},
		RoleInheritanceDenyPermission: {
			Name:    "RoleInheritanceDenyPermission",
			Message: "Role Inheritance Deny Permission!",
			Variant: RoleInheritanceDenyPermission,
		},
		User: {
			Name:    "User",
			Message: "User!",
			Variant: User,
		},
		UserRoleInvalid: {
			Name:    "UserRoleInvalid",
			Message: "User Role Invalid!",
			Variant: UserRoleInvalid,
		},
		RoleInvalid: {
			Name:    "RoleInvalid",
			Message: "Role Invalid!",
			Variant: RoleInvalid,
		},
		UserHaveNoRights: {
			Name:    "UserHaveNoRights",
			Message: "User have no rights to perform the action!",
			Variant: UserHaveNoRights,
		},
		RoleMissing: {
			Name:    "RoleMissing",
			Message: "Role missing!",
			Variant: RoleMissing,
		},
		UserRoleDenyPermission: {
			Name:    "UserRoleDenyPermission",
			Message: "User Role Deny Permission!",
			Variant: UserRoleDenyPermission,
		},
		PersistentRecordMismatch: {
			Name:    "PersistentRecordMismatch",
			Message: "Persistent Record mismatch!",
			Variant: PersistentRecordMismatch,
		},
		UnSync: {
			Name:    "UnSync",
			Message: "UnSync!",
			Variant: UnSync,
		},
		BlockIssue: {
			Name:    "BlockIssue",
			Message: "Block Issue!",
			Variant: BlockIssue,
		},
		BlockSyntaxIssue: {
			Name:    "BlockSyntaxIssue",
			Message: "Block Syntax Issue!",
			Variant: BlockSyntaxIssue,
		},
		BlockNotFound: {
			Name:    "BlockNotFound",
			Message: "Block not found!",
			Variant: BlockNotFound,
		},
		SubdomainNotFound: {
			Name:    "SubdomainNotFound",
			Message: "Subdomain not found!",
			Variant: SubdomainNotFound,
		},
		SubdomainMismatch: {
			Name:    "SubdomainMismatch",
			Message: "Subdomain mismatch!",
			Variant: SubdomainMismatch,
		},
		DomainOrSubdomainAlreadyAssigned: {
			Name:    "DomainOrSubdomainAlreadyAssigned",
			Message: "Domain or subdomain already assigned to someone!",
			Variant: DomainOrSubdomainAlreadyAssigned,
		},
		ResourceAlreadyAssigned: {
			Name:    "ResourceAlreadyAssigned",
			Message: "Resource already assigned to someone!",
			Variant: ResourceAlreadyAssigned,
		},
		RecordAlreadyAssigned: {
			Name:    "RecordAlreadyAssigned",
			Message: "Record already assigned to someone!",
			Variant: RecordAlreadyAssigned,
		},
		PathAlreadyAssigned: {
			Name:    "PathAlreadyAssigned",
			Message: "Path already assigned to someone!",
			Variant: PathAlreadyAssigned,
		},
		FileAlreadyAssigned: {
			Name:    "FileAlreadyAssigned",
			Message: "File already assigned to someone!",
			Variant: FileAlreadyAssigned,
		},
		DirAlreadyAssigned: {
			Name:    "FileAlreadyAssigned",
			Message: "File already assigned to someone!",
			Variant: FileAlreadyAssigned,
		},
		ServerAlreadyAssigned: {
			Name:    "ServerAlreadyAssigned",
			Message: "Server already assigned to someone!",
			Variant: ServerAlreadyAssigned,
		},
		SiteAlreadyAssigned: {
			Name:    "SiteAlreadyAssigned",
			Message: "Site already assigned to someone!",
			Variant: SiteAlreadyAssigned,
		},
		OwnerDifferent: {
			Name:    "OwnerDifferent",
			Message: "Owner different cannot access!",
			Variant: OwnerDifferent,
		},
		Username: {
			Name:    "Username",
			Message: "Username!",
			Variant: Username,
		},
		UsernameAlreadyExist: {
			Name:    "UsernameAlreadyExist",
			Message: "Username already exist!",
			Variant: UsernameAlreadyExist,
		},
		NameMismatch: {
			Name:    "NameMismatch",
			Message: "Category mismatch!",
			Variant: NameMismatch,
		},
		IdMismatch: {
			Name:    "IdMismatch",
			Message: "Id mismatch!",
			Variant: IdMismatch,
		},
		RequestInvalid: {
			Name:    "RequestInvalid",
			Message: "Request invalid!",
			Variant: RequestInvalid,
		},
		ResponseInvalid: {
			Name:    "ResponseInvalid",
			Message: "Response invalid!",
			Variant: ResponseInvalid,
		},
		SpecUndefined: {
			Name:    "SpecUndefined",
			Message: "Specification invalid or undefined!",
			Variant: SpecUndefined,
		},
		ValueNotFound: {
			Name:    "ValueNotFound",
			Message: "Value not found!",
			Variant: ValueNotFound,
		},
		QueryInvalid: {
			Name:    "QueryInvalid",
			Message: "Query invalid or query parameter(s) doesn't meet the expectation!",
			Variant: QueryInvalid,
		},
		QueryEmpty: {
			Name:    "QueryEmpty",
			Message: "Query empty!",
			Variant: QueryEmpty,
		},
		FieldMissing: {
			Name:    "FieldMissing",
			Message: "Field missing!",
			Variant: FieldMissing,
		},
		FieldsMissing: {
			Name:    "FieldsMissing",
			Message: "Field (s) missing!",
			Variant: FieldsMissing,
		},
		QueryParameterInvalid: {
			Name:    "QueryParameterInvalid",
			Message: "Query parameter invalid or missing or doesn't have expected data!",
			Variant: QueryParameterInvalid,
		},
		QueryParameterEmpty: {
			Name:    "QueryParameterEmpty",
			Message: "Query parameter(s) empty!",
			Variant: QueryParameterEmpty,
		},
		RequestParameterEmpty: {
			Name:    "RequestParameterEmpty",
			Message: "Request parameter(s) empty!",
			Variant: RequestParameterEmpty,
		},
		Response: {
			Name:    "Response",
			Message: "Response issue!",
			Variant: Response,
		},
		ResponseExpectedParameterMissing: {
			Name:    "ResponseExpectedParameterMissing",
			Message: "Response expected parameter(s) missing!",
			Variant: ResponseExpectedParameterMissing,
		},
		ArgumentInvalid: {
			Name:    "ArgumentInvalid",
			Message: "Argument invalid!",
			Variant: ArgumentInvalid,
		},
		ArgumentEmpty: {
			Name:    "ArgumentEmpty",
			Message: "Argument empty!",
			Variant: ArgumentEmpty,
		},
		ContentEmpty: {
			Name:    "ContentEmpty",
			Message: "Content empty!",
			Variant: ContentEmpty,
		},
		ContentDefined: {
			Name:    "ContentDefined",
			Message: "Content defined but it should be defined!",
			Variant: ContentDefined,
		},
		TraverseIssue: {
			Name:    "TraverseIssue",
			Message: "Traverse issue!",
			Variant: TraverseIssue,
		},
		RecursionIssue: {
			Name:    "RecursionIssue",
			Message: "Recursion issue!",
			Variant: RecursionIssue,
		},
		RecursionMaxDept: {
			Name:    "RecursionMaxDept",
			Message: "Recursion max dept reached!",
			Variant: RecursionMaxDept,
		},
		RecursionInvalidState: {
			Name:    "RecursionInvalidState",
			Message: "Recursion invalid state occurred!",
			Variant: RecursionInvalidState,
		},
		NodeInvalid: {
			Name:    "NodeInvalid",
			Message: "Node invalid!",
			Variant: NodeInvalid,
		},
		Level: {
			Name:    "Level",
			Message: "Level!",
			Variant: Level,
		},
		MaxDepth: {
			Name:    "MaxDepth",
			Message: "Max dept issue!",
			Variant: MaxDepth,
		},
		MaxLimit: {
			Name:    "MaxLimit",
			Message: "Max limit reached!",
			Variant: MaxLimit,
		},
		DirWalk: {
			Name:    "DirWalk",
			Message: "Dir walk or traverse issue!",
			Variant: DirWalk,
		},
		PathRecursiveTravel: {
			Name:    "PathRecursiveTravel",
			Message: "Path recursive walk or traverse issue!",
			Variant: PathRecursiveTravel,
		},
		DirNotFound: {
			Name:    "DirNotFound",
			Message: "Dir not found or permission issue!",
			Variant: DirNotFound,
		},
		FileNotFound: {
			Name:    "FileNotFound",
			Message: "File not found or permission issue!",
			Variant: FileNotFound,
		},
		ChownMismatch: {
			Name:    "ChownMismatch",
			Message: "Chown (owner: user or group) mismatch issue!",
			Variant: ChownMismatch,
		},
		ManyNullOrEmpty: {
			Name:    "ManyNullOrEmpty",
			Message: "Many items null or empty!",
			Variant: ManyNullOrEmpty,
		},
		RelatedRecordMissing: {
			Name:    "RelatedRecordMissing",
			Message: "Related record(s) missing!",
			Variant: RelatedRecordMissing,
		},
		ForeignKeyIssue: {
			Name:    "ForeignKeyIssue",
			Message: "Foreign key(s) issue!",
			Variant: ForeignKeyIssue,
		},
		PrimaryKeyIssue: {
			Name:    "PrimaryKeyIssue",
			Message: "Primary key(s) issue!",
			Variant: PrimaryKeyIssue,
		},
		SubRecordNotFound: {
			Name:    "SubRecordNotFound",
			Message: "Sub record(s) not found!",
			Variant: SubRecordNotFound,
		},
		SomeStatesMissing: {
			Name:    "SomeStatesMissing",
			Message: "Some state(s) missing or not found!",
			Variant: SomeStatesMissing,
		},
		SomePathsInvalidOrPermissionIssue: {
			Name:    "SomePathsInvalidOrPermissionIssue",
			Message: "Some path(s) invalid or missing or permission found!",
			Variant: SomePathsInvalidOrPermissionIssue,
		},
		RelatedRecordExpected: {
			Name:    "RelatedRecordExpected",
			Message: "Related record(s) expected but not found or missing!",
			Variant: RelatedRecordExpected,
		},
		ExpectedKeyMissing: {
			Name:    "ExpectedKeyMissing",
			Message: "Expected key(s) not found or missing!",
			Variant: ExpectedKeyMissing,
		},
		ExpectedNameMissing: {
			Name:    "ExpectedNameMissing",
			Message: "Expected name(s) not found or missing!",
			Variant: ExpectedNameMissing,
		},
		ExpectedIdMissing: {
			Name:    "ExpectedIdMissing",
			Message: "Expected identifier(s) not found or missing!",
			Variant: ExpectedIdMissing,
		},
		ExpectedUserMissing: {
			Name:    "ExpectedUserMissing",
			Message: "Expected user(s) not found or missing!",
			Variant: ExpectedUserMissing,
		},
		ExpectedValueMissing: {
			Name:    "ExpectedValueMissing",
			Message: "Expected value(s) not found or missing!",
			Variant: ExpectedValueMissing,
		},
		ExpectedContentMissing: {
			Name:    "ExpectedContentMissing",
			Message: "Expected content(s) not found or missing!",
			Variant: ExpectedContentMissing,
		},
		ExpectedStatusMissing: {
			Name:    "ExpectedStatusMissing",
			Message: "Expected status not found or missing!",
			Variant: ExpectedStatusMissing,
		},
		FileStateMismatch: {
			Name:    "FileStateMismatch",
			Message: "File state(s) mismatch!",
			Variant: FileStateMismatch,
		},
		PathStateMismatch: {
			Name:    "PathStateMismatch",
			Message: "Path state(s) mismatch!",
			Variant: PathStateMismatch,
		},
		DirStateMismatch: {
			Name:    "DirStateMismatch",
			Message: "Directory state(s) mismatch!",
			Variant: DirStateMismatch,
		},
		PathStateNotFound: {
			Name:    "PathStateNotFound",
			Message: "Path state(s) mismatch!",
			Variant: PathStateNotFound,
		},
		ExpectingDirButFile: {
			Name:    "ExpectingDirButFile",
			Message: "Expecting directory(ies) but received file(s)!",
			Variant: ExpectingDirButFile,
		},
		ExpectingFileButDir: {
			Name:    "ExpectingFileButDir",
			Message: "Expecting file(s) but received directory(ies)!",
			Variant: ExpectingFileButDir,
		},
		DirInvalidOrMissing: {
			Name:    "DirInvalidOrMissing",
			Message: "Expecting directory(ies) but invalid or missing or permission issue!",
			Variant: DirInvalidOrMissing,
		},
		UnsatisfyingCondition: {
			Name:    "UnsatisfyingCondition",
			Message: "Unsatisfying condition!",
			Variant: UnsatisfyingCondition,
		},
		AsyncIssue: {
			Name:    "AsyncIssue",
			Message: "Async issue!",
			Variant: AsyncIssue,
		},
		FetchingFailed: {
			Name:    "FetchingFailed",
			Message: "Fetching failed!",
			Variant: FetchingFailed,
		},
		VersionConflict: {
			Name:    "VersionConflict",
			Message: "Version conflict!",
			Variant: VersionConflict,
		},
		Conflict: {
			Name:    "Conflict",
			Message: "Conflict!",
			Variant: Conflict,
		},
		StatusConflict: {
			Name:    "StatusConflict",
			Message: "Status conflict!",
			Variant: StatusConflict,
		},
		IdConflict: {
			Name:    "IdConflict",
			Message: "Id conflict!",
			Variant: IdConflict,
		},
		NameConflict: {
			Name:    "NameConflict",
			Message: "Category conflict!",
			Variant: NameConflict,
		},
		KeyConflict: {
			Name:    "KeyConflict",
			Message: "Key conflict!",
			Variant: KeyConflict,
		},
		PrimaryKeyConflict: {
			Name:    "PrimaryKeyConflict",
			Message: "Primary Key(s) conflict!",
			Variant: PrimaryKeyConflict,
		},
		ForeignKeyConflict: {
			Name:    "ForeignKeyConflict",
			Message: "Foreign Key(s) conflict!",
			Variant: ForeignKeyConflict,
		},
		Merge: {
			Name:    "Merge",
			Message: "Merge!",
			Variant: Merge,
		},
		MergeConflict: {
			Name:    "MergeConflict",
			Message: "Merge conflict!",
			Variant: MergeConflict,
		},
		ProcessInvalid: {
			Name:    "ProcessInvalid",
			Message: "Process invalid!",
			Variant: ProcessInvalid,
		},
		EmptyId: {
			Name:    "EmptyId",
			Message: "Empty id or identifier!",
			Variant: EmptyId,
		},
		EmptyUser: {
			Name:    "EmptyUser",
			Message: "Empty user!",
			Variant: EmptyUser,
		},
		EmptyPath: {
			Name:    "EmptyPath",
			Message: "Empty (\"\") path!",
			Variant: EmptyPath,
		},
		EmptyDir: {
			Name:    "EmptyDir",
			Message: "Empty (\"\") directory!",
			Variant: EmptyDir,
		},
		EmptyResponse: {
			Name:    "EmptyResponse",
			Message: "Empty response!",
			Variant: EmptyResponse,
		},
		EmptyField: {
			Name:    "EmptyField",
			Message: "Empty field!",
			Variant: EmptyField,
		},
		EmptyRecords: {
			Name:    "EmptyRecords",
			Message: "Empty records!",
			Variant: EmptyRecords,
		},
		EmptyStates: {
			Name:    "EmptyStates",
			Message: "Empty states!",
			Variant: EmptyStates,
		},
		EmptyQuery: {
			Name:    "EmptyQuery",
			Message: "Empty query!",
			Variant: EmptyQuery,
		},
		EmptyRequest: {
			Name:    "EmptyRequest",
			Message: "Empty request!",
			Variant: EmptyRequest,
		},
		NoDiskSpace: {
			Name:    "NoDiskSpace",
			Message: "No disk space!",
			Variant: NoDiskSpace,
		},
		DiskSpace: {
			Name:    "DiskSpace",
			Message: "Disk space issue!",
			Variant: DiskSpace,
		},
		QuotaExpired: {
			Name:    "QuotaExpired",
			Message: "Quota expired!",
			Variant: QuotaExpired,
		},
		DiskSpaceQuotaExpired: {
			Name:    "DiskSpaceQuotaExpired",
			Message: "Quota expired!",
			Variant: DiskSpaceQuotaExpired,
		},
		LimitWarning: {
			Name:    "LimitWarning",
			Message: "Limit warning!",
			Variant: LimitWarning,
		},
		MailBoxExpired: {
			Name:    "MailBoxExpired",
			Message: "MailBox expired!",
			Variant: MailBoxExpired,
		},
		MailQuotaExpired: {
			Name:    "MailQuotaExpired",
			Message: "Mail quota expired!",
			Variant: MailQuotaExpired,
		},
		UserQuotaExpired: {
			Name:    "UserQuotaExpired",
			Message: "User quota expired!",
			Variant: UserQuotaExpired,
		},
		Expired: {
			Name:    "Expired",
			Message: "Expired!",
			Variant: Expired,
		},
		LimitExpired: {
			Name:    "LimitExpired",
			Message: "Limit expired!",
			Variant: LimitExpired,
		},
		ContentExpired: {
			Name:    "ContentExpired",
			Message: "Content expired!",
			Variant: ContentExpired,
		},
		TokenExpired: {
			Name:    "TokenExpired",
			Message: "Token expired!",
			Variant: TokenExpired,
		},
		VersionExpired: {
			Name:    "VersionExpired",
			Message: "Version expired!",
			Variant: VersionExpired,
		},
		CacheExpired: {
			Name:    "CacheExpired",
			Message: "Cache expired!",
			Variant: CacheExpired,
		},
		RefreshTokenExpired: {
			Name:    "RefreshTokenExpired",
			Message: "Refresh token expired!",
			Variant: RefreshTokenExpired,
		},
		AuthTokenExpired: {
			Name:    "AuthTokenExpired",
			Message: "Auth token expired!",
			Variant: AuthTokenExpired,
		},
		RouteExpired: {
			Name:    "RouteExpired",
			Message: "Route expired!",
			Variant: RouteExpired,
		},
		FileStateExpired: {
			Name:    "FileStateExpired",
			Message: "File(s) state(s) expired!",
			Variant: FileStateExpired,
		},
		UserExpired: {
			Name:    "UserExpired",
			Message: "User(s) expired!",
			Variant: UserExpired,
		},
		AccountExpired: {
			Name:    "AccountExpired",
			Message: "Account(s) expired!",
			Variant: AccountExpired,
		},
		Revoked: {
			Name:    "Revoked",
			Message: "Revoked!",
			Variant: Revoked,
		},
		SecurityIssue: {
			Name:    "SecurityIssue",
			Message: "Security issue!",
			Variant: SecurityIssue,
		},
		TaskFailed: {
			Name:    "TaskFailed",
			Message: "Task failed!",
			Variant: TaskFailed,
		},
		AsyncTaskFailed: {
			Name:    "AsyncTaskFailed",
			Message: "Async task failed!",
			Variant: AsyncTaskFailed,
		},
		ParallelTaskIssue: {
			Name:    "ParallelTaskIssue",
			Message: "Parallel task failed!",
			Variant: ParallelTaskIssue,
		},
		ApplicationMigrationFailed: {
			Name:    "ApplicationMigrationFailed",
			Message: "Application migration failed!",
			Variant: ApplicationMigrationFailed,
		},
		VersionMigrationFailed: {
			Name:    "VersionMigrationFailed",
			Message: "Version migration failed!",
			Variant: VersionMigrationFailed,
		},
		RecordMigrationFailed: {
			Name:    "RecordMigrationFailed",
			Message: "Record(s) migration failed!",
			Variant: RecordMigrationFailed,
		},
		SeedingMigrationFailed: {
			Name:    "SeedingMigrationFailed",
			Message: "Seeding migration failed!",
			Variant: SeedingMigrationFailed,
		},
		AddOrUpdate: {
			Name:    "AddOrUpdate",
			Message: "Add or update failed!",
			Variant: AddOrUpdate,
		},
		CreateOrUpdate: {
			Name:    "CreateOrUpdate",
			Message: "Create or update failed!",
			Variant: CreateOrUpdate,
		},
		DeleteOrSkip: {
			Name:    "DeleteOrSkip",
			Message: "Delete or skip failed!",
			Variant: DeleteOrSkip,
		},
		Group: {
			Name:    "Group",
			Message: "Group issue!",
			Variant: Group,
		},
		GroupCrud: {
			Name:    "GroupCrud",
			Message: "Group CRUD failed!",
			Variant: GroupCrud,
		},
		GroupAddOrUpdate: {
			Name:    "GroupAddOrUpdate",
			Message: "Group add or update failed!",
			Variant: GroupAddOrUpdate,
		},
		Status: {
			Name:    "Status",
			Message: "Status failed!",
			Variant: Status,
		},
		StatusCrud: {
			Name:    "StatusCrud",
			Message: "Status CRUD failed!",
			Variant: StatusCrud,
		},
		PasswordCrud: {
			Name:    "PasswordCrud",
			Message: "Password CRUD failed!",
			Variant: PasswordCrud,
		},
		PasswordIssue: {
			Name:    "PasswordIssue",
			Message: "Password issue!",
			Variant: PasswordIssue,
		},
		FileIssue: {
			Name:    "FileIssue",
			Message: "File issue!",
			Variant: FileIssue,
		},
		StackOverflow: {
			Name:    "StackOverflow",
			Message: "Stack overflow!",
			Variant: StackOverflow,
		},
		TaskIdNotFound: {
			Name:    "TaskIdNotFound",
			Message: "Task id not found!",
			Variant: TaskIdNotFound,
		},
		TaskCategoryNotFound: {
			Name:    "TaskCategoryNotFound",
			Message: "Task category not found!",
			Variant: TaskCategoryNotFound,
		},
		CategoryMissing: {
			Name:    "CategoryMissing",
			Message: "Category not found or missing!",
			Variant: CategoryMissing,
		},
		TaskMissing: {
			Name:    "TaskMissing",
			Message: "Task not found or missing!",
			Variant: TaskMissing,
		},
		TaskNotFound: {
			Name:    "TaskNotFound",
			Message: "Task not found or missing!",
			Variant: TaskNotFound,
		},
		ProcessNotFound: {
			Name:    "ProcessNotFound",
			Message: "Process not found or missing!",
			Variant: ProcessNotFound,
		},
		HandlerMissing: {
			Name:    "HandlerMissing",
			Message: "Handler not found or missing!",
			Variant: HandlerMissing,
		},
		StateMismatch: {
			Name:    "StateMismatch",
			Message: "State mismatch issue, probably restore or reset or revert may fix!",
			Variant: StateMismatch,
		},
		StateCorrupted: {
			Name:    "StateCorrupted",
			Message: "State corrupted issue, probably restore or reset or revert may fix!",
			Variant: StateCorrupted,
		},
		Search: {
			Name:    "Search",
			Message: "Search related issue!",
			Variant: Search,
		},
		History: {
			Name:    "History",
			Message: "History related issue!",
			Variant: History,
		},
		HistoryNotFound: {
			Name:    "HistoryNotFound",
			Message: "History not found or missing!",
			Variant: HistoryNotFound,
		},
		PackageNotFound: {
			Name:    "PackageNotFound",
			Message: "Package not found or missing!",
			Variant: PackageNotFound,
		},
		InternetConnection: {
			Name:    "InternetConnection",
			Message: "Internet connection related issue!",
			Variant: InternetConnection,
		},

		PayloadCastingFailed: {
			Name:    "PayloadCastingFailed",
			Message: "Payload casting failed!",
			Variant: PayloadCastingFailed,
		},
		PayloadIntegrity: {
			Name:    "PayloadIntegrity",
			Message: "Payload integrity lost or mismatch or corrupted!",
			Variant: PayloadIntegrity,
		},
		PayloadCategoryIssue: {
			Name:    "PayloadCategoryIssue",
			Message: "Payload category empty or mismatch or corrupted or has issues!",
			Variant: PayloadCategoryIssue,
		},
		PayloadEntityTypeIssue: {
			Name:    "PayloadEntityTypeIssue",
			Message: "Payload entity type empty or mismatch or corrupted or has issues!",
			Variant: PayloadEntityTypeIssue,
		},
		PayloadSubscriptionFailed: {
			Name:    "PayloadSubscriptionFailed",
			Message: "Payload subscription failed!",
			Variant: PayloadSubscriptionFailed,
		},
		PayloadIssue: {
			Name:    "PayloadIssue",
			Message: "Payload related generic or unknown issue!",
			Variant: PayloadIssue,
		},
		PayloadNotFound: {
			Name:    "PayloadNotFound",
			Message: "Payload not found or as not as expected!",
			Variant: PayloadNotFound,
		},
		SubscriptionRelated: {
			Name:    "SubscriptionRelated",
			Message: "Subscription related issue!",
			Variant: SubscriptionRelated,
		},
		SubscriptionInvalid: {
			Name:    "SubscriptionInvalid",
			Message: "Subscription invalid!",
			Variant: SubscriptionInvalid,
		},
		NewCreationFailed: {
			Name:    "NewCreationFailed",
			Message: "New creation failed!",
			Variant: NewCreationFailed,
		},
		NewOrCasting: {
			Name:    "NewOrCasting",
			Message: "New creation or casting failed!",
			Variant: NewOrCasting,
		},
		BytesNotFound: {
			Name:    "BytesNotFound",
			Message: "Bytes not found!",
			Variant: BytesNotFound,
		},
		CastingIssue: {
			Name:    "CastingIssue",
			Message: "Cannot cast properly or casting failed!",
			Variant: CastingIssue,
		},
		Pipeline: {
			Name:    "Pipeline",
			Message: "Pipeline related issue!",
			Variant: Pipeline,
		},
		PipelineConstructionFailed: {
			Name:    "PipelineConstructionFailed",
			Message: "Pipeline construction failed!",
			Variant: PipelineConstructionFailed,
		},
		PipelineCastingFailed: {
			Name:    "PipelineCastingFailed",
			Message: "Pipeline casting failed!",
			Variant: PipelineCastingFailed,
		},
		LoggerFailed: {
			Name:    "LoggerFailed",
			Message: "Logger failed!",
			Variant: LoggerFailed,
		},
		LoggerCastingFailed: {
			Name:    "LoggerCastingFailed",
			Message: "Logger casting failed!",
			Variant: LoggerCastingFailed,
		},
		LoggerMismatch: {
			Name:    "LoggerMismatch",
			Message: "Logger mismatch, not as expected!",
			Variant: LoggerMismatch,
		},
		LoggerFailedToSave: {
			Name:    "LoggerFailedToSave",
			Message: "Logger failed to save to file system or persistent layer failed transaction!",
			Variant: LoggerFailedToSave,
		},
		EntityFailed: {
			Name:    "EntityFailed",
			Message: "Entity failed!",
			Variant: EntityFailed,
		},
		EntityNotFound: {
			Name:    "EntityNotFound",
			Message: "Entity not found or missing or invalid!",
			Variant: EntityNotFound,
		},
		EntityCastingFailed: {
			Name:    "EntityCastingFailed",
			Message: "Entity casting failed!",
			Variant: EntityCastingFailed,
		},
		WithinRange: {
			Name:    "WithinRange",
			Message: "Given value is within the range, expected to be out of range!",
			Variant: MaxError,
		},
		OutOfRange: {
			Name:    "OutOfRange",
			Message: "Given value is out of range range, expected to be within the range!",
			Variant: MaxError,
		},
		RescheduleFailed: {
			Name:    "RescheduleFailed",
			Message: "Reschedule failed!",
			Variant: RescheduleFailed,
		},
		ScheduleFailed: {
			Name:    "ScheduleFailed",
			Message: "Schedule failed!",
			Variant: ScheduleFailed,
		},
		ScheduleSaveFailed: {
			Name:    "ScheduleSaveFailed",
			Message: "Schedule failed to save in persistent layer or file system.!",
			Variant: ScheduleSaveFailed,
		},
		ScheduleRemoveFailed: {
			Name:    "ScheduleRemoveFailed",
			Message: "Schedule failed to remove!",
			Variant: ScheduleRemoveFailed,
		},
		ScheduleEditFailed: {
			Name:    "ScheduleEditFailed",
			Message: "Schedule failed to edit!",
			Variant: ScheduleEditFailed,
		},
		ScheduleRetryFailed: {
			Name:    "ScheduleRetryFailed",
			Message: "Schedule retry failed!",
			Variant: ScheduleRetryFailed,
		},
		ConvertFailed: {
			Name:    "ConvertFailed",
			Message: "Convert or casting or processing failed!",
			Variant: ConvertFailed,
		},
		InjectionFailed: {
			Name:    "InjectionFailed",
			Message: "Injection failed!",
			Variant: InjectionFailed,
		},
		DependencyInjectionFailed: {
			Name:    "DependencyInjectionFailed",
			Message: "Dependency injection failed!",
			Variant: DependencyInjectionFailed,
		},
		DependencyInjectionMissing: {
			Name:    "DependencyInjectionMissing",
			Message: "Dependency injection missing failed!",
			Variant: DependencyInjectionMissing,
		},
		MissingComponent: {
			Name:    "MissingComponent",
			Message: "Missing component!",
			Variant: MissingComponent,
		},
		CronFailed: {
			Name:    "CronFailed",
			Message: "Cron failed!",
			Variant: CronFailed,
		},
		CronEditFailed: {
			Name:    "CronEditFailed",
			Message: "Cron edit failed!",
			Variant: CronEditFailed,
		},
		CronSaveFailed: {
			Name:    "CronSaveFailed",
			Message: "Cron save or update or persistent layer failed!",
			Variant: CronSaveFailed,
		},
		CronRemoveFailed: {
			Name:    "CronRemoveFailed",
			Message: "Cron remove failed!",
			Variant: CronRemoveFailed,
		},
		Shutdown: {
			Name:    "Shutdown",
			Message: "Shutdown failed!",
			Variant: Shutdown,
		},
		Reboot: {
			Name:    "Reboot",
			Message: "Reboot failed!",
			Variant: Reboot,
		},
		JsonParseIssue: {
			Name:    "JsonParseIssue",
			Message: "Json parse issue!",
			Variant: JsonParseIssue,
		},
		JsonError: {
			Name:    "JsonError",
			Message: "Json error!",
			Variant: JsonError,
		},
		Serialize: {
			Name:    "Serialize",
			Message: "Serialize or marshalling failed!",
			Variant: Serialize,
		},
		Deserialize: {
			Name:    "Deserialize",
			Message: "Deserialize or unmarshalling failed!",
			Variant: Deserialize,
		},
		ReflectSetTo: {
			Name:    "ReflectSetTo",
			Message: "Reflect set failed! Check attachments for more details.",
			Variant: ReflectSetTo,
		},
		MaxError: {
			Name:    "MaxError",
			Message: "Maxed out error!",
			Variant: MaxError,
		},
	}
)
