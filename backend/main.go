package main

import (
	"fmt"
	"github.com/ITegs/trippify/apiserver"
	"github.com/ITegs/trippify/database"
)

func main() {
	fmt.Println("Welcome to trippify!")

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
	db := database.NewDB(users, trips)

	api := apiserver.NewApiServer(db)

	api.Main()
}
