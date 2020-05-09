package portfolio

import (
	"encoding/json"
	u "investment-api/utils"
	"net/http"
)

//Create Create new portfolio
var Create = func(w http.ResponseWriter, r *http.Request) {
	id := u.RetrieveIDParameter(r, "user_id")
	data := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		u.Fail(w, "Error while decoding request body", "", http.StatusBadRequest)
		return
	}

	var model = NewPortfolio()
	service := NewPortfolioService(model)
	user, message, code := service.Create(id, data["name"].(string))

	if user == nil {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, user, http.StatusCreated)
	}
}

//Get get user portfolio
var Get = func(w http.ResponseWriter, r *http.Request) {
	id := u.RetrieveIDParameter(r, "user_id")
	portfolioID := u.RetrieveIDParameter(r, "portfolio_id")
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
var Update = func(w http.ResponseWriter, r *http.Request) {
	userID := u.RetrieveIDParameter(r, "user_id")
	id := u.RetrieveIDParameter(r, "portfolio_id")
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
var Delete = func(w http.ResponseWriter, r *http.Request) {
	id := u.RetrieveIDParameter(r, "user_id")
	portfolioID := u.RetrieveIDParameter(r, "portfolio_id")
	var model = NewPortfolio()
	service := NewPortfolioService(model)

	deleted, message, code := service.Delete(id, portfolioID)
	if deleted == false {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, nil, http.StatusNoContent)
	}
}
