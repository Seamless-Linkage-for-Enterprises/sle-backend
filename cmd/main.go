package main

import (
	"log"
	"sle/config"
	"sle/database"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Println(err.Error())
		return
	}

	// initialize database
	db, err := database.NewDatabase()
	if err != nil {
		log.Println("failed to connect database", err.Error())
		return
	}

	emailHand := config.Configuration(db.GetDB())
	go func() {
		for {
			if err := emailHand.DeleteOTPs(); err != nil {
				log.Println(err.Error())
			}
			time.Sleep(30 * time.Second)
		}
	}()
	config.RunServer(":8080")
	// defer db.CloseDB()
}
