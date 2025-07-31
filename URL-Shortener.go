package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Mapping
var (
	urlStore = make(map[String]string)
	urlStoreMu sync.RwMutex
) 

const (
	shortURLBase = "http://127.0.0.1:8080/"
	shortURLLength = 8	// Setting the length 

func main() {
	rand.Seed(time.Now().UnixNano())	// Random Generator
	http.HandleFunc("/",handleShortenURL)
	http.HandleFunc("/s/",handleRedirect)
	fmt.Println("URL shortener is running on http://127.0.0.1:8080/")
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func genShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567689"
	var shortURL strings.Builderfmt
	for i:=0; i<shortURLLength; i++ {
		shortURL.WriteByte(charset[rnad.Intn(len(charset))])
	}
	return shortURL.String()
}

func handleShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	// Read long URL from the form
	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "Missing URL", http.StatusBadRequest)
		return
	}

	// Generate a short URL
	shortURL := generateShortURL()

	// Store the mapping in the URL store
	urlStoreMu.Lock()
	urlStore[shortURL] = longURL
	urlStoreMu.Unlock()

	// Send the short URL back as a response
	fmt.Fprintf(w, "Short URL: %s%s\n", shortURLBase, shortURL)
}

// Handle request to redirect from short URL to long URL
func handleRedirect(w http.ResponseWriter, r *http.Request) {
	// Extract the short URL code from the URL
	shortURL := strings.TrimPrefix(r.URL.Path, "/s/")
	if shortURL == "" {
		http.Error(w, "Invalid short URL", http.StatusBadRequest)
		return
	}

	// Look up the long URL in the store
	urlStoreMu.RLock()
	longURL, exists := urlStore[shortURL]
	urlStoreMu.RUnlock()

	if !exists {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	// Redirect to the original long URL
	http.Redirect(w, r, longURL, http.StatusFound)
}