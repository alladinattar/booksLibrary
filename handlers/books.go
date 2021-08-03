package handlers

import (
	"encoding/json"
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
	switch r.Method {
	case "GET":
		b.getBooks(w,r)
	case "POST":
		b.addBook(w, r)
	}

}

func (b *BooksHandler) getBooks(w http.ResponseWriter, r *http.Request){
	b.l.Println("all books")
	books, err := b.repo.GetBooks()
	if err != nil {
		b.l.Println(err)
		w.WriteHeader(400)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		w.WriteHeader(400)
		b.l.Println("cannot encode data")
	}
}

func (b *BooksHandler) addBook (w http.ResponseWriter, r *http.Request){
	book := models.Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err!=nil{
		b.l.Println("invalid book data")
		w.WriteHeader(400)
		return
	}
	err = b.repo.AddBook(&book)
	if err!=nil{
		b.l.Println("error when add book")
		return
	}
	b.l.Println("book added")
}