package slices

type MinMax struct {
	Min int
	Max int
}

func FindMaxAndMinElement(nums []int) MinMax {
	minElement := nums[0]
	for _, value := range nums {
		if value < minElement {
			minElement = value
		}
	}
	maxElement := nums[0]
	for _, value := range nums {
		if value > maxElement {
			maxElement = value
		}
	}
	return MinMax{Min: minElement, Max: maxElement}
}
