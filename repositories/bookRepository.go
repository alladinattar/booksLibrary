package repositories

import (
	"database/sql"
	"fmt"
	"github.com/libraryGo/models"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}

}

func (r *BookRepository) Fetch() ([]*models.Book, error) {
	books := []*models.Book{}
	rows, err := r.db.Query("SELECT bookID, title, authors FROM books")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := models.Book{}
		err = rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			return []*models.Book{}, err
		}
		books = append(books, &book)
	}
	return books, nil
}

func (r *BookRepository) AddBook(book *models.Book) error {
	stmt, err := r.db.Prepare("INSERT INTO books(title, authors) values(?,?)")
	if err != nil {
		return err
	}
	result, err := stmt.Exec(book.Title, book.Author)
	if err != nil {
		return err
	}
	fmt.Print(result.RowsAffected())
	return nil
}

func (r *BookRepository) FindByAuthor(author string) ([]*models.Book, error) {
	books := []*models.Book{}
	rows, err := r.db.Query("SELECT bookID, title, authors FROM books WHERE authors='" + author + "'")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := models.Book{}
		err = rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			return []*models.Book{}, err
		}
		books = append(books, &book)
	}
	return books, nil
}
