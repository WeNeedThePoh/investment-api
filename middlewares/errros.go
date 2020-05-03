package middlewares

import (
	u "investment-api/utils"
	"net/http"
)

//NotFoundHandler Handler for route not found
func NotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u.Fail(w, "This resources was not found on our server", "", http.StatusNotFound)
	})
}
