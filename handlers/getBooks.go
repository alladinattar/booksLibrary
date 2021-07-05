package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type getBook struct {
	l *log.Logger
}

func NewGetBook(l *log.Logger) *getBook {
	return &getBook{l}
}

func (gb *getBook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gb.l.Println("getBook start")
	_, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oops, getBook cannot work", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "It's your book")
}
