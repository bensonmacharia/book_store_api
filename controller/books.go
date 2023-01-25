package controller

import (
	"net/http"

	"github.com/bensonmacharia/book_store_api/model"
	"github.com/bensonmacharia/book_store_api/util"
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
}

func GetAllUserBooks(context *gin.Context) {
	user, err := util.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Books})
}
