package middleware

import (
	"context"
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	// w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Loop over header names
		// for name, values := range r.Header {
		// 	// Loop over all values for the name.
		// 	for _, value := range values {
		// 		log.Println(name, value)
		// 	}
		// }
		log.Println("Checking Authentication")
		authorization := r.Header.Get(AUTHORIZATION_KEY)

		// Check that the header begins with a prefix of Bearer
		if !strings.HasPrefix(authorization, BEARER_KEY) {
			log.Println(BEARER_KEY + " not found: " + authorization)
			writeUnauthed(w)
			return
		}

		// Pull out the token
		encodedToken := strings.TrimPrefix(authorization, BEARER_KEY)

		// Decode the token from base 64
		token, err := base64.StdEncoding.DecodeString(encodedToken)
		if err != nil {
			log.Println("error: " + err.Error())
			writeUnauthed(w)
			return
		}

		// We're just assuming a valid base64 token is a valid user id.
		userID := string(token)

		log.Println("userID: " + userID)
		ctx := r.Context()
		ctx = context.WithValue(ctx, AUTH_USER_KEY, userID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
