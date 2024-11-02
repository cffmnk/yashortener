package main

import (
	"net/http"

	"github.com/cffmnk/yashortener/internal/app"
)

func main() {
	shortener := app.NewShortener()
	mux := http.NewServeMux()
	mux.HandleFunc("/", shortener.HandleShortenURL)
	mux.HandleFunc("/{id}", shortener.HandleRedirect)
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		panic(err)
	}
}
