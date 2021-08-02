package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/libraryGo/models"
	"log"
	"net/http"
)

type BooksHandler struct {
	l    *log.Logger
	repo models.DataProcesser
}

func NewBooksHandler(l *log.Logger, repository models.DataProcesser) *BooksHandler {
	return &BooksHandler{l, repository}
}

func (b *BooksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b.l.Println("all books")
	books, err := b.repo.GetBooks()
	if err != nil {
		fmt.Println("fds")
		b.l.Println(err)
		w.WriteHeader(400)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)


	err = json.NewEncoder(w).Encode(books)
	if err!=nil{
		w.WriteHeader(400)
		b.l.Println("cannot encode data")
	}
}
