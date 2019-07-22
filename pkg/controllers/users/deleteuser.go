package controllers

import (
	"Pet/pkg/driver"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	_, err := driver.Database.Exec("delete from productdb.Users where email = ?", email)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

