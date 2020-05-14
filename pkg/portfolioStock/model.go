package transaction

import (
	"errors"
	"investment-api/utils"
	"time"
)

//Model interface
type Model interface {
	Create(PortfolioID uint, StockID uint, Type string, Shares float64, CostPerShare float64) (*PortfolioStock, error)
	GetAll(portfolioID uint) ([]*PortfolioStock, error)
	Get(portfolioStockID uint) (*PortfolioStock, error)
	Update(Shares float64, AvgShareCost float64, Cost float64, Type string) error
	Delete() error
}

//PortfolioStock model
type PortfolioStock struct {
	ID                    uint       `json:"id"`
	PortfolioID           uint       `json:"portfolio_id" gorm:"column:portfolio_id"`
	StockID               uint       `json:"stock_id" gorm:"column:stock_id"`
	Type                  string     `json:"type" gorm:"not null"`
	Shares                float64    `json:"shares"`
	AvgShareCost          float64    `json:"avg_share_cost"`
	Cost                  float64    `json:"cost"`
	MarketValue           float64    `json:"market_value" gorm:"column:market_value;default:0"`
	TotalChange           float64    `json:"total_change" gorm:"column:total_change;default:0"`
	TotalChangePercentage float64    `json:"total_change_percentage" gorm:"column:total_change_percentage;default:0"`
	DailyChange           float64    `json:"daily_change" gorm:"column:daily_change;default:0"`
	DailyChangePercentage float64    `json:"daily_change_percentage" gorm:"column:daily_change_percentage;default:0"`
	UnrealisedGainLoss    float64    `json:"unrealised_gain_loss" gorm:"column:unrealised_gain_loss;default:0"`
	RealisedGainLoss      float64    `json:"realised_gain_loss" gorm:"column:realised_gain_loss;default:0"`
	ExpectedDivYield      float64    `json:"expected_div_yield" gorm:"column:expected_div_yield;default:0"`
	ExpectedDiv           float64    `json:"expected_div" gorm:"column:expected_div;default:0"`
	DivCollected          float64    `json:"div_collected" gorm:"column:div_collected;default:0"`
	CreatedAt             time.Time  `json:"-"`
	UpdatedAt             *time.Time `json:"-"`
	DeletedAt             *time.Time `json:"-"`
}

//TableName for GORM
func (PortfolioStock) TableName() string {
	return "portfolio_stocks"
}

//NewPortfolioStock instantiate new transaction model
func NewPortfolioStock() Model {
	return &PortfolioStock{}
}

//Create a new portfolioStock
func (portfolioStock *PortfolioStock) Create(PortfolioID uint, StockID uint, Type string, Shares float64, CostPerShare float64) (*PortfolioStock, error) {
	portfolioStock.PortfolioID = PortfolioID
	portfolioStock.StockID = StockID
	portfolioStock.Type = Type
	portfolioStock.Shares = Shares
	portfolioStock.AvgShareCost = CostPerShare
	portfolioStock.Cost = CostPerShare * Shares

	err := utils.GetDB().Create(portfolioStock).GetErrors()
	if len(err) != 0 {
		return nil, errors.New("something wrong happened while adding stock to portfolio")
	}

	return portfolioStock, nil
}

//GetAll portfolio stocks
func (portfolioStock *PortfolioStock) GetAll(portfolioID uint) ([]*PortfolioStock, error) {
	var stocks []*PortfolioStock

	err := utils.GetDB().Where("portfolio_id = ?", portfolioID).Find(&stocks).GetErrors()
	if len(err) != 0 {
		return stocks, errors.New("error while fetching portfolio stocks")
	}

	return stocks, nil
}

//Get portfolio stock
func (portfolioStock *PortfolioStock) Get(portfolioStockID uint) (*PortfolioStock, error) {
	err := utils.GetDB().Where("id = ?", portfolioStockID).First(portfolioStock).GetErrors()
	if len(err) != 0 {
		return portfolioStock, errors.New("portfolio stock not found")
	}

	return portfolioStock, nil
}

//Update portfolio stock
func (portfolioStock *PortfolioStock) Update(Shares float64, AvgShareCost float64, Cost float64, Type string) error {
	portfolioStock.Shares += Shares
	portfolioStock.AvgShareCost = AvgShareCost
	portfolioStock.Cost = Cost

	if Type != "" {
		portfolioStock.Type = Type
	}

	errs := utils.GetDB().Save(portfolioStock).GetErrors()
	if len(errs) != 0 {
		return errors.New("something went wrong while updating portfolio stock")
	}

	return nil
}

//Delete portfolio stock
func (portfolioStock *PortfolioStock) Delete() error {
	err := utils.GetDB().Delete(portfolioStock)
	if err == nil {
		return errors.New("something went wrong while deleting portfolio stock")
	}

	return nil
}
