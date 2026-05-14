package linuxservicecmd

func IsUnknownServiceExist(serviceName string) bool {
	return GetStatus(serviceName).IsUnknownService()
}
