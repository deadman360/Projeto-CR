package main

import (
	"net/http"

	"github.com/deadman360/projetoLPSol/routes"
)

func main() {
	routes.HandleRequest()
	http.ListenAndServe(":8000", nil)
}
