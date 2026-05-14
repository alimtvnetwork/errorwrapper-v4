package linuxservicecmd

import "github.com/alimtvnetwork/enum-v10/servicestate"

type CoreServicesInstruction struct {
	IsIgnoreUnknownService bool                `json:"IsIgnoreUnknownService,omitempty"`
	Action                 servicestate.Action `json:"Action"`
	ServicesNames          []string            `json:"ServicesNames,omitempty"`
}
