package controller

import (
	"fmt"
	"net/http"

	"book_store_api/model"
	"book_store_api/util"
	"github.com/gin-gonic/gin"
)

func AddBook(context *gin.Context) {
	var input model.Book
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := util.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//input.UserID = user.ID

	book := model.Book{
		Title:  input.Title,
		Author: input.Author,
		Genre:  input.Genre,
		UserID: user.ID,
	}

	savedBook, err := book.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedBook})
	var url = context.Request.Host + context.Request.URL.String()
	message := fmt.Sprintf("Book added: %s by %s", book.Title, user.Username)
	util.Logger(context.ClientIP(), url, 201, message)
}

func GetAllUserBooks(context *gin.Context) {
	user, err := util.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Books})
	var url = context.Request.Host + context.Request.URL.String()
	message := fmt.Sprintf("User books queried by %s", user.Username)
	util.Logger(context.ClientIP(), url, 200, message)
}
