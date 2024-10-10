package panics

import "fmt"

func readFile(path string) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic\n", err)
		}
	}()
	if path != "Hello.txt" {
		panic("Файл не найден!")
	}
	return "File opened! Content: Go is awesome!"
}
