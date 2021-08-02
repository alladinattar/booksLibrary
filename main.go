package main

import (
	"context"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/libraryGo/handlers"
	"github.com/libraryGo/repositories"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "", log.LstdFlags)
	db, err := sql.Open("sqlite3", "db/books.db")
	if err!=nil{
		l.Fatal(err)
	}
	repo := repositories.NewBookRepo(db)

	bookHandler := handlers.NewBookHandler(l,repo)
	booksHandler := handlers.NewBooksHandler(l, repo)
	r := mux.NewRouter()
	r.Handle("/books/{id:[0-9]+}", bookHandler).Methods("GET")

	r.Handle("/books", booksHandler).Methods("GET")

	s := http.Server{
		Addr:         ":9000",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	sig := <-sigChan
	l.Println("GraceFull shutdown ", sig)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
