package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bensonmacharia/book_store_api/model"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	newUser := model.Register{
		Username: "johndoe",
		Email:    "johndoe@example.com",
		Password: "12qwaszxasqw12",
	}
	writer := makeRequest("POST", "/auth/register", newUser, false)
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestLogin(t *testing.T) {
	user := model.Login{
		Username: "johndoe",
		Password: "12qwaszxasqw12",
	}

	writer := makeRequest("POST", "/auth/login", user, false)

	assert.Equal(t, http.StatusOK, writer.Code)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)
	_, exists := response["jwt"]

	assert.Equal(t, true, exists)
}
