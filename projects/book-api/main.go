package main

import (
	"book-api/controllers"
	"book-api/database"
	"log"
	"net/http"
	"os"

	_ "book-api/docs" // Import the generated docs

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()

	database.InitDatabase()

	r.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

	// Swagger UI route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Printf("Server is starting on port %s", port)

	log.Fatal(http.ListenAndServe(":8080", r))
}
