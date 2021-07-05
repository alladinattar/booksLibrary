package main

import (
	"context"
	"github.com/libraryGo/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "library", log.LstdFlags)
	gbHandler := handlers.NewGetBook(l)
	mux := http.NewServeMux()
	mux.Handle("/getBook", gbHandler)

	s := http.Server{
		Addr:         ":8080",
		Handler:      mux,
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
