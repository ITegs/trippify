package main

import (
	"fmt"
	"github.com/ITegs/trippify/apiserver"
	"github.com/ITegs/trippify/database"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Welcome to trippify!")

	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		database.RunMigrations("trippify")
		return
	}

	users, err := database.CollectionFactory("trippify", "users")
	if err != nil {
		fmt.Println("Couldn't get collection users")
		fmt.Println(err)
		return
	}
	trips, err := database.CollectionFactory("trippify", "trips")
	if err != nil {
		fmt.Println("Couldn't get collection trips")
		fmt.Println(err)
		return
	}
	spots, err := database.CollectionFactory("trippify", "spots")
	if err != nil {
		fmt.Println("Couldn't get collection spots")
		fmt.Println(err)
		return
	}

	db := database.NewDB(users, trips, spots)

	api := apiserver.NewApiServer(db)

	api.Main()
}
