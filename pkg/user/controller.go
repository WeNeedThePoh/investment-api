package user

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	u "investment-api/utils"
	"net/http"
)

//Create a new user
var Create = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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

//Get get user
var Get = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := u.RetrieveIDParameter(params, "user_id")
	var model = NewUser()
	service := NewUserService(model)

	user, message, code := service.Get(id)
	if user == nil {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, user, http.StatusOK)
	}
}

//Update update user
var Update = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := u.RetrieveIDParameter(params, "user_id")
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

//UpdatePassword update user password
var UpdatePassword = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := u.RetrieveIDParameter(params, "user_id")
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

//Delete delete user
var Delete = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := u.RetrieveIDParameter(params, "user_id")
	var model = NewUser()
	service := NewUserService(model)

	deleted, message, code := service.Delete(id)

	if deleted == false {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, nil, http.StatusNoContent)
	}
}
