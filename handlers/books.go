package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/libraryGo/repositories"
)

type BooksHandler struct{
	BooksRepository *repositories.BookRepository
}

func NewBooksHandler(repository *repositories.BookRepository)*BooksHandler {
	return &BooksHandler{repository}
}


func (h *BooksHandler) GetBooks(ctx *fiber.Ctx) error {
	if ctx.FormValue("author") != ""{
		books, err := h.BooksRepository.FindByAuthor(ctx.FormValue("author"))
		if err!=nil{
			return err
		}
		ctx.JSON(books)
		return nil
	}
	books, err := h.BooksRepository.Fetch()
	if err != nil {
		return err
	}
	return ctx.JSON(books)
}




