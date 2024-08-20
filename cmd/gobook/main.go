package main

import (
	"database/sql"
	"fmt"
	"gobooks/internal/cli"
	"gobooks/internal/service"
	"gobooks/internal/web"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	booksService := service.NewBookService(db)
	bookHandlers := web.NewBookHandlers(booksService)

	if len(os.Args) > 1 {
		bookCLI := cli.NewBookCLI(booksService)
		bookCLI.Run()
		return
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /books", bookHandlers.GetBooks)
	router.HandleFunc("GET /books/{id}", bookHandlers.GetBookById)
	router.HandleFunc("POST /books", bookHandlers.CreateBook)
	router.HandleFunc("PUT /books/{id}", bookHandlers.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandlers.DeleteBook)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
