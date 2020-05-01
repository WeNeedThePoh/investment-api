package models

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserModel interface {
	CreateUser(data map[string]interface{}) (*user, error)
	GetUser(id uint) (*user, error)
	GetByEmail(email string) *user
	UpdateUser(data map[string]interface{}) error
	UpdateUserPassword(newPassword string) error
	DeleteUser() error
}

type user struct {
	ID            uint       `json:"id"`
	Currency      *uint      `json:"currency_id" gorm:"column:currency_id"`
	Country       uint       `json:"country_id" gorm:"column:country_id"`
	Lang          *uint      `json:"lang_id" gorm:"column:lang_id"`
	Email         string     `json:"email" gorm:"size:75;not null"`
	FirstName     string     `json:"first_name" gorm:"size:40;not null"`
	LastName      *string    `json:"last_name" gorm:"size:40"`
	Password      string     `json:"password" gorm:"size:255;not null"`
	PasswordReset *string    `json:"password_reset" gorm:"size:255"`
	Active        bool       `json:"active"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

func NewUser() UserModel {
	return &user{}
}

func (user *user) CreateUser(data map[string]interface{}) (*user, error) {
	mapstructure.Decode(data, &user)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	err := GetDB().Create(user).GetErrors()
	if len(err) != 0 {
		return nil, errors.New("something wrong happened while creating the user")
	}

	return user, nil
}

func (user *user) GetUser(id uint) (*user, error) {
	err := GetDB().Table("users").Where("id = ?", id).First(user).GetErrors()
	if user.Email == "" && len(err) != 0 {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (user *user) GetByEmail(email string) *user {
	err := GetDB().Table("users").Where("email = ?", email).First(user).GetErrors()
	if user.Email == "" && len(err) != 0 {
		return nil
	}

	return user
}

func (user *user) UpdateUser(data map[string]interface{}) error {
	errs := GetDB().Model(user).Update(data).GetErrors()
	if len(errs) != 0 {
		return errors.New("something went wrong while updating user")
	}

	return nil
}

func (user *user) UpdateUserPassword(newPassword string) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	err := GetDB().Save(user).GetErrors()
	if len(err) != 0 {
		return errors.New("something wrong happened while updating user password")
	}

	return nil
}

func (user *user) DeleteUser() error {
	user.Active = false
	GetDB().Save(user)

	err := GetDB().Delete(user)
	if err == nil {
		user.Active = true
		GetDB().Save(user)
		return errors.New("something went wrong while deleting user")
	}

	return nil
}

func (user user) ToMap() map[string]interface{} {
	var data map[string]interface{}
	inrec, _ := json.Marshal(user)
	json.Unmarshal(inrec, &data)
	return data
}
