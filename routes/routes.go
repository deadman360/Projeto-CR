package routes

import (
	"github.com/deadman360/projetoLPSol/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/LP/{id}", controllers.LandingPage)

}
