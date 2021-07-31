package repositories

import (
	"database/sql"
	"github.com/libraryGo/models"
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
	rows, err := r.db.Query("SELECT * FROM books WHERE ID = " + string(ID))
	if err != nil {
		return nil, err
	}
	var id int
	var authors string
	var title string

	for rows.Next() {
		err = rows.Scan(&id, &authors, &title)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Close() //good habit to close
	if err != nil {
		return &models.Book{
			ID:     id,
			Title:  authors,
			Author: title,
		}, err
	}
	return &models.Book{}, nil
}
