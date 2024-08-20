package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fabiovige/gobooks/internal/service"
)

type BookHandlers struct {
	service *service.BookService
}

func NewBookHandlers(service *service.BookService) *BookHandlers {
	return &BookHandlers{service: service}
}

func (h *BookHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil || size < 1 {
		size = 10 // padrão para 10 itens por página
	}

	offset := (page - 1) * size

	// Capturando os filtros da URL
	filters := map[string]string{
		"id":     r.URL.Query().Get("id"),
		"title":  r.URL.Query().Get("title"),
		"author": r.URL.Query().Get("author"),
		"genre":  r.URL.Query().Get("genre"),
	}

	// Recupera o total de registros
	total, err := h.service.CountBooks(filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	books, err := h.service.GetBooks(offset, size, filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lastPage := (total + size - 1) / size

	pagination := service.PaginatedBooks{
		Total:        total,
		PerPage:      size,
		CurrentPage:  page,
		LastPage:     lastPage,
		FirstPageURL: fmt.Sprintf("/books?page=1&size=%d", size),
		LastPageURL:  fmt.Sprintf("/books?page=%d&size=%d", lastPage, size),
		NextPageURL:  "",
		PrevPageURL:  "",
		From:         offset + 1,
		To:           offset + len(books),
		Data:         books,
	}

	if page > 1 {
		pagination.PrevPageURL = fmt.Sprintf("/books?page=%d&size=%d", page-1, size)
	}
	if page < lastPage {
		pagination.NextPageURL = fmt.Sprintf("/books?page=%d&size=%d", page+1, size)
	}

	err = json.NewEncoder(w).Encode(pagination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (h *BookHandlers) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book service.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.CreateBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandlers) GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := h.service.GetBookByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandlers) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var book service.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book.ID = id
	err = h.service.UpdateBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandlers) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
