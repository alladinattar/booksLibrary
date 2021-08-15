package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/libraryGo/handlers"
	"github.com/libraryGo/repositories"
)

func Setup(app *fiber.App, env *Env) {
	repo := repositories.NewBookRepository(env.DB)
	booksHandler := handlers.NewBooksHandler(repo)

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "john" && pass == "doe" {
				return true
			}
			if user == "admin" && pass == "123456" {
				return true
			}
			return false
		},
		Unauthorized:    nil,
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))

	books := app.Group("/books")
	books.Get("/", booksHandler.GetBooks)

}
