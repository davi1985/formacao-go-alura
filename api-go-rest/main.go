package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Jonh Doe", Bio: "Jonh's Biography"},
		{Id: 2, Name: "Emilly Thor", Bio: "Emilly's Biography"},
	}

	database.ConnectDB()

	fmt.Println("Server is running")
	routes.HandleRequest()
}
