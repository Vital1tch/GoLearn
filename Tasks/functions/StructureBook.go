package functions

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func PrintBookInfo(b Book) {
	fmt.Printf("Название: %s, Автор: %s, Кол-во страниц: %d", b.Title, b.Author, b.Pages)
}
