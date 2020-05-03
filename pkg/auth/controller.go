package auth

import (
	"encoding/json"
	"investment-api/pkg/user"
	u "investment-api/utils"
	"net/http"
)

//Authenticate Login user
var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil || data["email"] == "" || data["password"] == "" || len(data) == 0 {
		u.Fail(w, "Missing required data", "The body payload can not be empty", http.StatusBadRequest)
		return
	}

	var model = user.NewUser()
	service := NewAuthService(model)
	resp, message, code := service.Login(data["email"], data["password"])
	if resp == nil {
		u.Fail(w, message, "", code)
		return
	}

	u.Success(w, resp, http.StatusOK)
}
