package routes

import (
	"go-rest-api/controllers"
	"go-rest-api/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middlewares.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.Index).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.FindById).Methods("GET")

	r.HandleFunc("/api/personalities", controllers.Create).Methods("POST")
	r.HandleFunc("/api/personalities/{id}", controllers.Edit).Methods("PUT")

	r.HandleFunc("/api/personalities/{id}", controllers.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
