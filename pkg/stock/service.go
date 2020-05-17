package stock

import (
	"net/http"
)

//Service dependencies
type Service struct {
	Stock Model
}

//NewStockService service construct
func NewStockService(model Model) *Service {
	return &Service{Stock: model}
}

//Add new stock
func (service *Service) Add(symbol string, price float64, company string, country uint) (interface{}, string, int) {
	_, err := service.Stock.GetBySymbol(symbol)
	if err == nil {
		return nil, "Symbol already exists", http.StatusBadRequest
	}

	newStock, err := service.Stock.Add(symbol, price, company, country)
	if err != nil {
		return nil, err.Error(), http.StatusBadRequest
	}

	return newStock, "", 0
}

//Get stock
func (service *Service) Get(stockID uint) (interface{}, string, int) {
	stock, err := service.Stock.Get(stockID)
	if err != nil {
		return nil, err.Error(), http.StatusNotFound
	}

	return stock, "", 0
}

//Update stock
func (service *Service) Update(symbol string, data map[string]interface{}) (bool, string, int) {
	stock, err := service.Stock.GetBySymbol(symbol)
	if err != nil {
		return false, err.Error(), http.StatusNotFound
	}

	err = stock.Update(data)
	if err != nil {
		return false, err.Error(), http.StatusBadRequest
	}

	return true, "", 0
}
