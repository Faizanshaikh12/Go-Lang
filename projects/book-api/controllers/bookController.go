package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"book-api/database"
	"book-api/helpers"
	"book-api/models"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var validate = validator.New()

// Fetch all books
// @Summary Get all books
// @Description Get a list of all books
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	database.DB.Find(&books)
	helpers.SendSuccessResponse(w, books, "Books fetched successfully", http.StatusOK)
	// json.NewEncoder(w).Encode(books)
}

// Add a new book
// @Summary Create a book
// @Description Create a new book
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book data"
// @Success 201 {object} models.Book
// @Router /books [post]
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	// Validate request body
	err := validate.Struct(book)
	if err != nil {
		helpers.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Create(&book)
	helpers.SendSuccessResponse(w, book, "Book created successfully", http.StatusCreated)
}

// Fetch a single book by ID
// @Summary Get a book by ID
// @Description Get a book by its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {string} string "Book not found"
// @Router /books/{id} [get]
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"]) // Validate that 'id' is an integer
	if err != nil || id <= 0 {
		helpers.SendErrorResponse(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		helpers.SendErrorResponse(w, "Book not found", http.StatusNotFound)
		return
	}

	helpers.SendSuccessResponse(w, book, "Book fetched successfully", http.StatusOK)
}

// Update a book by ID
// @Description Update a book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Book data"
// @Success 200 {object} models.Book
// @Failure 404 {string} string "Book not found"
// @Router /books/{id} [put]
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"]) // Validate 'id'
	if err != nil || id <= 0 {
		helpers.SendErrorResponse(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		helpers.SendErrorResponse(w, "Book not found", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&book)

	// Validate request body
	bodyErr := validate.Struct(book)
	if bodyErr != nil {
		helpers.SendErrorResponse(w, bodyErr.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&book)
	helpers.SendSuccessResponse(w, book, "Book updated successfully", http.StatusOK)
}

// Delete a book by ID
// @Summary Delete a book by ID
// @Description Delete a book by its ID
// @Tags books
// @Param id path int true "Book ID"
// @Success 204
// @Failure 404 {string} string "Book not found"
// @Router /books/{id} [delete]
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"]) // Validate 'id'
	if err != nil || id <= 0 {
		helpers.SendErrorResponse(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	database.DB.Delete(&models.Book{}, id)
	helpers.SendSuccessResponse(w, nil, "Book deleted successfully", http.StatusNoContent)

}
