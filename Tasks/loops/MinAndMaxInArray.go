package loops

import "fmt"

func FindMaxAndMinInArray(arr []int) {
	maxElement := arr[0]
	minElement := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxElement {
			maxElement = arr[i]
		}

		if arr[i] < minElement {
			minElement = arr[i]
		}
	}
	fmt.Printf("Цикл for:\nМаксимальный элемент: %d,\nМинимальный элемент:%d \n\n", maxElement, minElement)
}

func FindMinAndMaxInArrayRange(arr []int) {
	minElement := arr[0]
	maxElement := arr[0]

	for _, value := range arr {
		if value < minElement {
			minElement = value
		}
		if value > maxElement {
			maxElement = value
		}
	}
	fmt.Printf("Цикл range:\nМаксимальный элемент: %d,\nМинимальный элемент:%d \n", maxElement, minElement)
}
