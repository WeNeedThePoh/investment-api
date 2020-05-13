package transaction

import (
	"errors"
	"investment-api/utils"
	"time"
)

//Model interface
type Model interface {
	Create(PortfolioID uint, StockID uint, Type string, Shares float64, CostPerShare float64, Fees float64) (*Transaction, error)
	GetAll(portfolioID uint) ([]*Transaction, error)
	Get(transactionID uint) (*Transaction, error)
	Update(data map[string]interface{}) error
	Delete() error
}

//Transaction model
type Transaction struct {
	ID           uint       `json:"id"`
	PortfolioID  uint       `json:"portfolio_id" gorm:"column:portfolio_id"`
	StockID      uint       `json:"stock_id" gorm:"column:stock_id"`
	Type         string     `json:"type" gorm:"not null"`
	Shares       float64    `json:"shares"`
	Amount       float64    `json:"amount"`
	CostPerShare float64    `json:"cost_per_share" gorm:"column:cost_per_share"`
	Fees         float64    `json:"fees" gorm:"default:0"`
	CreatedAt    time.Time  `json:"-"`
	UpdatedAt    *time.Time `json:"-"`
	DeletedAt    *time.Time `json:"-"`
}

//TableName for GORM
func (Transaction) TableName() string {
	return "portfolio_transactions"
}

//NewTransaction instantiate new transaction model
func NewTransaction() Model {
	return &Transaction{}
}

//Create a new transaction
func (transaction *Transaction) Create(PortfolioID uint, StockID uint, Type string, Shares float64, CostPerShare float64, Fees float64) (*Transaction, error) {
	transaction.PortfolioID = PortfolioID
	transaction.StockID = StockID
	transaction.Type = Type
	transaction.Shares = Shares
	transaction.CostPerShare = CostPerShare
	transaction.Fees = Fees
	transaction.Amount = CostPerShare * Shares

	err := utils.GetDB().Create(transaction).GetErrors()
	if len(err) != 0 {
		return nil, errors.New("something wrong happened while creating the transaction")
	}

	return transaction, nil
}

//GetAll transactions
func (transaction *Transaction) GetAll(portfolioID uint) ([]*Transaction, error) {
	var transactions []*Transaction

	err := utils.GetDB().Table("portfolio_transactions").Where("portfolio_id = ?", portfolioID).Find(&transactions).GetErrors()
	if len(err) != 0 {
		return transactions, errors.New("error while fetching transactions")
	}

	return transactions, nil
}

//Get transaction by ID
func (transaction *Transaction) Get(transactionID uint) (*Transaction, error) {
	err := utils.GetDB().Table("portfolio_transactions").Where("id = ?", transactionID).First(transaction).GetErrors()
	if len(err) != 0 {
		return transaction, errors.New("transaction not found")
	}

	return transaction, nil
}

//Update transaction data
func (transaction *Transaction) Update(data map[string]interface{}) error {
	errs := utils.GetDB().Model(transaction).Update(data).GetErrors()
	if len(errs) != 0 {
		return errors.New("something went wrong while updating transaction")
	}

	return nil
}

//Delete transaction
func (transaction *Transaction) Delete() error {
	err := utils.GetDB().Delete(transaction)
	if err == nil {
		return errors.New("something went wrong while deleting transaction")
	}

	return nil
}
