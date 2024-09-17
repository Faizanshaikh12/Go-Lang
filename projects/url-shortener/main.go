package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// URLMapping stores the mapping from short URLs to original URLs.
type URLMapping struct {
	OriginalURL string
}

// URLStorage holds the in-memory storage for URL mappings and provides thread-safe access.
type URLStorage struct {
	store map[string]URLMapping
	mu    sync.RWMutex
}

// NewURLStorage initilizes the URL storages.
func NewURLStorage() *URLStorage {
	return &URLStorage{store: make(map[string]URLMapping)}
}

// ShortenURL generates a random string for the shortened URL and stores the mapping.
func (s *URLStorage) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}

	// Decode the request body.
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate a short URL (base64 encoded).
	shortURL := generateShortURL()

	// Store the original URL.
	s.mu.Lock()
	s.store[shortURL] = URLMapping{OriginalURL: req.URL}
	s.mu.Unlock()

	// Respond with the shortened URL.
	json.NewEncoder(w).Encode(map[string]string{"shortened_url": "http://localhost:8080/" + shortURL})
}

// RedirectURL finds the original URL for a shortened one and redirects.
func (s *URLStorage) RedirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	fmt.Println("shortURL", shortURL)

	// Look up the original URL.
	s.mu.RLock()
	mapping, exists := s.store[shortURL]
	s.mu.RUnlock()

	fmt.Println("mapping", mapping)
	fmt.Println("exists", exists)

	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Redirect to the original URL.
	http.Redirect(w, r, mapping.OriginalURL, http.StatusFound)
}

// generateShortURL creates a random base64 string as a short URL.
func generateShortURL() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 6)
	for i := range b {
		b[i] = byte(r.Intn(256)) // Generate a random byte (0-255)
	}
	return base64.URLEncoding.EncodeToString(b)
}

func main() {
	// Initilize the URL storage
	storage := NewURLStorage()

	// Create a new router
	r := mux.NewRouter()

	// Define the routes.
	r.HandleFunc("/shorten", storage.ShortenURL).Methods("POST")
	r.HandleFunc("/{shortURL}", storage.RedirectURL).Methods("GET")

	// Start the server.
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
