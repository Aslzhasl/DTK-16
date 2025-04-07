package main

import (
	"violation-type-service/config"
	"violation-type-service/database"
	"violation-type-service/routes"
	"log"
)

func main() {
	config.LoadEnv()
	db := database.ConnectDB()
	r := routes.SetupRouter(db)

	if err := r.Run(":8081"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
