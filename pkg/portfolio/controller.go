package portfolio

import (
	"encoding/json"
	u "investment-api/utils"
	"net/http"
)

//Create Create new portfolio
var Create = func(w http.ResponseWriter, r *http.Request) {
	id := u.RetrieveIDParameter(r, "id")
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
	id := u.RetrieveIDParameter(r, "id")
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
