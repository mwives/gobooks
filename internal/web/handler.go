package web

import (
	"encoding/json"
	"gobooks/internal/service"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type BookHandlers struct {
	service *service.BookService
}

func NewBookHandlers(service *service.BookService) *BookHandlers {
	return &BookHandlers{service}
}

func (h *BookHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	bookTitle := r.URL.Query().Get("title")
	var books []service.Book
	var err error

	if bookTitle != "" {
		books, err = h.service.GetBooksByName(bookTitle)
	} else {
		books, err = h.service.GetBooks()
	}

	if err != nil {
		http.Error(w, "failed to fetch books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandlers) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book service.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateBook(&book); err != nil {
		http.Error(w, "failed to create book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandlers) GetBookById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	book, err := h.service.GetBookById(id)
	if err != nil {
		http.Error(w, "failed to fetch book", http.StatusInternalServerError)
		return
	}
	if book == nil {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandlers) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusInternalServerError)
	}

	var book service.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	book.ID = id

	if err := h.service.UpdateBook(&book); err != nil {
		http.Error(w, "failed to update book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandlers) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusInternalServerError)
	}

	if err := h.service.DeleteBook(id); err != nil {
		http.Error(w, "failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *BookHandlers) SimulateReadingBooks(w http.ResponseWriter, r *http.Request) {
	bookIDsStr := r.URL.Query().Get("book_ids")
	if bookIDsStr == "" {
		http.Error(w, "missing book ids", http.StatusBadRequest)
		return
	}

	bookIDsList := strings.Split(bookIDsStr, ",")
	if len(bookIDsList) == 0 || (len(bookIDsList) == 1 && bookIDsList[0] == "") {
		http.Error(w, "invalid book ids", http.StatusBadRequest)
		return
	}

	bookIDs := make([]int, len(bookIDsList))
	for i, idStr := range bookIDsList {
		bookID, err := strconv.Atoi(strings.TrimSpace(idStr))
		if err != nil {
			http.Error(w, "invalid book id", http.StatusBadRequest)
			return
		}
		bookIDs[i] = bookID
	}

	responses := h.service.SimulateMultipleReading(bookIDs, 2*time.Second)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}
