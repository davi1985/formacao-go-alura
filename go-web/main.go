package main

import (
	"net/http"

	"github.com/formacao-go/go-web/routes"
)

func main() {
	routes.GetRoutes()
	http.ListenAndServe(":8000", nil)
}
