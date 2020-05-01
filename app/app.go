package app

import (
	"context"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"investment-api/services"
	u "investment-api/utils"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/users", "/login"}
		requestPath := r.URL.Path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			u.Fail(w, "Missing auth token", "", http.StatusForbidden)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			u.Fail(w, "Invalid/Malformed auth token", "", http.StatusForbidden)
			return
		}

		tokenPart := splitted[1]
		tk := &services.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			u.Fail(w, "Malformed authentication token", "", http.StatusForbidden)
			return
		}

		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			u.Fail(w, "Request parameter not found", "", http.StatusBadRequest)
			return
		}

		if  tk.UserId != uint(id) {
			u.Fail(w, "Request user id didn't match token user id", "", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
