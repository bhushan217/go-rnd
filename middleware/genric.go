package middleware

import (
	"log"
	"net/http"
	"strings"
)

func IsAdmin(next http.Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Authorizing Admin Role")
		ctx := r.Context()
		x := ctx.Value(AUTH_USER_KEY).(string)
		if !strings.Contains(x, "admin") {
			w.WriteHeader(http.StatusUnauthorized)
			// http.Error(w, http.StatusText(http.StatusUnauthorized), "Unauthorized", http.StatusUnauthorized)
			return
		}
		log.Println("Admin Role Authorized: " + x)
		next.ServeHTTP(w, r)
	}
}
func EnsureAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Authrizing Admin Role")
		ctx := r.Context()
		x := ctx.Value(AUTH_USER_KEY).(string)
		if !strings.Contains(x, "admin") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Printf("Unauthorized %s", x)
			// w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println("Admin Role Authorized: " + x)
		next.ServeHTTP(w, r)
	})
}

func LoadUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Loading User")
		next.ServeHTTP(w, r)
	})
}

func AllowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Enabling CORS")
		next.ServeHTTP(w, r)
	})
}

func CheckPermissions(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Checking Permissions of User")
		next.ServeHTTP(w, r)
	})
}

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Recover from panics and handle errors
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v", r)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
