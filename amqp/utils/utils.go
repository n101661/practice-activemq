package utils

func FQQN(address, queue string) string {
	if queue == "" {
		return address
	}
	return address + "::" + queue
}
