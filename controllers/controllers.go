package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/deadman360/projetoLPSol/database"
	"github.com/deadman360/projetoLPSol/models"
	"github.com/gorilla/mux"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func LandingPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var c models.Corretor
	database.Db.First(&c, id)
	json.NewEncoder(w).Encode(c)
}
