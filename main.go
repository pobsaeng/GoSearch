package main

import (
	"go-search-db/database"
	"go-search-db/handlers"
	"go-search-db/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.InitializeDB()
	if err != nil {
		panic("Could not connect to the database")
	}
	db.AutoMigrate(&models.Product{})

	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/api/products/populate", handlers.PopulateProducts(db))
	r.GET("/api/products/frontend", handlers.GetFrontendProducts(db))
	r.GET("/api/products/backend", handlers.GetBackendProducts(db))

	r.Run(":8000")
}
