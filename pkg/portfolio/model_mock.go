package portfolio

//MockPortfolioModel mock
type MockPortfolioModel struct {
	Portfolio    *Portfolio
	errorMessage error
	Portfolios   []*Portfolio
}

//Add mock
func (m MockPortfolioModel) Create(userID uint, name string) (*Portfolio, error) {
	return m.Portfolio, m.errorMessage
}

//GetAll mock
func (m MockPortfolioModel) GetAll(userID uint) ([]*Portfolio, error) {
	return m.Portfolios, m.errorMessage
}

//Get mock
func (m MockPortfolioModel) Get(userID uint, ID uint) (*Portfolio, error) {
	return m.Portfolio, m.errorMessage
}

//GetByName mock
func (m MockPortfolioModel) GetByName(userID uint, name string) (*Portfolio, error) {
	return m.Portfolio, m.errorMessage
}

//Update mock
func (m MockPortfolioModel) Update(data map[string]interface{}) error {
	return m.errorMessage
}

//Delete mock
func (m MockPortfolioModel) Delete() error {
	return m.errorMessage
}
