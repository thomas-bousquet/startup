package middlewares

import (
	"encoding/json"
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/thomas-bousquet/startup/errors"
	"github.com/thomas-bousquet/startup/utils/jwt"
	"net/http"
	"strings"
)

type AuthenticationMiddleware struct {
	jwt jwt.JWT
}

func NewAuthenticationMiddleware(jwt jwt.JWT) AuthenticationMiddleware {
	return AuthenticationMiddleware{
		jwt: jwt,
	}
}

func (m AuthenticationMiddleware) Execute(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {


		authorizationHeaderParts := strings.Fields(r.Header.Get("Authorization"))

		if len(authorizationHeaderParts) < 2 || strings.ToLower(authorizationHeaderParts[0]) != "bearer" {
			authorizationError := errors.NewAuthorizationError("Authorization header is not valid")

			body, marshalError := json.Marshal(authorizationError)

			if marshalError != nil {
				http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusUnauthorized)
			_, writeError := w.Write(body)

			if writeError != nil {
				http.Error(w, writeError.Error(), http.StatusInternalServerError)
				return
			}

			return
		}

		authorizationToken := authorizationHeaderParts[1]

		token, err := m.jwt.ParseToken(authorizationToken)

		if err != nil {
			body, marshalError := json.Marshal(err)

			if marshalError != nil {
				http.Error(w, marshalError.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusUnauthorized)
			_, writeError := w.Write(body)

			if writeError != nil {
				http.Error(w, writeError.Error(), http.StatusInternalServerError)
				return
			}

			return
		}

		claims := (*token).Claims.(jwtGo.StandardClaims)

		context.Set(r, "user_id", &claims.Subject)
		next.ServeHTTP(w, r)
	})
}
