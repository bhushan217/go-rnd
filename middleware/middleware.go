package middleware

import (
	"net/http"
)
const BEARER_KEY = "Bearer "
const AUTHORIZATION_KEY = "Authorization"
const AUTH_USER_KEY = "middleware.auth.userID"

type Middleware func(http.Handler) http.Handler

func Chain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
