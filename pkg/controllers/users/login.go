package controllers

import (
	"Pet/models"
	_ "Pet/pkg/controllers/sessions"
	"Pet/pkg/driver"
	"fmt"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		email := r.FormValue("email")
		password := r.FormValue("password")

		row := driver.Database.QueryRow("select * from productdb.users where (email = ? and password =?)", email, password)
		user := models.Users{}
		err1 := row.Scan(&user.Email,&user.Password)
		if err != nil {
			log.Println(err1)
			http.Error(w, http.StatusText(404), http.StatusNotFound)
		}
		if user.Email != email && user.Password != password {
			fmt.Println(user.Password,user.Email)
			http.ServeFile(w, r, "web/testLogin.html")
		} else {
			http.ServeFile(w, r, "web/failLogin.html")
		}

	} else {
		http.ServeFile(w, r, "web/login.html")
	}
}
