package shared

func Contains(arr []int, value int) bool {
	for _, a := range arr {
		if a == value {
			return true
		}
	}
	return false
}
