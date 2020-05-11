package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"investment-api/middlewares"
	"investment-api/pkg/auth"
	"investment-api/pkg/portfolio"
	"investment-api/pkg/stock"
	"investment-api/pkg/user"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	//AUTH
	router.HandleFunc("/login", auth.Authenticate).Methods("POST")

	//USERS
	router.HandleFunc("/users", user.Create).Methods("POST")
	router.HandleFunc("/users/{user_id:[0-9]+}", user.Get).Methods("GET")
	router.HandleFunc("/users/{user_id:[0-9]+}", user.Update).Methods("PATCH")
	router.HandleFunc("/users/{user_id:[0-9]+}/password", user.UpdatePassword).Methods("PATCH")
	router.HandleFunc("/users/{user_id:[0-9]+}", user.Delete).Methods("DELETE")

	//STOCKS
	router.HandleFunc("/stocks", stock.Search).Methods("GET")
	router.HandleFunc("/stocks/{symbol}", stock.Get).Methods("GET")

	//PORTFOLIOS
	router.HandleFunc("/users/{user_id:[0-9]+}/portfolios", portfolio.Create).Methods("POST")
	router.HandleFunc("/users/{user_id:[0-9]+}/portfolios", portfolio.GetAll).Methods("GET")
	router.HandleFunc("/users/{user_id:[0-9]+}/portfolios/{portfolio_id:[0-9]+}", portfolio.Get).Methods("GET")
	router.HandleFunc("/users/{user_id:[0-9]+}/portfolios/{portfolio_id:[0-9]+}", portfolio.Update).Methods("PATCH")
	router.HandleFunc("/users/{user_id:[0-9]+}/portfolios/{portfolio_id:[0-9]+}", portfolio.Delete).Methods("DELETE")

	//MIDDLEWARE
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
