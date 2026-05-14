package linuxservicecmd

func IsServiceExist(serviceName string) bool {
	return !GetStatus(serviceName).IsUnknownService()
}
