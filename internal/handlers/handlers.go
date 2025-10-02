package handlers

import (
	"html/template"
	"io/fs"
	"net/http"

	"github.com/kokweikhong/profilm/web"
)

type Handler struct {
	Templates *template.Template
	Static    http.Handler
	WarrantyHandler *WarrantyHandler
}

func New() *Handler {
	tmpl, err := template.ParseFS(web.Content, "templates/*.html")
	if err != nil {
		panic(err)
	}
	return &Handler{
		Templates: tmpl,
		Static:    StaticHandler(),
		WarrantyHandler: NewWarrantyHandler(tmpl),
	}
}

func StaticHandler() http.Handler {
	staticFS, err := fs.Sub(web.Content, "static")
	if err != nil {
		panic(err)
	}
	return http.FileServer(http.FS(staticFS))
}
