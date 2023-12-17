package controllers

import (
	"net/http"

	"github.com/ramnath.1998/vyaaparjunction/handlers"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func GetProducts(w http.ResponseWriter, r *http.Request) {

	productList := handlers.HandleGetProducts()
	json := GetJson(productList)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}

func PostProducts(w http.ResponseWriter, r *http.Request) {

}

func DeleteProducts(w http.ResponseWriter, r *http.Request) {

}

func UpdateProducts(w http.ResponseWriter, r *http.Request) {

}

func GetCategories(w http.ResponseWriter, r *http.Request) {

}

func PostCategories(w http.ResponseWriter, r *http.Request) {

}

func DeleteCategories(w http.ResponseWriter, r *http.Request) {

}

func UpdateCategories(w http.ResponseWriter, r *http.Request) {

}
