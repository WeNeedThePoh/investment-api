package user

//MockUserModel mock
type MockUserModel struct {
	User         *User
	ErrorMessage error
}

//Create mock
func (m MockUserModel) Create(data map[string]interface{}) (*User, error) {
	return m.User, m.ErrorMessage
}

//Get mock
func (m MockUserModel) Get(id uint) (*User, error) {
	return m.User, m.ErrorMessage
}

//GetByEmail mock
func (m MockUserModel) GetByEmail(email string) (*User, error) {
	return m.User, m.ErrorMessage
}

//Update mock
func (m MockUserModel) Update(data map[string]interface{}) error {
	return m.ErrorMessage
}

//UpdatePassword mock
func (m MockUserModel) UpdatePassword(newPassword string) error {
	return m.ErrorMessage
}

//Delete mock
func (m MockUserModel) Delete() error {
	return m.ErrorMessage
}
