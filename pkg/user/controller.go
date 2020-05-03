package user

import (
	"encoding/json"
	u "investment-api/utils"
	"net/http"
)

//CreateUser Create new user
var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		u.Fail(w, "Error while decoding request body", "", http.StatusBadRequest)
		return
	}

	var model = NewUser()
	service := NewUserService(model)
	user, message, code := service.Create(data)

	if user == nil {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, user, http.StatusCreated)
	}
}

//GetUser get user
var GetUser = func(w http.ResponseWriter, r *http.Request) {
	id := u.RetrieveIDParameter(r)
	var model = NewUser()
	service := NewUserService(model)

	user, message, code := service.Get(id)
	if user == nil {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, user, http.StatusOK)
	}
}

//UpdateUser update user
var UpdateUser = func(w http.ResponseWriter, r *http.Request) {
	id := u.RetrieveIDParameter(r)
	data := make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		u.Fail(w, "Missing required data", "The body payload can not be empty", http.StatusBadRequest)
		return
	}

	var model = NewUser()
	service := NewUserService(model)
	updated, message, code := service.Update(id, data)

	if updated == false {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, nil, http.StatusNoContent)
	}
}

//UpdateUserPassword update user password
var UpdateUserPassword = func(w http.ResponseWriter, r *http.Request) {
	id := u.RetrieveIDParameter(r)
	data := make(map[string]string)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil || data["password"] == "" || data["old_password"] == "" || len(data) == 0 {
		u.Fail(w, "Missing required data", "The body payload can not be empty", http.StatusBadRequest)
		return
	}

	var model = NewUser()
	service := NewUserService(model)
	updated, message, code := service.UpdatePassword(id, data["old_password"], data["password"])

	if updated == false {
		u.Fail(w, message, "", code)
	} else {
		u.Respond(w, nil, http.StatusNoContent)
	}
}

//DeleteUser delete user
var DeleteUser = func(w http.ResponseWriter, r *http.Request) {
	id := u.RetrieveIDParameter(r)
	var model = NewUser()
	service := NewUserService(model)

	deleted, message, code := service.Delete(id)

	if deleted == false {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, nil, http.StatusNoContent)
	}
}
