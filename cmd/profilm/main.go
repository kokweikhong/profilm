package main

import (
	"net/http"

	"github.com/kokweikhong/profilm/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	h := handlers.New()
	mux.Handle("/warranty", h.WarrantyHandler)
	mux.Handle("/warranty/", h.WarrantyHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", h.Static))
	http.ListenAndServe(":8080", mux)
}
