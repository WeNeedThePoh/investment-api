package portfolio

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	u "investment-api/utils"
	"net/http"
)

//Create new portfolio
var Create = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := u.RetrieveIDParameter(params, "user_id")
	data := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		u.Fail(w, "Error while decoding request body", "", http.StatusBadRequest)
		return
	}

	var model = NewPortfolio()
	service := NewPortfolioService(model)
	portfolio, message, code := service.Create(id, data["name"].(string))

	if portfolio == nil {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, portfolio, http.StatusCreated)
	}
}

//GetAll user portfolios
var GetAll = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := u.RetrieveIDParameter(params, "user_id")
	var model = NewPortfolio()
	service := NewPortfolioService(model)

	portfolios, message, code := service.GetAll(id)

	if portfolios == nil {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, portfolios, http.StatusOK)
	}
}

//Get get user portfolio
var Get = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := u.RetrieveIDParameter(params, "user_id")
	portfolioID := u.RetrieveIDParameter(params, "portfolio_id")
	var model = NewPortfolio()
	service := NewPortfolioService(model)

	portfolio, message, code := service.Get(id, portfolioID)
	if portfolio == nil {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, portfolio, http.StatusOK)
	}
}

//Update portfolio
var Update = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := u.RetrieveIDParameter(params, "user_id")
	id := u.RetrieveIDParameter(params, "portfolio_id")
	data := make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		u.Fail(w, "Missing required data", "The body payload can not be empty", http.StatusBadRequest)
		return
	}

	var model = NewPortfolio()
	service := NewPortfolioService(model)
	updated, message, code := service.Update(userID, id, data)

	if updated == false {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, nil, http.StatusNoContent)
	}
}

//Delete portfolio
var Delete = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := u.RetrieveIDParameter(params, "user_id")
	portfolioID := u.RetrieveIDParameter(params, "portfolio_id")
	var model = NewPortfolio()
	service := NewPortfolioService(model)

	deleted, message, code := service.Delete(id, portfolioID)
	if deleted == false {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, nil, http.StatusNoContent)
	}
}
