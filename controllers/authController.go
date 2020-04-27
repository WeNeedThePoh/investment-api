package controllers

import (
	"encoding/json"
	"investment-api/models"
	u "investment-api/utils"
	"net/http"
)

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Fail(w, "Missing required data", "", 400)
		return
	}

	resp, message, code := models.Login(user.Email, user.Password)
	if resp == nil {
		u.Fail(w, message, "", code)
		return
	}

	u.Success(w, resp, 200)
}
