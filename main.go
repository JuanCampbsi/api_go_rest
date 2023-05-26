package main

import (
	"fmt"

	"github.com/JuanCampbsi/api_go_rest/models"
	"github.com/JuanCampbsi/api_go_rest/routes"
)

func main() {
	models.Personalities = []models.Personalitie{
		{Name: "Name 1", History: "History 1"},
		{Name: "Name 2", History: "History 2"},
	}

	fmt.Println("Initial app")
	routes.HandleRequest()
}
