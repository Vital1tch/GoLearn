package methods

import (
	"fmt"
	"sort"
)

type Book struct { //Структура книга
	title  string //Поле
	author string
	year   int
}

type Library struct { //Структура библиотека, хранит в себе массив структур книга
	books []Book
}

func CreateBook(title string, author string, year int) Book { //Создаём книгу
	return Book{ //Параметры для создания экзмепляра
		title:  title,
		author: author,
		year:   year,
	}
}

func (l *Library) AddBook(book Book) { //Добавляем книгу в массив книг структуры Library
	l.books = append( //Добавляем в поле массивов books структуры Library книгу
		l.books,
		book,
	)
}

func (l *Library) DisplayBooks() { //Метод для Library
	for index, book := range l.books { //Пробигаемся циклом, для отображения всех книг в массиве. Index + 1 - для того, чтобы не начинать с 0
		fmt.Printf("Номер книги:%d, Название: %s, Автор:%s, Год:%d \n", index+1, book.title, book.author, book.year)
	}
}

func (l *Library) SortBooksByYear() {
	sort.Slice(l.books, func(i, j int) bool { //Используем встроенный пакет sort для срезов
		return l.books[i].year < l.books[j].year // < - по убыванию, > - по возрастанию
	})
	fmt.Println("\nОтсортированные книги по году:", l.books)
}

func (l *Library) SortBooksByTitle() {
	sort.Slice(l.books, func(i, j int) bool {
		return l.books[i].title < l.books[j].title
	})
	fmt.Println("\nОтсортированные по названию:", l.books)
}
