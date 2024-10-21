package main

import (
	"log"
	"sle/config"
	"sle/database"
)

func main() {

	// load .env file only for dev mode
	// if err := godotenv.Load(); err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }

	// initialize database
	db, err := database.NewDatabase()
	if err != nil {
		log.Println("failed to connect database", err.Error())
		return
	}

	config.Configuration(db.GetDB())
	config.RunServer(":8080")
	// defer db.CloseDB()
}
