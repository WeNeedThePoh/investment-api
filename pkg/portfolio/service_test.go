package portfolio

import (
	"errors"
	"testing"
)

func TestGetOnNotFound(t *testing.T) {
	var model = MockPortfolioModel{Portfolio: &Portfolio{}, errorMessage: errors.New("portfolio not found")}
	service := NewPortfolioService(model)

	_, _, code := service.Get(1, 2)
	if code != 404 {
		t.Errorf("Get(1, 2) = %d; want 404", code)
	}
}

func TestGet(t *testing.T) {
	var model = MockPortfolioModel{Portfolio: &Portfolio{Name: "portfolio name", UserID: 1, ID: 1}}
	service := NewPortfolioService(model)

	_, _, code := service.Get(1, 1)
	if code != 0 {
		t.Errorf("Get(1, 1) = %d; want 200", code)
	}
}

func TestUpdateOnNotFound(t *testing.T) {
	var model = MockPortfolioModel{Portfolio: &Portfolio{}, errorMessage: errors.New("portfolio not found")}
	service := NewPortfolioService(model)

	_, _, code := service.Update(1, 1, map[string]interface{}{"name": "new name"})
	if code != 404 {
		t.Errorf("Get(1, 1, map[string]interface{}{\"name\": \"new name\"}) = %d; want 404", code)
	}
}

func TestUpdate(t *testing.T) {
	var model = MockPortfolioModel{Portfolio: &Portfolio{Name: "portfolio name", UserID: 1, ID: 1}}
	service := NewPortfolioService(model)

	_, _, code := service.Update(1, 1, map[string]interface{}{"name": "new name"})
	if code != 0 {
		t.Errorf("Get(1, 1, map[string]interface{}{\"name\": \"new name\"}) = %d; want 200", code)
	}
}

func TestDeleteOnNotFound(t *testing.T) {
	var model = MockPortfolioModel{Portfolio: &Portfolio{}, errorMessage: errors.New("portfolio not found")}
	service := NewPortfolioService(model)

	_, _, code := service.Delete(1, 1)
	if code != 404 {
		t.Errorf("Get(1, 1) = %d; want 404", code)
	}
}

func TestDelete(t *testing.T) {
	var model = MockPortfolioModel{Portfolio: &Portfolio{Name: "portfolio name", UserID: 1, ID: 1}}
	service := NewPortfolioService(model)

	_, _, code := service.Delete(1, 1)
	if code != 0 {
		t.Errorf("Get(1, 1) = %d; want 0", code)
	}
}
