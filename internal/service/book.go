package service

import (
	"database/sql"
	"fmt"
	"time"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

type BookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{db}
}

func (s *BookService) CreateBook(book *Book) error {
	query := `INSERT INTO books (title, author, genre) VALUES (?, ?, ?)`
	result, err := s.db.Exec(query, book.Title, book.Author, book.Genre)
	if err != nil {
		return err
	}

	bookID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	book.ID = int(bookID)
	return nil
}

func (s *BookService) GetBooks() ([]Book, error) {
	query := `SELECT id, title, author, genre FROM books`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (s *BookService) GetBookById(bookId int) (*Book, error) {
	query := `SELECT id, title, author, genre FROM books WHERE id = ?`
	row := s.db.QueryRow(query, bookId)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *BookService) GetBooksByName(name string) ([]Book, error) {
	query := `SELECT id, title, author, genre FROM books WHERE title LIKE ?`
	rows, err := s.db.Query(query, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (s *BookService) UpdateBook(book *Book) error {
	query := `UPDATE books SET title = ?, author = ?, genre = ? WHERE id = ?`
	_, err := s.db.Exec(query, book.Title, book.Author, book.Genre, book.ID)
	return err
}

func (s *BookService) DeleteBook(bookId int) error {
	query := `DELETE FROM books WHERE id = ?`
	_, err := s.db.Exec(query, bookId)
	return err
}

func (s *BookService) SimulateReading(bookID int, duration time.Duration, results chan<- string) {
	book, err := s.GetBookById(bookID)
	if err != nil || book == nil {
		results <- fmt.Sprintf("Book with ID %d not found", bookID)
	}

	fmt.Printf("Reading %s by %s\n", book.Title, book.Author)
	time.Sleep(duration)
	results <- fmt.Sprintf("Finished reading %s by %s", book.Title, book.Author)
}

func (s *BookService) SimulateMultipleReading(bookIDs []int, duration time.Duration) []string {
	results := make(chan string, len(bookIDs))

	for _, id := range bookIDs {
		go func(bookID int) {
			s.SimulateReading(bookID, duration, results)
		}(id)
	}

	var responses []string
	// Append responses to the slice in the same order as the bookIDs
	for range bookIDs {
		responses = append(responses, <-results)
	}
	// This is an alternative way to get responses from the channel,
	// which would read from the channel until it is closed
	// for res := range results {
	// 	responses = append(responses, res)
	// }
	close(results)
	return responses
}
