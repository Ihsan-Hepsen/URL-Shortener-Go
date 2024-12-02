package main

import (
	"fmt"
	"net/http"
)

var encodeMap = map[string]string{}
var decodeMap = map[string]string{}

func shortenURL(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid request method: expecting POST", http.StatusMethodNotAllowed)
		return
	}

	originalURL := req.URL.Query().Get("url")
	if originalURL == "" {
		http.Error(w, "No URL provided", http.StatusBadRequest)
		return
	}

	if shortUrl, exists := encodeMap[originalURL]; exists {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Shortened URL already exists: http://localhost:8080/%s\n", shortUrl)
		return
	}

	// Generate a short key and store in maps
	shortKey := fmt.Sprintf("XXshort%dQ", len(encodeMap)+1)
	encodeMap[originalURL] = shortKey
	decodeMap[shortKey] = originalURL

	// return the response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Shortened URL: %s\n", encodeMap[originalURL])
}

func redirect(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Invalid request method: expecting GET", http.StatusMethodNotAllowed)
		return
	}

	shortKey := req.URL.Path[1:]
	originalURL, exists := decodeMap[shortKey]
	if !exists {
		fmt.Println("Can't redirect this url, No short code found or known")
		http.Error(w, "URL short code not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, req, originalURL, http.StatusFound)
}

func getURLs(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Invalid request method: expecting GET", http.StatusMethodNotAllowed)
		return
	}

	if len(encodeMap) == 0 {
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	for originalURL, shortKey := range encodeMap {
		fmt.Fprintf(w, "\"%s\": \"http://localhost:8080/%s\",\n", originalURL, shortKey)
	}
}

func main() {
	fmt.Println("Listening on 8080...")
	http.HandleFunc("/api/shorten", shortenURL)
	http.HandleFunc("/urls", getURLs)
	http.HandleFunc("/", redirect)

	http.ListenAndServe(":8080", nil)

}
