package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/formacao-go/go-web/models"
)

const HTTP_CODE_301 = 301

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()

	templates.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConvertedToFloat, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("Not possible convert price:", err)
		}

		quantityConvertedToInt, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("Not possible convert quantity:", err)
		}

		models.Create(name, description, priceConvertedToFloat, quantityConvertedToInt)
	}

	http.Redirect(w, r, "/", HTTP_CODE_301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")

	models.Delete(productID)

	http.Redirect(w, r, "/", HTTP_CODE_301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")

	product := models.Edit(productID)

	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConverted, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Not possible convert ID:", err)
		}

		priceConvertedToFloat, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("Not possible convert price:", err)
		}

		quantityConvertedToInt, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("Not possible convert quantity:", err)
		}

		models.Update(idConverted, name, description, priceConvertedToFloat, quantityConvertedToInt)
	}

	http.Redirect(w, r, "/", HTTP_CODE_301)

}
