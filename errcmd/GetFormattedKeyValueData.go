package errcmd

import "gitlab.com/evatix-go/core/constants"

// GetFormattedKeyValueData "MY_VAR=some_value"
func GetFormattedKeyValueData(
	varName string,
	varValue string,
) string {
	return varName + constants.EqualSymbol + varValue
}
