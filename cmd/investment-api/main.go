package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"investment-api/middlewares"
	"investment-api/pkg/auth"
	"investment-api/pkg/user"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	//AUTH
	router.HandleFunc("/login", auth.Authenticate).Methods("POST")

	//USERS
	router.HandleFunc("/users", user.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", user.GetUser).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", user.UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/{id:[0-9]+}/password", user.UpdateUserPassword).Methods("PATCH")
	router.HandleFunc("/users/{id:[0-9]+}", user.DeleteUser).Methods("DELETE")

	//MIDDLEWARES
	router.Use(middlewares.JwtAuthentication)
	router.NotFoundHandler = middlewares.NotFoundHandler()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
