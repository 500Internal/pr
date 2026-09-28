package main

import "fmt"

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
}

type Reader struct {
	ID       int
	Name     string
	IsActive bool
}

func (b Book) String() string {
	return fmt.Sprintf(`"%s" (%s, %d)`, b.Title, b.Author, b.Year)
}

func (b *Book) IssueBook(reader Reader) {
	if b.IsIssued {
		fmt.Printf("Книга %s уже кому-то выдана\n", b.Title)
		return
	}

	if !reader.IsActive {
		fmt.Printf("Читатель %s не активен и не может получить книгу.\n", reader.Name)
		return
	}

	b.IsIssued = true
	fmt.Printf("Книга %s была выдана читателю %s\n", b.Title, reader.Name)
}

func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга %s и так в библиотеке\n", b.Title)
		return
	}

	b.IsIssued = false
	fmt.Printf("Книга %s возвращена в библиотеку\n", b.Title)
}
