package errorwrapper

import "gitlab.com/evatix-go/core/constants"

const (
	defaultSkipInternal = constants.One
	prefixStackTrace    = constants.Hyphen + constants.Space
	MessagesJoiner      = constants.Space
)
