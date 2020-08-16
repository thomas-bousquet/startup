package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
)

func authenticationMiddleware(next http.Handler) http.Handler {



	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//func isNotAuthorised() {
		//	w.WriteHeader(http.StatusUnauthorized)
		//}

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
		}

		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(SECRETKEY), nil
		})


		next.ServeHTTP(w, r)
	})
}
