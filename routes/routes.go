package routes

import (
	"log"
	"net/http"

	"github.com/JuanCampbsi/api_go_rest/controllers"
	"github.com/JuanCampbsi/api_go_rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.TotalPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.ReturnOnePersonalities).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.CreateOnePersonalities).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.DeleteOnePersonalities).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.EditOndePersonalities).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
