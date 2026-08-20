package linuxservicecmdtestwrappers

import (
	"github.com/alimtvnetwork/enum-v10/linuxservicestate"
	"github.com/alimtvnetwork/errorwrapper-v4/linuxservicecmd"
)

var ServicesErrorValidationTestCases = []ServicesTestCaseWrapper{
	// {
	// 	Header: "[Detailed error] Restarting services [ufwx, cron, iptables, ufwx2] services and should fail on ufwx, ufwx2 and continue on error!",
	// 	ErrorValidation: []string{
	// 		"[Error (CommandExecution - #103): Command (cmd) execution (exe) failed, thus cannot process the request. " +
	// 			"Additional : \n\"CmdOnce - Error\"\n, CommandLine:\n/bin/bash -c systemctl restart ufwx, " +
	// 			"Failed to restart ufwx.service: Unit ufwx.service not found.\n, " +
	// 			"exit status 5\nService execution failed. " +
	// 			"Ref(s) {" +
	// 			"[Service (string): \"ufwx\"], " +
	// 			"[Action (string): \"restart\"]}]",
	// 		"[Error (CommandExecution - #103): Command (cmd) execution (exe) failed, " +
	// 			"thus cannot process the request. Additional : \n\"" +
	// 			"CmdOnce - Error\"\n, " +
	// 			"CommandLine:\n/bin/bash -c systemctl restart ufwx2, " +
	// 			"Failed to restart ufwx2.service: Unit ufwx2.service not found.\n, " +
	// 			"exit status 5\n" +
	// 			"Service execution failed. Ref(s) {" +
	// 			"[Service (string): \"ufwx2\"], " +
	// 			"[Action (string): \"restart\"]}]",
	// 	},
	// 	ServicesInstruction: linuxservicecmd.ServicesInstruction{
	// 		IsDetailedError:   true,
	// 		IsContinueOnError: true,
	// 		RestartServicesNames: []string{
	// 			ServiceNonExistUfwx,
	// 			ServiceExistCron,
	// 			ServiceExistIptables,
	// 			ServiceNonExistUfwx2,
	// 		},
	// 		Validations: &linuxservicecmd.StateValidateInstructions{
	// 			Validations: []linuxservicecmd.StateValidateInstruction{
	// 				{
	// 					ServiceName:      ServiceExistCron,
	// 					ExpectedExitCode: osserviceexit.ActiveRunning,
	// 				},
	// 				{
	// 					ServiceName:      ServiceExistIptables,
	// 					ExpectedExitCode: osserviceexit.ActiveRunning,
	// 				},
	// 			},
	// 		},
	// 	},
	// },
	// {
	// 	Header: "[Detailed error] Restarting services [ufwx, cron, iptables, ufwx2] services and should fail on ufwx and NOT continue on error!",
	// 	ErrorValidation: []string{
	// 		"[Error (CommandExecution - #103): Command (cmd) execution (exe) failed, " +
	// 			"thus cannot process the request. Additional : \n\"" +
	// 			"CmdOnce - Error\"\n, " +
	// 			"CommandLine:\n/bin/bash -c systemctl restart ufwx, " +
	// 			"Failed to restart ufwx.service: Unit ufwx.service not found.\n, " +
	// 			"exit status 5\nService execution failed. " +
	// 			"Ref(s) {" +
	// 			"[Service (string): \"ufwx\"], " +
	// 			"[Action (string): \"restart\"]}]",
	// 	},
	// 	ServicesInstruction: linuxservicecmd.ServicesInstruction{
	// 		IsDetailedError:   true,
	// 		IsContinueOnError: false,
	// 		RestartServicesNames: []string{
	// 			ServiceNonExistUfwx,
	// 			ServiceExistCron,
	// 			ServiceExistIptables,
	// 			ServiceNonExistUfwx2,
	// 		},
	// 		Validations: &linuxservicecmd.StateValidateInstructions{
	// 			Validations: []linuxservicecmd.StateValidateInstruction{
	// 				{
	// 					ServiceName:      ServiceExistCron,
	// 					ExpectedExitCode: osserviceexit.ActiveRunning,
	// 				},
	// 				{
	// 					ServiceName:      ServiceExistIptables,
	// 					ExpectedExitCode: osserviceexit.ActiveRunning,
	// 				},
	// 			},
	// 		},
	// 	},
	// },
	{
		Header: "[Simple error] Restarting services [ufwx, cron, iptables, ufwx2] services and should fail on ufwx and ufwx2 continue on error!",
		ErrorValidation: []string{
			"[Error (UnknownService - #627): Service is unknown in the operating system! Additional : UnknownService. Ref(s) {[Service Name (string): \"ufwx\"], [Applying Action (string): \"restart\"]}]",
			"[Error (UnknownService - #627): Service is unknown in the operating system! Additional : UnknownService. Ref(s) {[Service Name (string): \"ufwx2\"], [Applying Action (string): \"restart\"]}]",
			"[Error (ServiceStatusCodeMismatch - #632): service status code or exit code mismatch! ufwx is not in the expected status! - Expect [\"ActiveRunning[1]\"] != [\"UnknownService[5]\"] Actual.]",
			"[Error (ServiceStatusCodeMismatch - #632): service status code or exit code mismatch! ufwx2 is not in the expected status! - Expect [\"ActiveRunning[1]\"] != [\"UnknownService[5]\"] Actual.]",
		},
		ServicesInstruction: linuxservicecmd.ServicesInstruction{
			IsDetailedError:   false,
			IsContinueOnError: true,
			RestartServicesNames: []string{
				ServiceNonExistUfwx,
				ServiceExistCron,
				ServiceExistIptables,
				ServiceNonExistUfwx2,
			},
			Validations: &linuxservicecmd.StateValidateInstructions{
				Validations: []linuxservicecmd.StateValidateInstruction{
					{
						ServiceName:      ServiceNonExistUfwx,
						ExpectedExitCode: linuxservicestate.ActiveRunning,
					},
					{
						ServiceName:      ServiceNonExistUfwx2,
						ExpectedExitCode: linuxservicestate.ActiveRunning,
					},
				},
			},
		},
	},

	{
		Header: "[Simple error] Restarting services [ufwx, cron, iptables, ufwx2] services and should fail on ufw and NOT continue on error!",
		ErrorValidation: []string{
			"[Error (UnknownService - #627): Service is unknown in the operating system! Additional : UnknownService. Ref(s) {[Service Name (string): \"ufwx\"], [Applying Action (string): \"restart\"]}]",
		},
		ServicesInstruction: linuxservicecmd.ServicesInstruction{
			IsDetailedError:   false,
			IsContinueOnError: false,
			RestartServicesNames: []string{
				ServiceNonExistUfwx,
				ServiceExistCron,
				ServiceExistIptables,
				ServiceNonExistUfwx2,
			},
			Validations: &linuxservicecmd.StateValidateInstructions{
				Validations: []linuxservicecmd.StateValidateInstruction{
					{
						ServiceName:      ServiceNonExistUfwx,
						ExpectedExitCode: linuxservicestate.ActiveRunning,
					},
					{
						ServiceName:      ServiceNonExistUfwx2,
						ExpectedExitCode: linuxservicestate.ActiveRunning,
					},
				},
			},
		},
	},
}
