package models

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
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

func Login(email, password string) (map[string]interface{}, string, int) {
	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "User not found", 404
		}

		return nil, "Connection error. Please retry", 400
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, "Invalid login credentials", 401
	}

	tk := &Token{UserId: user.ID}

	exp, err := strconv.ParseInt(os.Getenv("JWT_EXPIRE_TIME"), 10, 64)
	if err != nil {
		return nil, "Couldn't parse jwt expire in", 400
	}

	tk.ExpiresAt = time.Now().Unix() + exp
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return map[string]interface{}{"token": tokenString, "exp_in": tk.ExpiresAt}, "", 200
}

func CreateUser(data map[string]interface{}) (*User, string, int) {
	user := &User{}
	err := GetDB().Table("users").Where("email = ?", data["email"]).First(user).GetErrors()
	if user.Email != "" && len(err) == 0 {
		return nil, "Email alreay in use", 400
	}

	mapstructure.Decode(data, &user)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	err = GetDB().Create(user).GetErrors()
	if len(err) != 0 {
		return nil, "Something wrong happened while creating the user", 400
	}

	return user, "", 200
}

func GetUser(id uint) *User {
	user := &User{}
	GetDB().Table("users").Where("id = ?", id).First(user)
	if user.Email == "" {
		return nil
	}

	return user
}

func UpdateUser(id uint, data map[string]interface{}) (bool, string, int) {
	user := GetUser(id)
	if user == nil {
		return false, "Not found", 404
	}

	delete(data, "password")
	delete(data, "password_reset")

	errors := GetDB().Model(user).Update(data).GetErrors()
	if len(errors) != 0 {
		return false, "Something wrong happened while updating user", 400
	}

	return true, "", 204
}

func UpdateUserPassword(id uint, oldPassword string, newPassword string) (bool, string, int) {
	user := GetUser(id)
	if user == nil {
		return false, "User not found", 404
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return false, "Invalid credentials", 401
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	errors := GetDB().Save(user).GetErrors()
	if len(errors) != 0 {
		return false, "Something wrong happened while updating user password", 400
	}

	return true, "", 204
}

func DeleteUser(id uint) (bool, string, int) {
	user := &User{}
	GetDB().Table("users").Where("id = ?", id).First(user)
	if user.Email == "" {
		return false, "Not Found", 404
	}

	user.Active = false
	GetDB().Save(user)

	err := GetDB().Delete(user)
	if err == nil {
		user.Active = true
		GetDB().Save(user)
		return false, "Error Deleting user", 400
	}

	return true, "", 200
}

func (user User) ToMap() map[string]interface{} {
	var data map[string]interface{}
	inrec, _ := json.Marshal(user)
	json.Unmarshal(inrec, &data)
	return data
}
