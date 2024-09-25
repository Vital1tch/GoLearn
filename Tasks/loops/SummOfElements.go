package loops

import "fmt"

func SummOfElements(arr []int) {
	summOfElements := 0

	for i := 0; i < len(arr); i++ {
		summOfElements += arr[i]
	}
	fmt.Println("Сумма, найденная с помощью for:", summOfElements)
}

func SummOfElementsRange(arr []int) {
	summOfElements := 0

	for _, value := range arr {
		summOfElements += value
	}
	fmt.Println("Сумма, найденная с помощью range:", summOfElements)
}
