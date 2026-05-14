package linuxservicecmd

func IsRunning(serviceName string) bool {
	return GetStatus(serviceName).IsActiveRunning()
}
