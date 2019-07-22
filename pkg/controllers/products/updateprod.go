package controllers

import (
	"Pet/pkg/driver"
	"log"
	"net/http"
)

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	id := r.FormValue("id")
	code := r.FormValue("code")
	vendor := r.FormValue("vendor")
	country := r.FormValue("country")
	warranty := r.FormValue("warranty")

	_, err = driver.Database.Exec("update productdb.Products set code=?, vendor=?,country=?, warranty = ? where id = ?",
		code, vendor,country, warranty, id)

	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", 301)
}
