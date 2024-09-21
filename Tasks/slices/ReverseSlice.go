package slices

func ReverseSlice(arr []int) []int {
	reversedSlice := make([]int, len(arr))
	for index, value := range arr {
		reversedSlice[len(arr)-1-index] = value
	}
	return reversedSlice
}
