package main

import (
	"fmt"
	"log"

	"book_store_api/controller"
	"book_store_api/model"
	"book_store_api/util"

	"github.com/gin-gonic/gin"

	"book_store_api/database"

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
	err := godotenv.Load(".env")
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

	router.Run(":8080")
	fmt.Println("Server running on port 8080")
}
