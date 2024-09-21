package slices

func SummOfElements(arr []int) int {

	summaryOfElements := 0
	for _, value := range arr {
		summaryOfElements += value
	}
	return summaryOfElements
}
