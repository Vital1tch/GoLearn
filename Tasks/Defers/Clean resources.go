package Defers

import "fmt"

func openFile(filename string) {
	defer fmt.Println("Файл закрыт")
	fmt.Printf("Файл %s открыт\n", filename)
	fmt.Println("Обработка файла.....")
}
