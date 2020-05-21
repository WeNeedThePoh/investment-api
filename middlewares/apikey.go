package middlewares

import (
	"github.com/julienschmidt/httprouter"
	u "investment-api/utils"
	"net/http"
	"os"
)

//APIKeyAuthentication API key middleware
var APIKeyAuthentication = func(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		APIKey := r.Header.Get("Api-Key")
		if APIKey == "" {
			APIKey = r.URL.Query().Get("apikey")
		}

		if APIKey == "" {
			u.Fail(w, "Apikey not provided", "", http.StatusForbidden)
			return
		}

		if APIKey == os.Getenv("APP_KEY") {
			h(w, r, ps)
			return
		}

		u.Fail(w, "Apikey is invalid", "", http.StatusUnauthorized)
		return
	}
}
