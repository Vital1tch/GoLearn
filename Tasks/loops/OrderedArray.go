package loops

import "fmt"

func SortedArray(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			fmt.Println("Массив не отсортирован")
			return false
		}
	}
	fmt.Println("Массив отсортирован")
	return true
}

func SortedArrayRange(arr []int) bool {
	for index, value := range arr {
		if value > arr[index+1] {
			return false
		}
	}
	return true
}
