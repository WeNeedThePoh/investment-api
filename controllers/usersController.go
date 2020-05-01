package controllers

import (
	"encoding/json"
	"investment-api/models"
	"investment-api/services"
	u "investment-api/utils"
	"net/http"
	_ "strconv"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		u.Fail(w, "Error while decoding request body", "", 400)
		return
	}

	var model = models.NewUser()
	service := services.NewUserService(model)
	user, message, code := service.CreateUser(data)

	if user == nil {
		u.Fail(w, message, "", code)
	} else {
		u.Success(w, user, http.StatusCreated)
	}
}

var GetUser = func(w http.ResponseWriter, r *http.Request) {
	id := u.RetrieveIdParameter(r)
	var model = models.NewUser()
	service := services.NewUserService(model)

	user, message, code := service.GetUser(id)
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
		u.Fail(w, "Missing required data", "The body payload can not be empty", 400)
		return
	}

	var model = models.NewUser()
	service := services.NewUserService(model)
	updated, message, code := service.UpdateUser(id, data)

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
		u.Fail(w, "Missing required data", "The body payload can not be empty", 400)
		return
	}

	var model = models.NewUser()
	service := services.NewUserService(model)
	updated, message, code := service.UpdateUserPassword(id, data["old_password"], data["password"])

	if updated == false {
		u.Fail(w, message, "", code)
	} else {
		u.Respond(w, nil, http.StatusNoContent)
	}
}

var DeleteUser = func(w http.ResponseWriter, r *http.Request) {
	id := u.RetrieveIdParameter(r)
	var model = models.NewUser()
	service := services.NewUserService(model)

	deleted, message, code := service.DeleteUser(id)

	if deleted == false {
		u.Fail(w, message, "", code)
		return
	} else {
		u.Success(w, nil, http.StatusNoContent)
	}
}
