package transaction

import (
	"net/http"
)

//Service dependencies
type Service struct {
	Transaction Model
}

//NewTransactionService service construct
func NewTransactionService(model Model) *Service {
	return &Service{Transaction: model}
}

//Create new transaction
func (service *Service) Create(portfolioID uint, data map[string]interface{}) (interface{}, string, int) {
	stockID := uint(data["stock_id"].(float64))
	transactionType := data["type"].(string)
	shares := data["shares"].(float64)
	costPerShare := data["cost_per_share"].(float64)
	fees := data["fees"].(float64)

	newTransaction, err := service.Transaction.Create(portfolioID, stockID, transactionType, shares, costPerShare, fees)
	if err != nil {
		return nil, err.Error(), http.StatusBadRequest
	}

	return newTransaction, "", 0
}

//GetAll portfolio transactions
func (service *Service) GetAll(portfolioID uint) ([]*Transaction, string, int) {
	portfolios, err := service.Transaction.GetAll(portfolioID)
	if err != nil {
		return nil, err.Error(), http.StatusNotFound
	}

	return portfolios, "", 0
}

//Get transaction
func (service *Service) Get(transactionID uint) (interface{}, string, int) {
	transaction, err := service.Transaction.Get(transactionID)
	if err != nil {
		return nil, err.Error(), http.StatusNotFound
	}

	return transaction, "", 0
}

//Update transaction
func (service *Service) Update(transactionID uint, data map[string]interface{}) (bool, string, int) {
	transaction, err := service.Transaction.Get(transactionID)
	if err != nil {
		return false, err.Error(), http.StatusNotFound
	}

	err = transaction.Update(data)
	if err != nil {
		return false, err.Error(), http.StatusBadRequest
	}

	return true, "", 0
}

//Delete transaction
func (service *Service) Delete(transactionID uint) (bool, string, int) {
	transaction, err := service.Transaction.Get(transactionID)
	if err != nil {
		return false, err.Error(), http.StatusNotFound
	}

	err = transaction.Delete()
	if err != nil {
		return false, err.Error(), http.StatusBadRequest
	}

	return true, "", 0
}
