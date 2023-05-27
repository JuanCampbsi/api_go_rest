package routes

import (
	"log"
	"net/http"

	"github.com/JuanCampbsi/api_go_rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.TotalPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.ReturnOnePersonalities).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
