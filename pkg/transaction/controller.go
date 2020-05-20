package transaction

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"investment-api/pkg/portfolioStock"
	"investment-api/pkg/stock"
	u "investment-api/utils"
	"net/http"
)

//Create new stock to portfolio
var Create = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	portfolioID := u.RetrieveIDParameter(params, "portfolio_id")
	data := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		u.Fail(w, "Error while decoding request body", "", http.StatusBadRequest)
		return
	}

	var portfolioStockModel = portfoliostock.NewPortfolioStock()
	var stockModel = stock.NewStock()
	portfolioStockService := portfoliostock.NewService(portfolioStockModel, stockModel)

	var model = NewTransaction()
	service := NewTransactionService(model, portfolioStockService)
	transaction, message, code := service.Create(portfolioID, data)

	if transaction == nil {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, transaction, http.StatusCreated)
	}
}

//GetAll transactions
var GetAll = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	portfolioID := u.RetrieveIDParameter(params, "portfolio_id")
	var portfolioStockModel = portfoliostock.NewPortfolioStock()
	var stockModel = stock.NewStock()
	portfolioStockService := portfoliostock.NewService(portfolioStockModel, stockModel)

	var model = NewTransaction()
	service := NewTransactionService(model, portfolioStockService)

	transactions, message, code := service.GetAll(portfolioID)

	if transactions == nil {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, transactions, http.StatusOK)
	}
}

//Get transaction
var Get = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	transactionID := u.RetrieveIDParameter(params, "transaction_id")
	var portfolioStockModel = portfoliostock.NewPortfolioStock()
	var stockModel = stock.NewStock()
	portfolioStockService := portfoliostock.NewService(portfolioStockModel, stockModel)

	var model = NewTransaction()
	service := NewTransactionService(model, portfolioStockService)

	transaction, message, code := service.Get(transactionID)
	if transaction == nil {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, transaction, http.StatusOK)
	}
}

//Update transaction
var Update = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	transactionID := u.RetrieveIDParameter(params, "transaction_id")
	data := make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		u.Fail(w, "Missing required data", "The body payload can not be empty", http.StatusBadRequest)
		return
	}

	var portfolioStockModel = portfoliostock.NewPortfolioStock()
	var stockModel = stock.NewStock()
	portfolioStockService := portfoliostock.NewService(portfolioStockModel, stockModel)

	var model = NewTransaction()
	service := NewTransactionService(model, portfolioStockService)
	updated, message, code := service.Update(transactionID, data)

	if updated == false {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, nil, http.StatusNoContent)
	}
}

//Delete transaction
var Delete = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	transactionID := u.RetrieveIDParameter(params, "transaction_id")
	var portfolioStockModel = portfoliostock.NewPortfolioStock()
	var stockModel = stock.NewStock()
	portfolioStockService := portfoliostock.NewService(portfolioStockModel, stockModel)

	var model = NewTransaction()
	service := NewTransactionService(model, portfolioStockService)

	deleted, message, code := service.Delete(transactionID)
	if deleted == false {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, nil, http.StatusNoContent)
	}
}
