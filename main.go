package main

import (
	"testGoLang/config"
	"testGoLang/model"
	"testGoLang/router"
)

func main() {
	db := config.DatabaseConfig()
	db.AutoMigrate(&model.User{})

	router.Router(db)
}
