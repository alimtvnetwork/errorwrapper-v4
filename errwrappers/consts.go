package errwrappers

import "gitlab.com/evatix-go/core/constants"

const (
	defaultSkipInternal    = constants.One
	DisplayJoiner          = "\n-----------------------------\n"
	finalMergedErrorPrefix = "\n-------------------------merged errors listed below---------------------------------\n"
)
