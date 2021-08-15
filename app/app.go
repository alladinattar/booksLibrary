package app

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

type Env struct {
	DB  *sql.DB
	Cfg *Config
}

func Run() error {
	app := fiber.New()
	cfg, err := NewConfig("./config.yaml")
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", cfg.Database.Path)
	if err != nil {
		return err
	}
	env := &Env{db, cfg}
	Setup(app, env)

	err = app.Listen(env.Cfg.Server.Port)
	if err != nil {
		return err
	}
	return nil
}
