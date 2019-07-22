package controllers

import (
	"Pet/pkg/driver"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := driver.Database.Exec("delete from productdb.Products where id = ?", id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}
