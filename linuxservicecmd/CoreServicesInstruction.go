package linuxservicecmd

import "gitlab.com/evatix-go/enum/servicestate"

type CoreServicesInstruction struct {
	IsIgnoreUnknownService bool                `json:"IsIgnoreUnknownService,omitempty"`
	Action                 servicestate.Action `json:"Action"`
	ServicesNames          []string            `json:"ServicesNames,omitempty"`
}
