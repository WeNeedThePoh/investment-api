package portfolio

import (
	"math/rand"
	"net/http"
	"strconv"
)

//Service dependencies
type Service struct {
	Portfolio Model
}

//NewPortfolioService service construct
func NewPortfolioService(model Model) *Service {
	return &Service{Portfolio: model}
}

//Create new portfolio
func (service *Service) Create(userID uint, name string) (interface{}, string, int) {
	portfolio, err := service.Portfolio.GetByName(userID, name)
	if err == nil {
		name = portfolio.Name + strconv.Itoa(rand.Intn(100))
	}

	newPortfolio, err := service.Portfolio.Create(userID, name)
	if err != nil {
		return nil, err.Error(), http.StatusBadRequest
	}

	return newPortfolio, "", 0
}

//GetAll user portfolios
func (service *Service) GetAll(userID uint) ([]*Portfolio, string, int) {
	portfolios, err := service.Portfolio.GetAll(userID)
	if err != nil {
		return nil, err.Error(), http.StatusNotFound
	}

	return portfolios, "", 0
}

//Get user portfolio
func (service *Service) Get(userID uint, portfolioID uint) (interface{}, string, int) {
	portfolio, err := service.Portfolio.Get(userID, portfolioID)
	if err != nil {
		return nil, err.Error(), http.StatusNotFound
	}

	return portfolio, "", 0
}

//Update portfolio
func (service *Service) Update(userID uint, id uint, data map[string]interface{}) (bool, string, int) {
	portfolio, err := service.Portfolio.Get(userID, id)
	if err != nil {
		return false, err.Error(), http.StatusNotFound
	}

	err = portfolio.Update(data)
	if err != nil {
		return false, err.Error(), http.StatusBadRequest
	}

	return true, "", 0
}

//Delete portfolio
func (service *Service) Delete(userID uint, portfolioID uint) (bool, string, int) {
	portfolio, err := service.Portfolio.Get(userID, portfolioID)
	if err != nil {
		return false, err.Error(), http.StatusNotFound
	}

	err = portfolio.Delete()
	if err != nil {
		return false, err.Error(), http.StatusBadRequest
	}

	return true, "", 0
}
