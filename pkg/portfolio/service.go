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
func (service *Service) Create(userID uint, name string) (map[string]interface{}, string, int) {
	portfolio, err := service.Portfolio.GetByName(userID, name)
	if err == nil {
		name = portfolio.Name + strconv.Itoa(rand.Intn(100))
	}

	newPortfolio, err := service.Portfolio.Create(userID, name)
	if err != nil {
		return nil, "asd", http.StatusBadRequest
	}

	resp := newPortfolio.ToMap()
	return resp, "", 0
}

//Get user portfolio
func (service *Service) Get(userID uint, portfolioID uint) (map[string]interface{}, string, int) {
	portfolio, err := service.Portfolio.Get(userID, portfolioID)
	if err != nil {
		return nil, "portfolio not found", http.StatusNotFound
	}

	resp := portfolio.ToMap()
	return resp, "", 0
}