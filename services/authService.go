package services

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"investment-api/models"
	"net/http"
	"os"
	"strconv"
	"time"
)

type AuthService struct {
	UserModel models.UserModel
}

type Token struct {
	UserId uint
	jwt.StandardClaims
}

func NewAuthService(model models.UserModel) *AuthService {
	return &AuthService{UserModel: model}
}

func (service *AuthService) Login (email string, password string) (map[string]interface{}, string, int){
	user := service.UserModel.GetByEmail(email)
	if user == nil {
		return nil, "user not found", http.StatusNotFound
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, "Invalid login credentials", http.StatusUnauthorized
	}

	tk := &Token{UserId: user.ID}

	exp, err := strconv.ParseInt(os.Getenv("JWT_EXPIRE_TIME"), 10, 64)
	if err != nil {
		return nil, "Couldn't parse jwt expire in", http.StatusBadRequest
	}

	tk.ExpiresAt = time.Now().Unix() + exp
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return map[string]interface{}{"token": tokenString, "exp_in": tk.ExpiresAt}, "", 0
}
