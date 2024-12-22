package library

import "fmt"

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	collection map[string]Book
}

func NewLibrary() *Library {
	return &Library{
		collection: make(map[string]Book),
	}
}

func (l *Library) AddBook(book Book) {
	l.collection[book.ID] = book
}

func (l *Library) BorrowBook(id string) {
	if book, exitsts := l.collection[id]; exitsts && !book.IsBorrowed {
		book.IsBorrowed = true
		l.collection[id] = book
		fmt.Println("Book borrowed:", book.Title)
	} else {
		fmt.Println("Book is not availible for borrowing.")
	}
}

func (l *Library) ReturnBook(id string) {
	if book, exists := l.collection[id]; exists && book.IsBorrowed {
		book.IsBorrowed = false
		l.collection[id] = book
		fmt.Println("Book returned:", book.Title)
	} else {
		fmt.Println("Book is not availible for return.")
	}
}

func (l *Library) ListBooks() {
	fmt.Println("Availible books:")
	for _, book := range l.collection {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}
