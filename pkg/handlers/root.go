package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// Root handles requests on main page
func Root(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/layout.html", "templates/generate.html")
	if err != nil {
		log.Fatalf("Cannot parse `generate.html` or `layout.html`: %s", err)
	}

	t.ExecuteTemplate(w, "layout", nil)
}
