package middleware

import (
	"net/http"
)

func AuthSlack(h http.Handler) http.Handler {
	return http.HandlerFunc(authSlackFn)
}

func authSlackFn(w http.ResponseWriter, r *http.Request) {
	// TODO: implement slack api authentication
}
