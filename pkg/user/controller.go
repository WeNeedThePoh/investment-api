package user

import (
	"encoding/json"
	u "investment-api/utils"
	"net/http"
	_ "strconv"
)

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

var GetUser = func(w http.ResponseWriter, r *http.Request) {
	id := u.RetrieveIdParameter(r)
	var model = NewUser()
	service := NewUserService(model)

	user, message, code := service.Get(id)
	if user == nil {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, user, http.StatusOK)
	}
}

var UpdateUser = func(w http.ResponseWriter, r *http.Request) {
	id := u.RetrieveIdParameter(r)
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

var UpdateUserPassword = func(w http.ResponseWriter, r *http.Request) {
	id := u.RetrieveIdParameter(r)
	data := make(map[string] string)

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

var DeleteUser = func(w http.ResponseWriter, r *http.Request) {
	id := u.RetrieveIdParameter(r)
	var model = NewUser()
	service := NewUserService(model)

	deleted, message, code := service.Delete(id)

	if deleted == false {
		u.Fail(w, message, "", code)
		return
	} else {
		u.Success(w, nil, http.StatusNoContent)
	}
}
