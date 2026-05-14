package errconsts

//goland:noinspection GoUnusedConst
const (
	RangeWithRangeFormat                                = "Range must be in between %+v and %+v. Range {%#v}"
	RangeWithOutRangeFormat                             = "%s Range must be in between %+v and %+v."
	Ranges                                              = "Ranges"
	VariantStructStringFormat                           = "%s (Code - %d) : %s"
	VariantStructMsgReferenceValuesFormat               = "%s\n%s %s"
	VariantStructReferenceValuesFormat                  = "%s %s"
	ValueHyphenValueFormat                              = "%s - %s"
	SimpleReferenceCompileFormat                        = "%s - %s (%s)"
	SimpleReferenceCompileOptimizedFormat               = "%s (%s)"
	ValueWithDoubleQuoteFormat                          = "\"%v\""
	ValueWithSingleQuoteFormat                          = "'%v'"
	CombineMessageNoReferencesAdditionalMessageFormat   = "[Error (%s - #%d): %s %s.]"
	CombineMessageNoReferencesNoAdditionalMessageFormat = "[Error (%s - #%d): %s]"
	SingleReferenceCompile                              = "[%s (%T): \"%+v\"]"
	CombineMessageAdditionalMessageFormat               = "[Error (%s - #%d): %s Additional : %s. Ref(s) {%v}]"
	CombineMessageNoAdditionalMessageFormat             = "[Error (%s - #%d): %s Ref(s) {%v}]"
	DoubleStringTogetherFormat                          = "%s\n%s"
	DoubleStringTogetherWithReferenceFormat             = "%s\n%s \nReference {%+v(%T): %+v}"
	ErrorCodeWithTypeNameFormat                         = "(Code - #%d) : %s"
	ErrorCodeHyphenTypeNameFormat                       = "(#%d - %s)"
	ErrorCodeHyphenTypeNameWithLineFormat               = "(#%d - %s) %s"
	ErrorCodeHyphenTypeNameWithReferencesFormat         = "(#%d - %s - {%v})"
	TypeNameWithReferencesFormat                        = "(%s - {%v})" // type name, csvReference
	DirectoryPath                                       = "Directory Path"
	FilePath                                            = "File Path"
	ErrorStart                                          = "Error - "
	ReferenceStart                                      = " Reference ( "
	SpaceParenthesisEnd                                 = " )"
	ReferenceWithTypeFormat                             = "%s (%s) : %+v"
)
