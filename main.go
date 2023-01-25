package main

import (
	"fmt"
	"log"

	"github.com/bensonmacharia/book_store_api/controller"
	"github.com/bensonmacharia/book_store_api/model"
	"github.com/bensonmacharia/book_store_api/util"
	"github.com/gin-gonic/gin"

	"github.com/bensonmacharia/book_store_api/database"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Book{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(util.JWTAuth())
	protectedRoutes.POST("/book", controller.AddBook)
	protectedRoutes.GET("/books", controller.GetAllUserBooks)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
