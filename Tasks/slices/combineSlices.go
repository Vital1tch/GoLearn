package slices

func CombineSlices(sl []int, sl2 []int) []int {
	result := append(sl, sl2...)
	return result
}
