package recover

import "fmt"

func processTasks(array []string) {
	for index, element := range array {
		func(index int, element string) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("Ошибка на задаче под индексом %d: %v\n", index, err)
				}
			}()
			if element == "" {
				panic("Задача пуста!")
			}
			fmt.Println("Выполнение", element)
		}(index, element)
	}
}
