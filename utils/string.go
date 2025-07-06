package utils

func DerefString(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}
