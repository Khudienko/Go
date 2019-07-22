package controllers

import (
	"Pet/models"
	"Pet/pkg/driver"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func ReceiveProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	row := driver.Database.QueryRow("select * from productdb.Products where id = ?", id)
	prod := models.Product{}
	err := row.Scan(&prod.Id, &prod.Code, &prod.Vendor,&prod.Country, &prod.Warranty)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		tmpl, _ := template.ParseFiles("web/edit.html")
		tmpl.Execute(w, prod)
	}
}
