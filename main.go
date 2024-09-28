package main

import (
	"net/http"

	"github.com/deadman360/projetoLPSol/routes"
	_ "github.com/lib/pq"
)

func main() {
	routes.HandleRequest()
	http.ListenAndServe(":8000", nil)
}
