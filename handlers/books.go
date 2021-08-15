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
	books, err := h.BooksRepository.Fetch()
	if err != nil {
		return err
	}
	return ctx.JSON(books)
}



