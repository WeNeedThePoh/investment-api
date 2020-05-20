package utils

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

//Success payload, all responses thar are 2xx
func Success(w http.ResponseWriter, data interface{}, statusCode int) {
	response := make(map[string]interface{})
	if data != nil {
		response = map[string]interface{}{"data": data}
	}

	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	Respond(w, response, statusCode)
}

//Fail payload, all responses that are 4xx
func Fail(w http.ResponseWriter, message string, detail string, statusCode int) {
	data := map[string]interface{}{"message": message, "detail": detail}

	Respond(w, data, statusCode)
}

//Respond prepare payload
func Respond(w http.ResponseWriter, data map[string]interface{}, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

//RetrieveIDParameter retrieve parameter id from request route
func RetrieveIDParameter(params httprouter.Params, parameter string) uint {
	param := params.ByName(parameter)
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return 0
	}

	return uint(id)
}
