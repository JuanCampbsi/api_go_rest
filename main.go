package main

import (
	"fmt"

	"github.com/JuanCampbsi/api_go_rest/database"
	"github.com/JuanCampbsi/api_go_rest/models"
	"github.com/JuanCampbsi/api_go_rest/routes"
)

func main() {
	models.Personalities = []models.Personalitie{
		{Id: 1, Name: "Name 1", History: "History 1"},
		{Id: 2, Name: "Name 2", History: "History 2"},
		{Id: 3, Name: "Name 3", History: "History 3"},
	}
	database.ConnectDataBase()
	fmt.Println("Initial App rest GO")
	routes.HandleRequest()
}
