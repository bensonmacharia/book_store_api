package controller

import (
	"errors"
	"fmt"
	"net/http"

	"book_store_api/model"
	"book_store_api/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Register(context *gin.Context) {
	var input model.Register

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		var url = context.Request.Host + context.Request.URL.String()
		message := fmt.Sprintf("Registration error for: %s %s", user.Username, err.Error())
		util.Logger(context.ClientIP(), url, 401, message)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})

	var url = context.Request.Host + context.Request.URL.String()
	message := fmt.Sprintf("User registered: %s", user.Username)
	util.Logger(context.ClientIP(), url, 201, message)

}

func Login(context *gin.Context) {
	var input model.Login

	if err := context.ShouldBindJSON(&input); err != nil {
		var errorMessage string
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			validationError := validationErrors[0]
			if validationError.Tag() == "required" {
				errorMessage = fmt.Sprintf("%s not provided", validationError.Field())
			}
		}
		context.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	user, err := model.FindUserByUsername(input.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		var url = context.Request.Host + context.Request.URL.String()
		message := fmt.Sprintf("Wrong password entered for: %s %s", user.Username, err.Error())
		util.Logger(context.ClientIP(), url, 401, message)
		return
	}

	jwt, err := util.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})

	var url = context.Request.Host + context.Request.URL.String()
	message := fmt.Sprintf("User logged in: %s", user.Username)
	util.Logger(context.ClientIP(), url, 200, message)

}
