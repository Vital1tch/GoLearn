package arrays

func ReverseArray(arr *[]int) []int {
	arrReversed := make([]int, len(*arr))

	for i := len(*arr) - 1; i >= 0; i-- {
		arrReversed = append(arrReversed, (*arr)[i])
	}
	return arrReversed
}
