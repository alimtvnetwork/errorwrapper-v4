package errcmd

import "github.com/alimtvnetwork/core-v9/constants"

// GetFormattedKeyValueData "MY_VAR=some_value"
func GetFormattedKeyValueData(
	varName string,
	varValue string,
) string {
	return varName + constants.EqualSymbol + varValue
}
