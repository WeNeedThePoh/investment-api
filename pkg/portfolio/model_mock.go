package portfolio

//MockPortfolioModel mock
type MockPortfolioModel struct {
	Portfolio *Portfolio
	errorMessage error
}

//Create mock
func (m MockPortfolioModel) Create(userID uint, name string) (*Portfolio, error) {
	return m.Portfolio, m.errorMessage
}

//Get mock
func (m MockPortfolioModel) Get(userID uint, ID uint) (*Portfolio, error) {
	return m.Portfolio, m.errorMessage
}

//GetByName mock
func (m MockPortfolioModel) GetByName(userID uint, name string) (*Portfolio, error) {
	return m.Portfolio, m.errorMessage
}
