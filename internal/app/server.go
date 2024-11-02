package app

import (
	"fmt"
	"io"
	"net/http"
)

type Shortener struct {
	store Storage
}

func NewShortener() *Shortener {
	return &Shortener{
		store: NewMemStorage(),
	}
}

func (s *Shortener) HandleShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error while reading body", http.StatusBadRequest)
		return
	}

	originURL := string(body)

	if len(originURL) == 0 {
		http.Error(w, "Empty body", http.StatusBadRequest)
		return
	}

	shortID := s.store.SaveURL(originURL)

	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("http://localhost:8080/%s", shortID)))
}

func (s *Shortener) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	shortID := r.URL.Path[1:]
	originURL, err := s.store.GetOriginURL(shortID)
	if err != nil {
		http.Error(w, "URL Not Found", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, originURL, http.StatusTemporaryRedirect)
}
