package user

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UsersService struct {
	User Model
}

func NewUserService(model Model) *UsersService {
	return &UsersService{User: model}
}

func (service *UsersService) Create(data map[string]interface{}) (map[string]interface{}, string, int) {
	user := service.User.GetByEmail(data["email"].(string))
	if user != nil {
		return nil, "email already in use", http.StatusNotFound
	}

	newUser, err := service.User.Create(data)
	if err != nil {
		return nil, "asd", http.StatusBadRequest
	}

	resp := newUser.ToMap()
	delete(resp, "password")
	delete(resp, "password_reset")
	return resp, "", 0
}

func (service *UsersService) Get(id uint) (map[string]interface{}, string, int) {
	user, err := service.User.Get(id)
	if err != nil {
		return nil, "user not found", http.StatusNotFound
	}

	resp := user.ToMap()
	delete(resp, "password")
	delete(resp, "password_reset")
	return resp, "", 0
}

func (service *UsersService) Update(id uint, data map[string]interface{}) (bool, string, int) {
	user, err := service.User.Get(id)
	if err != nil {
		return false, "user not found", http.StatusNotFound
	}

	delete(data, "password")
	delete(data, "password_reset")

	err = user.Update(data)
	if err != nil {
		return false, "asdasd", http.StatusBadRequest
	}

	return true, "", 0
}

func (service *UsersService) UpdatePassword(id uint, oldPassword string, password string) (bool, string, int) {
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
		return false, "asdasd", http.StatusBadRequest
	}

	return true, "", 0
}

func (service *UsersService) Delete(id uint) (bool, string, int) {
	user, err := service.User.Get(id)
	if err != nil {
		return false, "user not found", http.StatusNotFound
	}

	err = user.Delete()
	if err != nil {
		return false, "asdasd", http.StatusBadRequest
	}

	return true, "", 0
}
