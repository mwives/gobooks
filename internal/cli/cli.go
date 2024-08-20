package cli

import (
	"fmt"
	"gobooks/internal/service"
	"os"
	"strconv"
	"time"
)

type BookCLI struct {
	service *service.BookService
}

func NewBookCLI(service *service.BookService) *BookCLI {
	return &BookCLI{service}
}

func (cli *BookCLI) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gobooks <command> [<args>]")
		return
	}

	command := os.Args[1]
	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gobooks search <book title>")
			return
		}
		bookName := os.Args[2]
		cli.searchBooks(bookName)
	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gobooks simulate <book id> <book id> <book id> ...")
			return
		}
		bookIds := os.Args[2:]
		cli.simulateReading(bookIds)
	}
}

func (cli *BookCLI) searchBooks(name string) {
	books, err := cli.service.GetBooksByName(name)
	if err != nil {
		fmt.Println("Error searching books:", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("No books found")
		return
	}

	fmt.Printf("%d books found:\n", len(books))
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Genre: %s\n", book.ID, book.Title, book.Author, book.Genre)
	}
}

func (cli *BookCLI) simulateReading(bookIDsStr []string) {
	bookIDs := make([]int, len(bookIDsStr))
	for i, idStr := range bookIDsStr {
		bookID, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("Invalid book ID: %s\n", idStr)
			return
		}
		bookIDs[i] = bookID
	}

	responses := cli.service.SimulateMultipleReading(bookIDs, 2*time.Second)
	for _, response := range responses {
		fmt.Println(response)
	}
}
