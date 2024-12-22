package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/trippingcoin/Pugashbek_Torekhan/Assaignment_1/Assaignment_1/Library/library"
)

func main() {
	lib := library.NewLibrary()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Choose an option: Add, Borrow, Return, List, Exit")
		fmt.Print("> ")
		command, _ := reader.ReadString('\n')
		command = command[:len(command)-1]

		switch strings.ToUpper(command) {
		case "ADD":
			var id, title, author string

			fmt.Print("Enter Book ID: ")
			id, _ = reader.ReadString('\n')
			id = id[:len(id)-1]

			fmt.Print("Enter Title: ")
			title, _ = reader.ReadString('\n')
			title = title[:len(title)-1]

			fmt.Print("Enter Author: ")
			author, _ = reader.ReadString('\n')
			author = author[:len(author)-1]

			book := library.Book{
				ID:     id,
				Title:  title,
				Author: author,
			}
			lib.AddBook(book)
			fmt.Println("Book added successfully!")

		case "BORROW":
			fmt.Print("Enter Book ID to Borrow: ")
			id, _ := reader.ReadString('\n')
			id = id[:len(id)-1]

			lib.BorrowBook(id)

		case "RETURN":
			fmt.Print("Enter Book ID to Return: ")
			id, _ := reader.ReadString('\n')
			id = id[:len(id)-1]

			lib.ReturnBook(id)

		case "LIST":
			lib.ListBooks()

		case "EXIT":
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
