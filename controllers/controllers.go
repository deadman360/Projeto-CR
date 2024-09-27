package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/deadman360/projetoLPSol/models"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))
func Hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello")
}
func Home(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	corretor := models.Search(id)
	templ.ExecuteTemplate(w, "Home", corretor)
}
