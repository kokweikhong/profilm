package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

var (
	warrantySearchIndexURL = regexp.MustCompile(`^\/warranty[\/]*$`)
	warrantyInfoURL        = regexp.MustCompile(`^\/warranty\/info[\/]*$`)
)

type WarrantyHandler struct{
	Templates *template.Template
}

func NewWarrantyHandler(tmpl *template.Template) *WarrantyHandler {
	return &WarrantyHandler{
		Templates: tmpl,
	}
}

func (wh *WarrantyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case warrantySearchIndexURL.MatchString(r.URL.Path) && r.Method == http.MethodGet:
		wh.WarrantySearchIndex(w, r)
		return
	case warrantyInfoURL.MatchString(r.URL.Path) && r.Method == http.MethodGet:
		fmt.Print("warranty info page\n")
		wh.WarrantyInfoIndex(w, r)
		return
	default:
		http.NotFound(w, r)
	}
}

// warranty search page
func (wh *WarrantyHandler) WarrantySearchIndex(w http.ResponseWriter, r *http.Request) {
	wh.Templates.ExecuteTemplate(w, "warranty_search.html", nil)
}

// warranty information page
func (wh *WarrantyHandler) WarrantyInfoIndex(w http.ResponseWriter, r *http.Request) {
	wh.Templates.ExecuteTemplate(w, "warranty_information.html", nil)
}