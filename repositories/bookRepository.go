package repositories

import (
	"database/sql"
	"github.com/libraryGo/models"
	"strconv"
)

type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *BookRepo {
	return &BookRepo{
		db: db,
	}

}

func (r *BookRepo) FindByID(ID int) (*models.Book, error) {
	//request to db realization
	rows, err := r.db.Query("SELECT * FROM books WHERE ID = " + strconv.Itoa(ID))
	if err != nil {
		return nil, err
	}
	book := models.Book{}
	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Author, &book.Title)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Close()
	if err != nil {
		return &book, err
	}
	return &book, nil
}

func (r *BookRepo) GetBooks() ([]*models.Book, error){
	books := []*models.Book{}

	rows, err := r.db.Query("SELECT bookID, title, authors FROM books")
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		book := models.Book{}
		err = rows.Scan(&book.ID, &book.Author, &book.Title)
		if err!=nil{
			return []*models.Book{}, err
		}
		books = append(books, &book)
	}
	return books, nil
}