package stock

import (
	"encoding/json"
	"errors"
	"investment-api/utils"
	"time"
)

//Model interface
type Model interface {
	Create(symbol string, price float64, company string, country uint) (*Stock, error)
	GetBySymbol(symbol string) (*Stock, error)
	Get(stockID uint) (*Stock, error)
	Update(data map[string]interface{}) error
	Delete() error
}

//Stock model
type Stock struct {
	ID                    uint       `json:"id"`
	Country               uint       `json:"country_id" gorm:"column:country_id"`
	Symbol                string     `json:"symbol"`
	CompanyName           string     `json:"company_name"`
	Price                 float64   `json:"price;default:0"`
	MinPrice              float64   `json:"min_price" gorm:"column:min_price;default:0"`
	MaxPrice              float64   `json:"max_price" gorm:"column:max_price;default:0"`
	DailyChange           float64   `json:"daily_change" gorm:"column:daily_change;default:0"`
	DailyChangePercentage float64   `json:"daily_change_percentage" gorm:"column:daily_change_percentage;default:0"`
	YearChange            float64   `json:"year_change" gorm:"column:year_change;default:0"`
	YearChangePercentage  float64   `json:"year_change_percentage" gorm:"column:year_change_percentage;default:0"`
	DivYield              float64   `json:"div_yield" gorm:"column:div_yield;default:0"`
	DivShare              float64   `json:"div_share" gorm:"column:div_share;default:0"`
	Eps                   float64   `json:"eps" gorm:"column:eps;default:0"`
	UpdatedAt             *time.Time `json:"updated_at"`
}

//NewStock instantiate new stock model
func NewStock() Model {
	return &Stock{}
}

//Create a new stock
func (stock *Stock) Create(symbol string, price float64, company string, country uint) (*Stock, error) {
	stock.Symbol = symbol
	stock.Price = price
	stock.CompanyName = company
	stock.Country = country

	err := utils.GetDB().Create(stock).GetErrors()
	if len(err) != 0 {
		return nil, errors.New("something wrong happened while creating the stock")
	}

	return stock, nil
}

//GetBySymbol get stock by symbol
func (stock *Stock) GetBySymbol(symbol string) (*Stock, error) {
	err := utils.GetDB().
		Table("stocks").
		Where("symbol LIKE ?", symbol).
		First(stock).GetErrors()

	if stock.Symbol == "" && len(err) != 0 {
		return stock, errors.New("stock not found")
	}

	return stock, nil
}

//Get get stock by ID
func (stock *Stock) Get(stockID uint) (*Stock, error) {
	err := utils.GetDB().Table("stocks").Where("id = ?", stockID).First(stock).GetErrors()
	if stock.Symbol == "" && len(err) != 0 {
		return stock, errors.New("stock not found")
	}

	return stock, nil
}

//Update stock data
func (stock *Stock) Update(data map[string]interface{}) error {
	errs := utils.GetDB().Model(stock).Update(data).GetErrors()
	if len(errs) != 0 {
		return errors.New("something went wrong while updating stock")
	}

	return nil
}

//Delete stock
func (stock *Stock) Delete() error {
	err := utils.GetDB().Delete(stock)
	if err == nil {
		return errors.New("something went wrong while deleting stock")
	}

	return nil
}

//ToMap transformer struct to map
func (stock Stock) ToMap() map[string]interface{} {
	var data map[string]interface{}
	inrec, _ := json.Marshal(stock)
	json.Unmarshal(inrec, &data)
	return data
}
