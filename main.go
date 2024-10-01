package main

import (
	"net/http"

	"github.com/deadman360/projetoLPSol/database"
	"github.com/deadman360/projetoLPSol/routes"
)

func main() {
	database.DbConnect()
	routes.HandleRequest()
	http.ListenAndServe(":8000", nil)
}
