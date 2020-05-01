package services

import (
	"golang.org/x/crypto/bcrypt"
	"investment-api/models"
	"net/http"
)

type UsersService struct {
	UserModel models.UserModel
}

func NewUserService(model models.UserModel) *UsersService {
	return &UsersService{UserModel: model}
}

func (service *UsersService) CreateUser(data map[string]interface{}) (map[string]interface{}, string, int) {
	user := service.UserModel.GetByEmail(data["email"].(string))
	if user != nil {
		return nil, "email already in use", http.StatusNotFound
	}

	newUser, err := service.UserModel.CreateUser(data)
	if err != nil {
		return nil, "asd", http.StatusBadRequest
	}

	resp := newUser.ToMap()
	delete(resp, "password")
	delete(resp, "password_reset")
	return resp, "", 0
}

func (service *UsersService) GetUser(id uint) (map[string]interface{}, string, int) {
	user, err := service.UserModel.GetUser(id)
	if err != nil {
		return nil, "user not found", http.StatusNotFound
	}

	resp := user.ToMap()
	delete(resp, "password")
	delete(resp, "password_reset")
	return resp, "", 0
}

func (service *UsersService) UpdateUser(id uint, data map[string]interface{}) (bool, string, int) {
	user, err := service.UserModel.GetUser(id)
	if err != nil {
		return false, "user not found", http.StatusNotFound
	}

	delete(data, "password")
	delete(data, "password_reset")

	err = user.UpdateUser(data)
	if err != nil {
		return false, "asdasd", http.StatusBadRequest
	}

	return true, "", 0
}

func (service *UsersService) UpdateUserPassword(id uint, oldPassword string, password string) (bool, string, int) {
	user, err := service.UserModel.GetUser(id)
	if err != nil {
		return false, "user not found", http.StatusNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return false, "Invalid credentials", http.StatusUnauthorized
	}

	err = user.UpdateUserPassword(password)
	if err != nil {
		return false, "asdasd", http.StatusBadRequest
	}

	return true, "", 0
}

func (service *UsersService) DeleteUser(id uint) (bool, string, int) {
	user, err := service.UserModel.GetUser(id)
	if err != nil {
		return false, "user not found", http.StatusNotFound
	}

	err = user.DeleteUser()
	if err != nil {
		return false, "asdasd", http.StatusBadRequest
	}

	return true, "", 0
}
