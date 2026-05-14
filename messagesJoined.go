package errorwrapper

import (
	"strings"
)

func MessagesJoined(messages ...string) string {
	return strings.Join(
		messages, MessagesJoiner)
}
