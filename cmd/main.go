package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/Golukpal/gin-crud/internal/db"
	"github.com/Golukpal/gin-crud/internal/handlers"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	db.RunMigrations()

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/users", handlers.CreateUser)
		api.GET("/users", handlers.GetUsers)
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}