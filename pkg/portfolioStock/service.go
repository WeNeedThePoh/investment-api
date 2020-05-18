package portfoliostock

import (
	"investment-api/pkg/stock"
	"net/http"
)

//Service dependencies
type Service struct {
	PortfolioStock Model
	Stock          stock.Model
}

//NewService service construct
func NewService(model Model, stock stock.Model) *Service {
	return &Service{PortfolioStock: model, Stock: stock}
}

//Add stock to portfolio
func (service *Service) Add(portfolioID uint, symbol string, shares float64, costPerShare float64, stockType string) (interface{}, string, int) {
	stockToAdd, _ := service.Stock.GetBySymbol(symbol)
	marketValue := stockToAdd.Price * shares

	cost := shares * costPerShare
	data := map[string]interface{}{
		"portfolio_id":            portfolioID,
		"symbol":                  symbol,
		"type":                    stockType,
		"shares":                  shares,
		"avg_share_cost":          costPerShare,
		"market_value":            marketValue,
		"cost":                    cost,
		"total_change":            marketValue - cost,
		"total_change_percentage": (marketValue - cost) / marketValue * 100,
	}
	newPortfolioStock, err := service.PortfolioStock.Add(data)
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
	portfolioStock, err := service.PortfolioStock.Get(symbol, portfolioID)
	if err != nil {
		return nil, err.Error(), http.StatusNotFound
	}

	return portfolioStock, "", 0
}

//UpdateOrAdd portfolio stock
func (service *Service) UpdateOrAdd(portfolioID uint, symbol string, shares float64, costPerShare float64, stockType string) (interface{}, string, int) {
	stock, err := service.PortfolioStock.Get(symbol, portfolioID)
	if err != nil {
		return service.Add(portfolioID, symbol, shares, costPerShare, stockType)
	}

	stockToAdd, _ := service.Stock.GetBySymbol(symbol)
	marketValue := stockToAdd.Price * shares

	cost := (shares * costPerShare) + stock.Cost
	data := map[string]interface{}{
		"portfolio_id":            portfolioID,
		"symbol":                  symbol,
		"type":                    stockType,
		"shares":                  stock.Shares + shares,
		"avg_share_cost":          costPerShare,
		"market_value":            marketValue,
		"cost":                    cost,
		"total_change":            marketValue - cost,
		"total_change_percentage": (marketValue - cost) / marketValue * 100,
	}

	err = stock.Update(data)
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
