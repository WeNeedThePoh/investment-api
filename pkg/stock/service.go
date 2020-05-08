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

//Create new stock
func (service *Service) Create(symbol string, price float64, company string, country uint) (map[string]interface{}, string, int) {
	_, err := service.Stock.GetBySymbol(symbol)
	if err == nil {
		return nil, "Symbol already exists", http.StatusBadRequest
	}

	newStock, err := service.Stock.Create(symbol, price, company, country)
	if err != nil {
		return nil, err.Error(), http.StatusBadRequest
	}

	resp := newStock.ToMap()
	return resp, "", 0
}

//Get stock
func (service *Service) Get(stockID uint) (map[string]interface{}, string, int) {
	stock, err := service.Stock.Get(stockID)
	if err != nil {
		return nil, "stock not found", http.StatusNotFound
	}

	resp := stock.ToMap()
	return resp, "", 0
}

//Update stock
func (service *Service) Update(stockID uint, data map[string]interface{}) (bool, string, int) {
	stock, err := service.Stock.Get(stockID)
	if err != nil {
		return false, "stock not found", http.StatusNotFound
	}

	err = stock.Update(data)
	if err != nil {
		return false, err.Error(), http.StatusBadRequest
	}

	return true, "", 0
}

//Delete stock
func (service *Service) Delete(stockID uint) (bool, string, int) {
	stock, err := service.Stock.Get(stockID)
	if err != nil {
		return false, "stock not found", http.StatusNotFound
	}

	err = stock.Delete()
	if err != nil {
		return false, err.Error(), http.StatusBadRequest
	}

	return true, "", 0
}
