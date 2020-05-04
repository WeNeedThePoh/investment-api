package portfolio

import (
	portfolio "investment-api/pkg/portfolio/mock"
	"testing"
)

func TestPortfolioNotFoundGet(t *testing.T) {
	var model = portfolio.MockPortfolioModel{Portfolio: nil}
	service := NewPortfolioService(model)

	_, _, code := service.Get(1, 2)
	if code != 404 {
		t.Errorf("Get(1, 2) = %d; want 404", code)
	}
}

func TestGet(t *testing.T) {
	var model = portfolio.MockPortfolioModel{Portfolio: &Portfolio{Name: "portfolio name", UserID: 1, ID: 1}}
	service := NewPortfolioService(model)

	_, _, code := service.Get(1, 1)
	if code != 0 {
		t.Errorf("Get(1, 1) = %d; want 200", code)
	}
}
