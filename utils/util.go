package utils

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Success(w http.ResponseWriter, data map[string]interface{}, statusCode int) {
	response := make(map[string]interface{})
	if data != nil {
		response = map[string]interface{}{"data": data}
	}

	if statusCode == 0 {
		statusCode = 200
	}
	Respond(w, response, statusCode)
}

func Fail(w http.ResponseWriter, message string, detail string, statusCode int) {
	data := map[string]interface{}{"message": message, "detail": detail}

	Respond(w, data, statusCode)
}

func Respond(w http.ResponseWriter, data map[string]interface{}, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func RetrieveIdParameter(r *http.Request) uint {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		return 0
	}

	return uint(id)
}
