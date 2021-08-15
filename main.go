package main

import (
	"github.com/libraryGo/app"
	_ "github.com/mattn/go-sqlite3"
	"log"
)



func main() {
	err := app.Run()
	if err!=nil{
		log.Fatal(err)
	}
}
