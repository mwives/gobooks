package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"gobooks/internal/cli"
	"gobooks/internal/service"
	"gobooks/internal/web"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConnStr := os.Getenv("DB_CONN_STRING")
	if dbConnStr == "" {
		dbConnStr = "gobooks:gobooks@tcp(localhost:3306)/gobooks" // Default to localhost
	}

	db, err := sql.Open("mysql", dbConnStr)
	if err != nil {
		log.Fatal(err)
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
	router.HandleFunc("POST /books/read", bookHandlers.SimulateReadingBooks)
	router.HandleFunc("PUT /books/{id}", bookHandlers.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandlers.DeleteBook)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
