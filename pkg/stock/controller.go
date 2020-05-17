package stock

import (
	"github.com/gorilla/mux"
	"investment-api/external/alphavantage"
	u "investment-api/utils"
	"net/http"
	"strconv"
)

//Get stock
var Get = func(w http.ResponseWriter, r *http.Request) {
	symbol := mux.Vars(r)["symbol"]

	stock, err := alphavantage.GetStock(symbol)
	if err != nil {
		u.Fail(w, err.Error(), "", 400)
	}

	stocksMatches := stock.(map[string]interface{})
	stockMap := stocksMatches["Global Quote"].(map[string]interface{})

	var model = NewStock()
	service := NewStockService(model)
	price, _ := strconv.ParseFloat(stockMap["05. price"].(string), 64)
	dailyChange, _ := strconv.ParseFloat(stockMap["09. change"].(string), 64)
	dailyChangePercent, _ := strconv.ParseFloat(stockMap["10. change percent"].(string), 64)
	data := map[string]interface{}{
		"price": price,
		"daily_change": dailyChange,
		"daily_change_percentage": dailyChangePercent,
	}
	service.Update(symbol, data)

	u.Success(w, stock, http.StatusOK)
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
	}

	stocksMatches := stocks.(map[string]interface{})
	stocksMap := stocksMatches["bestMatches"].([]interface{})
	if len(stocksMap) > 0 {
		stock := stocksMap[0].(map[string]interface{})
		var model = NewStock()
		service := NewStockService(model)
		service.Add(stock["1. symbol"].(string), 0, stock["2. name"].(string), 1)
	}

	u.Success(w, stocks, http.StatusOK)
}
