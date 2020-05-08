package stock

import (
	"encoding/json"
	u "investment-api/utils"
	"net/http"
)

//Create new stock
var Create = func(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		u.Fail(w, "Error while decoding request body", "", http.StatusBadRequest)
		return
	}

	var model = NewStock()
	service := NewStockService(model)
	stock, message, code := service.Create(data["symbol"].(string), data["price"].(float64), data["company"].(string), uint(data["country"].(float64)))

	if stock == nil {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, stock, http.StatusCreated)
	}
}

//Get stock
var Get = func(w http.ResponseWriter, r *http.Request) {
	stockID := u.RetrieveIDParameter(r, "id")
	var model = NewStock()
	service := NewStockService(model)

	portfolio, message, code := service.Get(stockID)
	if portfolio == nil {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, portfolio, http.StatusOK)
	}
}

//Update stock
var Update = func(w http.ResponseWriter, r *http.Request) {
	stockID := u.RetrieveIDParameter(r, "id")
	data := make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		u.Fail(w, "Missing required data", "The body payload can not be empty", http.StatusBadRequest)
		return
	}

	var model = NewStock()
	service := NewStockService(model)
	updated, message, code := service.Update(stockID, data)

	if updated == false {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, nil, http.StatusNoContent)
	}
}

//Delete stock
var Delete = func(w http.ResponseWriter, r *http.Request) {
	stockID := u.RetrieveIDParameter(r, "id")
	var model = NewStock()
	service := NewStockService(model)

	deleted, message, code := service.Delete(stockID)
	if deleted == false {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, nil, http.StatusNoContent)
	}
}
