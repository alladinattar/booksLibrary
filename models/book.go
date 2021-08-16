package models

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"authors"`
}

type BooksRepository interface {
	FindByAuthor(Author string) ([]*Book, error)
	Fetch() ([]*Book, error)
	Add(*Book) error
}

