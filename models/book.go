package models

type Book struct{
	ID int
	Title string
	Author string
}

type DataProcesser interface {
	FindByID(ID int) (*Book, error)
}
//bookID,title,authors,average_rating,isbn,isbn13,language_code,  num_pages,ratings_count,text_reviews_count,publication_date,publisher


