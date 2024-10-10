package panics

import "fmt"

func getElement(slice []string, element int) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic")
		}
	}()
	for index, _ := range slice {
		if index == element {
			return slice[element]
		}
	}
	panic("Index traceback")
}
