package utils

func ItemExists(arrayType []int, item int) bool {
	for i := 0; i < len(arrayType); i++ {
		if arrayType[i] == item {
			return true
		}
	}

	return false
}
