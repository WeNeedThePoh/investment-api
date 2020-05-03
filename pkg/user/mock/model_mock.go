package user

import "investment-api/pkg/user"

//MockUserModel mock
type MockUserModel struct {
	User         *user.User
	errorMessage error
}

//Create mock
func (m MockUserModel) Create(data map[string]interface{}) (*user.User, error) {
	return m.User, m.errorMessage
}

//Get mock
func (m MockUserModel) Get(id uint) (*user.User, error) {
	return m.User, m.errorMessage
}

//GetByEmail mock
func (m MockUserModel) GetByEmail(email string) *user.User {
	return m.User
}

//Update mock
func (m MockUserModel) Update(data map[string]interface{}) error {
	return m.errorMessage
}

//UpdatePassword mock
func (m MockUserModel) UpdatePassword(newPassword string) error {
	return m.errorMessage
}

//Delete mock
func (m MockUserModel) Delete() error {
	return m.errorMessage
}
