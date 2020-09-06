package middlewares

import (
	"encoding/json"
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	. "github.com/thomas-bousquet/startup/errors"
	"github.com/thomas-bousquet/startup/repositories"
	"github.com/thomas-bousquet/startup/utils/jwt"
	"net/http"
	"strings"
)

type AuthenticationMiddleware struct {
	jwt jwt.JWT
	userRepository repositories.UserRepository
}

func NewAuthenticationMiddleware(jwt jwt.JWT, userRepository repositories.UserRepository) AuthenticationMiddleware {
	return AuthenticationMiddleware{
		jwt: jwt,
		userRepository: userRepository,
	}
}

func (m AuthenticationMiddleware) ExecuteWithRole(role string) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authorizationHeaderParts := strings.Fields(r.Header.Get("Authorization"))

			if len(authorizationHeaderParts) < 2 || strings.ToLower(authorizationHeaderParts[0]) != "bearer" {
				authorizationError := NewAuthorizationError("Authorization header is not valid")

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

			claims := token.Claims.(*jwtGo.StandardClaims)
			userId := claims.Subject

			user, err := m.userRepository.FindUserWithRole(userId, role)

			if err != nil {
				// TODO: Log properly / find a good logger (ELK ?)
				unexpectedError := NewUnexpectedError()

				body, marshalError := json.Marshal(unexpectedError)

				if marshalError != nil {
					http.Error(w, unexpectedError.Message, http.StatusInternalServerError)
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

			if user == nil {
				authorizationError := NewAuthorizationError("Authorization header is not valid")

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

			context.Set(r, "user_id", &claims.Subject)
			next.ServeHTTP(w, r)
		})
	}

}
