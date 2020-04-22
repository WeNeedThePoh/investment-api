package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"investment-api/controllers"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	//AUTH
	//router.HandleFunc("/login", controllers.Authenticate).Methods("POST")

	//USERS
	//router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", controllers.GetUser).Methods("GET")
	//router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/{id:[0-9]+}", controllers.DeleteUser).Methods("DELETE")

	//MIDDLEWARE
	//router.Use(app.JwtAuthentication)

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
