package slices

func removeElement(slice []int, index int) []int {
	result := append(slice[:index], slice[index+1:]...)
	return result
}
