package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/libraryGo/models"
	"log"
	"net/http"
	"strconv"
)

type BooksHandler struct {
	l *log.Logger
	repo models.DataProcesser
}

func NewBookHandler(l *log.Logger, repository models.DataProcesser) *BooksHandler {
	return &BooksHandler{l, repository}
}

func (b *BooksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err!=nil{
		w.WriteHeader(404)
	}
	book, err := b.repo.FindByID(id)
	if err!=nil{
		b.l.Fatal(err)
	}
	fmt.Fprintf(w, "It's your book", book.Title)
}


