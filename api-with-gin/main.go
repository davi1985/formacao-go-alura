package main

import (
	"api-with-gin/database"
	"api-with-gin/routes"
)

func main() {
	database.ConnectDB()

	routes.HandleRequests()
}
