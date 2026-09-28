package main

import "fmt"

func main() {
	fmt.Println("Проект 'Простая библиотека' запущен.")

	book := Book{
		ID:     1,
		Title:  "Мастер и Маргарита",
		Author: "Михаил Булгаков",
		Year:   1967,
	}

	reader := Reader{
		ID:       1,
		Name:     "Иван",
		IsActive: true,
	}

	fmt.Println(book)

	book.IssueBook(reader)

	book.IssueBook(reader)

	book.ReturnBook()

	book.ReturnBook()
}
