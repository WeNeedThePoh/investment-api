package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"investment-api/models"
	u "investment-api/utils"
	"net/http"
	"strconv"
	_ "strconv"
)

/*var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Respond(w, resp)
}*/

var GetUser = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	user := models.GetUser(uint(id))
	if user == nil {
		u.Fail(w, "User not found", "", 404)
		return
	}

	u.Success(w, user.ToMap())
}
