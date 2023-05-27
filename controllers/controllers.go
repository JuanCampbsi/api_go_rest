package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JuanCampbsi/api_go_rest/database"
	"github.com/JuanCampbsi/api_go_rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TotalPersonalities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var personalities []models.Personalitie
	database.DB.Find(&personalities)
	json.NewEncoder(w).Encode(personalities)
}

func ReturnOnePersonalities(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalities models.Personalitie
	database.DB.First(&personalities, id)

	json.NewEncoder(w).Encode(personalities)
}

func CreateOnePersonalities(w http.ResponseWriter, r *http.Request) {
	var newPersonalitie models.Personalitie
	json.NewDecoder(r.Body).Decode(&newPersonalitie)

	database.DB.Create(&newPersonalitie)
	json.NewEncoder(w).Encode(&newPersonalitie)
}

func DeleteOnePersonalities(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalities models.Personalitie

	database.DB.Delete(&personalities, id)
	json.NewEncoder(w).Encode(personalities)
}

func EditOndePersonalities(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalities models.Personalitie

	database.DB.First(&personalities, id)
	json.NewDecoder(r.Body).Decode(&personalities)
	database.DB.Save(&personalities)
	json.NewEncoder(w).Encode(personalities)
}
