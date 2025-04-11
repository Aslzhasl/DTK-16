package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"violation-type-service/config"
	"violation-type-service/database"
	"violation-type-service/internal/auth"
	"violation-type-service/internal/handler"
	"violation-type-service/internal/middleware"
	"violation-type-service/internal/repository"
	"violation-type-service/internal/service"
)

func main() {
	// Загрузка переменных окружения из .env
	config.LoadEnv()

	// Подключение к БД
	db := database.InitDB()

	// Создание зависимостей
	repo := repository.NewViolationTypeRepository(db)
	svc := service.NewViolationTypeService(repo)
	h := handler.NewViolationTypeHandler(svc, repo)
	authClient := auth.NewJavaAuthClient("http://172.20.10.4:8081") // укажи актуальный IP

	// Настройка роутера
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)

	adminRouter := router.PathPrefix("/api/violation-types").Subrouter()
	adminRouter.Use(auth.JWTWithAuth(authClient, "ROLE_ADMIN"))

	adminRouter.Handle("", middleware.JWTAdminOnly(http.HandlerFunc(h.GetAll))).Methods(http.MethodGet)
	adminRouter.Handle("", middleware.JWTAdminOnly(http.HandlerFunc(h.Create))).Methods(http.MethodPost)
	adminRouter.Handle("/{id}", middleware.JWTAdminOnly(http.HandlerFunc(h.Update))).Methods(http.MethodPut)
	adminRouter.Handle("/{id}", middleware.JWTAdminOnly(http.HandlerFunc(h.Delete))).Methods(http.MethodDelete)
	adminRouter.Handle("/import", middleware.JWTAdminOnly(http.HandlerFunc(h.ImportExcel))).Methods(http.MethodPost)

	// Запуск сервера
	server := &http.Server{
		Addr:    ":8082", // или любой другой порт
		Handler: router,
	}

	log.Printf("Сервер запущен на порту %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
