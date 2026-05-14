package linuxservicecmdtestwrappers

import (
	"github.com/alimtvnetwork/enum-v10/linuxservicestate"
	"github.com/alimtvnetwork/enum-v10/servicestate"
	"github.com/alimtvnetwork/errorwrapper-v3/linuxservicecmd"
)

var ServicesTestCases = []ServicesTestCaseWrapper{
	{
		Header: "Stop [cron, iptables] service and remains not running and ufwx service is unknown",
		ServicesInstruction: linuxservicecmd.ServicesInstruction{
			StopServicesNames: []string{ServiceExistCron, ServiceExistIptables},
			Services:          linuxservicecmd.ManyServicesInstructions{},
			Validations: &linuxservicecmd.StateValidateInstructions{
				Validations: []linuxservicecmd.StateValidateInstruction{
					{
						ServiceName:      ServiceExistCron,
						ExpectedExitCode: linuxservicestate.NotRunning,
					},
					{
						ServiceName:      ServiceExistIptables,
						ExpectedExitCode: linuxservicestate.NotRunning,
					},
					{
						ServiceName:      ServiceNonExistUfwx,
						ExpectedExitCode: linuxservicestate.UnknownService,
					},
				},
			},
		},
	},
	{
		Header: "Start [cron, iptables] started and active running!",
		ServicesInstruction: linuxservicecmd.ServicesInstruction{
			StartServicesNames: []string{ServiceExistCron, ServiceExistIptables},
			Services:           linuxservicecmd.ManyServicesInstructions{},
			Validations: &linuxservicecmd.StateValidateInstructions{
				Validations: []linuxservicecmd.StateValidateInstruction{
					{
						ServiceName:      ServiceExistCron,
						ExpectedExitCode: linuxservicestate.ActiveRunning,
					},
					{
						ServiceName:      ServiceExistIptables,
						ExpectedExitCode: linuxservicestate.ActiveRunning,
					},
				},
			},
		},
	},
	{
		Header: "Stop [cron, iptables] service and remains not running",
		ServicesInstruction: linuxservicecmd.ServicesInstruction{
			StopServicesNames: []string{ServiceExistCron, ServiceExistIptables},
			Services:          linuxservicecmd.ManyServicesInstructions{},
			Validations: &linuxservicecmd.StateValidateInstructions{
				Validations: []linuxservicecmd.StateValidateInstruction{
					{
						ServiceName:      ServiceExistCron,
						ExpectedExitCode: linuxservicestate.NotRunning,
					},
					{
						ServiceName:      ServiceExistIptables,
						ExpectedExitCode: linuxservicestate.NotRunning,
					},
				},
			},
		},
	},
	{
		Header: "Restarting [cron, iptables] stopped and started and active running!",
		ServicesInstruction: linuxservicecmd.ServicesInstruction{
			RestartServicesNames: []string{ServiceExistCron, ServiceExistIptables},
			Services:             linuxservicecmd.ManyServicesInstructions{},
			Validations: &linuxservicecmd.StateValidateInstructions{
				Validations: []linuxservicecmd.StateValidateInstruction{
					{
						ServiceName:      ServiceExistCron,
						ExpectedExitCode: linuxservicestate.ActiveRunning,
					},
					{
						ServiceName:      ServiceExistIptables,
						ExpectedExitCode: linuxservicestate.ActiveRunning,
					},
				},
			},
		},
	},
	{
		Header: "Stop and disable [cron, iptables] services!",
		ServicesInstruction: linuxservicecmd.ServicesInstruction{
			IsDetailedError:      true,
			StopServicesNames:    []string{ServiceExistCron, ServiceExistIptables},
			DisableServicesNames: []string{ServiceExistCron, ServiceExistIptables},
			Services:             linuxservicecmd.ManyServicesInstructions{},
			Validations: &linuxservicecmd.StateValidateInstructions{
				Validations: []linuxservicecmd.StateValidateInstruction{
					{
						ServiceName:      ServiceExistCron,
						ExpectedExitCode: linuxservicestate.NotRunning,
					},
					{
						ServiceName:      ServiceExistIptables,
						ExpectedExitCode: linuxservicestate.NotRunning,
					},
				},
			},
		},
	},
	{
		Header: "enable and start [cron, iptables] services and active running!",
		ServicesInstruction: linuxservicecmd.ServicesInstruction{
			IsDetailedError:     true,
			StopServicesNames:   []string{ServiceExistCron, ServiceExistIptables},
			EnableServicesNames: []string{ServiceExistCron, ServiceExistIptables},
			Services: linuxservicecmd.ManyServicesInstructions{
				IsDetailedError:   false,
				IsContinueOnError: false,
				Instructions: []linuxservicecmd.CoreServicesInstruction{
					{
						IsIgnoreUnknownService: true,
						Action:                 servicestate.Start,
						ServicesNames: []string{
							ServiceExistCron,
							ServiceExistIptables,
							ServiceNonExistUfwx, // ignored from adding error
						},
					},
				},
			},
			Validations: &linuxservicecmd.StateValidateInstructions{
				Validations: []linuxservicecmd.StateValidateInstruction{
					{
						ServiceName:      ServiceExistCron,
						ExpectedExitCode: linuxservicestate.ActiveRunning,
					},
					{
						ServiceName:      ServiceExistIptables,
						ExpectedExitCode: linuxservicestate.ActiveRunning,
					},
				},
			},
		},
	},
	{
		Header: "Restarting [cron, iptables] services and active running!",
		ServicesInstruction: linuxservicecmd.ServicesInstruction{
			IsDetailedError:   true,
			IsContinueOnError: true,
			RestartServicesNames: []string{
				ServiceExistCron,
				ServiceExistIptables,
			},
			Validations: &linuxservicecmd.StateValidateInstructions{
				Validations: []linuxservicecmd.StateValidateInstruction{
					{
						ServiceName:      ServiceExistCron,
						ExpectedExitCode: linuxservicestate.ActiveRunning,
					},
					{
						ServiceName:      ServiceExistIptables,
						ExpectedExitCode: linuxservicestate.ActiveRunning,
					},
				},
			},
		},
	},
}
