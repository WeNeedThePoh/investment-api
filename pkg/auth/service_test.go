package auth

import (
	"investment-api/pkg/user"
	mock "investment-api/pkg/user/mock"
	"os"
	"testing"
)

func TestUserNotFoundLogin(t *testing.T) {
	var model = mock.MockUserModel{User: nil}
	service := NewAuthService(model)

	_, _, code := service.Login("test@email.com", "password")
	if code != 404 {
		t.Errorf("Login(\"test@email.com\", \"password\") = %d; want 404", code)
	}
}

func TestInvalidCredentialsLogin(t *testing.T) {
	mockedUser := user.User{Email: "test@email.com", Password: "$2a$10$AV7uo698BQZ8NdCa9m3v0uRLm79mRthXrq3owP7AbSyjZphXZjDjy"}
	var model = mock.MockUserModel{User: &mockedUser}
	service := NewAuthService(model)

	_, _, code := service.Login("test@email.com", "wrong_password")
	if code != 401 {
		t.Errorf("Login(\"test@email.com\", \"password\") = %d; want 401", code)
	}
}

func TestLogin(t *testing.T) {
	os.Setenv("JWT_EXPIRE_TIME", "6000")
	os.Setenv("JWT_SECRET", "secret")

	mockedUser := user.User{Email: "test@email.com", Password: "$2a$10$AV7uo698BQZ8NdCa9m3v0uRLm79mRthXrq3owP7AbSyjZphXZjDjy"}
	var model = mock.MockUserModel{User: &mockedUser}
	service := NewAuthService(model)

	_, _, code := service.Login("test@email.com", "asdasdasd")
	if code != 0 {
		t.Errorf("Login(\"test@email.com\", \"password\") = %d; want 200", code)
	}
}
