package utils

import (
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter, data map[string]interface{}) {
	response := map[string]interface{}{"data": data}

	Respond(w, response, 200)
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
