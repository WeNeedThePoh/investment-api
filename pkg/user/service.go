package user

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

//Service dependencies
type Service struct {
	User Model
}

//NewUserService service construct
func NewUserService(model Model) *Service {
	return &Service{User: model}
}

//Create new user
func (service *Service) Create(data map[string]interface{}) (map[string]interface{}, string, int) {
	user := service.User.GetByEmail(data["email"].(string))
	if user != nil {
		return nil, "email already in use", http.StatusNotFound
	}

	newUser, err := service.User.Create(data)
	if err != nil {
		return nil, err.Error(), http.StatusBadRequest
	}

	resp := newUser.ToMap()
	delete(resp, "password")
	delete(resp, "password_reset")
	return resp, "", 0
}

//Get user
func (service *Service) Get(id uint) (map[string]interface{}, string, int) {
	user, err := service.User.Get(id)
	if err != nil {
		return nil, "user not found", http.StatusNotFound
	}

	resp := user.ToMap()
	delete(resp, "password")
	delete(resp, "password_reset")
	return resp, "", 0
}

//Update user
func (service *Service) Update(id uint, data map[string]interface{}) (bool, string, int) {
	user, err := service.User.Get(id)
	if err != nil {
		return false, "user not found", http.StatusNotFound
	}

	delete(data, "password")
	delete(data, "password_reset")

	err = user.Update(data)
	if err != nil {
		return false, err.Error(), http.StatusBadRequest
	}

	return true, "", 0
}

//UpdatePassword update user password
func (service *Service) UpdatePassword(id uint, oldPassword string, password string) (bool, string, int) {
	user, err := service.User.Get(id)
	if err != nil {
		return false, "user not found", http.StatusNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return false, "Invalid credentials", http.StatusUnauthorized
	}

	err = user.UpdatePassword(password)
	if err != nil {
		return false, err.Error(), http.StatusBadRequest
	}

	return true, "", 0
}

//Delete user
func (service *Service) Delete(id uint) (bool, string, int) {
	user, err := service.User.Get(id)
	if err != nil {
		return false, "user not found", http.StatusNotFound
	}

	err = user.Delete()
	if err != nil {
		return false, err.Error(), http.StatusBadRequest
	}

	return true, "", 0
}
