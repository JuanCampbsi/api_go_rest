package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JuanCampbsi/api_go_rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TotalPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}
