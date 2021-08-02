package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/libraryGo/models"
	"log"
	"net/http"
	"strconv"
)

type BookHandler struct {
	l    *log.Logger
	repo models.DataProcesser
}

func NewBookHandler(l *log.Logger, repository models.DataProcesser) *BookHandler {
	return &BookHandler{l, repository}
}

func (b *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(404)
		b.l.Println("convert error")
	}

	b.l.Println("find by id")
	book, err := b.repo.FindByID(id)
	if err != nil {
		b.l.Println(err)
		w.WriteHeader(400)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)


	err = json.NewEncoder(w).Encode(book)
	if err!=nil{
		w.WriteHeader(400)
		b.l.Println("cannot encode data")
	}
}
