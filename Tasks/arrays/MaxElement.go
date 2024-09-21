package arrays

func FindMaxElement(numbers []int) int {

	maxElement := numbers[0]

	for _, number := range numbers {
		if number > maxElement {
			maxElement = number
		}
	}

	return maxElement
}

func FindMinElement(numbers []int) int {
	minElement := numbers[0]
	for _, number := range numbers {
		if number < minElement {
			minElement = number
		}
	}

	return minElement
}
