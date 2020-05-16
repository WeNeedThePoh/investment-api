package portfoliostock

import (
	"net/http"
)

//Service dependencies
type Service struct {
	PortfolioStock Model
}

//NewService service construct
func NewService(model Model) *Service {
	return &Service{PortfolioStock: model}
}

//Add stock to portfolio
func (service *Service) Add(portfolioID uint, symbol string, shares float64, costPerShare float64, stockType string) (interface{}, string, int) {
	newPortfolioStock, err := service.PortfolioStock.Add(portfolioID, symbol, stockType, shares, costPerShare)
	if err != nil {
		return nil, err.Error(), http.StatusBadRequest
	}

	return newPortfolioStock, "", 0
}

//GetAll portfolio stocks
func (service *Service) GetAll(portfolioID uint) ([]*PortfolioStock, string, int) {
	stocks, err := service.PortfolioStock.GetAll(portfolioID)
	if err != nil {
		return nil, err.Error(), http.StatusNotFound
	}

	return stocks, "", 0
}

//Get portfolio stock
func (service *Service) Get(symbol string, portfolioID uint) (interface{}, string, int) {
	stock, err := service.PortfolioStock.Get(symbol, portfolioID)
	if err != nil {
		return nil, err.Error(), http.StatusNotFound
	}

	return stock, "", 0
}

//UpdateOrAdd portfolio stock
func (service *Service) UpdateOrAdd(portfolioID uint, symbol string, shares float64, costPerShare float64, stockType string) (interface{}, string, int) {
	stock, err := service.PortfolioStock.Get(symbol, portfolioID)
	if err != nil {
		return service.Add(portfolioID, symbol, shares, costPerShare, stockType)
	}

	cost := shares * costPerShare
	err = stock.Update(shares, cost, stockType)
	if err != nil {
		return nil, err.Error(), http.StatusBadRequest
	}

	return stock, "", 0
}

//Remove stock from portfolio
func (service *Service) Remove(symbol string, portfolioID uint) (bool, string, int) {
	stock, err := service.PortfolioStock.Get(symbol, portfolioID)
	if err != nil {
		return false, err.Error(), http.StatusNotFound
	}

	err = stock.Remove()
	if err != nil {
		return false, err.Error(), http.StatusBadRequest
	}

	return true, "", 0
}
