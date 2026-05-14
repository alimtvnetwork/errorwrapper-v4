package linuxservicecmdtestwrappers

import "gitlab.com/evatix-go/errorwrapper/linuxservicecmd"

type ServicesTestCaseWrapper struct {
	Header string
	linuxservicecmd.ServicesInstruction
	ErrorValidation []string
}
