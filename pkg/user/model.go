package user

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"investment-api/utils"
	"time"
)

//Model interface
type Model interface {
	Create(data map[string]interface{}) (*User, error)
	Get(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(data map[string]interface{}) error
	UpdatePassword(newPassword string) error
	Delete() error
}

//User model
type User struct {
	ID            uint       `json:"id"`
	Currency      *uint      `json:"currency_id" gorm:"column:currency_id"`
	Country       uint       `json:"country_id" gorm:"column:country_id"`
	Lang          *uint      `json:"lang_id" gorm:"column:lang_id"`
	Email         string     `json:"email" gorm:"size:75;not null"`
	FirstName     string     `json:"first_name" gorm:"size:40;not null"`
	LastName      *string    `json:"last_name" gorm:"size:40"`
	Password      string     `json:"-" gorm:"size:255;not null"`
	PasswordReset *string    `json:"-" gorm:"size:255"`
	Active        bool       `json:"active"`
	CreatedAt     time.Time  `json:"-"`
	UpdatedAt     *time.Time `json:"-"`
	DeletedAt     *time.Time `json:"-"`
}

//NewUser instantiate new user model
func NewUser() Model {
	return &User{}
}

//Create a new user
func (user *User) Create(data map[string]interface{}) (*User, error) {
	mapstructure.Decode(data, &user)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	err := utils.GetDB().Create(user).GetErrors()
	if len(err) != 0 {
		return nil, errors.New("something wrong happened while creating the user")
	}

	return user, nil
}

//Get user by ID
func (user *User) Get(id uint) (*User, error) {
	err := utils.GetDB().Table("users").Where("id = ?", id).First(user).GetErrors()
	if user.Email == "" && len(err) != 0 {
		return user, errors.New("user not found")
	}

	return user, nil
}

//GetByEmail get user by email
func (user *User) GetByEmail(email string) (*User, error) {
	err := utils.GetDB().Table("users").Where("email = ?", email).First(user).GetErrors()
	if user.Email == "" && len(err) != 0 {
		return user, errors.New("user not found")
	}

	return user, nil
}

//Update user data
func (user *User) Update(data map[string]interface{}) error {
	errs := utils.GetDB().Model(user).Update(data).GetErrors()
	if len(errs) != 0 {
		return errors.New("something went wrong while updating user")
	}

	return nil
}

//UpdatePassword updates user password
func (user *User) UpdatePassword(newPassword string) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	err := utils.GetDB().Save(user).GetErrors()
	if len(err) != 0 {
		return errors.New("something wrong happened while updating user password")
	}

	return nil
}

//Delete user
func (user *User) Delete() error {
	user.Active = false
	utils.GetDB().Save(user)

	err := utils.GetDB().Delete(user)
	if err == nil {
		user.Active = true
		utils.GetDB().Save(user)
		return errors.New("something went wrong while deleting user")
	}

	return nil
}

//ToMap transformer struct to map
func (user User) ToMap() map[string]interface{} {
	var data map[string]interface{}
	inrec, _ := json.Marshal(user)
	json.Unmarshal(inrec, &data)
	return data
}
