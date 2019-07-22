package controllers

import (
	"Pet/pkg/driver"
	"log"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		name := r.FormValue("name")
		mobile := r.FormValue("mobile")
		email := r.FormValue("email")
		password := r.FormValue("password")

		_, err = driver.Database.Exec("CREATE TABLE IF NOT EXISTS `productdb`.`users`(`id` INT NOT NULL AUTO_INCREMENT COMMENT '',`name` VARCHAR(90) NOT NULL COMMENT '',\n`mobile` INT NOT NULL COMMENT '',`email` VARCHAR(45) NOT NULL COMMENT '',`password` VARCHAR(45) NOT NULL COMMENT '',PRIMARY KEY (`id`)  COMMENT '',UNIQUE INDEX `password_UNIQUE` (`password` ASC)  COMMENT '',UNIQUE INDEX `email_UNIQUE` (`email` ASC)  COMMENT '',UNIQUE INDEX `mobile_UNIQUE` (`mobile` ASC)  COMMENT '')")
		_, err = driver.Database.Exec("INSERT INTO productdb.Users(name,mobile,email,password)values (?,?,?,?)",
			name, mobile, email, password)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "web/registration.html")
	}
}
