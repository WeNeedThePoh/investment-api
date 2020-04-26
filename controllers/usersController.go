package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"investment-api/models"
	u "investment-api/utils"
	"net/http"
	"strconv"
	_ "strconv"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		u.Fail(w, "Error while decoding request body", "", 400)
		return
	}

	user, message, code := models.CreateUser(data)
	if user == nil {
		u.Fail(w, message, "", code)
		return
	}

	resp := user.ToMap()
	delete(resp, "password")
	delete(resp, "password_reset")
	u.Success(w, resp, 200)
}

var GetUser = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		u.Fail(w, "Request parameter not found", "", 400)
	}

	user := models.GetUser(uint(id))
	if user == nil {
		u.Fail(w, "User not found", "", 404)
		return
	}

	resp := user.ToMap()
	delete(resp, "password")
	delete(resp, "password_reset")
	u.Success(w, resp, 200)
}

var UpdateUser = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		u.Fail(w, "Request parameter not found", "", 400)
	}

	data := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		u.Fail(w, "Missing required data", "The body payload can not be empty", 400)
		return
	}

	updated, message, code := models.UpdateUser(uint(id), data)
	if updated == false {
		u.Fail(w, message, "", code)
		return
	}

	u.Success(w, nil, 204)
}

var UpdateUserPassword = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		u.Fail(w, "Request parameter not found", "", 400)
	}

	data := make(map[string] string)
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil || data["password"] == "" || data["old_password"] == "" || len(data) == 0 {
		u.Fail(w, "Missing required data", "The body payload can not be empty", 400)
		return
	}

	updated, message, code := models.UpdateUserPassword(uint(id), data["old_password"], data["password"])
	if updated == false {
		u.Fail(w, message, "", code)
		return
	}

	u.Respond(w, nil, 204)
}

var DeleteUser = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		u.Fail(w, "Request parameter not found", "", 400)
	}

	deleted, message, code := models.DeleteUser(uint(id))
	if deleted == false {
		u.Fail(w, message, "", code)
		return
	}

	u.Success(w, nil, 204)
}
