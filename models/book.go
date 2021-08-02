package models

type Book struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

type DataProcesser interface {
	FindByID(ID int) (*Book, error)
	GetBooks()([]*Book, error)
}
//bookID,title,authors,average_rating,isbn,isbn13,language_code,  num_pages,ratings_count,text_reviews_count,publication_date,publisher


