package utils

import (
	"encoding/json"
	"net/http"
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
