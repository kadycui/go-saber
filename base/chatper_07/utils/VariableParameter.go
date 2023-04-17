package utils

func Add(slice ...int) int {
	sum := 0
	for _, value := range slice {
		sum = sum + value
	}
	return sum
}
