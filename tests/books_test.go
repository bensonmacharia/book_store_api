package tests

import (
	"net/http"
	"testing"

	"github.com/bensonmacharia/book_store_api/model"
	"github.com/stretchr/testify/assert"
)

func TestAddBook(t *testing.T) {
	newBook := model.Book{
		Title:  "Test Title",
		Author: "Test Author",
		Genre:  "Test Genre",
		UserID: 1,
	}
	writer := makeRequest("POST", "/api/book", newBook, true)
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestGetAllUserBooks(t *testing.T) {
	writer := makeRequest("GET", "/api/books", nil, true)
	assert.Equal(t, http.StatusOK, writer.Code)
}
