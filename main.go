package main

import (
	"fmt"

	"github.com/JuanCampbsi/api_go_rest/routes"
)

func main() {
	fmt.Println("Initial app")
	routes.HandleRequest()
}
