package linuxservicecmdtestwrappers

import "github.com/alimtvnetwork/errorwrapper-v4/linuxservicecmd"

type ServicesTestCaseWrapper struct {
	Header string
	linuxservicecmd.ServicesInstruction
	ErrorValidation []string
}
