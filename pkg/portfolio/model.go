package portfolio

import (
	"encoding/json"
	"errors"
	"investment-api/utils"
	"time"
)

//Model interface
type Model interface {
	Create(userID uint, name string) (*Portfolio, error)
	GetByName(userID uint, name string) (*Portfolio, error)
	GetAll(userID uint) ([]*Portfolio, error)
	Get(userID uint, ID uint) (*Portfolio, error)
	Update(data map[string]interface{}) error
	Delete() error
}

//Portfolio model
type Portfolio struct {
	ID                    uint       `json:"id"`
	Currency              uint       `json:"currency_id" gorm:"column:currency_id"`
	UserID                uint       `json:"-" gorm:"column:user_id"`
	Name                  string     `json:"name" gorm:"not null"`
	Cost                  *float64   `json:"cost"`
	MarketValue           *float64   `json:"market_value" gorm:"column:market_value"`
	TotalChange           *float64   `json:"total_change" gorm:"column:total_change"`
	TotalChangePercentage *float64   `json:"total_change_percentage" gorm:"column:total_change_percentage"`
	DailyChange           *float64   `json:"daily_change" gorm:"column:daily_change"`
	DailyChangePercentage *float64   `json:"daily_change_percentage" gorm:"column:daily_change_percentage"`
	UnrealisedGainLoss    *float64   `json:"unrealised_gain_loss" gorm:"column:unrealised_gain_loss"`
	RealisedGainLoss      *float64   `json:"realised_gain_loss" gorm:"column:realised_gain_loss"`
	ExpectedDivYield      *float64   `json:"expected_div_yield" gorm:"column:expected_div_yield"`
	ExpectedDiv           *float64   `json:"expected_div" gorm:"column:expected_div"`
	DivCollected          *float64   `json:"div_collected" gorm:"column:div_collected"`
	CreatedAt             time.Time  `json:"-"`
	UpdatedAt             *time.Time `json:"-"`
	DeletedAt             *time.Time `json:"-"`
}

//NewPortfolio instantiate new user model
func NewPortfolio() Model {
	return &Portfolio{}
}

//Create a new user
func (portfolio *Portfolio) Create(userID uint, name string) (*Portfolio, error) {
	portfolio.Name = name
	portfolio.UserID = userID

	err := utils.GetDB().Create(portfolio).GetErrors()
	if len(err) != 0 {
		return nil, errors.New("something wrong happened while creating the portfolio")
	}

	return portfolio, nil
}

//GetByName get user's portfolio by name
func (portfolio *Portfolio) GetByName(userID uint, name string) (*Portfolio, error) {
	err := utils.GetDB().
		Table("portfolios").
		Where("user_id = ?", userID).
		Where("name = ?", name).
		First(portfolio).GetErrors()

	if portfolio.Name == "" && len(err) != 0 {
		return portfolio, errors.New("portfolio not found")
	}

	return portfolio, nil
}

//GetAll user's portfolios
func (portfolio *Portfolio) GetAll(userID uint) ([]*Portfolio, error) {
	var portfolios []*Portfolio

	err := utils.GetDB().Table("portfolios").Where("user_id = ?", userID).Find(&portfolios).GetErrors()
	if len(err) != 0 {
		return portfolios, errors.New("error while fetching user portfolios")
	}

	return portfolios, nil
}

//Get get user's portfolio by ID
func (portfolio *Portfolio) Get(userID uint, ID uint) (*Portfolio, error) {
	err := utils.GetDB().Table("portfolios").Where("user_id = ?", userID).Where("id = ?", ID).First(portfolio).GetErrors()
	if portfolio.Name == "" && len(err) != 0 {
		return portfolio, errors.New("portfolio not found")
	}

	return portfolio, nil
}

//Update portfolio data
func (portfolio *Portfolio) Update(data map[string]interface{}) error {
	errs := utils.GetDB().Model(portfolio).Update(data).GetErrors()
	if len(errs) != 0 {
		return errors.New("something went wrong while updating portfolio")
	}

	return nil
}

//Delete portfolio
func (portfolio *Portfolio) Delete() error {
	err := utils.GetDB().Delete(portfolio)
	if err == nil {
		return errors.New("something went wrong while deleting portfolio")
	}

	return nil
}

//ToMap transformer struct to map
func (portfolio Portfolio) ToMap() map[string]interface{} {
	var data map[string]interface{}
	inrec, _ := json.Marshal(portfolio)
	json.Unmarshal(inrec, &data)
	return data
}
