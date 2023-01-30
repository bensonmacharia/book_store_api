package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"book_store_api/controller"
	"book_store_api/database"
	"book_store_api/model"
	"book_store_api/util"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setup()
	exitCode := m.Run()
	teardown()

	os.Exit(exitCode)
}

func router() *gin.Engine {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(util.JWTAuth())
	protectedRoutes.POST("/book", controller.AddBook)
	protectedRoutes.GET("/books", controller.GetAllUserBooks)

	return router
}

func setup() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectTest()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Book{})
}

func teardown() {
	migrator := database.Database.Migrator()
	migrator.DropTable(&model.User{})
	migrator.DropTable(&model.Book{})
}

func makeRequest(method, url string, body interface{}, isAuthenticatedRequest bool) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if isAuthenticatedRequest {
		request.Header.Add("Authorization", "Bearer "+bearerToken())
	}
	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}

func bearerToken() string {
	user := model.Login{
		Username: "johndoe",
		Password: "12qwaszxasqw12",
	}

	writer := makeRequest("POST", "/auth/login", user, false)
	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)
	return response["jwt"]
}
