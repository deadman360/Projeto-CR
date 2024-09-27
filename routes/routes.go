package routes

import (
	"net/http"

	"github.com/deadman360/projetoLPSol/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Hello)
	http.HandleFunc("/home", controllers.Home)
	
}
