package stock

import (
	"github.com/gorilla/mux"
	"investment-api/external/alphavantage"
	u "investment-api/utils"
	"net/http"
)

//Get stock
var Get = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	stock, err := alphavantage.GetStock(vars["symbol"])
	if err != nil {
		u.Fail(w, err.Error(), "", 400)
	} else {
		u.Success(w, stock, http.StatusOK)
	}
}

//Search stock symbol
var Search = func(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("search")
	if searchTerm == "" {
		u.Fail(w, "Provide a search term", "", 400)
		return
	}

	stocks, err := alphavantage.SearchStock(searchTerm)
	if err != nil {
		u.Fail(w, err.Error(), "", 400)
	} else {
		u.Success(w, stocks, http.StatusOK)
	}
}
