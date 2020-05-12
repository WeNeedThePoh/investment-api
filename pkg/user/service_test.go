package user

import (
	"errors"
	"testing"
)

func TestGetOnNotFound(t *testing.T) {
	var model = MockUserModel{User: &User{}, ErrorMessage: errors.New("user not found")}
	service := NewUserService(model)

	_, _, code := service.Get(1)
	if code != 404 {
		t.Errorf("Get(1) = %d; want 404", code)
	}
}

func TestGet(t *testing.T) {
	var model = MockUserModel{User: &User{Email: "test@test.co,", ID: 1}}
	service := NewUserService(model)

	_, _, code := service.Get(1)
	if code != 0 {
		t.Errorf("Get(1) = %d; want 200", code)
	}
}

func TestUpdateOnNotFound(t *testing.T) {
	var model = MockUserModel{User: &User{}, ErrorMessage: errors.New("user not found")}
	service := NewUserService(model)

	_, _, code := service.Update(1, map[string]interface{}{"first_name": "new name"})
	if code != 404 {
		t.Errorf("Get(1, map[string]interface{}{\"first_name\": \"new name\"}) = %d; want 404", code)
	}
}

func TestUpdate(t *testing.T) {
	var model = MockUserModel{User: &User{Email: "test@test.co,", ID: 1}}
	service := NewUserService(model)

	_, _, code := service.Update(1, map[string]interface{}{"first_name": "new name"})
	if code != 0 {
		t.Errorf("Get(1, map[string]interface{}{\"first_name\": \"new name\"}) = %d; want 200", code)
	}
}

func TestDeleteOnNotFound(t *testing.T) {
	var model = MockUserModel{User: &User{}, ErrorMessage: errors.New("user not found")}
	service := NewUserService(model)

	_, _, code := service.Delete(1)
	if code != 404 {
		t.Errorf("Get(1) = %d; want 404", code)
	}
}

func TestDelete(t *testing.T) {
	var model = MockUserModel{User: &User{Email: "test@test.co,", ID: 1}}
	service := NewUserService(model)

	_, _, code := service.Delete(1)
	if code != 0 {
		t.Errorf("Get(1) = %d; want 0", code)
	}
}
