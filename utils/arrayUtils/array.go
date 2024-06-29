package arrayUtils

func Contains[T string | int | bool](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
