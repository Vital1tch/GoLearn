package slices

func FillSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = i + 1
	}
	slice = append(slice, 10, 123, 111)

}
