package main

import (
	"log"
	"violation-type-service/config"
	"violation-type-service/database"
	"violation-type-service/internal/handler"
	"violation-type-service/internal/repository"
	"violation-type-service/internal/service"
	"violation-type-service/routes"
)

func main() {
	config.LoadEnv()
	db := database.ConnectDB()
	repo := repository.NewViolationTypeRepository(db)
	svc := service.NewViolationTypeService(repo)
	handler := handler.NewViolationTypeHandler(svc, repo)

	r := routes.SetupRouter(handler)

	if err := r.Run(":8081"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
