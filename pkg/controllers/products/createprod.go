package controllers

import (
	"Pet/pkg/driver"
	"log"
	"net/http"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		code := r.FormValue("code")
		vendor := r.FormValue("vendor")
		country := r.FormValue("country")
		warranty := r.FormValue("warranty")

		_, err = driver.Database.Exec("insert into productdb.Products (code, vendor,country, warranty) values (?,?,?,?)",
			code, vendor,country, warranty)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "web/create.html")
	}
}
