package main

import (
	controllers2 "Pet/pkg/controllers"
	controllers3 "Pet/pkg/controllers/products"
	"Pet/pkg/controllers/sessions"
	controllers "Pet/pkg/controllers/users"
	"Pet/pkg/driver"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	driver.OpenConnection()
	defer driver.Database.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", controllers2.IndexHandler)
	router.HandleFunc("/create", controllers3.CreateProductHandler)
	router.HandleFunc("/edit/{id:[0-9]+}", controllers3.ReceiveProductHandler).Methods("GET")
	router.HandleFunc("/edit/{id:[0-9]+}", controllers3.UpdateProductHandler).Methods("POST")
	router.HandleFunc("/delete/{id:[0-9]+}", controllers3.DeleteProductHandler)
	router.HandleFunc("/registration", controllers.CreateUserHandler)
	router.HandleFunc("/login", controllers.LoginHandler)

	router.HandleFunc("/delete/{email}",controllers.DeleteUserHandler)


	router.HandleFunc("/secret", sessions.Secret)
	//router.HandleFunc("/login", sessions.Login)
	router.HandleFunc("/logout", sessions.Logout)


	http.Handle("/", router)

	fmt.Println("Server is listening...")
	_ = http.ListenAndServe(":8181", handlers.LoggingHandler(os.Stdout, router))
}

