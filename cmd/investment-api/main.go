package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	m "investment-api/middlewares"
	"investment-api/pkg/auth"
	"investment-api/pkg/portfolio"
	"investment-api/pkg/stock"
	"investment-api/pkg/transaction"
	"investment-api/pkg/user"
	"net/http"
	"os"
)

func main() {
	router := httprouter.New()

	//AUTH
	router.POST("/login", auth.Authenticate)

	//USERS
	router.POST("/users", m.JwtAuthentication(user.Create))
	router.GET("/users/:user_id", m.JwtAuthentication(user.Get))
	router.PATCH("/users/:user_id", m.JwtAuthentication(user.Update))
	router.PATCH("/users/:user_id/password", m.JwtAuthentication(user.UpdatePassword))
	router.DELETE("/users/:user_id", m.JwtAuthentication(user.Delete))

	//STOCKS
	router.GET("/stocks", m.JwtAuthentication(stock.Search))
	router.GET("/stocks/:symbol", m.JwtAuthentication(stock.Get))

	//PORTFOLIOS
	router.POST("/users/:user_id/portfolios", m.JwtAuthentication(portfolio.Create))
	router.GET("/users/:user_id/portfolios", m.JwtAuthentication(portfolio.GetAll))
	router.GET("/users/:user_id/portfolios/:portfolio_id", m.JwtAuthentication(portfolio.Get))
	router.PATCH("/users/:user_id/portfolios/:portfolio_id", m.JwtAuthentication(portfolio.Update))
	router.DELETE("/users/:user_id/portfolios/:portfolio_id", m.JwtAuthentication(portfolio.Delete))

	//TRANSACTIONS
	router.POST("/users/:user_id/portfolios/:portfolio_id/transactions", m.JwtAuthentication(transaction.Create))
	router.GET("/users/:user_id/portfolios/:portfolio_id/transactions", m.JwtAuthentication(transaction.GetAll))
	router.GET("/users/:user_id/portfolios/:portfolio_id/transactions/:transaction_id", m.JwtAuthentication(transaction.Get))
	router.PATCH("/users/:user_id/portfolios/:portfolio_id/transactions/:transaction_id", m.JwtAuthentication(transaction.Update))
	router.DELETE("/users/:user_id/portfolios/:portfolio_id/transactions/:transaction_id", m.JwtAuthentication(transaction.Delete))

	//MIDDLEWARE
	router.NotFound = m.NotFoundHandler()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
