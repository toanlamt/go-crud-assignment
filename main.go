package main

import (
    "log"
	"github.com/gin-gonic/gin"
	"github.com/toanlamt/go-crud-assignment/config"
	"github.com/toanlamt/go-crud-assignment/internal/handlers"
	"github.com/toanlamt/go-crud-assignment/internal/infrastructure"
	"github.com/toanlamt/go-crud-assignment/internal/repositories"
	"github.com/toanlamt/go-crud-assignment/internal/routes"
	"github.com/toanlamt/go-crud-assignment/internal/services"
)

func main() {
	cfg := config.Load()

	db, err := infrastructure.NewDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository := repositories.NewProductRepository(db)
	service := services.NewProductService(repository)
	handler := handlers.NewProductHandler(service)

	router := gin.Default()

	routes.RegisterRoutes(router, handler)

	log.Println("Server started at :8082")

	router.Run(":8082")
}