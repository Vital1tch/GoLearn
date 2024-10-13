package Maps

import "fmt"

func CheckAgeOfHuman(name string, humans map[string]int) int {
	if result, ok := humans[name]; ok {
		fmt.Printf("Возраст равен:%d", result)
	} else {
		fmt.Println("Человек не найден(")
	}
	return -1
}
