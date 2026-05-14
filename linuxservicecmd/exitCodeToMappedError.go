package linuxservicecmd

import (
	"github.com/alimtvnetwork/enum-v10/linuxservicestate"
	"github.com/alimtvnetwork/errorwrapper-v3"
	"github.com/alimtvnetwork/errorwrapper-v3/errnew"
	"github.com/alimtvnetwork/errorwrapper-v3/errtype"
)

func exitCodeToMappedError(result *Result) *errorwrapper.Wrapper {
	if result == nil {
		return errnew.Null.WithMessage(
			"exitCodeToMappedError received nil result",
			result)
	}

	if result.ExitCode.IsSuccess() {
		return nil
	}

	request := result.Request
	refValues := request.ReferenceValues()

	switch result.ExitCode {
	case linuxservicestate.DeadButPidExists:
		return errnew.Ref.ManyWithMsg(
			errtype.ServiceDeadButPidExist,
			"dead but /var/run pid file exists",
			refValues...)

	case linuxservicestate.DeadButVarLockFileExists:
		return errnew.Ref.ManyWithMsg(
			errtype.ServiceDead,
			"dead but /var/lock file exists",
			refValues...)
	case linuxservicestate.NotRunning:
		return errnew.Ref.ManyWithMsg(
			errtype.ServiceNotRunning,
			"ServiceNotRunning",
			refValues...)
	case linuxservicestate.UnknownService:
		return errnew.Ref.ManyWithMsg(
			errtype.UnknownService,
			"UnknownService",
			refValues...)
	case linuxservicestate.InvalidService:
		return errnew.Ref.ManyWithMsg(
			errtype.UnknownService,
			"UnknownService",
			refValues...)
	}

	return errnew.Ref.ManyWithMsg(
		errtype.InvalidIdentifier,
		"invalid exit code given! Exit code : "+result.ExitCode.Name(),
		refValues...)
}
